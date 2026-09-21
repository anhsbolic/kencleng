# Review findings — WU-S1-002

> Phase             : Independent Techplan Review
> Work Unit         : WU-S1-002
> Run               : TPR-001
> Role              : Reviewer
> Model             : `gpt-5.6-sol`
> Reasoning         : `medium`
> Created           : 2026-09-21
> Target revision   : `83c930885549d287675f6b48a85303b294ef36a5`
> Workflow revision : `396b9ba664aaab9cb959786a97fd2594346c700e`

**Gate:** Complex — Techplan membawa breaking API contract, melintasi spec/OpenAPI/generated frontend contract, dan menyentuh security-sensitive public projection, anti-enumeration, serta media-revocation boundaries.

**Sections resolved:** Rules & Validation = §4; Decision Log = §5; Interface Contract = §8; Testing Checklist + Test Focus Pointer = §12; Open Items = §13.

### Blocking

- **[SECURITY / INTERFACE] Public allowlist belum ditutup secara executable di OpenAPI** — Lokasi: §4 R3/R8, §8 `PublicCampaignDetail` dan seluruh nested public object schemas, serta §12 checklist R3/R8. Defect: Techplan menetapkan bahwa public projection berisi *exactly* field yang didaftarkan dan menyatakan semua property sebagai `required`, tetapi tidak menetapkan `additionalProperties: false` untuk `PublicCampaignDetail` dan nested public projections. Pada OpenAPI 3.0.3, `additionalProperties` default ke `true`; `required` hanya mewajibkan kehadiran property dan tidak melarang property lain. Karena itu Build masih harus mengarang keputusan closed-object semantics, dan contract yang dihasilkan dapat tetap menerima field internal/operasional di luar allowlist meski seluruh required field hadir. Source evidence: `docs/product/probes/01-public-campaign-detail-contract-reconciliation.md#2-safety-principle--public-continuity-is-not-public-data-inheritance` mewajibkan explicit allowlist yang aman saat internal model berkembang; `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#decision-3--use-a-minimum-explicit-public-projection` menetapkan dedicated mapping boundary dan forbidden-field absence; carry-forward risk item 1 mewajibkan regression evidence; OpenAPI Specification 3.0.3 §4.7.24.1 menegaskan default `additionalProperties: true`. Material concern: boundary pencegah internal-field leakage belum menjadi contract yang unambiguous/executable. Resolution pass perlu mengunci `additionalProperties: false` pada setiap object schema di public response graph yang dimaksud sebagai exact projection, lalu memperbarui verification agar memeriksa closed-object behavior pada dereferenced bundle/generated correspondence.

### Non-blocking

- none.

### Clean

- Seluruh 14 Rules & Validation entries mempunyai coverage di Testing Checklist; tidak ada orphan rule atau orphan checklist item.
- Decision Log mempertahankan pilihan dan rejected alternatives material dari Exploration tanpa membuka ulang keputusan yang sudah settled.
- Open Items mempunyai lifecycle yang jelas: satu Active item non-blocking dan lima Resolved items dengan konsekuensi yang dipertahankan.
- Test Focus Pointer membawa forward projection leakage, existence disclosure/auth variance, media revocation/cache staleness, money truth, dan organizer-content safety; runtime concurrency/performance diberi `N/A` eksplisit karena WU ini tidak mengubah runtime.
- Diagram Mermaid tidak digunakan. Dua text-flow diagrams konsisten dengan sequencing dan dependency boundary di prose, tanpa branch yang overlap atau mustahil.
- Technical-fact spot checks cocok dengan live repository: `api/openapi/index.yaml` memang belum mengekspos `common.yaml` `bearerAuth`; `npm run validate` masih menghasilkan baseline 1 error/130 warnings; `api/openapi/campaign.yaml` masih memakai mixed `CampaignDetail`, public attachment list, dan float progress; `openapi-typescript` serta `tsc` tersedia di frontend; backend belum memiliki Campaign/Organization domain route. Relevant authority/source files tidak berubah antara target revision Techplan dan review HEAD.

## Phase handoff

- Completed: independent review gate + review when warranted
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TPR-001/review-findings.md`
- Human decision: revise — satu blocking security/interface finding perlu resolution pass sebelum Human Techplan approval
- Open / deferred: closed-object semantics dan verification untuk seluruh public response graph; non-blocking follow-up none
- Recommended next step: one resolution pass then human gate; re-review diperlukan bila resolution mengubah material interface/security semantics, kecuali Human gate secara eksplisit me-waive
- Session transition: gunakan fresh Planner/resolver Session agar reviewer tidak menulis ulang Draft Techplan; Build tetap fresh-preferred hanya setelah Approval
- Context pointers: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md` §4 R3/R8, §8, §12 R3/R8; `docs/product/probes/01-public-campaign-detail-contract-reconciliation.md#2-safety-principle--public-continuity-is-not-public-data-inheritance`; `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#decision-3--use-a-minimum-explicit-public-projection`; `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 1
