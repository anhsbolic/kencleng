# Tech Plan: Slice 2 Donation Domain & Contract Reconciliation

> Phase             : Techplan
> Work Unit         : WU-S2-002
> Run               : TP-S2-002-002
> Role              : Planner
> Specialization    : Techplan review resolution
> Participant       : Codex Planner
> Author            : Codex Planner
> Model             : gpt-6-luna
> Reasoning         : high
> Session transition: FRESH — Techplan re-entry setelah independent review; Session ID tidak tersedia.
> Created           : 2026-09-26
> Updated           : 2026-09-26
> Target revision   : 416e60415c51d0be7444638581ad206add24992e
> Workflow revision : cb5dca028b1d6ba49d43300dd06ec1bf2a9c984b
> Status            : Draft / In Review
> Approach          : Resolution terfokus pada invariant atomic coupling success/funding; mekanisme dan Open Items lain tetap terbuka.
> Refs              : `WU-S2-002/manifest.md`; `RV-S2-002-001/review-findings.md`; Product/MVP authority; current-effective Exploration `EXP-S2-001-001`; `docs/spec/README.md`; `api/README.md`

---

## 1. Background

Slice 2 harus memungkinkan pengunjung berdonasi sebagai guest dan memahami state sandbox yang nyata. Product/MVP authority menetapkan floor untuk eligibility, integritas uang, proteksi duplikasi, hasil yang tidak dapat dipalsukan melalui HTTP, safe guest revisit, dan representasi sandbox yang tidak mengaku sebagai settlement eksternal. Donation belum memiliki implementasi runtime aktif; checkout saat ini mendaftarkan Public Campaign Detail, tetapi belum mendaftarkan submit/status/settlement Donation. OpenAPI dan spec Donation yang ada berstatus historis/draft dan memuat operasi di luar Slice 2 serta pilihan yang belum direvalidasi.

Work Unit ini menghasilkan rekonsiliasi delivery spec dan shared contract yang cukup tegas untuk menentukan apakah `CONTRACT_READY` dapat diterima dan menurunkan delivery topology. Jika keputusan pemilik authority belum tersedia, artifact harus mempertahankannya sebagai blocker; Techplan maupun Build tidak mengarang nilai pengganti.

## 2. Scope

**In scope:**

- Menetapkan batas Slice 2 untuk guest submission, persisted truthful donation state, safe status revisit, dan funding yang merefleksikan settlement sukses tepat satu kali.
- Mengklasifikasikan detail historis Donation sebagai `KEEP`, `ADAPT`, `REPLACE`, atau `DEFER` dalam spec yang direkonsiliasi.
- Menyiapkan perubahan spec Donation yang diperlukan: domain invariants, threat model, task list, dan feature spec Slice 2 yang koheren; sinkronkan Campaign hanya untuk boundary eligibility/`max_amount` yang benar-benar diperlukan.
- Merekomendasikan dan, setelah keputusan authority yang diperlukan tersedia, merekonsiliasi split API source (`api/openapi/donation.yaml`, `api/openapi/index.yaml`, dan komponen `common.yaml` hanya bila diperlukan), serta menghasilkan kembali bundle dan frontend types sesuai `api/README.md`.
- Menetapkan bukti validasi kontrak dan Open Items yang menentukan penerimaan/penolakan `CONTRACT_READY`.

**Out of scope (explicit):**

