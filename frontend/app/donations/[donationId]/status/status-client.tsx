"use client";

import { useParams } from "next/navigation";
import { useCallback, useEffect, useRef, useState } from "react";
import { getDonationStatus } from "@/lib/api/donation";
import styles from "./status.module.css";

type ViewState =
  | { kind: "loading" }
  | { kind: "pending" }
  | { kind: "success" }
  | { kind: "failed" }
  | { kind: "unavailable" };

function StatusContent({ donationId }: { donationId: string }) {
  const [credential, setCredential] = useState<string | null>(null);
  const [view, setView] = useState<ViewState>({ kind: "loading" });
  const [checking, setChecking] = useState(false);
  const startedFor = useRef<string | null>(null);

  const checkStatus = useCallback(async (token: string) => {
    setChecking(true);
    try {
      const result = await getDonationStatus(donationId, token);
      if (result.kind === "unavailable") {
        setView({ kind: "unavailable" });
      } else {
        setView({ kind: result.data.status });
      }
    } catch {
      setView({ kind: "unavailable" });
    } finally {
      setChecking(false);
    }
  }, [donationId]);

  useEffect(() => {
    if (startedFor.current === donationId) return;
    startedFor.current = donationId;
    const fragment = window.location.hash.slice(1);
    const cleanUrl = window.location.pathname + window.location.search;
    window.history.replaceState(window.history.state, "", cleanUrl);
    let token = "";
    try {
      token = decodeURIComponent(fragment);
    } catch {
      token = "";
    }
    if (!token) {
      queueMicrotask(() => setView({ kind: "unavailable" }));
      return;
    }
    queueMicrotask(() => setCredential(token));
    queueMicrotask(() => { void checkStatus(token); });
  }, [checkStatus, donationId]);

  const terminal = view.kind === "success" || view.kind === "failed";
  const label = view.kind === "pending"
    ? "Menunggu hasil simulasi"
    : "Hasil simulasi donasi: " + (view.kind === "success" ? "berhasil" : "gagal");
  const announcement = checking || view.kind === "loading"
    ? "Memeriksa status donasi…"
    : view.kind === "unavailable"
      ? "Link status tidak tersedia atau mungkin kedaluwarsa."
      : label;
  return (
    <main aria-busy={view.kind === "loading" || checking} className={styles.page}>
      <p
        aria-live="polite"
        role="status"
        style={{
          position: "absolute",
          width: 1,
          height: 1,
          padding: 0,
          margin: -1,
          overflow: "hidden",
          clip: "rect(0, 0, 0, 0)",
          whiteSpace: "nowrap",
          border: 0,
        }}
      >
        {announcement}
      </p>
      {view.kind === "unavailable" ? (
        <section aria-labelledby="status-title" className={styles.state}>
          <p className={styles.kicker}>Status donasi</p>
          <h1 id="status-title">Link status tidak tersedia atau mungkin kedaluwarsa.</h1>
        </section>
      ) : view.kind === "loading" ? (
        <section aria-labelledby="status-title" className={styles.state}>
          <p className={styles.kicker}>Status donasi</p>
          <h1 id="status-title">Memeriksa status donasi…</h1>
        </section>
      ) : (
        <section aria-labelledby="status-title" className={styles.state}>
          <p className={styles.kicker}>Status donasi</p>
          <h1 id="status-title">{label}</h1>
          <p>{terminal ? "Status ini berasal dari hasil simulasi sandbox." : "Belum ada hasil simulasi yang tersedia."}</p>
          {!terminal && credential ? (
            <button disabled={checking} onClick={() => void checkStatus(credential)} type="button">
              {checking ? "Memeriksa…" : "Periksa status lagi"}
            </button>
          ) : null}
        </section>
      )}
    </main>
  );
}

export default function StatusClient() {
  const params = useParams<{ donationId: string }>();
  return <StatusContent donationId={params.donationId} />;
}
