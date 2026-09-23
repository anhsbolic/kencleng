# Patch Plan — CR-FE-001

- Phase: Code Review / Patch Plan
- Work Unit / Run: `WU-S1-004` / `CR-FE-001`
- Author: Codex CLI agent
- Role / Specialization: Reviewer / Independent Code Review — Frontend public Campaign detail
- Model / Reasoning / Session: `gpt-5.6-terra` / high / Fresh independent Code Review session
- Created: 2026-09-23
- Target revision: `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
- Current reviewed revision: `1b767c94732f95649299e33e8f4478c2c781f690`
- Workflow revision: not exposed by this Run

Tujuan patch hanya menutup blocking F1 dan F2 dari `review-findings.md`. Jangan mengubah OpenAPI, backend, topology, kontrak public projection, atau scope donation.

## P1 — Batasi lifecycle browser MSW pada route detail

1. Di `frontend/app/campaigns/[campaignId]/mock-service-worker.tsx`, pertahankan worker aktif hanya selama provider yang enabled masih mounted.
2. Setelah dynamic import selesai, periksa flag mounted sebelum memanggil `worker.start()`; ini mencegah unmount race memulai intersepsi baru.
3. Simpan referensi worker yang telah dimulai dan panggil `worker.stop()` pada cleanup. Cleanup harus idempotent terhadap Strict Mode/dev lifecycle dan tidak menggunakan `terminate()`.
4. Pertahankan behavior current: worker tetap dimulai sebelum query children dirender ketika `NEXT_PUBLIC_MSW_ENABLED=true`; production request function tetap tanpa mock-mode branch.
5. Tambahkan focused test untuk lifecycle provider (mock browser worker secara lokal bila diperlukan): unmount sesudah start menghentikan intersepsi satu kali, dan unmount sebelum import/start tidak kemudian memulai worker. Bila unit mock tidak mampu membuktikan service-worker behavior, catat alasan dan jalankan representative browser navigation check di Testing.

## P2 — Perbaiki announcement dan focus untuk recovery retryable

1. Di `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx`, perluas recovery condition sehingga success setelah `temporarily-unavailable` **maupun** `request-failure` memindahkan focus ke landmark/heading success yang berlabel. Jangan menggeser focus untuk initial successful load atau route transition biasa.
2. Di `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx`, beri perubahan loading → unavailable/request-failure semantics live-status yang sesuai. Pesan harus tetap fixed, public-safe, dan tidak memuat response body/error message.
3. Tambahkan RTL/MSW coverage yang mengubah handler `unavailable` menjadi successful fixture setelah `Coba lagi`, lalu assert success heading dan focus pada landmark/heading yang relevan. Tambahkan assertion role/live status untuk retryable state sesuai semantic yang dipilih.
4. Pertahankan native button, named retry control, dan behavior no-retry untuk safe `404`.

## Scoped verification for Build/Patch

- Jalankan focused Vitest tests yang diubah/ditambahkan untuk lifecycle mock dan async retry/focus.
- Jalankan `cd frontend && npm run verify` setelah patch.
- Jangan jalankan broad Testing-owned build/browser matrix pada patch loop kecuali perlu untuk membuktikan finding. Independent Testing tetap harus menjalankan scope yang telah ditetapkan Approved `TP-FE-001`, termasuk responsive/state inspection dan Human rendered acceptance sebelum milestone claim.

## Completion criteria

- Navigasi keluar dari enabled detail route tidak meninggalkan MSW intersepsi client aktif dari provider route tersebut.
- Success sesudah retry `503` tidak meninggalkan focus pada node yang telah dihapus, dan state async retryable mempunyai announcement semantics.
- Tidak ada mock fixture/environment branch baru di `frontend/lib/api/`; API path, generated types, public-safe 404 behavior, serta non-activating donation context tetap tidak berubah.
