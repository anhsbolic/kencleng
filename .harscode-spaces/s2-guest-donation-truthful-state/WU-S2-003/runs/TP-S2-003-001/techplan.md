# Tech Plan: Guest Donation Backend — Truthful Persisted State

> Phase             : Techplan
> Ticket            : WU-S2-003 / TP-S2-003-001
> Author            : P-S2-003-TP-001-1 (Planner)
> Participant ID    : P-S2-003-TP-001-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : Invocation configures `gpt-6-luna`; active runtime model not independently exposed
> Reasoning          : Invocation configures `high`; active runtime effort not independently exposed
> Session           : not exposed
> Created           : 2026-10-01
> Updated           : 2026-10-01
> Target revision   : `7fd8b473b239b20bda3990ab29c51440d321a796` (Invocation baseline; live authorities/code re-read)
> Workflow revision : Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`
> Status            : Draft / In Review
> Approach          : Bentuk capability Donation backend yang memenuhi kontrak Slice 2; tahan D1 implementation sampai seam Campaign dan implikasi Tier-0 diverifikasi.
> Refs              : Approved TP-S2-002-015; Product Slice 2; agreed Donation specs; `api/openapi/donation.yaml` + `common.yaml`; `docs/project/kencleng-monetary-data-standard.md`; EXP-S2-003-001 Stage 2–3.

---

## 1. Background

Slice 2 menjanjikan donasi guest dengan status sandbox yang benar-benar tersimpan, akses status sementara yang aman, dan pendanaan Campaign yang mencerminkan keberhasilan tepat satu kali. Kontrak Donation dan spesifikasi Slice 2 telah diterima, tetapi backend saat ini belum memiliki Donation route, domain, persistence, simulator, status handler, atau test Donation. Kontrak tersebut adalah authority perilaku, bukan bukti runtime.

Risiko utamanya berada pada batas Campaign–Donation untuk D1, akses status bearer tanpa autentikasi, dan email guest yang membawa PII. Techplan ini menetapkan perilaku dan bukti yang wajib dijaga, serta mencatat keputusan owner yang belum tersedia tanpa memilih mekanisme transaksi, kebijakan credential, atau penerimaan residual risk.

## 2. Scope

**In scope:**

- Backend guest submission sesuai kontrak; validasi IDR whole-Rupiah, request idempotency, dan persistensi awal `pending`.
- Backend-controlled simulator yang menulis hasil terminal sandbox `success` atau `failed`; tidak ada hasil terminal yang dapat dipilih caller.
- Guest status revisit yang status-only, memakai credential sementara sesuai kontrak, serta error publik seragam.
- Donation email opsional sesuai acceptance: verifikasi kepemilikan sebelum pengiriman, paling banyak satu notice status-only di state terminal, dan lifecycle bounded/recoverable sesuai O11. Implementasi pemenuhan notice tetap bergantung pada keputusan O3 dan kemampuan sender yang nyata.
- Minimum backend capability untuk D1 di dalam WU-S2-003. Campaign tetap memiliki kepemilikan lifecycle, eligibility/close coordination, dan winning close reason. Submission-vs-close, settlement, funding, dan threshold mengikuti invariant yang diterima; mekanisme tidak dipilih di sini.
- Reconcile producer `donation_action` pada Campaign GET hanya setelah kontrak Campaign/API WU-S2-005 diterima. Dependency ini tidak menahan perencanaan atau implementasi Donation yang tidak bergantung padanya; GET snapshot tidak memberi otorisasi submit.
- Migrations, unit/HTTP tests, serta PostgreSQL integration/concurrency evidence yang diperlukan oleh perilaku di atas. Aplikasi migration/index tetap mengikuti gate Human/manual yang berlaku.

**Out of scope (explicit):**

- `frontend/`, perubahan Product/MVP/spec/API authority, atau perubahan Campaign contract sebelum WU-S2-005 selesai dan diterima.
- Account prerequisite, claim/history, public donor list, guest-email reveal, Campaign-wide email, dan real payment/settlement rails.
- Full Campaign closure/result lifecycle, deadline scheduler, force-close, atau Slice 3; hanya D1 boundary yang diperlukan.
- Memilih atau mengubah mekanisme transaction/locking/isolation untuk balance updates pada jalur yang dilindungi root `AGENTS.md`; tidak ada protected write yang diotorisasi.
- Menerapkan PII/credential control yang memerlukan keputusan Security/PII/API yang belum ada, menerima residual risk, atau mengklaim `BACKEND_VERIFIED`.
- Menetapkan supported currency tambahan, monetary range, universal scale/precision, derived rounding/tax, atau mengadopsi Campaign `NUMERIC(19,2)` sebagai standar Donation.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Implement hanya guest Donation submit/status, simulator, notification yang diwajibkan, dan batas Campaign yang perlu untuk Slice 2. Kehadiran operasi historis lain di OpenAPI tidak memperluas scope. | `docs/product/mvp-delivery-slices.md` §5; Donation invariants INV-donation-01–11; TP-S2-002-015 §§2–3; EXP Stage 2 Area 1. |
| Q2 | Input Donation adalah IDR whole Rupiah, minimum Rp5.000 dengan Rp1 increment (Rp5.001 valid). API/wire money berupa major-unit decimal string plus currency; kalkulasi/persistensi exact decimal, tanpa `float`/`float64`. Parameter range/scale yang belum disetujui tetap terbuka. | INV-donation-01; `features/01-submit-donation-settlement.md`; monetary data standard; TP-015 Q2/R2/R12. |
| Q3 | Untuk satu intent, key dan payload sama mengembalikan Donation yang sama setelah retry ambigu; key yang sama dengan payload berbeda ditolak. Key baru hanya untuk intent baru yang disengaja. Ini terpisah dari replay settlement. | INV-donation-10; Donation OpenAPI `Idempotency-Key`; TP-015 R8; approved O9 direction. |
| Q4 | Submission yang diterima tersimpan sebagai `pending`; hanya simulator backend dapat menentukan hasil terminal. QRIS satu-satunya metode aktif dan bermakna simulasi; caller tidak dapat memilih `success`/`failed`. Pending tidak memicu resubmission dan tidak menjanjikan ETA. | INV-donation-07/11; `features/01-submit-donation-settlement.md`; Slice 2 Product decisions; TP-015 R4. |
| Q5 | Optional guest name tidak public. Email hanya untuk notice status Donation, opt-in, terpisah dari Account verification. Verifikasi harus mendahului notice/access message; maksimal satu notice status-only yang berlabel simulasi saat terminal, tidak saat initial pending. Verified address tetap eligible sampai notice dipenuhi melalui bounded/recoverable terminalization. | INV-donation-03/04/06; O3/O11 pada TP-015; agreed threat model; `features/01-submit-donation-settlement.md`. |
| Q6 | Status credential adalah bearer credential sulit ditebak, dibawa di URL fragment untuk frontend handoff/URL cleanup, diverifikasi dengan one-way HMAC, dan kedaluwarsa keras 24 jam sejak issuance tanpa perpanjangan. Result memuat status saja. Credential strength, key/comparison, exposure dan abuse controls masih owner/evidence gates. | INV-donation-05/06; `features/02-donation-status-check.md`; Donation OpenAPI status operation; TP-015 Q7/R6. |
| Q7 | Absent Donation, credential hilang/salah/kedaluwarsa memberikan public `404` dengan Problem Details body, headers, dan cache behavior yang sama, termasuk `Cache-Control: private, no-store`. Kesamaan timing/abuse behavior harus dibuktikan runtime, bukan diasumsikan dari kontrak. | INV-donation-05; `features/02-donation-status-check.md`; Donation OpenAPI; O5. |
| Q8 | D1: eligibility submission diurutkan secara atomik terhadap Campaign close. Close menang → submission baru ditolak; submission menang saat eligible → Donation diterima dan tetap dapat settle penuh setelah close. Settlement sukses dan full funding increment commit atomik tepat sekali; close reason pemenang tidak berubah; total funding boleh melewati threshold. | `docs/spec/4-campaign/invariants.md#inv-campaign-13`; Donation INV-02/08/09; Product Slice 2 §§5–6; O6/D1 dan TP-015 Q8/R7. |
| Q9 | Campaign memiliki lifecycle dan close reason. WU-S2-003 mengerjakan minimum D1 integration capability sesuai route Orchestrator; tidak mengambil full Campaign lifecycle. Detail seam harus diverifikasi terhadap live code dan root Tier-0 fence sebelum D1 Build. | WU-S2-003 manifest/current state; Work Graph; EXP Stage 2 Area 3/F-02; Stage 3; `backend/internal/domain/campaign/` live code; root `AGENTS.md` §3. |
| Q10 | Campaign GET `donation_action` producer bergantung pada hasil contract reconciliation WU-S2-005 yang accepted. Keputusan arah availability-only belum merupakan contract acceptance; snapshot GET tidak menjamin eligibility/POST authorization. | WU-S2-005 Stage 2–3 evidence; WU-S2-005 manifest; Work Graph scoped HARD dependency; current `campaign.yaml` and campaign projection. |
| Q11 | Seluruh behavior backend Tier 1 menjalani independent review dan Testing yang relevan. Klaim uang/close/concurrency memerlukan invariant assertions dan PostgreSQL evidence saat fake boundary tidak dapat membuktikan transaction/rollback/query semantics. O2–O5 tidak dianggap terbukti oleh OpenAPI validation atau unit/race saja. | `docs/spec/5-donation/tasks.md` Verification ownership; `docs/kencleng-agentic-workflow.md` §§4/15; `backend/AGENTS.md`; EXP Stage 2 Area 6/F-05; TP-015 §12. |