- Perubahan Product Authority, MVP scope/sequencing, atau keputusan material Product/Design/Security tanpa owner yang berwenang.
- Backend/frontend runtime, database migration/application, job/scheduler/runbook implementation, test execution untuk runtime Donation, atau downstream FE/BE Work Unit topology.
- Account sebagai prasyarat, guest claim/history, public donor/social-proof list, payment-method breadth demi kelengkapan, dan real payment rails.
- Mengubah protected Tier-0 money ledger/transaction-locking implementation atau protected crypto/auth core.
- Mengklaim `CONTRACT_READY`, Slice-2 implementation, atau delivery milestone sebelum evidence dan authority yang diwajibkan tersedia.
- Perubahan orchestration state, Work Graph, Control Surface, Events, maupun project tracker oleh Participant Run ini.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Hasil Slice 2 adalah alur guest donation dan state pemrosesan/hasil sandbox yang nyata; Account bukan prasyarat. | `docs/product/mvp-scope.md` §§4–5, 9; `docs/product/mvp-delivery-slices.md` §5; `WU-S2-002/manifest.md` |
| Q2 | Amount semantics, satu jalur sandbox yang cukup, status pending/success/failure, safe guest revisit, dan copy non-real-settlement harus direkonsiliasi untuk slice ini. | `docs/product/mvp-delivery-slices.md` §5; Exploration Stage 2 `Area 1`; Stage 3 `Recommended material direction` |
| Q3 | Duplicate submission ditangani sesuai kebutuhan; transisi settlement tidak dapat dipalsukan lewat HTTP; funding sukses tepat satu kali dan aman terhadap concurrency; eligibility Campaign diputuskan server-side. | `docs/product/mvp-scope.md` §7; `docs/product/mvp-delivery-slices.md` §5; root `AGENTS.md` §2; Stage 2 `Area 3–4` |
| Q4 | Status guest tidak membuka data donor lain; credential/identifier sensitif tidak bocor melalui URL, cache, log, response, atau penyimpanan yang tidak aman. | `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` §7; Stage 2 `Area 2–3, 5`; Design `patterns.md` §7 |
| Q5 | Contract dan spec mengikuti Product/MVP serta UI/UX authority; detail historis bukan authority hanya karena sudah rinci. | `docs/product/README.md`; `docs/spec/README.md` §§1, 7; root `AGENTS.md` §1; Stage 2 `Scope and method` |
| Q6 | Split OpenAPI domain source dan referenced shared components adalah authored contract; aggregate `api/openapi.yaml` dan generated frontend types adalah turunan yang dihasilkan mengikuti API workflow. | `api/README.md` (`Structure`, `Editing workflow`) |
| Q7 | Slice 2 tidak mengaktifkan closure/result behavior kecuali bagian yang perlu untuk mempertahankan eligibility; `max_amount` interaction harus direkonsiliasi pada boundary Campaign/Donation. | `docs/product/mvp-delivery-slices.md` §§5–6; Stage 2 `Area 3`; Stage 3 `Direction considered` dan `Open / deferred questions` |

## 4. Rules & Validation

