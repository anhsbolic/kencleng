"use client";

import { type ReactNode, useEffect, useState } from "react";
import styles from "./campaign-detail.module.css";

export default function MockServiceWorker({ children }: { children: ReactNode }) {
  const enabled = process.env.NEXT_PUBLIC_MSW_ENABLED === "true";
  const [ready, setReady] = useState(!enabled);
  const [failed, setFailed] = useState(false);

  useEffect(() => {
    if (!enabled) {
      return;
    }

    let active = true;
    let startedWorker: (typeof import("@/mocks/browser"))["worker"] | undefined;

    void import("@/mocks/browser")
      .then(async ({ worker }) => {
        if (!active) {
          return;
        }

        await worker.start({ onUnhandledRequest: "bypass" });

        if (!active) {
          worker.stop();
          return;
        }

        startedWorker = worker;
      })
      .then(() => {
        if (active) {
          setReady(true);
        }
      })
      .catch(() => {
        if (active) {
          setFailed(true);
        }
      });

    return () => {
      active = false;
      startedWorker?.stop();
      startedWorker = undefined;
    };
  }, [enabled]);

  if (failed) {
    return (
      <main aria-labelledby="mock-startup-title" className={styles.page}>
        <section className={styles.state} role="alert">
          <p className={styles.kicker}>Mode pengembangan</p>
          <h1 id="mock-startup-title">Mock tidak dapat disiapkan.</h1>
          <p>Coba muat ulang halaman atau periksa konfigurasi mock lokal.</p>
        </section>
      </main>
    );
  }

  if (!ready) {
    return (
      <main aria-busy="true" aria-live="polite" className={styles.page}>
        <section className={styles.state} role="status">
          <p className={styles.kicker}>Menyiapkan halaman</p>
          <h1>Memuat detail campaign…</h1>
        </section>
      </main>
    );
  }

  return children;
}
