import type { RefObject } from "react";
import type { PublicCampaignDetail } from "@/lib/api/public-campaign";
import styles from "./campaign-detail.module.css";

type ViewProps = {
  campaign?: PublicCampaignDetail;
  isRetrying?: boolean;
  mainRef: RefObject<HTMLElement | null>;
  onRetry?: () => void;
  state: "loading" | "not-found" | "temporarily-unavailable" | "request-failure" | "success";
};

const relationshipCopy = {
  below_target: "Masih di bawah target",
  target_reached: "Target telah tercapai",
  above_target: "Melebihi target",
  not_computable: "Perbandingan belum dapat ditampilkan",
} as const;

function formatIdr(amount: string) {
  const [whole, fraction] = amount.split(".");
  return `Rp ${whole.replace(/\B(?=(\d{3})+(?!\d))/g, ".")},${fraction}`;
}

function formatDate(value: string) {
  return new Intl.DateTimeFormat("id-ID", {
    day: "numeric",
    month: "long",
    year: "numeric",
  }).format(new Date(value));
}

function StatePage({
  isRetrying,
  mainRef,
  onRetry,
  state,
}: Omit<ViewProps, "campaign">) {
  const content = {
    loading: {
      eyebrow: "Memuat campaign",
      title: "Menyiapkan informasi campaign…",
      description: "Susunan halaman sedang dipersiapkan.",
    },
    "not-found": {
      eyebrow: "Campaign tidak tersedia",
      title: "Campaign ini tidak dapat ditampilkan.",
      description: "Tautan mungkin sudah tidak tersedia untuk dilihat publik.",
    },
    "temporarily-unavailable": {
      eyebrow: "Informasi belum siap",
      title: "Detail campaign sedang tidak tersedia.",
      description: "Coba lagi beberapa saat lagi.",
    },
    "request-failure": {
      eyebrow: "Permintaan belum berhasil",
      title: "Detail campaign belum dapat dimuat.",
      description: "Coba lagi untuk meminta informasi campaign.",
    },
  }[state as Exclude<ViewProps["state"], "success">];

  return (
    <main
      aria-busy={state === "loading"}
      aria-labelledby="campaign-state-title"
      className={styles.page}
      ref={mainRef}
      tabIndex={-1}
    >
      <section className={styles.state} role={state === "loading" ? "status" : undefined}>
        <p className={styles.kicker}>{content.eyebrow}</p>
        <h1 id="campaign-state-title">{content.title}</h1>
        <p>{content.description}</p>
        {onRetry ? (
          <button className={styles.retry} onClick={onRetry} type="button">
            {isRetrying ? "Mencoba lagi…" : "Coba lagi"}
          </button>
        ) : null}
      </section>
    </main>
  );
}

function Funding({ campaign }: { campaign: PublicCampaignDetail }) {
  const { funding } = campaign;

  if (funding.availability === "unavailable") {
    return (
      <section aria-labelledby="funding-title" className={styles.funding}>
        <p className={styles.kicker}>Penggalangan dana</p>
        <h2 id="funding-title">Informasi dana belum tersedia.</h2>
        <p>Jumlah terkumpul dan target belum dapat ditampilkan saat ini.</p>
      </section>
    );
  }

  return (
    <section aria-labelledby="funding-title" className={styles.funding}>
      <p className={styles.kicker}>Penggalangan dana</p>
      <h2 id="funding-title">Dana terkumpul</h2>
      <p className={styles.amount}>{formatIdr(funding.collected_amount)}</p>
      <dl className={styles.fundingFacts}>
        <div>
          <dt>Target dana</dt>
          <dd>{formatIdr(funding.target_amount)}</dd>
        </div>
        <div>
          <dt>Perbandingan dengan target</dt>
          <dd>{relationshipCopy[funding.progress.relationship]}</dd>
        </div>
        {funding.progress.percentage !== null ? (
          <div>
            <dt>Persentase dari sumber data</dt>
            <dd>{funding.progress.percentage}%</dd>
          </div>
        ) : null}
      </dl>
    </section>
  );
}

function Media({ campaign }: { campaign: PublicCampaignDetail }) {
  if (campaign.media.state === "absent") {
    return (
      <section aria-labelledby="media-title" className={styles.mediaState}>
        <p className={styles.kicker}>Media campaign</p>
        <h2 id="media-title">Belum ada media untuk ditampilkan.</h2>
        <p>Informasi campaign tetap tersedia tanpa gambar pendukung.</p>
      </section>
    );
  }

  if (campaign.media.state === "unavailable") {
    return (
      <section aria-labelledby="media-title" className={styles.mediaState}>
        <p className={styles.kicker}>Media campaign</p>
        <h2 id="media-title">Media sementara belum tersedia.</h2>
        <p>Informasi media sedang tidak dapat ditampilkan; ini berbeda dari belum ada media.</p>
      </section>
    );
  }

  return (
    <section aria-labelledby="media-title" className={styles.media}>
      <div className={styles.sectionHeading}>
        <p className={styles.kicker}>Media campaign</p>
        <h2 id="media-title">Gambar dari pengelola</h2>
      </div>
      <div className={styles.mediaGrid}>
        {campaign.media.items.map((item) => (
          <figure key={item.id}>
            {/* The API-owned path is an opaque controlled reference. */}
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img alt={item.alt_text} src={item.content_url} />
            {item.caption ? <figcaption>{item.caption}</figcaption> : null}
            <p className={styles.provenance}>Dari pengelola</p>
          </figure>
        ))}
      </div>
    </section>
  );
}

function CampaignSuccess({ campaign, mainRef }: Required<Pick<ViewProps, "campaign" | "mainRef">>) {
  return (
    <main aria-labelledby="campaign-title" className={styles.page} ref={mainRef} tabIndex={-1}>
      <article className={styles.detail}>
        <header className={styles.hero}>
          <p className={styles.kicker}>Campaign publik</p>
          <h1 id="campaign-title">{campaign.title}</h1>
          <p className={styles.steward}>Dikelola oleh {campaign.steward.name}</p>
          <p className={styles.lifecycle}>Penggalangan sedang berlangsung hingga {formatDate(campaign.lifecycle.fundraising_ends_at)}.</p>
        </header>

        <div className={styles.primaryFacts}>
          <Funding campaign={campaign} />
          <aside aria-labelledby="action-title" className={styles.action}>
            <p className={styles.kicker}>Langkah berikutnya</p>
            <h2 id="action-title">Dukungan belum dapat dilakukan dari halaman ini.</h2>
            <p>Alur donasi untuk campaign ini belum tersedia. Informasi campaign tetap dapat dibaca di sini.</p>
          </aside>
        </div>

        <section aria-labelledby="purpose-title" className={styles.story}>
          <p className={styles.kicker}>Tujuan</p>
          <h2 id="purpose-title">Apa yang sedang diupayakan</h2>
          <p>{campaign.purpose.content}</p>
          <p className={styles.provenance}>Dari pengelola</p>
        </section>

        <section aria-labelledby="story-title" className={styles.story}>
          <p className={styles.kicker}>Cerita pengelola</p>
          <h2 id="story-title">Konteks campaign</h2>
          <p>{campaign.story.content}</p>
          <p className={styles.provenance}>Dari pengelola</p>
        </section>

        <Media campaign={campaign} />
      </article>
    </main>
  );
}

export default function CampaignDetailView(props: ViewProps) {
  if (props.state !== "success") {
    return <StatePage {...props} />;
  }

  return <CampaignSuccess campaign={props.campaign!} mainRef={props.mainRef} />;
}
