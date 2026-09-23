# Laporan Pengujian — TST-FE-001

> Fase: Testing independen  
> Work Unit / Run: `WU-S1-004` / `TST-FE-001`  
> Peran / spesialisasi: Verifier / Public Campaign frontend mock-verified experience  
> Peserta: Codex CLI agent  
> Model / reasoning / sesi: `gpt-5.6-terra` / high / Fresh independent Testing session  
> Dibuat: 2026-09-23  
> Revisi yang diuji: `a545083`  
> Revisi workflow: tidak diekspos oleh Run

## 0. Ringkasan sweep

- **Terkonfirmasi:** 13 test Vitest yang telah ada, termasuk lifecycle owner MSW Strict Mode serta recovery `503 → success` dan `request-failure → success`, dijalankan ulang melalui `npm run test` dan lulus.
- **Terkonfirmasi dari patch/review sebelumnya:** lifecycle singleton browser MSW tidak menghentikan owner aktif sesudah completion startup stale; RTL membuktikan focus landmark kembali pada kedua recovery retryable. Tidak ada regresi pada verifikasi ulang independen.
- **Ditutup dari deferred Build:** inspeksi browser nyata pada MSW, handler controlled media, desktop/mobile, states public-safe, focus/announcement, dan overflow dijalankan pada Run ini.
- **Masih memerlukan Testing manusia:** penerimaan rendered oleh manusia serta persetujuan wording Indonesia provenance/action yang final. Keduanya adalah gate Human; tidak dapat ditandai lulus oleh verifier ini.

## 0a. Eksekusi Test Focus Pointer

| Area | Evidence anchor dibuka | Verifikasi terspesialisasi | Hasil |
|---|---|---|---|
| Public projection / disclosure / money truth | `EXP-FE-001/evidence/gap-analysis.md#area-1--product-reconciled-campaign-behavior-and-public-api-contract` | Browser MSW memuat fixtures available, funding unavailable, zero, above-target, dan not-computable; memastikan nilai/relationship supplied tampil tanpa CTA donation. Safe 404 diperiksa sebagai satu state non-retryable. | Lulus |
| Controlled media boundary | `EXP-FE-001/evidence/gap-analysis.md#area-4--generated-types-network-boundary-msw-and-live-integration-dependency` | Browser merekam `GET /api/campaigns/{id}` lalu `GET` exact opaque `content_url`; keduanya ditangani MSW `200`. Tidak ada mock branch pada `frontend/lib/api/`; live proxy/storage/cache tetap sengaja di luar bukti ini. | Lulus dalam batas mock-parallel |
| Hostile organizer content / a11y / rendered state | `EXP-FE-001/evidence/gap-analysis.md#area-5--observable-state-hostile-content-accessibility-and-rendered-evidence` | Browser dan RTL memeriksa literal markup-like text tanpa elemen `strong`, live status `polite`, tombol retry bernama dengan focus-visible outline, loading/error semantics, serta desktop/mobile state inspection. | Lulus |
| Concurrency / performance runtime | `EXP-FE-001/evidence/solutioning.md#consequences-and-carried-risks` | Tidak ada test load/concurrency baru: D5 meniadakan polling/retry otomatis, tidak ada mutation/shared mutable state/hot-path baru. Lifecycle overlap yang relevan sudah dijalankan kembali dalam test MSW terarah. | N/A — proporsional dan tidak ada drift Techplan |

## 1. Cakupan pengujian

| Rule / skenario | Kategori | Verifikasi observable | Hasil |
|---|---|---|---|
| R1 — request detail typed/path | Contract/network | `public-campaign.test.ts` yang dijalankan melalui suite penuh membuktikan path encoded `/api/campaigns/{campaignId}` dan fixture generated-type. Browser MSW juga merekam request detail real route. | Lulus |
| R2 — jalur mock-ke-real tunggal | Contract/network | Browser `localhost` dengan `NEXT_PUBLIC_MSW_ENABLED=true` menampilkan log MSW untuk exact detail path dan media `content_url`; tidak ada request ke endpoint alternatif. Inspeksi `lib/api/` mengonfirmasi tidak ada import fixture/flag mock/URL alternatif. | Lulus |
| R3 — 404, 503, generic failure aman | Negative/error | RTL full suite mencakup 404, 503, network failure dan retry recovery. Browser mobile memeriksa 404 tanpa retry maupun alasan internal dan 503 dengan retry bernama. | Lulus |
| R4 — public projection dan funding truth | Truth/edge | Browser memverifikasi `Rp 0,00`, target besar `Rp 99.999.999.999.999.999,99`, `0%`, `Rp 6.250.000,00`, `125.50%`, `Melebihi target`, funding unavailable tanpa nilai rekayasa, serta `not_computable` tanpa baris persentase. | Lulus |
| R5 — media truth | Contract/rendered | Browser memverifikasi image memakai path opaque supplied dan handler PNG menjawab `200`; desktop/mobile memeriksa available, absent, unavailable sebagai state berlainan. RTL membuktikan unavailable berbeda dari absent. Caption non-null tampil sesuai fixture; cabang caption nullable diimplementasikan kondisional dan tidak memiliki fixture browser khusus. | Lulus — cabang nullable tidak mendapat regression fixture khusus |
| R6 — tidak ada aksi donasi simulasi | Scope/negative | Pada wide/mobile available, absent, unavailable-media, funding unavailable, 404, dan 503, query browser menemukan nol link/button bernama donate/donasi; action region memberi konteks non-aktif. | Lulus |
| R7 — async UI aksesibel | Accessibility | Suite penuh menguji announcement/retry/focus recovery. Browser memeriksa `aria-live="polite"` pada unavailable, focus aktual pada retry, dan outline `solid`; hostile text tampil literal tanpa HTML. | Lulus |
| R8 — hierarchy responsif | Rendered | Screenshot/inspeksi browser pada 1440×900 available serta 375×812 long-content, absent, unavailable-media, 404, dan 503 menghasilkan `scrollWidth - innerWidth = 0`; heading, source, funding, media, dan action tetap dapat dibaca. | Lulus |
| R9 — batas milestone jujur | Handoff/human | Bukti ini hanya membuktikan mock network/route UI. Backend projection, proxy, controlled-media enforcement, storage/retraction, cache, serta response/timing parity tetap deferred; Human gate belum ditutup. | Lulus secara dokumentasi; gate Human masih terbuka |

