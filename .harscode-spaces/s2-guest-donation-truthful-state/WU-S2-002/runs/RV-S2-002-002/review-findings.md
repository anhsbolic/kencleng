# Review findings — WU-S2-002

> Phase             : Independent Techplan Review
> Author            : Codex Reviewer
> Created           : 2026-09-26
> Model             : `gpt-6-luna`
> Reasoning         : `high`
> Session           : Fresh independent Reviewer Session (ID not exposed)
> Target revision   : `416e60415c51d0be7444638581ad206add24992e`
> Workflow revision : `cb5dca028b1d6ba49d43300dd06ec1bf2a9c984b`

**Gate:** Complex — review wajib karena Techplan memuat 7 Rules & Validation, merekonsiliasi spec/API lintas domain dan batas layanan, serta menyangkut uang, PII, credential status guest, dan settlement.

**Sections resolved:** Background = §1; Scope = §2; Requirements = §3; Rules & Validation = §4; Decision Log = §5; Backward Compatibility = §6; Edge Cases & Risks = §7; Interface Contract = §8; Architecture / Plan = §9; Implementation Details = §10; Files Changed / Files NOT Changed = §11; Testing Checklist + Test Focus Pointer = §12; Open Items = §13 (mapping diperiksa pada template aktif).

### Blocking

- **[MONEY / VERIFICATION] Kebijakan idempotensi untuk retry/double-submit guest belum menjadi rule atau Open Item yang dapat diverifikasi** — Lokasi: §3 Q3, §4 Rules & Validation (R3), §5 D3, §12 checklist R3 dan Test Focus Pointer baris settlement. Defect: Q3 hanya menyatakan duplicate submission ditangani “sesuai kebutuhan”; D3 menyebut idempotency “bila diperlukan”; R3 dan checklist menguji replay settlement serta settlement konkuren, tetapi tidak menentukan/menguji apakah pengulangan POST setelah timeout/ambiguous response atau double activation dengan key baru dapat membuat kontribusi kedua. Pointer F1 berfokus pada transisi settlement duplikat dan tidak menunjuk risiko submission retry dari Exploration Area 1/5. Karena kebutuhan idempotensi memang masih bersyarat, Techplan harus mencatat keputusan/verification ini sebagai Active Open Item atau menetapkan rule dengan sumber authority yang tepat; saat ini Build/spec reconciliation masih harus mengarang batas perilaku dan bukti penerimaannya. Source evidence: Product/MVP `docs/product/mvp-delivery-slices.md` §5 mensyaratkan perlindungan duplicate client submission “where idempotency is required”; Stage 2 `Area 1 — Product/MVP and slice boundary` mencatat retry setelah respons ambigu, `Area 3 — Donation delivery specs and API evidence` mencatat same-key retry, dan `Area 5 — Frontend live state and cross-stack surface` mencatat double activation/fresh keys; Stage 3 `Recommended material direction` mengarahkan idempotent handling untuk repeated submissions bila diperlukan. Material concern: potensi kontribusi/funding kedua akibat satu niat donor tidak tercakup secara terpisah dari replay settlement; keputusan ini memengaruhi business/API semantics dan verification. Rujukan: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 1 — Product/MVP and slice boundary`, `#Area 3 — Donation delivery specs and API evidence`, `#Area 5 — Frontend live state and cross-stack surface`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md#Recommended material direction`.

### Non-blocking

- none.

### Clean

- **Atomic coupling finding RV-S2-002-001 ditutup.** §4 R3 dan §8 Persistence/data shape kini menetapkan satu hasil bisnis atomik: setiap committed/observable state memuat status `success` bersama refleksi funding, atau keduanya tidak committed; kontribusi sukses hanya dihitung sekali. §7 RISK-3 dan §12 R3 meminta bukti kegagalan parsial, replay, serta concurrency/race. Ini mempertahankan arah Stage 3 `Recommended material direction` tanpa menetapkan transaction/lock/ledger primitive. Batas Tier-0 tetap eksplisit sebagai area protected dan tidak ada otorisasi implementasi.
- **Rule fidelity:** Ada 7 rule R1–R7 dan masing-masing memiliki checklist row. R1–R7 secara umum mengikuti scope, authority gaps, truthful sandbox, backend settlement, guest security, Campaign eligibility, serta source/generated API workflow. Catatan blocking di atas adalah bagian duplicate-submission yang belum terpisah dari settlement verification.
- **Decision fidelity:** D1–D8 konsisten dengan `Direction considered` dan `Material rejected alternatives`; pilihan yang telah settled tidak dibuka ulang, sementara payment, credential, closure, dan compatibility boundaries tetap dibatasi sesuai Exploration.
- **Open Items lifecycle:** O1–O8 berstatus Active dan tiga item resolved tetap dicatat beserta resolusinya. Ambiguitas retry/idempotency pada finding blocking perlu dibawa masuk ke lifecycle ini.
- **Diagram:** Tidak ada diagram; pemeriksaan sintaks dan semantik diagram tidak berlaku.
- **Technical-fact / guardrail spot-check:** `api/README.md` mengonfirmasi split source sebagai authored authority, pembaruan `index.yaml` saat path berubah, dan bundle/types sebagai output turunan; `backend/cmd/server/main.go` mendaftarkan Campaign detail/media tanpa route Donation; `backend/internal/domain/campaign/service.go` menetapkan `donation_action` unavailable dan memetakan funding memakai `shopspring/decimal`. Klaim Techplan pada anchor terkait cocok dengan target source `416e604…`; tidak ditemukan implementasi Donation runtime yang diklaim tersedia.
- **Test Focus Pointer:** Empat baris fokus settlement/concurrency, status credential, guest PII, dan eligibility/`max_amount` memiliki anchor ke heading Exploration spesifik serta status relevansi eksplisit. Anchor file-area dan Stage 3 heading yang dicantumkan dapat ditemukan; gap yang tersisa adalah fokus submission retry/double-submit pada finding blocking.

## Phase handoff

- Completed: independent review gate + review when warranted
- Artifacts: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-002/review-findings.md`
- Human decision: revise — satu finding material tentang kebijakan/verifikasi idempotensi guest submission masih terbuka sebelum Human Techplan gate
- Open / deferred: tentukan/route apakah retry dan double-submit guest wajib idempotent, lalu tambahkan rule, checklist, dan Test Focus pointer/anchor yang sesuai atau catat keputusan N/A dengan alasan authority
- Recommended next step: satu resolution pass lalu Human gate; re-review bila resolusi mengubah semantik uang atau strategi verifikasi secara material, kecuali Human gate secara eksplisit me-waive
- Session transition: gunakan fresh Resolver Session independen dari Planner; Build tetap fresh-preferred setelah Approval
- Context pointers: Techplan §3 Q3, §4 R3, §5 D3, §12 checklist R3/Test Focus Pointer; Exploration Stage 2 `Area 1 — Product/MVP and slice boundary`, `Area 3 — Donation delivery specs and API evidence`, `Area 5 — Frontend live state and cross-stack surface`; Stage 3 `Recommended material direction`