## 4. Rules & Validation

- **R1 — Scope and guest access:** Submission tidak memerlukan Account. Hanya operasi guest submit/status dalam kontrak aktif yang diimplementasikan; operasi donor-list, history/claim, reveal email, real rails, dan broader Slice 3 tetap tidak aktif.
- **R2 — Exact amount:** Terima hanya IDR whole Rupiah >= Rp5.000; Rp5.001 valid; tolak fractional/invalid amount dan method non-QRIS. Jaga decimal wire/storage/calculation tanpa floating point. Jangan menetapkan parameter range/scale/rounding yang belum disetujui.
- **R3 — Request retry:** Retry dengan key+payload sama menghasilkan satu Donation dan mengembalikan hasil yang sama; key sama dengan payload berbeda ditolak. Client/API handling tidak merotasi key saat outcome ambigu. Ini tidak menggantikan settlement replay protection.
- **R4 — Simulator authority:** New Donation persist sebagai `pending`; hanya behavior/config/fixture backend yang dapat memindahkannya ke satu hasil terminal `success`/`failed`. Request/UI tidak dapat memilih state. Pending tidak auto-resubmit atau memberikan instruksi/estimasi settlement nyata.
- **R5 — Guest email purpose and lifecycle:** Email tidak digunakan sebelum ownership verification. Hanya opt-in; tidak untuk campaign updates. Maksimal satu notice status-only, berlabel simulasi, pada terminal success/failed dan tidak pada pending. Verified-address eligibility dipertahankan sampai fulfillment; delivery lifecycle bounded/recoverable sesuai keputusan O3/O11. Unverified email dihapus hanya mengikuti window Security/PII yang disetujui.
- **R6 — Guest status credential:** Issuance dan lookup memakai credential dengan arah fragment/handoff/cleanup, one-way HMAC verifier, hard expiry 24 jam dari issuance, tanpa extension, dan status-only projection. Tidak ada PII, token, atau access link dalam response status atau ordinary logs.
- **R7 — Uniform public failure:** Absent Donation dan absent/wrong/expired credential menghasilkan `404` yang sama untuk Problem Details body, status, headers, dan cache directives; `private, no-store` berlaku pada hasil `404`. Cache policy response `200` masih perlu keputusan API/Security owner di O4/O5. Credential check, exposure, abuse, dan response/timing parity mengikuti kontrol owner-approved dan evidence gate.
- **R8 — D1 eligibility order:** Submission dan Campaign close memiliki satu hasil urutan yang konsisten: close-first ditolak, submission-first saat eligible diterima. Membaca snapshot `published` saja tidak memenuhi aturan ini.
- **R9 — Atomic exact-once funding:** Settlement pending→success bersama full funding increment adalah satu hasil atomic; failure/rollback tidak meninggalkan salah satu sisi ter-commit. Replay dan settlement concurrent tidak menggandakan increment atau kehilangan update. Pending/failed tidak dihitung sebagai funding.
- **R10 — Post-close and threshold:** Accepted pending Donation tetap dapat sukses penuh setelah close. Settlement sesudah close tidak membuka Campaign atau mengubah winning `closed_reason`; funding tidak di-cap pada `max_amount` dan dapat overshoot.
- **R11 — Campaign ownership and bounded scope:** Donation memakai seam eksplisit antar-domain; Campaign tetap pemilik lifecycle/eligibility/close reason. D1 implementation dibatasi kebutuhan yang accepted dan tidak membawa scheduler, force-close atau public result Slice 3. Jika implementasi menyentuh path transaction/locking balance yang root-fenced, berhenti untuk Human-paired/Tier-0 authorization.
- **R12 — Campaign GET producer dependency:** Implementasi producer `donation_action` menunggu Campaign/API contract WU-S2-005 diterima dan mengikuti scope/ownership di sana. Nilai GET adalah informasi availability; POST tetap melakukan otoritatif eligibility check sendiri.
- **R13 — Public sandbox truth:** Name/email guest tidak public. QRIS diberi label sandbox simulation; metode lain yang tampil nonaktif. Tidak ada UI/API state yang menyiratkan provider settlement atau pembayaran nyata.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | WU-S2-003 memiliki minimum backend capability untuk D1 dalam scope-nya; Campaign tetap pemilik coordination/lifecycle/close reason. | Chosen — Orchestrator route | Mengikuti kebutuhan WU dan TP-015; bukan keputusan atas mechanism atau file-level seam. Sebelum D1 Build, Techplan perlu concrete seam/Tier-0 impact check. |
| D1-alt | Memindahkan seluruh closure ke WU Campaign baru atau menarik full closure lifecycle ke Donation. | Rejected for this route | Yang pertama tidak diperlukan untuk rute delivery minimum yang diputuskan Orchestrator; yang kedua memperluas ke Slice 3 scheduler/force-close/result. Jika live seam tidak aman dalam scope saat ini, buka kembali melalui Orchestrator/Human, jangan diam-diam mengubah route. |
| D2 | Jangan memakai query `status='published'` sebagai bukti eligibility atomik; harus ada coordination boundary untuk submit-vs-close. | Chosen — D1 invariant | Read model saat ini adalah publikasi, bukan transaksi eligibility. Menyederhanakan menjadi check biasa melanggar race requirement. |
| D2-alt | Donation menerima berdasarkan Campaign GET atau read predicate `published` saja. | Rejected | Snapshot tidak serialisasi submission dengan close; dapat menerima donation setelah close menang. |
| D3 | Pertahankan O1 exact-decimal direction; jangan salin `NUMERIC(19,2)` Campaign sebagai Donation storage contract. | Chosen — shared monetary standard | Supported currencies/range/scale masih belum diputuskan dan Campaign precision hanya precedent domain. Storage parameter yang diperlukan harus diselesaikan di owner gate, bukan ditebak. |
| D4 | FakeSender/dev outbox bukan fulfillment produksi untuk terminal notice. O3 perlu jalur delivery yang benar-benar memenuhi behavior dan bounded recovery. | Chosen — product truth / Stage 3 | FakeSender hanya log/no-op di konfigurasi non-development dan dev outbox adalah inbox simulasi. Tidak ada provider/sender yang dipilih oleh Techplan. |
| D5 | Fragment handoff, one-way HMAC, hard 24-hour expiry, status-only, uniform 404, dan `private, no-store` pada uniform 404 merupakan contract direction yang settled; cache policy success dan implementation controls tetap owner-gated. | Chosen — TP-015 / agreed contract | Techplan tidak mengganti key-purpose policy, secret comparison, abuse controls, atau residual-risk acceptance; crypto/auth Tier-0 path tetap fenced. |
| D6 | Campaign GET action producer menunggu contract WU-S2-005 yang accepted; availability-only tidak menjadi authorization untuk POST. | Chosen — Work Graph dependency / owner direction | Arah owner sudah tercatat tetapi authored Campaign contract belum selesai. Donation submit/status/D1 plan tetap runnable tanpa producer ini. |
| D7 | Scope hanya guest submit/status, simulator, required email, dan minimum D1; operasi historis lainnya deferred. | Chosen — Product/MVP | API historis yang lebih luas tidak mengubah scope aktif. |

