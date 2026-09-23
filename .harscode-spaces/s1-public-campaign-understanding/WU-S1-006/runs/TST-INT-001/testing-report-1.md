# Laporan Pengujian Integrasi — TST-INT-001

> Fase: Testing independen  
> Work Unit / Run: `WU-S1-006` / `TST-INT-001`  
> Peran / spesialisasi: Verifier / Cross-stack integration — Slice 1 product outcome  
> Peserta: Codex CLI agent  
> Sesi: Fresh independent integration Testing session  
> Dibuat: 2026-09-23  
> Revisi diuji: `3e166fe696b19efe5ae2e01941942f687d770d2a`  
> Model / reasoning: `gpt-5.6-terra` / `high`

## 0. Ringkasan sweep

- **Terkonfirmasi:** laporan `TST-BE-002`, `TST-FE-001`, dan `TST-TOP-002` dipakai sebagai bukti prasyarat; run ini tidak mengulang matriks internal mereka. Run ini membuktikan korespondensi nyata frontend → Caddy root → backend → state PostgreSQL persisted/MinIO private.
- **Ditutup dari deferred lintas-WU:** browser production dengan MSW tidak aktif mencapai detail Campaign nyata melalui `localhost:8080`; detail dan image controlled keduanya meminta path `/api/...` yang sama dan menerima `200` serta `Cache-Control: private, no-store` dari Caddy.
- **Masih memerlukan Testing/Human:** penerimaan rendered/product oleh Human serta keputusan `SLICE_FINALIZED` tetap gate Human. Tidak ada Techplan khusus WU-S1-006/Test Focus Pointer yang disediakan; invocation, kontrak Slice 1, dan exact focus pointers WU-S1-003/004/005 dipakai sebagai oracle. Ini gap artefak orkestrasi non-blocking, bukan gap produk.

## 0a. Eksekusi Test Focus Pointer

| Area | Evidence anchor dibuka | Verifikasi terspesialisasi | Hasil |
|---|---|---|---|
| Public projection, anti-enumeration, dan funding | `WU-S1-003/EXP-BE-001/evidence/gap-analysis.md` Area 3; `TST-BE-002` | Fixture persisted public read melalui root Caddy dan browser; compare wire value/funding ke DOM; public unknown Campaign. | Lulus |
| Controlled media dan retraction | `WU-S1-003/EXP-BE-001/evidence/gap-analysis.md` Area 4; `WU-S1-005/EXP-TOP-001/evidence/solutioning.md` “Runtime verification required” item 2, 5–6 | Fetch fresh `content_url` sebelum/selepas withdrawal parent dan membership; cek no-store, type, non-redirect, dan tidak ada byte baru. | Lulus |
| Proxy path/cache/error preservation | `WU-S1-005/EXP-TOP-001/evidence/gap-analysis.md` Area 1 dan Stage 2 carry-forward; `TST-TOP-002` | `/api/healthz`, detail, media, `404`, dan induced eligible dependency `503` via root Caddy. | Lulus |
| Frontend production data boundary dan rendered hierarchy | `WU-S1-004/EXP-FE-001/evidence/gap-analysis.md` Area 4–5; `TST-FE-001` | Chromium headless memeriksa response resource, DOM desktop/mobile, state media, not-found, CTA, dan overflow. | Lulus |

## 1. Cakupan pengujian