## 2. Verifikasi error

| Error case | Perilaku/kategori diharapkan | Aktual | Actionable/propagated benar? |
|---|---|---|---|
| `404` | State public-safe tanpa disclosure, tanpa retry | Heading `Campaign ini tidak dapat ditampilkan.`; tidak ada retry atau detail malformed/non-public/diagnostic. | Ya |
| `503` | Unavailable retryable, announced | Heading unavailable, `role=status`, `aria-live=polite`, tombol `Coba lagi`. RTL membuktikan recovery ke success dan focus landmark. | Ya |
| Network/generic failure | Pesan umum retryable tanpa body/error mentah | RTL menjalankan `HttpResponse.error()` lalu recovery success/focus; test API membuktikan body `500` tidak diekspos lewat error. | Ya |
| Browser mock startup | Tidak mengirim query sebelum worker siap; failure tetap aman | Browser startup berhasil lalu query/controlled-media ditangani MSW. Test lifecycle mencakup unmount/pending/Strict Mode. | Ya |

## 3. Verifikasi final

- **Perintah final target repo:**
  - `cd frontend && npm run verify` → lulus: ESLint tanpa error dan Vitest 4 files / 13 tests lulus. Satu warning pre-existing pada generated `public/mockServiceWorker.js` (`Unused eslint-disable directive`), bukan kegagalan.
  - `cd frontend && npm run build` → lulus setelah build diizinkan mengakses Google Fonts; Next 16.2.12 mengompilasi, type-check, dan menghasilkan route dinamis `/campaigns/[campaignId]`.
  - `git diff --check` → lulus.
- **Browser rendered evidence:** `NEXT_PUBLIC_MSW_ENABLED=true npm run dev`, kemudian Chromium headless pada `localhost`. Host `127.0.0.1` sempat ditolak Next dev untuk HMR karena `allowedDevOrigins`; verifikasi final memakai host yang dikonfigurasi (`localhost`) dan lulus. Ini batas tooling development, bukan defect route.
- **Broad checks yang sengaja tidak dijalankan:** `npm run test:browser` tidak dijalankan. Techplan D8 tidak meminta committed Playwright regression; browser nyata tetap dijalankan secara terarah untuk risk R2/R5/R8/R7. Tidak ada test browser committed yang perlu dijalankan.
- **Migration/schema collision:** N/A — tidak ada persistence, migration, atau perubahan schema/OpenAPI pada Work Unit frontend ini.
- **Backward compatibility:** N/A — clean-start frontend tidak memiliki consumer Campaign terdahulu; route memakai operation/generated type yang ada tanpa mengubah kontrak.
- **Broader-suite cross-cutting:** `npm run verify` dan production build telah dijalankan. Tidak ada perubahan lintas stack/proxy/API yang memerlukan suite backend atau integration.
- **Konsistensi Techplan segar:** dibaca ulang end-to-end; tidak ditemukan kontradiksi atau Techplan drift material.

## 4. Pola bug berulang baru

Tidak ada.

## Verdict

**Lulus dengan tindak lanjut yang ditandai.**

Tidak ada kegagalan produksi atau patch plan Testing. Implementasi frontend memenuhi bukti mock-parallel independen. Namun Run ini **tidak** menutup dua gate Human berikut dan karenanya tidak mempromosikan `FRONTEND_MOCK_VERIFIED`:

1. Persetujuan Human Design untuk wording Indonesia final bagi provenance dan unavailable action.
2. Human rendered acceptance pada route mock representative desktop/mobile (success/media, not-found, retryable unavailable/failure).

Selain itu, live backend/proxy/media/cache/storage/timing parity tetap deferred ke WU-S1-003/WU-S1-005/WU-S1-006 dan bukan bukti yang dapat diberikan frontend mock ini.

## Phase handoff

- **Completed:** verifikasi independen rule R1–R9 melalui suite Vitest, production build, network-boundary browser MSW, dan inspeksi rendered responsif.
- **Artifacts:** `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TST-FE-001/testing-report-1.md`
- **Human decision:** Human Design perlu menyetujui wording final; human perlu menjalankan rendered acceptance yang proporsional.
- **Open / deferred:** gate Human di atas; real backend projection, same-origin proxy, controlled-media enforcement, cache/storage/retraction, dan timing parity.
- **Recommended next step:** jalankan Human rendered acceptance. Setelah gate Human tertutup, orchestration owner dapat menentukan status milestone berdasarkan laporan ini; jangan klaim integrated verification.
- **Session transition:** fresh PR/finalization session berguna setelah bukti Human tersedia; tidak ada Build/Patch session karena tidak ditemukan defect produksi.
- **Context pointers:** current-effective `TP-FE-001/techplan.md`; laporan ini; final frontend diff; Human acceptance record saat tersedia.