## 6. Backward Compatibility

- **Existing data:** Backend belum memiliki Donation table, route, atau Donation tests. Migration/schema baru perlu dirancang dari authority aktif. Jangan membuat inferensi storage Donation dari tabel Campaign atau data seed. Jika inventory live sebelum Build menemukan data/consumer baru, hentikan perubahan terkait dan evaluasi compatibility/O8 sebelum destructive migration.
- **API/contracts/clients:** Authored split `api/openapi/donation.yaml` dan komponen `common.yaml` adalah contract baseline. WU-S2-002 menyatakan submit/status historis belum didistribusikan untuk scope repo/current Slice 2; batas konfirmasi itu tidak berlaku ke consumer lain. Jangan ubah split authority dalam backend Build; jika drift/bug kontrak material ditemukan, route ke API owner. `api/openapi.yaml` adalah generated bundle, tidak diedit manual.
- **Campaign GET:** Contract producer masih menunggu WU-S2-005. Jangan mengubah response producer dari keputusan availability-only yang belum direkonsiliasi/diterima. Penambahan producer kelak harus kompatibel dengan contract yang diterima dan tidak menggantikan POST eligibility check.
- **Migration/deprecation:** PostgreSQL migrations mengikuti `golang-migrate`; update harus reversible bila sesuai data transition. Tidak ada migration/index manual yang boleh diterapkan dari Run ini. Jangan menghapus data/kolom/operasi historis tanpa O8/authority gate yang berlaku.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| E1 | D1 race antara submit dan close atau dua close contenders; accepted pending settlement sesudah close. | High | Critical | Jangan gunakan GET/read-only `published` check sebagai gate. D1 runtime butuh dua ordering dan post-close assertions. Mechanism belum dipilih; concrete seam dan Tier-0 check adalah Active Open Item. Tidak ada residual risk diterima. |
| E2 | Settlement concurrent/replay atau rollback di antara Donation terminal status dan funding update menyebabkan duplikasi, kehilangan increment, atau partial commit. | Medium | Critical | Atomicity, replay, concurrent successful Donation, rollback, dan threshold overshoot dibuktikan dengan DB-backed/invariant tests; race detector sendiri tidak cukup. Protected ledger/transaction locking tidak boleh diedit tanpa Human authorization. |
| E3 | Credential ditebak/dicuri atau bocor melalui URL/referrer/browser/proxy/log/cache; response membocorkan keberadaan Donation melalui body/header/cache/timing. | Medium | High | Terapkan settled contract; owner harus menentukan key/comparison, strength evidence, cache success, exposure, abuse/rate controls sebelum status capability diterima. Uji status/header/body/cache/timing pada topology yang relevan. Tidak ada risk acceptance. |
| E4 | O3 notice hilang/duplikat setelah commit, terkirim sebelum email verified, atau verified PII terhapus/retained tanpa terminalisasi finite. | Medium | High | Owner tentukan mekanisme retry/terminalization, delivery reality, verification/deletion window dan controls; buktikan idempotent delivery dan lifecycle races. FakeSender/dev inbox bukan bukti delivery produksi. |
| E5 | Campaign `NUMERIC(19,2)` dipakai sebagai standar Donation sehingga nilai valid/parameter money berubah diam-diam. | Medium | High | Gunakan exact decimal; putuskan batas/scale hanya berdasarkan supported-currency dan calculation requirement melalui owner. Test no-float dan round-trip pada nilai yang disetujui. |
| E6 | Campaign GET producer diubah sebelum contract WU-S2-005 diterima atau UI menganggap availability snapshot sebagai otorisasi. | Medium | High | Dependency HARD hanya untuk producer. Tahan perubahan Campaign GET sampai acceptance; POST tetap otoritatif. |
| E7 | `donation_action` belum tersedia/contract belum accepted menghambat Campaign Detail → flow walaupun Donation backend siap. | High | Medium | Catat sebagai dependency scoped WU-S2-005; lanjutkan submit/status dan D1 planning yang tidak membutuhkan producer. Integrasi end-to-end menunggu contract accepted. |
| E8 | O2 simulator mekanisme/timing belum ditentukan dan disalahartikan sebagai ETA atau settlement nyata. | Medium | High | Jangan tampilkan estimate/instruksi bayar; backend config/fixture mengontrol skenario gagal. Owner menyelesaikan scenario/timing sebelum simulator Build mengunci detail yang material. |
| E9 | `tasks.md` status tracker rows tertinggal `draft / owner review required` meski accepted specs/event dan `CONTRACT_READY` menyatakan sebaliknya. | Low | Low | EXP F-06 mengklasifikasikan ini sebagai stale tracking, bukan conflict acceptance. Jangan edit tracker di Run ini; rujuk event/manifest authority. |