| Rule / skenario | Category | Verifikasi observable | Hasil |
|---|---|---|---|
| Same-origin API path | Contract/topology | `GET localhost:8080/api/healthz` memberi `200`; detail Campaign root juga `200`, `Via: 1.1 Caddy`, dan `private, no-store`. | Lulus |
| Persisted public detail dan wire correspondence | Integration/truth | Fixture Campaign test-only berisi steward/purpose/story, target `5000000.00`, collected `6250000.00`, available PNG. Root detail memberi exact data dan `content_url` opaque `/api/campaigns/.../media/.../content`; browser menampilkan steward, provenance, `Rp 6.250.000,00`, `125.00%`, dan `Melebihi target`. | Lulus |
| Tidak ada alur donasi simulasi | Scope/negative | DOM browser menunjukkan “Dukungan belum dapat dilakukan dari halaman ini”; tidak ada control/link/form donation. | Lulus |
| Controlled media bytes | Integration/security | Browser resource log mencatat detail `200` lalu exact supplied `content_url` `200`; keduanya no-store. Fresh root media fetch memberi `image/png`, 70 byte, `Via: 1.1 Caddy`, tanpa `Location`. | Lulus |
| State media absent vs unavailable | Integration/rendered | Dua persisted Campaign nyata memberi `media.state: absent` dan `unavailable`; browser mobile menampilkan masing-masing “Belum ada media untuk ditampilkan” dan “Media sementara belum tersedia”, tanpa overflow. | Lulus |
| Safe public not-found | Security/error | Unknown UUID melalui root memberi `404`, `application/problem+json`, no-store, dan `Via`. Browser menampilkan “Campaign ini tidak dapat ditampilkan” tanpa alasan non-public/malformed/internal dan tanpa retry. | Lulus |
| Retryable dependency failure | Error/rendered | Setelah precondition sukses, PostgreSQL fixture-only dihentikan. Backend mengembalikan root `503` Problem Details, `private, no-store`, `Via`; browser menampilkan “Detail campaign sedang tidak tersedia”, satu tombol “Coba lagi”, tanpa detail database/object/internal. | Lulus |
| Retraction: parent public eligibility | Security/retraction | `content_url` yang sama memberi precondition `200 image/png`; setelah hanya status parent fixture dibuat `unpublished`, fresh root fetch memberi `404` Problem Details/no-store/`Via`, 165 byte, bukan PNG. | Lulus |
| Retraction: media membership | Security/retraction | Fixture dipulihkan, precondition URL sama kembali `200 image/png`; setelah hanya row membership media fixture dihapus, fresh root fetch memberi identik `404` Problem Details/no-store/`Via`, bukan PNG. | Lulus |
| Responsive rendered hierarchy | Rendered | Chromium 1440×900 dan 375×812: desktop menunjukkan identity/funding/action/provenance/media; mobile success, absent, unavailable, not-found, dan retryable error memiliki `scrollWidth - innerWidth = 0`. | Lulus |

## 2. Verifikasi error

| Error case | Perilaku/kategori diharapkan | Aktual | Actionable/propagated benar? |
|---|---|---|---|
| Unknown/non-public public resource | Safe `404`, no disclosure, no retry | Root `404` Problem Details/no-store; UI safe state tanpa retry/internal detail. | Ya |
| Declared absent media | Factual absent state, bukan unavailable | UI “Belum ada media untuk ditampilkan.” | Ya |
| Declared unavailable media | State distinct dari absent | UI “Media sementara belum tersedia.” | Ya |
| Backend dependency unavailable | Root `503`, no-store; UI retryable | Root `503` Campaign Problem Details/no-store/`Via`; UI memiliki named retry dan tidak membocorkan detail. | Ya |
| Known URL after parent/member withdrawal | Fresh request tidak mengirim byte baru; public `404`, no-store | Kedua jalur memberi `404` Problem Details, bukan PNG/redirect. | Ya |

## 3. Verifikasi final