- **R1 — Batas slice:** Kontrak Slice 2 hanya memuat kemampuan yang diperlukan bagi guest submit, state sandbox yang persisten, safe revisit, dan funding truth; operasi Account, claim/history, public donor list, serta real rails tidak menjadi requirement aktif tanpa bukti enabling-critical dan perubahan scope melalui authority.
- **R2 — Input dan pemrosesan:** Amount floor/precision/currency, field guest, sandbox path, timing/outcome semantics, dan makna status ditetapkan oleh authority yang tepat. Angka atau perilaku historis tidak boleh disalin sebagai keputusan aktif tanpa revalidasi.
- **R3 — State dan uang:** Donation state berasal dari persisted backend truth; hanya jalur internal yang berwenang dapat menyelesaikan pending state; failed/pending tidak dihitung sebagai funding terkumpul. Untuk setiap settlement sukses, perubahan Donation ke `success` dan refleksi jumlahnya pada funding adalah satu hasil bisnis atomik: pada setiap state committed/observable keduanya hadir bersama, atau keduanya tidak committed. Setiap kontribusi sukses tercermin tepat satu kali, termasuk pada replay dan settlement konkuren tanpa lost update. Mekanisme detail mengikuti authority dan tidak ditulis ke kontrak sebagai asumsi implementasi.
- **R4 — Guest revisit/security:** Kontrak menjelaskan credential dan akses status yang aman, tidak mengembalikan credential sebagai data status, dan menetapkan kegagalan publik yang tidak membocorkan keberadaan Donation. Lifetime, transport/storage, URL/log/cache/referrer, perbandingan secret, abuse control, dan residual-risk acceptance harus diputuskan sebelum kontrak security dinyatakan siap.
- **R5 — Eligibility dan batas Campaign:** Server menegakkan eligibility saat submit, termasuk stale public detail. Interaksi successful settlement dengan `max_amount`, Campaign close, dan penolakan submission berikutnya dinyatakan konsisten oleh owner Campaign/Donation sebelum delivery dapat mengimplementasikannya.
- **R6 — Konsistensi spec/contract:** Invariants, threat model, feature acceptance, API operation/schema/error semantics, dan Campaign boundary tidak saling bertentangan. Hanya split source yang diedit; index path, generated bundle, serta frontend generated types diperbarui/diturunkan sesuai workflow bila contract berubah.
- **R7 — Truthful experience:** Status, konsekuensi, next action, dan penjelasan sandbox yang dibawa contract/spec tidak menyiratkan real external settlement, tidak mengubah pending menjadi sukses melalui presentasi, dan membedakan fakta platform dari hal yang belum diketahui. Terminologi visual/copy yang material menunggu Design Authority setelah makna state disepakati.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Reconcile satu capability guest donation + truthful state yang sempit; jangan bawa seluruh historical Donation domain ke Slice 2. | Chosen | Sesuai scope MVP dan manifest WU. History/claim, donor list, rails, dan Account bukan baseline slice. Alternatif full legacy domain ditolak karena melampaui scope; Account-first ditolak karena guest adalah aktor utama. |
| D2 | State harus persisted dan settlement authority server-side/internal; frontend-only result dan settlement HTTP callback tidak diterima. | Chosen | Hasil harus benar di dalam sandbox dan tidak meniru bukti settlement nyata. Alternatif client simulation serta endpoint callback publik ditolak karena misleading/forgery risk. |
| D3 | Pertahankan prinsip money safety, eligibility, idempotency bila diperlukan, PII protection, dan log safety; detail mekanisme lama tetap perlu direvalidasi. | Chosen | Product/security floor dan root convention berlaku; historical detail tidak membuktikan mekanisme runtime yang aktif. Nilai minimum, API fields, locking mechanism, dan token design tidak dikunci oleh pilihan ini. |
| D4 | Operasi Account history/claim dan public donor list diklasifikasikan `DEFER` untuk Slice 2; payment-method breadth dan real rails juga `DEFER`. | Chosen | Ditetapkan eksplisit sebagai out-of-scope pada Product/MVP. Jika ditemukan enabling-critical evidence, rute ke owning authority sebelum scope berubah. |
| D5 | Jangan mengadopsi otomatis token status historis yang non-expiring/query-param, lifetime/risk acceptance-nya, atau error shape historical `401`/OpenAPI `404`. | Chosen | Credential adalah akses status; Exploration mencatat paparan credential dan konflik contract. Alternatif meneruskan legacy shape tanpa keputusan ditolak sebagai asumsi keamanan/kontrak. Owner harus menyelesaikan Active Open Items sebelum kontrak dikunci. |
| D6 | Rekonsiliasi contract/spec dahulu, lalu Orchestrator menentukan backend-first atau contract-parallel dari kontrak stabil dan dependency nyata. | Chosen | Belum ada runtime Donation atau shared contract yang direkonsiliasi. Memulai FE/BE independen terhadap kontrak lama ditolak; parallelism hanya optimasi setelah contract ready. |
| D7 | Tutup `max_amount` hanya sebatas eligibility yang dibutuhkan Slice 2; tidak mengimpor seluruh Campaign closure/result scope. | Chosen as boundary; behavior unresolved | Product mengecualikan closure/result kecuali menjaga eligibility. Historical hook cukup untuk menandai boundary, tidak untuk memilih crossing/over-target policy. Owner Campaign/Donation harus memutuskan rincian (Open Item O6). |
| D8 | Bundle OpenAPI dan generated client types adalah derived artifacts; authored source tetap domain split file dan index wiring sesuai `api/README.md`. | Chosen | Menghindari drift dan hand-edit generated output. |

## 6. Backward Compatibility