## 8. Interface Contract

**Persistence/data shape:** Tambahkan model/persistence Donation hanya dari agreed specs dan authored OpenAPI. Detail schema/scale, retention fields, retry record lifetime, credential verifier storage, terminal notice tracking, serta verification token lifecycle harus mengikuti owner decisions yang masih terbuka. Monetary path exact-decimal dan tidak menggunakan `float64`; jangan menambahkan fields untuk claim/history/public donor list. Setiap SQL memakai `goqu` parameterized style. Persyaratan persistence harus mendukung submission idempotency dan settlement replay sebagai dua guarantees berbeda. Jangan memuat network/email call di dalam database transaction.

**API/external interface:**

- Submit: `POST /campaigns/{campaignId}/donations`; guest; idempotency key dari `common.yaml`; body/schema, success `201` dengan state awal `pending`; validasi amount dan method dari `SubmitDonationRequest`. `409` untuk campaign ineligible atau key reused dengan payload berbeda; detail tetap mengikuti authored contract.
- Status: `GET /donations/{donationId}/status`; optional `X-Donation-Status-Credential` karena missing credential juga harus uniform `404`; `200` memuat hanya status; `404` seragam untuk absent Donation/missing/wrong/expired credential dan membawa `Cache-Control: private, no-store`. Cache policy untuk `200` belum dipilih: API/Security owner harus menentukannya sebelum route diterima, lalu implementation dan tests mengikuti keputusan itu.
- Bearer token issuance dalam submit response digunakan frontend untuk fragment handoff lalu URL cleanup; backend tidak mengharapkan fragment dikirim ke server. Token, PII, request bodies sensitif, dan raw wrapped errors tidak masuk logs.
- Donation action di Campaign GET adalah contract surface terpisah; implementasi producer menunggu hasil WU-S2-005. Snapshot availability tidak menjamin POST.

