# Task 02 — Rekonsiliasi authored Donation OpenAPI

> Phase: Techplan decomposition task
> Author: Codex Planner
> Created / Updated: 2026-09-30
> Model: gpt-6-luna
> Reasoning: high
> Session: Fresh Planner Session; Session ID not exposed
> Target revision: `6891341a050982e14174ab5af132a200f24e71d9` plus durable working-tree artifacts; reopen current sources
> Workflow revision: `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`
> Work Unit: `WU-S2-002`
> Run: `TPD-S2-002-001`
> Participant: `P-S2-002-PLD-001-1`

## Purpose and scoped outcome

Rekonsiliasi authored Donation OpenAPI untuk guest submission dan status lookup Slice 2 agar selaras dengan Product/MVP, Design, spec Donation yang direkonsiliasi, serta Techplan Approved. Hasil mencakup hanya path, field, encoding, error/response, header, cache, simulator control, dan komponen bersama yang dapat ditetapkan oleh authority yang berlaku. Pertahankan seluruh bagian yang bergantung pada Open Item sebagai gate; task boleh memajukan rincian independen yang cukup ditentukan, tetapi tidak boleh menyatakan contract final/ready selama prasyarat material belum dipenuhi.

## Parent Techplan

- Authority: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md`
- Version: current-effective artifact pada target revision Run ini
- Status: `Approved`
- Techplan tetap menjadi spine lintas task untuk semua scope, keputusan, risiko, contract, verification, dan Open Item.

## Governing parent IDs

- Requirements: Q1–Q11.
- Rules: R1–R11, khususnya R2–R9 dan R11.
- Decisions: D1–D16 dan D16-alt sebagai riwayat; jangan membuka ulang keputusan settled.
- Risks: RISK-1–RISK-10.
- Verification: kewajiban §12 untuk R1–R11; schema validation/generation tidak menggantikan owner review, security evidence, Human rendered acceptance, atau runtime Testing.
- Open Items: O1–O5, O8 conditional, O11; setiap field/shape yang bergantung padanya tetap gated sesuai §13.

## Scope dan code/spec anchors

- Prasyarat spesifikasi: Task `01-donation-domain-spec-reconciliation.md`; hasilnya menjadi current delivery-spec input. Tetap baca Techplan spine dan live authority secara mandiri.
- `api/openapi/donation.yaml` — authored Donation paths dan schemas yang menjadi target utama rekonsiliasi.
- `api/openapi/common.yaml` — ubah hanya bila shared component benar-benar diperlukan dan disetujui oleh kontrak aktif.
- `api/openapi/index.yaml` — perbarui path registry secara mekanis hanya untuk path yang ditambah/dihapus.
- `api/README.md` — ikuti workflow split source, validation, bundle, dan generated types.
- `docs/spec/5-donation/` serta Campaign cross-reference yang disetujui — sumber delivery behavior setelah Task 01; jangan menganggap bentuk historis OpenAPI sebagai authority.
- `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6, `docs/ui-ux/README.md`, `AGENTS.md` §§1–3, dan Techplan §§3–13 — authoritative scope/decisions/boundaries.

## Detail pelaksanaan dan batas gate