- **Existing data:** Tidak ada live Donation module, migration, atau stored Donation runtime contract yang ditemukan pada target checkout; skema storage aktif Slice 2 belum ditentukan. Jangan merancang data migration dari draft schema tanpa data/source evidence.
- **API/contracts/clients:** `api/openapi/donation.yaml` saat ini mendeklarasikan submit, status, donor-list, dan Account operations, tetapi Exploration tidak menemukan server routes atau runtime client flow. Frontend generated OpenAPI types tetap memuat operasi Donation historis; repository search tidak menemukan pemakaian aktif di luar generated artifact. Penghapusan/perubahan shape tetap memengaruhi generated API surface; Build wajib memastikan tidak ada consumer aktif dan merutekan consumer eksternal/kontrak yang ditemukan sebagai Open Item sebelum breaking change. Public Campaign Detail Slice 1 tetap memiliki `donation_action` unavailable sampai kontrak/implementasi Slice 2 mengoordinasikan aktivasi.
- **Migration/deprecation compatibility:** Tidak ada perubahan runtime/backfill/migration pada WU ini. Karena Donation contract historis berstatus draft dan runtime tidak tersedia di checkout, jangan menganggap ada compatibility guarantee; verifikasi status distribusi/consumer sebelum menghapus atau mengganti operasi lama.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Nilai lama seperti minimum Rp 5.000, enam payment methods, 2–5 detik, 5% failure, atau non-expiring token secara tidak sengaja dianggap approved. | High | High | R2/R4; O1–O4; semua nilai historical diberi status evidence saja sampai owner memutuskan. |
| RISK-2 | Status credential dapat bocor atau lookup membedakan Donation tidak ada vs credential salah; threat model lama menerima residual risk tanpa otorisasi MVP saat ini. | Medium | High | R4; keputusan security dan response parity wajib eksplisit; jangan menyalin penerimaan risiko lama. |
| RISK-3 | Retry/settlement concurrency menyebabkan duplikasi/lost funding update, atau kegagalan di antara perubahan state dan funding meninggalkan hasil bisnis parsial; akar mekanisme menyentuh Tier-0 protected ledger/locking. | Medium | Critical | R3 dan §8 menetapkan hasil bisnis atomik; Test Focus Pointer F1 meminta bukti partial-failure/replay/race. Batas write Tier-0 tetap berlaku dan tidak diotorisasi Techplan ini. |
| RISK-4 | Campaign berubah tidak eligible saat submit/settlement, atau `max_amount` terlampaui tanpa perilaku yang konsisten. | Medium | High | R5; O6; owner Campaign/Donation menentukan ordering dan dampak tanpa mengimpor closure/result scope Slice 3. |
| RISK-5 | Menghapus operasi historis dari OpenAPI merusak konsumen yang tidak terdeteksi atau membiarkan generated bundle/types stale. | Low / unknown | High | O8; consumer audit dan validasi workflow API sebelum kontrak final; perubahan source, index, bundle, generated types harus konsisten. |
| RISK-6 | Wording/status design menyiratkan settlement eksternal atau fakta yang belum diketahui. | Medium | High | R7; semantik state diselesaikan lebih dulu; minta Design Authority hanya untuk keputusan istilah/label yang material. |
| RISK-7 | Beberapa owner belum memutuskan dimensi kontrak; WU tidak dapat membuktikan `CONTRACT_READY`. | Medium | High | Catat owner, bukti, dan blocker secara eksplisit; WU boleh merekomendasikan `CONTRACT_READY` ditolak/ditunda jika kontrak tak bisa dieksekusi tanpa asumsi. |

## 8. Interface Contract

**Persistence/data shape:** Belum ada Donation runtime schema/migration aktif di target checkout. Kontrak/spec harus menetapkan state Donation minimum, data yang memang dibutuhkan untuk submission/status, pemisahan field private dari public response, serta retention/PII policy setelah O1–O4/O7 dijawab. Untuk settlement sukses, state Donation dan funding yang merefleksikan kontribusinya harus menjadi satu hasil bisnis atomik dan concurrency-safe: tidak boleh ada state committed/observable yang menunjukkan salah satunya tanpa yang lain; kegagalan sebelum hasil lengkap harus meninggalkan keduanya tidak committed. Setiap kontribusi sukses dihitung sekali saja meski settlement diulang atau bersamaan. Ini menetapkan invariant observable, bukan bentuk transaksi, lock, algoritma ledger, atau struktur kode. Jangan mewarisi kolom/token/Account/Event/public-display historis hanya karena sudah ada pada draft.

**API/event/external interface:** Sumber authored adalah `api/openapi/donation.yaml` dengan path registry `api/openapi/index.yaml` dan hanya shared components yang benar-benar direferensikan dari `api/openapi/common.yaml`. Contract perlu menyatakan guest submission, response/status state, kegagalan validation/business, dan safe revisit. Endpoint/path, header, body fields, status credential carrier/lifetime, response code, serta caching tidak dikunci di Techplan karena keputusan O1–O5 menentukan shape. Settlement/result transition wajib tidak dipaparkan sebagai HTTP action. `api/openapi.yaml` serta `frontend/lib/api/generated/openapi.ts` adalah output yang diturunkan dari source, bukan tempat authoring manual.

**Cross-layer/business boundary:** Campaign Detail saat ini mengembalikan `donation_action` unavailable; FE hanya mempresentasikan kontrak, tidak menetapkan amount validity, eligibility, atau state finansial. Backend menjadi otoritas eligibility dan state; settlement outcome tetap jelas sebagai sandbox. Campaign max/closure boundary perlu konsisten dengan Donation eligibility tanpa mengambil cakupan Slice 3 yang tidak diperlukan.