**Cross-layer/business boundary:** Donation domain mengatur guest submission dan simulator; Campaign domain tetap memiliki lifecycle/eligibility/close reason. D1 membutuhkan koordinasi atomik untuk eligibility/close dan success/funding, tetapi port, transaction boundary, lock order/isolation, dan lokasi implementasi belum dipilih. Dilarang menyimpulkan bahwa Campaign public read repository adalah port mutation. Jika D1 memerlukan file yang menjalankan transaction/locking untuk balance updates atau Tier-0 crypto/auth, tahan perubahan tersebut untuk Human-paired authorization.

## 9. Architecture / Plan

Bentuk backend Donation capability sebagai domain package dan HTTP transport sesuai domain-driven monolith convention yang ada. Migrasi tabel/constraints mendukung request idempotency, persisted state, credential verifier/expiry, optional email verification dan terminal-notice lifecycle hanya sesuai keputusan owner yang aktif. Interface/field spesifik tidak dibekukan di sini jika keputusan masih terbuka.

Flow submit menerima request yang tidak dipercaya, memvalidasi contract/amount/method, mencocokkan key-payload untuk request idempotency, mengoordinasikan eligibility dengan Campaign melalui seam yang eksplisit, lalu menyimpan `pending` dan mengembalikan contract response. Simulator backend-only memilih hasil sesuai konfigurasi/fixture. Pada `success`, settlement dan keseluruhan funding update harus menjadi satu exact-once outcome dalam batas Campaign/Donation yang disepakati; `failed` tidak menambah funding. Donation yang telah diterima sebelum close dapat settle penuh sesudahnya tanpa mengubah close reason.

Status flow menerima `donationId` dan optional bearer credential header; mengecek verifier/expiry dan mengembalikan status-only response atau public uniform not-found. URL fragment ditangani frontend, bukan server. Cache, logging, browser exposure, comparison, abuse dan response parity controls menunggu O4/O5 owner decision dan harus diverifikasi terhadap stack/proxy yang sebenarnya.

Email flow opsional, jika dipilih/diwajibkan fulfillment path, memverifikasi kepemilikan sebelum detail/access notice, tidak mengirim saat initial pending, dan menjamin terminal notice paling banyak satu kali setelah terminal state. Verified address tetap memenuhi kewajiban sampai delivery berhasil atau terminalization bounded/recoverable yang disetujui. Tidak ada sender production yang sekarang memenuhi promise itu; O3 harus menetapkan path sebelum fulfillment diklaim.