- Terjemahkan same-key/same-payload, same-key/different-payload, deliberate new-key, serta pemisahan idempotency submission dari settlement replay. Retry-record lifetime/serialization harus mengikuti keputusan dan kontrak berwenang; jangan menciptakan policy baru.
- Pertahankan QRIS sebagai satu-satunya sandbox simulation aktif; tidak ada rail nyata atau client-callable settlement. Status berasal dari backend simulator.
- Status response hanya mengungkap status yang diizinkan. Generic copy untuk absent/invalid/expired sudah dipilih, tetapi transport parity, token controls, abuse behavior, dan residual-risk acceptance tetap membutuhkan keputusan/evidence Security/API O4/O5.
- Pertahankan D1/R7: submit-vs-close ordering, accepted Donation tetap settleable penuh setelah close, settlement/funding atomik exact-once, winning close reason stabil, dan funding dapat melampaui threshold. Jangan memilih mekanisme database/locking atau memperluas Slice 3.
- O1 `AUTHORITY_SYNC`: jangan finalkan currency code/wire representation, supported currencies, global precision/range, DB scale, atau derived precision/rounding sebelum owner/scope dan keputusan yang diperlukan tersedia. Aturan whole-Rupiah input, min/increment, exact-decimal/no-float tetap berlaku.
- O2/O3 dan O11 `HUMAN_DECISION`: jangan finalkan simulator timing/scenario detail yang belum diputuskan atau verified-email retention/terminal-notice contract detail sebelum gap O11 direkonsiliasi. Jangan memilih cap, alternatif retensi, numeric terminal bound, timeout-as-`failed`, atau menerima risk.
- O3–O5: verification/delivery, retention, encryption/HMAC, token lifecycle/exposure, parity, abuse, dan risk acceptance membutuhkan owner evidence yang berlaku; jangan menyamarkan keputusan tersebut sebagai schema defaults.
- O8: sebelum menghapus/mengganti operasi historis, API/Orchestrator mencatat bukti consumer/distribution; jika tidak dilakukan remove/replace, pertahankan status conditional.
- Update authored source saja. Jalankan `cd api && npm run validate`; jika authored source berubah, ikuti `api/README.md` untuk bundle dan frontend types, lalu periksa source/index/bundle/generated consistency. Jangan hand-edit generated outputs.

## Hard dependency dan gate

- Hard dependency task: `01-donation-domain-spec-reconciliation.md` harus tersedia untuk memastikan OpenAPI mengekspresikan behavior domain yang telah diselaraskan.
- External gates untuk rincian terkait: O1 `AUTHORITY_SYNC`; O2/O3-O11 `HUMAN_DECISION` dan owner/security decisions/evidence; O4/O5 API/Security decisions; O8 API/Orchestrator evidence hanya sebelum remove/replace.
- Bagian API yang tidak bergantung pada gate dapat dikerjakan setelah Task 01. Jangan menahan seluruh task bila hanya field/behavior tertentu yang gated, tetapi catat dengan jelas kontrak yang belum final. Tidak ada klaim `CONTRACT_READY` sampai syarat Techplan §9/§13 dipenuhi dan owner review selesai.

## Verification sebelum handoff

- Owner/Human memeriksa scope Slice 2 dan batas Slice 3 (R1); amount/method direction dan exact-decimal boundaries (R2); state/funding contract (R3); simulator semantics (R4); email decisions/retention gates (R5); token/status parity/security gates (R6); D1 translation (R7); request idempotency (R8); compatibility evidence bila remove/replace (R9); exact O7 wording/guidance (R10–R11).
- Jalankan validasi OpenAPI dan documented generation checks sesuai `api/README.md`; periksa authored source, path index, aggregate bundle, dan generated types bila berubah.
- Simpan R3/R4/R5/R6/R7/R8 runtime evidence untuk downstream Testing. Schema lint bukan bukti atomicity, concurrency, simulator runtime, PII handling, atau anti-enumeration.
- Bila authority atau material contract shape belum ada, laporkan sebagai Decision/gate dan hentikan hanya rincian terdampak; jangan mengarang atau menutup Open Item.

## NOT dalam task ini

Tidak mengubah Product/MVP, Design authority, Techplan, runtime code, tests, migrations, atau orchestration projections. Tidak mengedit `api/openapi.yaml` maupun generated frontend types dengan tangan; tidak mengubah Tier-0. Tidak menetapkan O1/O2/O3/O4/O5/O8/O11, tidak menerima residual risk, dan tidak menyatakan `CONTRACT_READY` tanpa gate dan evidence yang diwajibkan.