## 9. Architecture / Plan

1. **Re-open authorities dan klasifikasikan detail lama.** Build mengonfirmasi scope/owner dari Product/MVP, Design, spec, threat model, API, dan code anchors terkini. Untuk setiap detail legacy yang relevan tandai `KEEP`, `ADAPT`, `REPLACE`, atau `DEFER`; hilangkan dari active-slice contract fitur yang jelas out-of-scope (Account/claim/history/public donor list/real rails) dengan mempertimbangkan consumer audit O8.
2. **Selesaikan keputusan prasyarat sebelum menulis kontrak final.** Route O1–O8 kepada Product, Donation/Campaign domain, Security, API, atau Design owner yang disebut di §13. Jika owner belum menjawab, spec mempertahankan `OPEN`/blocker yang eksplisit dan kontrak hanya mengunci fakta yang memang sudah berwenang. Jangan isi field atau nilai historis untuk membuat dokumen terlihat lengkap.
3. **Susun spec Slice-2 yang koheren.** Reconcile Donation invariants, threat model, task list, dan feature acceptance dari perilaku yang diputuskan. Pastikan settlement non-HTTP, amount decimal-safe, PII/token boundaries, duplicate protection, exact-once funding, concurrency, status parity, eligibility, dan threat-driven evidence tercermin tanpa menyalin fitur domain yang ditunda. Sinkronkan Campaign closure invariant/feature hanya jika owner menyetujui perubahan pada `max_amount` boundary.
4. **Turunkan shared API contract.** Ubah domain split Donation source dan index; ubah `common.yaml` hanya untuk komponen bersama yang benar-benar dibutuhkan. Ubah Campaign public donation action hanya melalui koordinasi contract Campaign yang eksplisit. Jangan ubah API implementation/code dalam Work Unit reconciliation.
5. **Hasilkan dan validasi turunannya.** Ikuti `api/README.md`: jalankan `npm run validate`; jika source berubah, bundle ulang dari authored split source dan generate frontend API types melalui command yang didefinisikan repo. Audit diff untuk memastikan tidak ada hand edit bundle/types atau perubahan warning di luar aturan README. Pengecekan ini hanya memvalidasi artifact contract, bukan runtime capability.
6. **Tentukan handoff evidence.** Catat tiap keputusan resolved/open, classification legacy, exact contract sources, verification result, consumer impact, dan residual risks. Beri rekomendasi kepada Orchestrator apakah syarat `CONTRACT_READY` terpenuhi; jangan ubah milestone/tracker sebagai bagian Participant Run.

