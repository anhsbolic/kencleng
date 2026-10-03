"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { QueryClient, QueryClientProvider, useQuery } from "@tanstack/react-query";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import type { SubmitDonationRequest } from "@/lib/api/donation";
import { submitDonation } from "@/lib/api/donation";
import { getPublicCampaignDetail } from "@/lib/api/public-campaign";
import styles from "./donation.module.css";

const donationFormSchema = z.object({
  amount: z.string().regex(/^\d+$/, "Masukkan jumlah dalam Rupiah utuh.")
    .refine((value) => value.length > 4 || (value.length === 4 && value >= "5000"), "Minimum donasi adalah Rp 5.000."),
  guest_name: z.string().max(120, "Nama terlalu panjang."),
  emailOptIn: z.boolean(),
  guest_email: z.string(),
}).superRefine((values, context) => {
  if (values.emailOptIn && !z.email().safeParse(values.guest_email).success) {
    context.addIssue({
      code: "custom",
      path: ["guest_email"],
      message: "Masukkan alamat email yang valid untuk menerima pemberitahuan status.",
    });
  }
});

type FormValues = z.infer<typeof donationFormSchema>;
type RetryIntent = { payload: SubmitDonationRequest; key: string };

function formatCap(amount: string) {
  return "Rp " + amount.replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

function DonationContent({ campaignId }: { campaignId: string }) {
  const router = useRouter();
  const campaignQuery = useQuery({
    queryKey: ["campaign", "detail", campaignId],
    queryFn: () => getPublicCampaignDetail(campaignId),
    retry: false,
    refetchOnWindowFocus: false,
  });
  const [submitting, setSubmitting] = useState(false);
  const [retryIntent, setRetryIntent] = useState<RetryIntent | null>(null);
  const [requestError, setRequestError] = useState<string | null>(null);
  const [emailOptIn, setEmailOptIn] = useState(false);
  const form = useForm<FormValues>({
    resolver: zodResolver(donationFormSchema),
    defaultValues: { amount: "", guest_name: "", emailOptIn: false, guest_email: "" },
  });

  async function send(payload: SubmitDonationRequest, key: string) {
    if (submitting) return;
    setSubmitting(true);
    setRequestError(null);
    try {
      const result = await submitDonation(campaignId, payload, key);
      if (result.kind === "accepted") {
        router.push(
          "/donations/" + encodeURIComponent(result.data.id) + "/status#" +
            encodeURIComponent(result.data.status_token),
        );
        return;
      }
      if (result.kind === "ambiguous") {
        setRetryIntent({ payload, key });
        setRequestError("Hasil pengiriman belum dapat dipastikan. Coba kirim ulang dengan permintaan yang sama.");
        return;
      }
      setRetryIntent(null);
      if (result.kind === "validation-error") {
        if (result.fields.length === 0 || result.fields.includes("amount")) {
          form.setError("amount", {
            type: "server",
            message: "Jumlah belum dapat diproses. Periksa jumlah dan batas per donasi yang ditampilkan.",
          });
        } else {
          setRequestError("Data donasi belum dapat diproses. Periksa kembali isian Anda.");
        }
      } else if (result.kind === "conflict") {
        setRequestError("Campaign ini sudah tidak menerima donasi baru atau permintaan perlu dimulai ulang.");
      } else {
        setRequestError("Donasi belum dapat dikirim. Periksa kembali detail campaign sebelum mencoba lagi.");
      }
    } catch {
      setRetryIntent({ payload, key });
      setRequestError("Hasil pengiriman belum dapat dipastikan. Coba kirim ulang dengan permintaan yang sama.");
    } finally {
      setSubmitting(false);
    }
  }

  function onValid(values: FormValues) {
    const payload: SubmitDonationRequest = {
      amount: values.amount,
      currency_code: "IDR",
      payment_method: "qris",
      guest_email_status_opt_in: values.emailOptIn,
      is_anonymous: false,
      ...(values.guest_name.trim() ? { guest_name: values.guest_name.trim() } : {}),
      ...(values.emailOptIn ? { guest_email: values.guest_email.trim() } : {}),
    };
    void send(payload, globalThis.crypto.randomUUID());
  }

  if (campaignQuery.isPending) {
    return <State title="Memuat detail campaign…" message="Informasi donasi sedang disiapkan." />;
  }
  if (campaignQuery.isError || campaignQuery.data.kind !== "success") {
    return <State title="Donasi belum dapat dimulai." message="Detail campaign tidak tersedia saat ini." />;
  }

  const campaign = campaignQuery.data.data;
  if (campaign.donation_action.availability !== "available") {
    return <State
      title="Campaign ini tidak menerima donasi baru."
      message="Kelayakan diperiksa kembali oleh server saat donasi dikirim."
      backHref={"/campaigns/" + encodeURIComponent(campaignId)}
    />;
  }

  const locked = submitting || retryIntent !== null;
  return (
    <main className={styles.page}>
      <div className={styles.back}><Link href={"/campaigns/" + encodeURIComponent(campaignId)}>Kembali ke campaign</Link></div>
      <article className={styles.content}>
        <header className={styles.header}>
          <p className={styles.kicker}>Donasi tamu · simulasi</p>
          <h1>Berikan dukungan untuk {campaign.title}</h1>
          <p>Pengelola: {campaign.steward.name}</p>
        </header>
        <section aria-labelledby="cap-title" className={styles.cap}>
          <p className={styles.kicker}>Batas per donasi</p>
          <h2 id="cap-title">{formatCap(campaign.max_donation_amount.amount)} {campaign.max_donation_amount.currency_code}</h2>
          <p>Batas maksimum tiap donasi pada detail campaign ini. Jumlah yang dapat diterima tetap ditentukan saat permintaan diproses.</p>
        </section>

        <form className={styles.form} noValidate onSubmit={form.handleSubmit(onValid)}>
          <h2>Detail donasi</h2>
          <div className={styles.field}>
            <label htmlFor="donation-amount">Jumlah donasi (IDR)</label>
            <p id="amount-help">Minimum Rp 5.000. Masukkan Rupiah utuh; batas pada detail adalah konteks, bukan jaminan penerimaan.</p>
            <div className={styles.amountInput}>
              <span aria-hidden="true">Rp</span>
              <input
                aria-describedby={form.formState.errors.amount ? "amount-error" : "amount-help"}
                aria-invalid={Boolean(form.formState.errors.amount)}
                autoComplete="off"
                disabled={locked}
                id="donation-amount"
                inputMode="numeric"
                {...form.register("amount")}
              />
            </div>
            {form.formState.errors.amount ? <p className={styles.fieldError} id="amount-error">{form.formState.errors.amount.message}</p> : null}
          </div>

          <div className={styles.field}>
            <label htmlFor="guest-name">Nama (opsional)</label>
            <input disabled={locked} id="guest-name" maxLength={120} {...form.register("guest_name")} />
            <p>Nama tidak ditampilkan sebagai informasi publik.</p>
          </div>

          <fieldset className={styles.methods}>
            <legend>Metode yang ditampilkan</legend>
            <p className={styles.methodActive}>QRIS <span>Simulasi sandbox · aktif</span></p>
            <ul>
              <li>GoPay <span>Tidak tersedia</span></li>
              <li>ShopeePay <span>Tidak tersedia</span></li>
              <li>Transfer bank <span>Tidak tersedia</span></li>
            </ul>
          </fieldset>

          <div className={styles.emailOptIn}>
            <label>
              <input
                disabled={locked}
                type="checkbox"
                {...form.register("emailOptIn", { onChange: (event) => setEmailOptIn(event.target.checked) })}
              />
              Kirim pemberitahuan status donasi melalui email (opsional)
            </label>
            <p>Verifikasi email dalam 24 jam sejak alamat dicatat. Jika tidak diverifikasi, alamat dihapus dan pemberitahuan tidak dikirim. Donasi tetap berjalan.</p>
            {emailOptIn ? (
              <div className={styles.field}>
                <label htmlFor="guest-email">Alamat email</label>
                <input
                  aria-invalid={Boolean(form.formState.errors.guest_email)}
                  aria-describedby={form.formState.errors.guest_email ? "guest-email-error" : undefined}
                  autoComplete="email"
                  disabled={locked}
                  id="guest-email"
                  type="email"
                  {...form.register("guest_email")}
                />
                {form.formState.errors.guest_email ? <p className={styles.fieldError} id="guest-email-error">{form.formState.errors.guest_email.message}</p> : null}
              </div>
            ) : null}
          </div>

          {requestError ? <p className={styles.requestError} role="alert">{requestError}</p> : null}
          {retryIntent ? (
            <button className={styles.primary} disabled={submitting} type="button" onClick={() => void send(retryIntent.payload, retryIntent.key)}>
              {submitting ? "Mengirim ulang…" : "Coba kirim ulang"}
            </button>
          ) : (
            <button className={styles.primary} disabled={submitting} type="submit">
              {submitting ? "Mengirim donasi…" : "Kirim donasi simulasi"}
            </button>
          )}
          <p className={styles.sandboxNote}>Ini simulasi sandbox. QRIS di sini tidak memproses pembayaran atau penyelesaian dana sungguhan.</p>
        </form>
      </article>
    </main>
  );
}

function State({ title, message, backHref }: { title: string; message: string; backHref?: string }) {
  return (
    <main className={styles.page}>
      <section className={styles.state} role="status">
        <p className={styles.kicker}>Donasi tamu</p>
        <h1>{title}</h1>
        <p>{message}</p>
        {backHref ? <Link href={backHref}>Kembali ke campaign</Link> : null}
      </section>
    </main>
  );
}

export default function DonationClient() {
  const params = useParams<{ campaignId: string }>();
  const [queryClient] = useState(() => new QueryClient({ defaultOptions: { queries: { retry: false } } }));
  return (
    <QueryClientProvider client={queryClient}>
      <DonationContent campaignId={params.campaignId} />
    </QueryClientProvider>
  );
}