Dua integration gates tetap berbeda: (1) D1 domain seam yang harus diselesaikan sebelum D1 implementation; (2) Campaign GET producer yang menunggu accepted contract WU-S2-005. GET producer bukan authorization boundary dan tidak menggantikan eligibility saat POST.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `backend/cmd/server/main.go` — `run` route registration | Router sekarang hanya mendaftarkan Campaign public, Auth, Account; tidak ada Donation route. | Tambahkan hanya endpoint Donation yang sesuai kontrak setelah status/security controls siap; jangan menambahkan operasi deferred. |
| `backend/internal/domain/campaign/repository.go` — `Repository`, `SeedRepository` | Repository yang ada hanya public read/media dan seed; tidak ada Donation-facing mutation seam. | Gunakan sebagai bukti seam masih harus dirancang. Jangan memperluas read port menjadi asumsi transaksi D1 tanpa ownership/transaction review. |
| `backend/internal/domain/campaign/repository_db.go` — `FindPublicDetail` | Prepared `goqu` query menyaring `status='published'`; read query bukan ordering terhadap close. | Tidak boleh dijadikan eligibility proof. Bangun seam hanya setelah D1/Tier-0 gate; pertahankan query parameterized. |
| `backend/internal/domain/campaign/service.go` — `toPublicDetail`, `mapFunding` | Public read model mengembalikan action unavailable; map funding menggunakan `shopspring/decimal` untuk Campaign format saat ini. | Campaign GET producer hanya setelah WU-S2-005 accepted. Decimal dependency adalah precedent, bukan Donation scale decision. |
| `backend/internal/transport/http/campaign_public.go` — public handler/cache/error writers | Menunjukkan pola handler, JSON projection, dan `private, no-store`; public campaign tidak memenuhi Donation credential semantics. | Gunakan pola error/projection yang sesuai; jangan menyalin business behavior tanpa reconciliation. Cache policy untuk Donation status success tetap perlu keputusan owner. |
| `backend/internal/transport/http/middleware.go` — `RateLimit` | Existing limiter menggunakan `RemoteAddr`, idle-key eviction; wiring saat ini bukan untuk public Donation. Reverse-proxy caveat berlaku. | O4/O5 owner perlu menentukan abuse control dan trust boundary sebelum route exposure; jangan menganggap middleware saat ini cukup. |
| `backend/internal/transport/http/errors.go` — `MapServiceError` | Generic public 500, namun unhandled internal error dicatat mentah. | Audit error chain/log output untuk token/email payload sebelum Donation route aktif; sanitize structured logs tanpa melemahkan internal error chain. |
| `backend/internal/platform/crypto/crypto.go` — `HMAC`; `keys.go` — `Keys`/`New` | Existing HMAC machinery dan key-purpose allocation hanya untuk established encrypted-PII lookup; root-fenced. | Read-only precedent. Jangan reuse untuk credential verifier atau mengubah crypto/key files tanpa keputusan key-purpose Security/PII dan Human-paired authorization. |
| `backend/internal/platform/notification/sender.go` / `dev_sender.go` — `Sender`, `FakeSender`, `DevSender` | Sender sekarang Account-only; production FakeSender tidak deliver dan dev inbox tersimulasi. | O3 harus menentukan guest delivery interface/provider and recovery semantics. Jangan klaim sender sekarang memenuhi external notice. |
| `backend/migrations/000011_create_public_campaigns.up.sql` — `campaigns` schema | Schema Slice 1 memiliki collected/target `NUMERIC(19,2)`, tidak ada Donation flow/`closed_reason` workflow. | Gunakan hanya untuk memahami Campaign saat ini. Jangan jadikan skala Campaign sebagai standar Donation; jangan tambah full closure lifecycle melalui migration Donation. |
| `backend/internal/domain/campaign/repository_integration_test.go` — `integration` build tag / isolated Postgres helper | Existing pattern menggunakan testcontainers, migrations dan PostgreSQL 16. | Gunakan pola ini untuk bukti query/transaction/migration yang hanya dapat dibuktikan dengan Postgres; integrasi D1 harus menguji invariant sebenarnya. |
| `api/openapi/donation.yaml` — submit/status operations, request/response schemas | Authored source contract; menghasilkan operation list lebih luas dari active scope. | Jangan edit dari Backend Build. Gunakan sebagai field/response authority; contract discrepancy dikembalikan ke API owner. |
| `api/openapi/campaign.yaml` — `GET /campaigns/{campaignId}`, `PublicCampaignDonationAction` | Bentuk Campaign GET sekarang membawa donation action availability; rekonsiliasi WU-S2-005 masih berjalan. | Tunggu contract owner acceptance sebelum mengubah producer. |

Tidak ada file produksi yang dapat diubah dengan aman untuk keseluruhan D1 sebelum Open Item 1 ditutup. Open Items 2–4 juga menahan detail/acceptance yang bergantung pada keputusan O3/O4/O5/O2; implementer tidak boleh mengisi celah tersebut dengan asumsi.

## 11. Files Changed / Files NOT Changed

### Expected downstream files (indicative; confirm against Open Items before Build)

| File / area | Change type | Description |
|---|---|---|
| `backend/internal/domain/donation/` | New | Donation entity/service/repository; exact modules depend on approved execution seam and owner decisions. |
| `backend/internal/transport/http/` | New/modify | Guest submit/status handlers, explicit projections and errors. Status route remains gated by O4/O5 controls. |
| `backend/migrations/` | New | Donation persistence/constraints designed to exact-decimal and approved lifecycle needs; no speculative universal scale. |
| `backend/cmd/server/main.go` | Modify | Wire only active Donation dependencies/routes after gates. |
| `backend/internal/platform/notification/` | Conditional modify | Only if O3 chooses a supported guest delivery path; do not treat Account-only sender or dev outbox as production fulfilment. |
| Campaign backend domain and migration | Conditional modify | Minimum D1-owned seam is unresolved; no broader closure lifecycle. Any protected transaction/locking file requires Human authorization. Campaign GET producer requires accepted WU-S2-005 contract. |
| `backend/**_test.go` | New | Unit/handler and integration/concurrency evidence mapped to §12. |

### Files / areas intentionally untouched

| File / area intentionally untouched | Why |
|---|---|
| `frontend/` | Backend-scoped WU; separate FE delivery boundary. |
| `docs/product/**`, `docs/spec/**`, `api/**` | Authority/spec/contract writes are outside this backend Build. Open discrepancies route to the owning authority. |
| Root-fenced Donation balance transaction/locking logic; `backend/internal/platform/crypto/`; protected auth core; disbursement state machine | Tier-0 Human-paired boundary; no write is authorized. |
| WU-S2-005 specs/API or producer before acceptance | Scoped HARD dependency for Campaign GET donation-action producer. |
| Existing EXP/TP/Review evidence and orchestration projections (`manifest`, `events`, `control-surface`, `outcome`, `work-graph`) | Current Run may write only its own planning and phase provenance artifacts; orchestration projections belong to Orchestrator. |
| `api/openapi.yaml` / generated frontend types | Generated artifacts, outside backend Build; no direct edits. |

## 12. Testing Checklist