- **Runtime method:** PostgreSQL `16-alpine` disposable pada `:15435`, semua migration sampai `000011`, backend native pada `:8090` dengan `DATABASE_URL` hanya ke database tersebut, frontend Next pada `:3000` tanpa `NEXT_PUBLIC_MSW_ENABLED`, dan Caddy root Compose pada `:8080`. MinIO private Compose dipakai untuk object fixture. Tidak ada database Compose bersama dimigrasikan atau dimutasi.
- **Pembersihan:** container PostgreSQL disposable dihentikan dengan `--rm`; proses backend/frontend sementara dihentikan; prefix object fixture `campaign/4f9d83f0-3c71-4f04-93d9-90bfbf868bf3` di `kencleng-private` kosong setelah cleanup. Tidak ada perubahan production source/worktree dari run ini.
- **Command final yang dijalankan:** `cd frontend && npm run verify` → **lulus**, 4 file/13 test. Hanya warning existing generated `public/mockServiceWorker.js`; `cd frontend && npm run build` → **lulus**. `git diff --check` → **lulus**.
- **Aggregate target command:** root `make verify` → **tidak clean** dan berhenti pada backend `gosec`. Hasilnya 11 temuan lintas-scope pre-existing (OAuth URL/logging, cookie Account, SHA-1 breach-check, key file reads), ditambah diagnostic cache Go (`staticcheck ./...` memperingatkan tidak menemukan package; gosec SSA cache misses). Tidak ada finding Campaign Slice-1 baru; command tidak melanjutkan ke suite backend/frontend setelah lint gagal. Ini konsisten secara substantif dengan baseline `TST-BE-002`, tetapi jumlah baseline berubah 12 → 11 karena kondisi tool/cache saat ini.
- **Broad checks yang sengaja tidak diulang:** full backend unit/contract/race dan disposable Postgres/MinIO adapter matrix tidak diulang karena telah lulus di `TST-BE-002`; topology policy/retraction baseline tidak diulang karena telah lulus di `TST-TOP-002`. Run ini menggantikannya dengan observasi cross-stack nyata yang sebelumnya deferred.
- **Migration/schema collision:** lulus untuk integration fixture: all migrations fresh apply sampai `000011`. No shared/manual DB migration dilakukan.
- **Backward compatibility:** route frontend tetap memakai API generated `/api/campaigns/{campaignId}` dan supplied opaque `content_url`; Caddy menjalankan path tersebut tanpa frontend alternate data branch. Tidak ada legacy public Campaign API yang perlu diadaptasi.
- **Fresh consistency read:** outcome/invocation, Product Slice 1, integration map, split OpenAPI Campaign route, dan prerequisite reports konsisten. Satu perbedaan dokumenter dicatat: `TST-FE-001` menyebut literal absent-media “Belum ada media untuk campaign ini”, sedangkan UI nyata saat ini berbunyi “Belum ada media untuk ditampilkan.” Semantik required tetap distinct; laporan lama perlu dikoreksi bila literal tersebut dipakai sebagai evidence exact-copy.

## 4. Pola bug berulang baru

Tidak ada pola defect produk baru. Baseline aggregate `make verify` tetap tidak bersih di area Account/OAuth yang bukan scope Slice 1; cache Go yang tidak konsisten juga membuat static analysis saat ini kurang dapat direproduksi.

## Verdict

**Pass with flagged follow-ups.** Bukti runtime ini mendukung rekomendasi `INTEGRATED_VERIFIED`: detail persisted, projection/funding truthful, no-fake-donation, controlled byte delivery, same-origin forwarding, public-safe error handling, media-state distinction, responsive rendering, dan fresh-fetch retraction seluruhnya teramati melalui root Caddy dengan MSW nonaktif.

Follow-up non-blocking:

1. Owner tooling/repository harus menstabilkan Go cache dan men-triage baseline `gosec` lintas-scope agar `make verify` dapat kembali hijau.
2. Jika wording exact menjadi evidence yang dikutip, perbarui `TST-FE-001` agar literal absent-media cocok dengan UI.

## Phase handoff

- **Completed:** verifikasi integrasi real Slice 1 melalui frontend production → root Caddy → backend persisted/private storage, termasuk negative, responsive, dan retraction states.
- **Artifacts:** laporan ini; tidak ada patch plan karena tidak ditemukan defect production.
- **Human decision:** Human harus melakukan integrated rendered/product acceptance, kemudian Orchestration/Human memutuskan `INTEGRATED_VERIFIED` dan `SLICE_FINALIZED`; run ini tidak menyetujui keduanya sendiri.
- **Open / deferred:** baseline `make verify` lintas-scope/cache dan dua follow-up di atas; tidak ada failure Slice-1 yang memerlukan Build/Patch.
- **Recommended next step:** lakukan Human acceptance pada representative desktop/mobile real route, lalu finalization/PR session segar bila gate Human terpenuhi.
- **Session transition:** fresh finalization/PR session direkomendasikan karena keputusan Human dan review evidence harus independen dari runtime setup ini.
- **Context pointers:** invocation ini; `TST-BE-002`, `TST-FE-001`, `TST-TOP-002`; `docs/project/kencleng-integration-map.md`; `docs/product/mvp-delivery-slices.md`.
