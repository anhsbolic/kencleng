import { CheckCircle2, CreditCard, FileText } from "lucide-react";

const STEPS = [
  {
    icon: CheckCircle2,
    title: "Temukan kampanye terverifikasi",
    body: "Setiap organisasi diperiksa tim verifikasi Kencleng sebelum kampanyenya tampil.",
  },
  {
    icon: CreditCard,
    title: "Donasi aman dengan berbagai metode",
    body: "Transfer bank, e-wallet, atau QRIS. Minimal Rp 10.000, tanpa biaya tersembunyi.",
  },
  {
    icon: FileText,
    title: "Pantau transparansi penyaluran dana",
    body: "Laporan penggunaan dana terbit berkala dan dapat Anda baca kapan saja.",
  },
];

/**
 * Static 3-step explainer — carried over from the Tier 1 prototype.
 * Not in `page-map.md`/`patterns.md`'s resolved scope for `/`, but kept
 * (techplan Decision 3): accurate static copy describing the real,
 * already-documented product flow, no data dependency — unlike
 * `TrustStrip`, there's no fabricated-numbers risk here.
 */
export function HowItWorks() {
  return (
    <div id="cara-kerja" className="border-y border-neutral-200 bg-white">
      <div className="mx-auto flex max-w-[1360px] flex-col gap-5 px-4 py-7 md:gap-7 md:px-6 md:py-14">
        <div className="flex flex-col gap-1.5">
          <span className="text-[11px] font-bold tracking-wide text-neutral-400 uppercase">
            Cara kerja
          </span>
          <h2 className="text-h2 font-bold text-neutral-900">
            Tiga langkah, dari niat sampai laporan
          </h2>
        </div>

        <div className="grid grid-cols-1 gap-3 md:grid-cols-3 md:gap-4.5">
          {STEPS.map((step, index) => (
            <div
              key={step.title}
              className="flex flex-col gap-3 rounded-lg border border-neutral-200 bg-neutral-50 p-4"
            >
              <div className="flex items-center gap-2.5">
                <span className="inline-flex size-9 items-center justify-center rounded-md bg-primary-50 text-primary-700">
                  <step.icon aria-hidden="true" className="size-4.5" />
                </span>
                <span className="text-[11px] font-bold tracking-wide text-neutral-400 uppercase">
                  Langkah {index + 1}
                </span>
              </div>
              <h3 className="text-h3 font-bold text-neutral-900">{step.title}</h3>
              <p className="text-body text-neutral-700">{step.body}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