This is a planning Run: no tests, contract validation, migration, runtime/security checks, or UI acceptance were executed. Each rule below has downstream evidence owner(s); no check is claimed complete.

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Human/spec/API scope review plus route inventory confirms guest-only active endpoints and deferred historical operations remain absent. | Human / Testing | Scope mistakes expose unsupported data/actions and widen MVP. |
| R2 | Table-driven amount tests: Rp4.999 reject; Rp5.000/Rp5.001 accept; fractional Rupiah and wrong currency reject; exact-decimal round-trip/calculation tests and code review show no float conversion. | Testing | Bad monetary validation or binary rounding corrupts donation/funding meaning. Owner resolves range/storage details first; don't invent test values outside accepted bounds. |
| R3 | HTTP/repository tests for same key+same payload after ambiguous retry, same key+different payload rejection, concurrent duplicate submission and deliberate fresh key; assert one Donation. | Testing | Retry semantics prevent accidental duplicate intent; settlement replay is separately covered by R9. |
| R4 | Service/API security tests prove request fields cannot set terminal status, initial state persists `pending`, simulator-only transitions, failure only via approved backend scenario, and no auto-resubmit/ETA semantics. | Testing | Contract text cannot prove authority boundaries or persisted result; false success/failure can mislead guests. O2 scenario/timing decision must precede locking those details. |
| R5 | Security/PII owner review of approved controls/lifecycle; integration tests cover no email before verification, no pending notice, terminal notice at most once, retry/recovery, retention and deletion/verification races, and verified eligibility through fulfillment. | Human / Testing | A false send, duplicate notice, lost terminal obligation, or indefinite PII retention is material. O3 mechanisms and risk acceptance remain owner-gated. |
| R6 | API/Security review of generation/strength/key-purpose/comparison/exposure/expiry controls; runtime tests for 24-hour boundary and no expiry extension; inspect logs/browser/proxy/cache handling. | Human / Testing | Credential possession unlocks private status; OpenAPI direction alone is not safe implementation proof. Avoid protected crypto/auth writes absent authorization. |
| R7 | Handler integration tests compare absent Donation/missing/wrong/expired cases byte-for-byte for status, Problem body, headers and cache behavior including `private, no-store`; test 200 cache behavior against the API/Security owner decision; topology-aware timing/abuse assessment by Security/Testing. | Testing | Response parity and cache behavior are externally observable anti-enumeration controls; code-only equality cannot establish proxy/cache/timing behavior. |
| R8 | PostgreSQL concurrency integration test for submit-vs-close both ways, with outcome assertions for winner; include close-first rejection and submit-first acceptance. | Testing | A read predicate/fake repository cannot establish cross-transaction ordering; errors accept new donations after close or reject accepted intent. Mechanism requires Open Item 1. |
| R9 | PostgreSQL integration tests for atomic status+funding rollback, settlement replay, concurrent settlements/concurrent donations, no lost increment; focused invariant assertions plus scoped `go test -race` where changed code has concurrent Go state. | Testing | DB atomicity and business exact-once are not proven by race detector or unit fakes. D1 transaction/locking surface is Tier-0 fenced and needs Human route before implementation. |
| R10 | PostgreSQL/runtime tests for accepted-pending donation settling full after close, unchanged winning close reason, no reopening, and threshold overshoot; test pending/failed exclusion. | Testing | Prevents lost accepted funds, closure-truth mutation, or unauthorized threshold cap. |
| R11 | Independent Code Review checks domain ownership, route scope, query parameterization, failure paths and protected-path diff; Human verifies Tier-0 authorization before any protected write. | Human / Testing | Cross-domain balance code and security boundaries require independent authority; easy implementation could breach root fence or couple internals. |
| R12 | After WU-S2-005 contract acceptance, contract and handler tests verify the Campaign GET producer matches accepted availability-only schema; separately prove POST checks eligibility at submission. | Testing | Stale/invented GET semantics break FE integration; treating snapshot as auth creates a race/security flaw. This row is blocked only for producer work. |
| R13 | Human rendered/product acceptance for active vs unavailable method labels, simulation disclosure, no real-payment instruction, and no public guest PII; API projection tests assert no guest name/email. | Human / Testing | Automation cannot judge truthful meaning of the rendered experience, and public projection tests prevent PII disclosure. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Concurrent submit retries and simulator settlement/replay | Duplicate requests or concurrent/replayed transitions can create duplicate Donations or terminal/funding effects. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md#area-2--backend-donation-submission-amount-idempotency-and-simulator`; Stage 2 Area 6. | Yes — request and settlement idempotency remain separate runtime obligations; no Donation implementation/evidence exists. |
| Campaign close, Donation acceptance, settlement/funding atomicity and concurrency | Interleavings can accept after close, lose accepted pending donations, duplicate or lose funding, or change winning close reason. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md#area-3--campaign-eligibility-close-ordering-and-funding-integration`; F-02; Stage 3 `Direction considered` / D1. | Yes — D1 policy settled; mechanism and executable seam remain open. |
| Guest credential, anti-enumeration, logs/cache/referrer and abuse | Bearer theft/guessing can reveal private status; uniform contract does not prove runtime parity or safe key purpose. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md#area-4--guest-status-credential-private-response-and-abuse-boundary`; F-03. | Yes — settled directions survive; controls, empirical proof and residual-risk decision remain open. |
| Guest email verification, terminal notice, retention/deletion and recoverability | Guest PII/access-adjacent message may be sent unverified, lost/duplicated, or retained without bounded lifecycle. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md#area-5--guest-email-verification-and-terminal-status-notification`; F-04. | Yes — O11 direction survives; O3 controls/mechanism/evidence remain open. |
| Independent money, DB, and concurrency evidence | Race-free code alone cannot prove exact-once funding/atomic rollback or D1 behavior against PostgreSQL. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md#area-6--risk-tier-verification-ownership-and-available-test-evidence`; F-05. | Yes — no runtime evidence exists; independent Testing remains required. |

## 13. Open Items

### Active — needs external input or verification

