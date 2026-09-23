# Code Review Findings — Slice 1 Public Campaign Detail Frontend Mock-Parallel

- Phase: Code Review
- Work Unit / Run: `WU-S1-004` / `CR-FE-001`
- Author: Codex CLI agent
- Role / Specialization: Reviewer / Independent Code Review — Frontend public Campaign detail
- Model / Reasoning / Session: `gpt-5.6-terra` / high / Fresh independent Code Review session
- Created: 2026-09-23
- Target revision: `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
- Current reviewed revision: `1b767c94732f95649299e33e8f4478c2c781f690`
- Workflow revision: not exposed by this Run

Review dilakukan terhadap current frontend scope yang disebut oleh invocation: route detail, API/query boundary, MSW, tests, dan dokumentasi. Perubahan backend/topology serta Harscode Space lain pada current revision tidak termasuk scope Run ini.

## 1. Safety

### F1 — Worker MSW tetap mengintersepsi setelah route detail dibongkar

- **Location:** `frontend/app/campaigns/[campaignId]/mock-service-worker.tsx:16-33`
- **Problem:** Cleanup effect hanya mengubah `active = false`. Setelah `worker.start()` berhasil, tidak ada `worker.stop()`. Bila pengguna meninggalkan route ini pada development mock runtime, worker yang sudah aktif tetap mengintersepsi handler Campaign pada client yang sama. Jika unmount terjadi sebelum dynamic import selesai, chain saat ini masih dapat memulai worker sesudah unmount.
- **Why it matters:** Ini adalah lifecycle/resource leak dan membuat mock opt-in yang seharusnya route-local memengaruhi request setelah pengguna meninggalkan Campaign detail. Hal tersebut dapat menyamarkan request live/route lain pada sesi development dan melemahkan bukti mock-to-real boundary R2.
- **Suggested resolution:** Simpan referensi worker yang benar-benar dimulai; jangan panggil `start()` bila effect sudah dibersihkan; panggil `worker.stop()` sekali pada cleanup bila worker telah aktif. Jangan gunakan `terminate()`, karena Run hanya perlu menghentikan intersepsi client, bukan menghapus registrasi worker.
- **Blocking:** Blocking.
- **Authority:** Approved `TP-FE-001` §11 scope fencing (`mock provider must not become global`); `frontend/AGENTS.md` §2 (MSW pada network boundary) dan Harscode Code Review Safety checklist (resource lifecycle).

## 2. Quality

No findings. Batas API terpusat, query key factory, pemisahan route-local component, serta tidak adanya mirror server state/Zustand proporsional dengan satu surface ini.

## 3. Stack-Specific Best Practices

### F2 — Transisi async tidak memenuhi announcement/focus recovery untuk state retryable

- **Location:** `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx:34-39,60-68`; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx:62-80`
- **Problem:** Ketika `503` menghasilkan `temporarily-unavailable`, user dapat memilih `Coba lagi`. Jika retry berikutnya berhasil, `previousState.current` bernilai `temporarily-unavailable`, tetapi effect hanya memfokuskan `<main>` jika state sebelumnya `request-failure`. Tombol retry kemudian dihapus saat success view menggantikannya, sehingga focus keyboard dapat menjadi lost. Selain itu, state result `temporarily-unavailable` dan `request-failure` tidak berada dalam live region; hanya loading yang memakai `role="status"`, sehingga perubahan loading → hasil retryable tidak diumumkan secara programatis.
- **Why it matters:** R7 Techplan secara eksplisit mewajibkan announced state changes dan retry transition tanpa focus hilang. Pengguna keyboard/assistive technology dapat kehilangan konteks sesudah recovery `503`, atau tidak mengetahui bahwa loading telah berubah ke state unavailable/failure.
- **Suggested resolution:** Perlakukan kedua state retryable sebagai recovery predecessor yang memindahkan focus ke landmark/heading success saat request berhasil, tanpa memindahkan focus pada initial load atau route navigation biasa. Tambahkan live-status semantics yang tepat untuk perubahan async retryable/loading (dengan copy tetap aman dan tanpa menyiarkan raw error), lalu lindungi jalur `503 → retry → success` dengan RTL focus assertion.
- **Blocking:** Blocking.
- **Best-practice authority:** `../harscode-workspace/best-practices/react/accessibility-fundamentals.md` (explicit focus pada async replacement); `../harscode-workspace/best-practices/react/component-test-mocking-discipline.md` (dedicated loading/error tests); `frontend/AGENTS.md` §§10–11 dan Approved `TP-FE-001` R7.

## 4. Consistency

No additional findings. Di luar F1 dan F2, current diff konsisten dengan `AGENTS.md` dan `frontend/AGENTS.md`: memakai `PublicCampaignDetail` generated type, satu request boundary same-origin, MSW pada network boundary tanpa mock branch di request production, TanStack Query tanpa Zustand mirror, rendering plain React text untuk organizer content, serta tidak menambahkan Donate action.

## Verification executed during Review

- `git diff --check a053b48ef3fef35a073c5fe57e5ee4581f848e07 -- frontend/` → passed; tidak ada whitespace error pada frontend scope.
- Tidak ada runtime suite/reproduction dijalankan. F1 dan F2 dapat dipastikan dari lifecycle/state control flow current diff; broad/final frontend verification tetap Testing-owned menurut Approved `TP-FE-001`.

## Verdict

**Request changes**

Blocking findings: F1, F2.

## Phase handoff

- Completed: four-pass independent Code Review + verdict.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-FE-001/review-findings.md`; `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-FE-001/patch-plan.md`.
- Human decision: none.
- Open / deferred: final Indonesian provenance/action wording and Human rendered acceptance tetap requirement milestone; live backend/proxy/media/cache parity tetap deferred ke integration owner. Keduanya bukan pengganti F1/F2.
- Recommended next step: Build/Patch menggunakan patch plan spesifik ini; jangan mulai Testing sebelum patch convergence.
- Session transition: mulai fresh Build/Patch session yang re-grounded pada Techplan, findings ini, dan patch plan. Original Build session sudah selesai; fresh session membuat scope patch tetap eksplisit dan independen dari review.
- Context pointers: `TP-FE-001/techplan.md` R2/R7; F1/F2; `mock-service-worker.tsx`; `campaign-detail-client.tsx`; `campaign-detail-view.tsx`.