Tidak ada diagram karena alur rekonsiliasi linear dan belum ada branching/state transition runtime yang settled untuk digambarkan.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `docs/product/mvp-delivery-slices.md` — `## 5. Slice 2 — Guest Donation + Truthful Donation State` | Batas aktif, correctness/security floor, exclusion, completion evidence. | Build harus memakai ini sebagai scope owner; jangan promote detail legacy. |
| `docs/spec/README.md` — `Authority relationship`, `Four document types`, `Rules for filling these out` | Status spec dan urutan rekonsiliasi; domain-first locations. | Reconcile only spec artifacts needed for Slice 2; label `draft`/open until owner review. |
| `docs/spec/5-donation/invariants.md` — `INV-donation-01`…`INV-donation-15`, state machines | Historical invariant candidates, several cover deferred operations. | Classify individually; keep only Slice-2 relevant and revalidate; do not bulk retain or delete evidence. |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` — `Submission`, `Settlement`, `Concurrency & correctness notes` | Historical amount/payment/email/timing/closure decisions and internal transition. | Retain internal-only/guarded transition as requirement direction only where reaffirmed by Product; defer specifics until O1–O4/O6. |
| `docs/spec/5-donation/features/02-donation-status-check.md` — `Behavior`, `Validation & error cases` | Historical credential and uniform-401 semantics. | Resolve against OpenAPI 401/404 and security guidance; don't import the spec alone. |
| `docs/spec/5-donation/threat-model.md` — `Submit donation`, `Token-based status check`, `Internal settlement process` | Sensitive guest submission/status and forged-settlement risks. | Refresh threat model for active surface; re-evaluate historical accepted residual risks with current owner. |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` — `Behavior` | Reconciled Slice-1 `donation_action` is required/unavailable, no activation target. | Preserve current Slice-1 meaning pending a coordinated, explicit contract change. |
| `docs/spec/4-campaign/features/09-closure.md` — `max_amount trigger`; `docs/spec/4-campaign/invariants.md` — `INV-campaign-13` | Historical cross-domain close hook conflicts with narrow Slice-2 closure boundary unless reconciled. | Coordinate O6; do not edit Campaign behavior solely to match old Donation rule. |
| `api/README.md` — `Structure`, `Editing workflow` | Authored source, path registry, bundle, generated client workflow. | Follow canonical split-source and generation sequence; no generated artifact hand editing. |
| `api/openapi/donation.yaml` — `/campaigns/{campaignId}/donations`, `/donations/{donationId}/status`, `/account/donations*`; `SubmitDonationRequest`, `Donation`, `PaymentMethod` | Historical API surface and direct 401/404 conflict/out-of-scope operations. | Reconcile narrow active contract after decisions; check consumers before removals. |
| `api/openapi/index.yaml` — Donation path refs; `api/openapi/common.yaml` — `IdempotencyKeyHeader`, `Unauthorized`, `NotFound` | Registry and shared error/key definitions currently referenced by legacy Donation contract. | Keep path inventory and shared references synchronized only as needed by final source. |
| `backend/cmd/server/main.go` — `run` route registrations around `GET /campaigns/{campaignId}` | Live server route surface; confirms Donation runtime is absent at target revision. | Evidence only; no backend wiring is part of this reconciliation Work Unit. |
| `backend/internal/domain/campaign/service.go` — `toPublicDetail` | Current Campaign service sets donation action to unavailable and maps decimal funding. | Use as current behavior evidence; do not mutate service in this contract WU. |
| `backend/internal/transport/http/campaign_public.go` — `publicCampaignDetailResponse`, `toPublicCampaignDetailResponse` | Current closed public wire projection including unavailable Donation action. | Use as contract coordination anchor; any future activation needs coordinated spec/API/runtime work. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` — `CampaignSuccess`, `Funding`; `frontend/lib/api/public-campaign.ts` — `getPublicCampaignDetail` | Current FE only consumes Campaign detail and shows non-activating Donation context. | Evidence of no current donation flow; no frontend behavior changes in this Work Unit. |
| `frontend/lib/api/generated/openapi.ts` — Donation path declarations | Generated historical types remain even though no active usage was found outside generated file. | Regenerate after source contract update; audit consumers before removing declarations. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `docs/spec/5-donation/invariants.md` | Adapt | Pertahankan invariant Slice 2; klasifikasikan detail historis yang tidak berlaku dan catat hal yang belum diputuskan. |
| `docs/spec/5-donation/threat-model.md` | Adapt | Perbarui threat model untuk batas guest submit/status/sandbox yang aktif dan residual risk yang benar-benar diterima saat ini. |
| `docs/spec/5-donation/tasks.md` | Adapt | Selaraskan task view dengan delivery spec Slice 2 yang koheren setelah scope direkonsiliasi. |
| `docs/spec/5-donation/features/*` | Adapt / tambah bila perlu | Rekonsiliasi perilaku guest submit/state/status; jangan menulis operasi Account/list yang tidak terkait. File tepatnya mengikuti `docs/spec/README.md` dan pembagian feature final. |
| `docs/spec/4-campaign/invariants.md`, `docs/spec/4-campaign/features/09-closure.md` | Adapt bersyarat | Ubah hanya jika owner Campaign/Donation menyepakati perubahan boundary yang diperlukan; selain itu cukup rujuk aturan yang berlaku dan biarkan file tetap. |
| `api/openapi/donation.yaml` | Adapt | Menjadi authored Slice-2 contract setelah Open Items yang menentukan bentuknya diputuskan. |
| `api/openapi/index.yaml` | Perbarui sumber turunan | Selaraskan path ref Donation dengan split domain source jika operasi berubah. |
| `api/openapi/common.yaml` | Adapt bersyarat | Ubah hanya jika kontrak final memerlukan shared component; jangan sentuh kontrak umum yang tidak terkait. |
| `api/openapi.yaml` | Generated | Bundle ulang dari source bila berubah. |
| `frontend/lib/api/generated/openapi.ts` | Generated | Generate ulang melalui workflow API canonical bila berubah; audit permukaan generated dan dampak ke consumer. |

**File / area intentionally untouched**

| File / area intentionally untouched | Why |
|---|---|
| `docs/product/**`, `docs/ui-ux/**` | Product/Design owner memegang authority; rute keputusan kepada mereka, jangan ubah berkasnya dalam contract WU ini. |
| `.harscode-spaces/**` di luar `techplan.md` Run ini | Orchestrator memiliki state/history koordinasi; Run ini hanya menulis artifact Techplan miliknya. |
| `docs/project/kencleng-development-tracker.md` | Pembaruan tracker dimiliki Orchestrator; Participant Run tidak mengubah state. |
| Berkas runtime `backend/**`, `frontend/**` | Work Unit ini merekonsiliasi spec/contract, bukan implementasi. Generated types hanya boleh berubah sebagai output turunan jika API source berubah. |
| `backend/internal/domain/donation/ledger.go` dan implementasi transaction/locking; `backend/internal/platform/crypto/`; area auth/state-machine yang diproteksi | Batas write Tier-0; Work Unit ini tidak memberi otorisasi dan tidak memerlukan perubahan implementasi tersebut. |
| Edit manual `api/openapi.yaml` | Ini aggregate generated; workflow source yang mengatur regenerasinya. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Human meninjau scope akhir spec/API Donation terhadap Slice 2; pastikan Account history/claim, donor-list, dan real rails ditunda kecuali ada dasar enabling-critical yang disetujui authority. | Human | Mencegah scope historis memperluas kontrak MVP; batas Product/MVP memerlukan penilaian authority manusia. |
| R2 | Untuk setiap field amount/payment/guest/timing di spec dan API, telusuri keputusan authority yang sudah diselesaikan; nilai yang belum diputuskan harus tetap terbuka dan menghalangi kesiapan kontrak. | Human | Nilai ini mengubah eligibility dan janji kepada pengguna; lint tidak dapat menetapkan authority semantik. |
| R3 | Cocokkan invariant/state machine dan API dengan hasil bisnis atomik: setiap perubahan ke `success` committed/observable bersama refleksi kontribusi tepat sekali pada funding; kegagalan sebelum hasil lengkap tidak meninggalkan hanya satu perubahan; replay/settlement konkuren tidak menggandakan kontribusi atau menghilangkan increment. Pada delivery runtime, verifikasi perilaku ini lewat bukti kegagalan saat coupling belum lengkap serta bukti replay dan concurrency/race; jangan mengunci mekanisme pengujian ke primitive tertentu di Techplan ini. | Testing | Pemeriksaan kontrak membuktikan keselarasan acceptance/API; bukti runtime berada di luar WU ini dan dimiliki Testing lanjutan. Tanpanya, state sukses dapat berbeda dari funding yang ditampilkan, atau retry/konkurensi dapat menggandakan maupun menghilangkan nilai uang. |
| R4 | Security owner meninjau transport/lifetime/storage/log/cache/referrer credential, response parity, perbandingan secret, dan keputusan rate-limit/risk acceptance; cocokkan contract dengan threat model. | Human | Pilihan ini mengatur akses tanpa login ke status Donation privat; schema lint atau unit test biasa tidak dapat menerima residual risk. |
| R5 | Owner Campaign/Donation meninjau eligibility saat submit dan konsistensi `max_amount`/closure; cocokkan invariant/feature spec Campaign dan Donation yang direkonsiliasi. | Human | Asumsi eligibility/closure yang keliru mengubah perilaku fundraising dan memerlukan authority lintas-domain. |
| R6 | Jalankan `cd api && npm run validate`; jika API source berubah, bundle dan generate types dengan command di `api/README.md`, lalu periksa kesesuaian source/index/bundle/types serta perubahan warning. Testing memiliki pemeriksaan final; Build boleh menjalankannya sebagai bukti edit-loop. | Testing | Memvalidasi ref dan konsistensi artifact turunan. Tanpanya kontrak mungkin gagal dibundle atau consumer memakai output basi; pemeriksaan ini tidak membuktikan runtime. |
| R7 | Human Design/Product meninjau makna state, disclosure sandbox, dan istilah material setelah semantik diputuskan; pastikan copy selaras dengan truth sandbox yang persisted dan tidak menyiratkan settlement eksternal. | Human | Bahasa sukses/pending yang menyesatkan merusak tujuan trust slice; pemeriksaan schema otomatis tidak menilai kebenaran pengalaman. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Settlement authority, transisi duplikat, atomic coupling success/funding, dan donasi konkuren | Integritas uang dan settlement internal-only adalah boundary security/concurrency. Testing runtime perlu membuktikan invariant observable lintas state/funding pada kegagalan parsial, replay, serta race; tidak menentukan primitive implementasi. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence`; file yang sama `#Area 4 — Backend live state and security/correctness boundary`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md#Recommended material direction` | Ya — Techplan memperjelas constraint coupling yang sudah dipilih Exploration; bukti runtime dimiliki Testing implementasi lanjutan. |
| Credential status guest, response parity, paparan URL/log/cache/referrer, dan risiko abuse | Credential tanpa login memberi akses ke status privat; kebijakan token historis dan mismatch 401/404 belum diputuskan. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence`; file yang sama `#Area 5 — Frontend live state and cross-stack surface` | Ya — desain credential dan risk acceptance masih Open Item; tetapkan tes khusus setelah keputusan kontrak. |
| Guest email/PII, enkripsi dan logging aman | Data guest dapat bocor bila disimpan, dikembalikan, atau dicatat tanpa mengikuti pola PII/token aktif. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence`; invariant legacy yang dirujuk tetap evidence, bukan keputusan final. | Ya — bila guest email dipertahankan, downstream spec dan Testing harus membuktikan penyimpanan/penggunaan aman; bila tidak dibutuhkan, tandai N/A dengan alasan pada spec. |
| Eligibility saat submit dan interaksi `max_amount` sukses | Detail Campaign yang stale atau close konkuren dapat menerima Donation yang tidak eligible atau mengubah perilaku Slice 3. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md#Open / deferred questions` | Ya — eligibility masuk scope; aturan max/closure tepatnya menunggu O6. |

## 13. Open Items

### Active — needs external input or verification

1. **Amount and payment contract — Product + Donation delivery owner.** Decide currency/precision/rounding constraints and amount floor, and whether one named sandbox representation suffices. Historical Rp 5,000 and six-method enum are not approval. Until answered, exact request schema and amount validation acceptance cannot be finalized.
2. **Sandbox result semantics — Product Authority with Donation delivery owner.** Decide what the user is promised about pending duration and how pending/success/failure arise; historical 2–5s/5% is not approved. Until answered, processing/status behavior and truthful UI copy remain incomplete.
3. **Guest data and retention — Product Authority; Security/PII owner for handling.** Decide which guest fields are necessary/optional, retention/deletion and notification consequences. Any email uses established encryption/HMAC and safe-logging conventions; do not infer that email is mandatory or optional from legacy spec alone.
4. **Safe revisit credential and risk acceptance — Security/project authority + Donation/API owner; Design owner for material flow interaction.** Decide credential type/carrier, lifetime/revocation, generation/storage/comparison, URL/history/log/referrer/cache exposure, anti-enumeration, and abuse-control/residual-risk posture. Historical non-expiring query token and its old accepted risks are not authorization. Blocks final public access contract.
5. **Status failure contract — Donation/API owner with Security review.** Resolve historical feature requirement for identical `401` on absent/wrong/missing credential versus OpenAPI `404` for missing Donation; include response/body/timing/caching semantics appropriate to anti-enumeration. Blocks final status operation.
6. **`max_amount` crossing and Campaign eligibility — Campaign + Donation owners; Product Authority if semantics alter lifecycle meaning.** Decide whether a successful donation can cross the cap, how the crossing settlement and close interact, and when future submissions become ineligible. Do not import full historical closure behavior; blocks final cross-domain invariants/acceptance.
7. **Status terminology/source labels — Design Authority after O2 semantic meaning is settled.** Confirm whether canonical principles are sufficient for exact labels or whether a material design decision is needed. Do not use copy/color to imply settlement certainty or source verification.
8. **Historical API consumer/distribution status — API/Orchestrator owner.** Before removing/replacing old Donation operations, verify no active in-repo or external consumer relies on this draft shape. In-repo scan found generated declarations but no active runtime usage; external distribution is not established by the current evidence. Preserve or stage a compatibility path if a consumer is found, and route any contract break through the owning authority.

### Resolved — retained as decision history

1. ~~**Baseline actor and product outcome**~~ **RESOLVED — guest donation without Account prerequisite, with truthful real-in-sandbox status.** Product/MVP approved scope; retained in Q1 and D1.
2. ~~**Whole historical Donation scope**~~ **RESOLVED — defer Account claim/history, public donor list, real rails, and unnecessary payment breadth for this slice.** Explicit Product/MVP exclusions; see D4.
3. ~~**Frontend-only or HTTP-triggered settlement**~~ **RESOLVED — settlement outcome must be backend-owned/persisted and no client-callable transition can forge it.** Product correctness floor and Exploration direction; see D2/R3.