1. **D1 concrete seam and Tier-0 gate (blocks D1 implementation and `BACKEND_VERIFIED` only).** Orchestrator has routed the minimum D1 capability to WU-S2-003 while Campaign retains lifecycle/eligibility/close/funding coordination ownership. Live Campaign code is presently read-only public detail/media plus seed and has no Donation mutation/close seam. Before D1 Build, Planner/Orchestrator must record the concrete cross-domain ownership/interface and determine whether the required transaction/locking balance path falls under root `AGENTS.md` Tier-0 fence. Do not choose an isolation/locking mechanism here or route around the fence. If protected implementation is required, Human-paired scope and explicit authorization are prerequisites. Owner: Orchestrator with Campaign/Donation owner and Human for any protected write. Evidence: EXP Stage 2 Area 3/F-02, Stage 3 `Remaining owner and evidence routes`, current `campaign/repository.go`, `campaign/repository_db.go`, migration `000011`, INV-campaign-13.
2. **O3 email fulfillment and lifecycle (blocks email fulfillment claim).** Resolve Security/PII-approved verification/deletion windows and controls, delivery sender capability, retries/idempotency, bounded/recoverable terminalization and timeout semantics, and retention/deletion-race evidence. Product behavior stays as accepted; FakeSender/dev outbox do not count as fulfillment. Owner: Donation Delivery + Security/PII; Product/Human for any material meaning or residual-risk decision. Evidence: TP-015 Q6/Q13/R5/R13; Donation INV-03/04; Stage 2 Area 5/F-04; Stage 3 remaining routes.
3. **O4/O5 status credential controls and residual-risk gates (blocks status exposure acceptance).** Security/API owners must determine concrete credential strength/generation, key-purpose and comparison handling, expiry enforcement, browser/referrer/log/cache protections, abuse/rate controls, and success-response cache directive; Security/PII/Human handles residual-risk acceptance. Then Testing demonstrates response/body/header/cache/timing parity and abuse behavior against relevant runtime topology. Preserve settled fragment/HMAC/24-hour/status-only/uniform 404 direction. Owner: Security/PII + API; Testing; Human for risk acceptance. Evidence: TP-015 Q7/R6; Donation INV-05; Stage 2 Area 4/F-03; Stage 3 `Public status privacy and abuse`.
4. **O2 simulator details (blocks final simulator behavior lock only).** Resolve backend-controlled demo failure scenario/configuration and any timing needed for implementation; never create a donor-selected result, ETA, or real-payment implication. Owner: Donation Delivery with Product owner for any user-visible promise. Evidence: TP-015 Q4/R4; Donation INV-07; Stage 2 Area 2; Stage 3 remaining routes.
5. **O1 monetary concrete persistence parameters (blocks final schema/type choice only where required).** If persistence/schema needs range, scale, supported-currency or derived-value precision, obtain owner decision/evidence first. Major-unit decimal string, explicit currency, exact decimal, whole-Rupiah IDR input, and no-float remain settled. Owner: shared monetary-standard/Product owner with Donation Delivery. Evidence: `kencleng-monetary-data-standard.md`; TP-015 Q2/R12; Donation INV-01; Stage 2 Area 3.
6. **Campaign GET `donation_action` producer (scoped dependency; does not block Donation submit/status/D1 planning).** Wait for WU-S2-005 authored spec/API reconciliation and owner acceptance before implementing the producer. Then confirm producer scope/owner in WU-S2-003 and update only through its authorized backend route; do not treat availability-only GET as submit authorization. Owner: WU-S2-005/API owner first, then Campaign backend/Orchestrator. Evidence: WU-S2-005 manifest and EXP Stage 2–3; Work Graph scoped HARD edge; current `campaign.yaml` and `campaign/service.go#toPublicDetail`.
7. **O8 consumer/distribution audit (conditional).** Existing TP-015 clears replacement of historical submit/status operations only within this repository/current Slice 2 based on owner confirmation. Reopen only if new in-scope consumers or a destructive compatibility operation are found before Build/migration. Owner: API owner/Human. Evidence: TP-015 Q10/R9; Donation spec O8.
8. **Donation task status tracker discrepancy (non-blocking, outside this Run).** `docs/spec/5-donation/tasks.md` tracker rows still say draft despite agreed headers/Event/CONTRACT_READY. EXP F-06 classifies stale rows as tracking inconsistency. Route a status-only reconciliation to tracker/spec status owner; do not reinterpret accepted criteria. Owner: project/spec tracker owner.

### Resolved — retained as decision history

1. ~~**D1 delivery route**~~ **RESOLVED —** Orchestrator assigned minimum D1 integration capability within WU-S2-003; Campaign retains lifecycle and close/funding coordination. This resolves work-unit routing only; concrete seam/mechanism/Tier-0 authorization remain Active above.
2. ~~**Campaign GET direction**~~ **RESOLVED —** Human selected availability-only direction in WU-S2-005 coordination. Authored Campaign contract is still Slice 1 until reconciliation and owner acceptance; producer dependency remains Active above.
3. ~~**D1 product ordering**~~ **RESOLVED —** Close-first rejects new submission; accepted Donation remains settleable in full after close; success and full funding commit atomically exactly once; later settlement does not reopen Campaign or change winning close reason; funding may exceed threshold. No mechanism selected.
4. ~~**O11 verified email eligibility**~~ **RESOLVED —** Verified opted-in address stays eligible until required terminal notice fulfillment; Delivery must provide bounded/recoverable terminalization. Numeric bound, timeout meaning, mechanism and risk acceptance remain Active under O3.
5. ~~**O4/O5 contract direction**~~ **RESOLVED —** Fragment handoff/cleanup, one-way HMAC verifier, 24-hour hard expiry, status-only, uniform public 404 and `private, no-store` are settled contract directions. Runtime controls, cache-success expression, empirical parity and residual-risk acceptance remain Active.
6. ~~**Donation backend existence**~~ **RESOLVED —** EXP confirmed capability absent; this is the delivery gap, not a Product/API uncertainty. Build remains a later authorized phase.
