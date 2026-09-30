# Task 01 — Rekonsiliasi spec domain Donation

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

Rekonsiliasi sumber spec Donation untuk Slice 2 terhadap Product/MVP, Design, dan Techplan Approved. Hasilnya adalah invariants, threat model, task list, dan feature acceptance yang konsisten untuk guest submission, status simulasi yang truthful, safe status revisit, serta batas threshold/eligibility Campaign yang disetujui. Klasifikasikan materi historis `KEEP`, `ADAPT`, `REPLACE`, atau `DEFER` dengan alasan berbasis authority. Pertahankan semua Active Open Item sebagai batas yang eksplisit; pekerjaan spec yang tidak bergantung pada keputusan tersebut boleh maju.

Task ini tidak membuat atau mengubah OpenAPI dan tidak membuat keputusan owner. Penyelesaian task ini tidak berarti `CONTRACT_READY`.

## Parent Techplan

- Authority: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md`
- Version: current-effective artifact pada target revision Run ini
- Status: `Approved`
- Techplan adalah spine lengkap dan otoritatif. Task ini hanya menambahkan batas eksekusi; bila detail bertentangan atau material tidak tersedia di spine, hentikan bagian terdampak dan laporkan gap melalui jalur yang berwenang.

## Governing parent IDs

- Requirements: Q1–Q11.
- Rules: R1–R11.
- Decisions: D1–D16 dan D16-alt sebagai riwayat keputusan yang ditolak.
- Risks: RISK-1–RISK-10.
- Verification: seluruh kewajiban §12, terutama owner review yang mengikat isi spec (R1–R7, R10–R11); kewajiban runtime pada §12 tetap untuk fase Testing berikutnya.
- Open Items: O1–O5, O8 (conditional), dan O11; status/penanganannya tetap mengikuti §13.

## Scope dan code/spec anchors

- `docs/spec/5-donation/invariants.md` — evaluasi INV-donation-01…15 dan state machine historis; adaptasi hanya aturan Slice 2 yang didukung spine.
- `docs/spec/5-donation/threat-model.md` — evaluasi ancaman submit, status, email, token, abuse, dan residual risk; jangan menganggap risiko diterima.
- `docs/spec/5-donation/tasks.md` dan `docs/spec/5-donation/features/` — bentuk task list dan feature acceptance yang hanya mencakup guest submit, truthful simulator state, status/recovery, email sesuai keputusan, dan perilaku batas Campaign yang diperlukan.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` — bukti historis untuk submission, settlement, atomicity, dan concurrency.
- `docs/spec/5-donation/features/02-donation-status-check.md` — bukti historis status-link dan konflik transport `401`/`404`.
- `docs/spec/5-donation/features/03-public-donor-list.md` sampai `06-guest-email-reveal.md` — klasifikasikan ruang lingkup historis sebagai `DEFER` atau status lain hanya jika dibenarkan oleh Product/MVP; jangan bawa fitur tertunda ke Slice 2.
- `docs/spec/4-campaign/invariants.md` `INV-campaign-13` dan `docs/spec/4-campaign/features/09-closure.md` — conditional reference/perubahan sempit bila pemilik Campaign memerlukannya untuk eligibility/threshold. Jangan merekonsiliasi lifecycle closure Slice 3.
- Authority dan aturan: `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6, `docs/ui-ux/README.md` dan `docs/ui-ux/patterns.md` §§7, 14–15, `docs/spec/README.md`, `AGENTS.md` §§1–3, serta Techplan §§3–13.

## Detail pelaksanaan dan batas gate

- Pertahankan IDR whole-Rupiah min Rp5.000 dan increment Rp1, exact-decimal/no-float, state simulator `pending`/`success`/`failed`, same-key retry policy, guest privacy, generic unavailable-link behavior, dan keputusan D1 persis seperti spine.
- Bawa copy Design Q11/R11 persis ke acceptance yang relevan. Ini bukan penerimaan visual UI.
- Catat bahwa detail representasi/wire/storage/precision O1 belum dapat difinalkan sebelum `AUTHORITY_SYNC` dan keputusan owner terkait. Jangan memilih format mata uang atau presisi turunan.
- Catat O2 dan O3 sebagai Active untuk timing/scenario serta email verification/retention/retry/control. Catat konflik O11 sebagai scoped `HUMAN_DECISION`; jangan memilih cap, retensi alternatif, batas terminal, makna timeout, atau risk acceptance.
- Catat O4/O5 sebagai Active untuk kontrol token, anti-enumeration, response parity, abuse, dan residual-risk decision. Generic copy yang sudah dipilih tidak menyelesaikan transport/control.
- Pertahankan O8 conditional: consumer/distribution audit dibutuhkan sebelum remove/replace operasi historis.
- Risiko security/PII dan batas Tier-0 tetap sesuai Techplan dan `AGENTS.md`; task ini tidak mengubah implementation authority.

## Hard dependency

Tidak ada hard dependency task. Task ini dapat mulai sekarang dari Techplan Approved dan sumber live saat ini. Keputusan eksternal yang masih terbuka harus tetap ditandai; keputusan itu hanya menghalangi rincian spec yang bergantung padanya.

## Verification sebelum handoff

- Tunjukkan traceability spec ke Q1–Q11 dan R1–R11; jangan menghilangkan kewajiban atau memperluas Slice 2.
- Pastikan threshold/eligibility/funding menerjemahkan D1/R7 tanpa memilih mekanisme locking/isolation dan tanpa mengubah close reason atau membawa Slice 3.
- Pastikan exact decimal, no `float64`, idempotency submit versus replay settlement, backend-only terminal state, privacy, serta O1/O11/Open Item boundaries tercermin akurat.
- Pastikan threat model tidak menandai residual risk Security/PII sebagai diterima tanpa owner authority.
- Owner review dan runtime Testing di §12 tetap kewajiban downstream; jangan klaim sudah terpenuhi hanya dari penulisan spec.

## NOT dalam task ini

Tidak mengubah Product/MVP, Design authority, Techplan, API/OpenAPI, runtime code, tests, migrations, generated artifacts, atau orchestration projections. Tidak mengubah file Tier-0. Tidak memutuskan O1/O2/O3/O4/O5/O8/O11, tidak menerima residual risk, dan tidak mengklaim `CONTRACT_READY`.
