# Kencleng — Development Tracker

> Status: Living project status
> Last reconciled: 2026-09-30
> Current Kencleng authoritative Product Authority baseline: `main@e32916b597412094976e3e6263095e861ec18391` (`Promote Product Authority and MVP delivery model (#27)`)
> Harscode operational baseline: `main@b64fa11082a094d0e1b6e9488c20eac1c7f9777b`
> Purpose: Keep cross-domain and product-slice delivery state visible without turning dated progress into workflow policy.

## 1. What this tracker owns

This file owns current project delivery state across product slices, domains, and cross-cutting frontend/backend work.

It does not replace:

- `docs/product/` for Product Authority, MVP scope, or MVP sequencing;
- `docs/spec/<domain-dir>/tasks.md` for domain-local task views;
- reconciled feature specs for detailed acceptance criteria;
- UI/UX authorities for design truth;
- `docs/project/kencleng-integration-map.md` for cross-stack structural dependency mapping;
- Harscode artifacts for one development run;
- `docs/kencleng-agentic-workflow.md` for project orchestration policy.

When status cannot be proven confidently, use `NEEDS_RECONCILIATION` instead of guessing.

## 2. Status vocabulary

Ordinary progress:

```text
NOT_STARTED
IN_PROGRESS
BLOCKED
NEEDS_RECONCILIATION
READY_TO_START
```

Evidence-backed milestones:

```text
CONTRACT_READY
BACKEND_VERIFIED
FRONTEND_MOCK_VERIFIED
INTEGRATED_VERIFIED
SLICE_FINALIZED
DELIVERED
```

Historical rows may use `DOMAIN_FINALIZED`; new MVP delivery should prefer the real product unit (`SLICE_FINALIZED`) where applicable.

Code existing in history is not sufficient evidence for a verified milestone.

## 3. Product Authority / MVP status

| Concern | Status | Evidence / notes |
|---|---|---|
| Whole-product Product Authority | `ACTIVE / AUTHORITATIVE` | `docs/product/product-overview.md`; promoted to `main` by PR #27 at `e32916b597412094976e3e6263095e861ec18391`. |
| Product Design / Brand Authority | `READY` | Canonical `docs/ui-ux/`; Sunlit Editorial / Evidence-Led Optimism approved. |
| MVP release scope | `APPROVED` | `docs/product/mvp-scope.md`; human approval 2026-09-17. |
| MVP delivery sequencing | `APPROVED` | `docs/product/mvp-delivery-slices.md`; human approval 2026-09-17. |
| Product Authority routing | `ACTIVE / AUTHORITATIVE` | Root/scoped AGENTS, product/spec routing, and orchestration were promoted with PR #27. |
| Probe 01 — Public Campaign Detail | `PASS` | Forward derivation + public lifecycle decision + narrow contract reconciliation recorded. |
| Probe 02 — Account Registration + Email Verification | `PAUSED_REFRAMED` | Useful salvage evidence; Account is outside baseline MVP critical path. |

Approved MVP loop:

```text
Slice 1 — Public Campaign Understanding
→ Slice 2 — Guest Donation + Truthful Donation State
→ Slice 3 — Campaign Closure + Persistent Public Result
→ Slice 4 — Accountability Follow-up
```

## 4. Current domain snapshot

Domain rows remain useful for semantic/implementation evidence, but they do **not** define current MVP delivery order.

| Domain | Delivery/spec state | Backend | Frontend | Current notes |
|---|---|---|---|---|
| Account | `NEEDS_RECONCILIATION` when next needed | Historical implementation exists | `NOT_STARTED` for new generation | Outside baseline MVP critical path. Preserve security/correctness evidence; do not resume historical roadmap by inertia. |
| Notification | Historical/draft reference | `NOT_STARTED` as standalone delivery | `NOT_STARTED` | Include only when a real slice requires active notification behavior. |
| Organization | Historical/draft reference | `NOT_STARTED` | `NOT_STARTED` | Slice 1 needs only minimum persisted/public-safe steward context; full self-service is deferred. |
| Campaign | `SLICE_FINALIZED` for Slice 1 | `BACKEND_VERIFIED` | `FRONTEND_MOCK_VERIFIED` | Contract, backend, frontend mock-parallel experience, topology/private media, and real cross-stack integration are verified; Human finalization approved for Slice 1. |
| Donation | `TP-S2-002-011` adalah current-effective Approved setelah Human meninjau report `TP-S2-002-012` dan status direkonsiliasi pada `TP-S2-002-013`. Independent review `RV-S2-002-006` bersih; D1 dipertahankan. | `NOT_STARTED` active generation | `NOT_STARTED` | Human menerima split `TPD-S2-002-001`: Task 01 Donation domain specs, lalu Task 02 authored OpenAPI. Build/review loop Task 01 selesai; domain specs masih `draft` menunggu acceptance owner/Human yang berlaku. Build Task 02 `BLD-S2-002-003` lulus validasi OpenAPI (126 warnings, tanpa error), tetapi tidak menghasilkan diff authored/generated API. O8 compatibility gate clear: Human/API owner mengonfirmasi operasi historis submit/status belum pernah didistribusikan eksternal; replacement terotorisasi dalam scope repo ini. O11 direction direkonsiliasi: route B superseded; verified email tidak boleh dihapus sebelum terminal notice selesai; Delivery harus menghasilkan terminalization policy bounded/recoverable, tanpa angka, architecture, timeout-as-failed, atau residual-risk acceptance. Anhar adalah named owner project-wide yang reassignable untuk Shared Currency Standard dan menyetujui direction: major-unit decimal-string + explicit currency code, exact-decimal calculation/persistence, Product-controlled active currency, dan current Slice-2 whole-Rupiah IDR input. Campaign `NUMERIC(19,2)` hanya precedent; additional currencies, numeric range, per-currency fraction rules, universal precision/scale, dan migration detail tetap belum diputuskan. WU kini `ACTIVE / QUEUED` untuk rekonsiliasi Task 02; TP-011 tetap current-effective Approved sampai direkonsiliasi melalui jalur planning/review/approval yang berlaku. Runtime proof O2–O5 tetap downstream. Belum ada contract atau milestone delivery Slice 2; money-core paths tetap Tier-0 fenced. |
| Disbursement | Historical/draft reference | `NOT_STARTED` | `NOT_STARTED` | Not baseline MVP critical path; do not pull in merely to make accountability look complete. |

## 5. Historical Account/backend evidence

Historical Account backend work remains useful and is not erased by the Product Authority reframe.

Known historical commits include:

| Scope | Historical main evidence |
|---|---|
| Account backend task 01 | `14834e5` — register/email verification |
| Account backend task 02 | `efc1111` → `ce61841` — Google OAuth sequence |
| Account backend task 03 + old frontend tasks 01–02 | `16a4bf9` |
| Account backend task 04 + old frontend tasks 03–04 | `ea7d5bc` |
| Account backend/frontend task 05 | `6f036c6` |
| Account backend/frontend task 06 | `50b9a18` |
| Account backend task 07 | `6a846bd` |
| Account backend task 08 | `0798c5d` — exploration only at the older checkpoint |

These entries prove historical work happened. They do not automatically prove current product relevance/completion. Probe 02 records the current salvage/reframe posture.

## 6. UI/UX and frontend readiness

The upstream Product Brand + UI Design Exploration is complete and human-approved.

Active authority includes:

- `docs/ui-ux/README.md` — routing map;
- `docs/ui-ux/brand-product-ui-brief.md` — Sunlit Editorial / Evidence-Led Optimism;
- `docs/ui-ux/product-design-principles.md` — design judgment/readiness/autonomy boundary;
- `docs/ui-ux/design-guidelines.md` — concrete visual system;
- `docs/ui-ux/patterns.md` — recurring interaction behavior;
- `docs/ui-ux/asset-governance.md` — asset lifecycle/truthfulness;
- `docs/ui-ux/page-map.md` — persona/surface intent;
- `docs/ui-ux/visual-references/selected-direction/` — approved direction evidence.

Removed legacy design/prototype generations are Git history only.

Frontend reboot:

```text
COMPLETE
```

Frontend Experience Foundation Task 01:

```text
DELIVERED
```

Evidence:

- PR #24 / `71093b687cd7135495bc6ed62d520a621f96f586` implemented the representative `/` calibration surface and received human rendered acceptance PASS;
- PR #25 / `fca5a8f3178b53e5eec006064d8bcf2b078771b3` closed Validation 02;
- PR #26 / `15a3e02cc88d00e8dee70f8dfb07c36cb1e2fe5a` aligned frontend autonomy and contract-parallel delivery.

Cross-cutting frontend readiness remains sufficient for real vertical product work.

## 7. Current development selection

Active product outcome is now:

```text
Slice 2 — Guest Donation + Truthful Donation State
```

Slice 1 remains finalized:

```text
SLICE_FINALIZED
```

Evidence orchestration Slice 2:

- `WU-S2-001 / EXP-S2-001-001` completed Stage 2 gap analysis and Stage 3 solutioning; the Stage 3 Human gate is recorded in its artifact provenance.
- `WU-S2-002` — Donation Domain & Contract Reconciliation — berstatus `ACTIVE / QUEUED` untuk rekonsiliasi Task 02 setelah O1 direction project-wide disetujui: major-unit decimal string + explicit currency code, exact-decimal calculation/persistence, active currency mengikuti Product, dan current Slice-2 whole-Rupiah IDR input. Campaign `NUMERIC(19,2)` hanya precedent; additional currencies, range, per-currency fraction rules, universal scale, dan migration detail belum diputuskan. O8 clear; O11 direction mengharuskan terminal notice terpenuhi melalui bounded/recoverable terminalization. TP-011 tetap Approved sampai direction baru dan batas `CONTRACT_READY` direkonsiliasi melalui jalur planning/review/approval yang berlaku. Review spec Task 01 tetap independen; evidence timing/recovery O2 dan security/parity O3/O4/O5 dirutekan ke delivery/Testing berikutnya. Belum ada contract atau milestone delivery yang earned.
- The Product/MVP amendment resolved Slice-2 product policy O1–O6/O9. Anhar superseded the independent pending-email cap route B: verified opt-in email may not be deleted before terminal notice is fulfilled, and Delivery must produce bounded, recoverable terminalization so the email lifecycle remains finite. No numeric bound, architecture, timeout-as-failed meaning, or residual-risk acceptance was selected. O3–O5 contract/control detail and evidence remain open. O1 representation direction is now approved in the project-wide monetary data standard; Task 02 API translation and O2 simulator details remain. OIR-006 has resolved ordering decision D1, now reflected in Task 01 domain specs and awaiting independent review. O8 is cleared for the current replacement based on Human/API owner confirmation that the historical operations were never externally distributed. O9 and O7 wording are reflected in the domain-spec diff; API translation remains. No contract or implementation milestone has been earned.
- Anhar confirmed current-Slice-2 UX disclosure near the optional guest-email choice: verify within 24 hours from email capture or the unverified address is deleted without a status email. O7 Design review selected the exact near-opt-in label/helper and source-first terminal status/notice labels; it does not settle verified-email retention while Donation remains `pending`.
- Pilot #2 bounded readiness assessment is at `.harscode-spaces/s2-guest-donation-truthful-state/readiness-reconciliation.md`; the Authority Map names Anhar Solehudin for current-Slice-2 areas plus project-wide, reassignable ownership for the Shared Currency Standard effective 2026-09-30. O1 owner/scope and representation direction are complete; concrete currency precision/range/fraction/storage parameters remain intentionally deferred. The five baseline Participant Profiles and Registry are present. OIR-S2-002-005 is reconciled; OIR-S2-002-006 completed and recorded D1. TP-011 is Approved and still current-effective pending authorized reconciliation. RV-006 completed cleanly and TP-012 generated the report. Human accepted the TPD-001 Task 01 → Task 02 split; BLD-S2-002-001 and its review/patch loop completed, with RV-S2-002-008 confirming F-01 closed; Task 02 is the active queued frontier.
- Backend/frontend delivery Work Units remain underived until the shared Slice-2 contract is reconciled.

Slice-1 completion evidence:

- `WU-S1-001 / EXP-001` established current authority and delivery gaps;
- `WU-S1-002 / TP-001` reconciled the public Campaign contract and was Human-approved;
- independent planning review/resolution closed the public closed-object gap;
- Build/Code Review/targeted patch cycles closed the TypeScript verification and controlled-media URL findings;
- `TST-001` independently verified R1–R13 at the artifact/contract/generated-contract boundary;
- Human Authority accepted `CONTRACT_READY` on 2026-09-22.

Current boundary:

```text
CONTRACT_READY
✓ reconciled

BACKEND_VERIFIED
✓ earned

FRONTEND_MOCK_VERIFIED
✓ earned

TOPOLOGY_VERIFIED
✓ earned

INTEGRATED_VERIFIED
✓ earned

SLICE_FINALIZED
✓ Human-approved
```

Slice 1 has completed technical, integrated, and Human finalization. Slice 2 is in progress: Exploration, current-frontier owner attribution, focused O3–O5 decisions, O2/O3 gap analysis, bounded Planner delivery proposal, O7 Design review, O1 evidence gathering and project-wide direction approval, OIR-006 owner ordering decision, TP-011 resolution and approval, clean RV-006 review, TP-012 report generation, and TPD-S2-002-001 decomposition are complete. O8 compatibility gate is cleared by Human/API owner confirmation that the historical operations have never been distributed externally. O11 direction is resolved: route B is superseded and verified email must remain eligible through terminal notice under a bounded, recoverable terminalization policy; implementation details and evidence remain open. O1 direction is approved: major-unit decimal-string + explicit currency code, exact-decimal calculation/persistence, Product-controlled active currency, and Slice-2 whole-Rupiah IDR input. Additional currencies, numeric range, per-currency fraction precision, universal storage scale, and migration details remain open until supported currencies and computation needs are established. Human accepted the task split; Task 01's Build/Patch and Review loop is complete; Task 02 is the active queued frontier, with TP-011 reconciliation required through the applicable planning/review/approval route before `CONTRACT_READY`.

## 8. Continuous Real-Task Validation posture

Harscode authority:

```text
main@b64fa11082a094d0e1b6e9488c20eac1c7f9777b
Operational Default / Level 2
```

The historical `workflow-v2` branch is experimental lineage, not authority for new Kencleng work.

CRTV resumes through ordinary MVP development after Product Authority promotion is merged.

Critical CRTV rule:

> Use the canonical Harscode Exploration kickoff with normal variables/context. Do not add a custom task-specific solution-steering prompt that tells the agent which authorities, gaps, or conclusions it should discover.

If an agent cannot discover information that the repository should make derivable, record that as validation evidence rather than rescuing the run with hidden conclusions.

Session-level telemetry should preserve both efficiency and correctness evidence.

## 9. Current delivery gate

The Product Authority promotion gate is closed. PR #27 is historical promotion evidence, not an active blocker.

Current gate:

```text
Exploration Slice 2, atribusi owner saat ini, persiapan keputusan O3–O5, gap analysis O2/O3, proposal Planner terbatas, review Design O7, evidence O1, D1 OIR-006, amendmen TP-011, independent review, report Planner, dan decomposition yang diterima selesai
→ Build/Patch Task 01 `BLD-S2-002-001` / `BLD-S2-002-002` selesai; Review loop `RV-S2-002-007` / `RV-S2-002-008` menutup F-01
→ Build Task 02 `BLD-S2-002-003` selesai tanpa diff contract; independent Code Review N/A karena tidak ada diff
→ O8 clear setelah konfirmasi owner; O11 direkonsiliasi: route B superseded dan terminalization harus bounded/recoverable tanpa menetapkan angka/architecture; Anhar mengambil owner project-wide O1
→ O1 direction disetujui project-wide: major-unit decimal string + explicit currency code, exact-decimal calculation/persistence, Product-controlled currency, current Slice-2 whole-Rupiah IDR; numeric/range/fraction/scale/migration details remain open
→ lanjutkan Task 02 reconciliation dengan mempropagasikan O1/O11 dan semantik O2/O4/O5 yang settled; tidak perlu menunggu runtime evidence
→ selesaikan source spec Donation/Campaign dan authored API contract yang terdampak setelah keputusan authority material tersedia; pertahankan timing/recovery O2 serta kontrol/parity empiris O3/O4/O5 sebagai kewajiban delivery/Testing downstream
→ rekonsiliasi wording acceptance `CONTRACT_READY` TP-011 melalui amendmen/review/approval berwenang agar tidak mensyaratkan evidence yang hanya bisa diperoleh setelah implementasi
→ Planner merekonsiliasi O1/O11 serta batas `CONTRACT_READY` dalam TP-011 melalui amendmen/review/approval yang berlaku; TP-011 tetap current-effective Approved sampai rekonsiliasi selesai
→ selesaikan field contract yang masih bergantung pada keputusan tersebut, lalu rutekan review owner yang diwajibkan
→ rekonsiliasi detail domain Donation dan shared contract (`WU-S2-002`)
→ raih Slice-2 `CONTRACT_READY` dari artifact contract/spec yang sudah direkonsiliasi dan diterima owner, tanpa menunggu runtime proof downstream
→ turunkan Work Unit backend/frontend dari contract yang direkonsiliasi
→ raih milestone verifikasi implementasi/security independen yang berlaku dengan evidence O2–O5
→ integrasikan terhadap contract yang sama
→ selesaikan acceptance Human/Product Slice 2 sebelum `SLICE_FINALIZED`
```

The current-effective approved Techplan distinguishes settled Product/MVP policy from remaining amount representation, simulator timing, guest email/security, status-access/response parity, Design review, and Campaign/Donation ordering detail. Route each item to its named owner or applicable specialist evidence without promoting historical draft values. Do not infer runtime completion from contract readiness.

## 10. Update discipline

Update this file when project-level state materially changes.

For each update:

- point to concrete evidence when available;
- distinguish operator-reported verification from checks actually executed by an agent/tool;
- avoid copying detailed acceptance criteria from feature specs;
- avoid copying Harscode phase reports;
- keep blockers/provisional dependencies visible;
- do not mark work complete because implementation merely exists;
- keep structural cross-stack mappings in `kencleng-integration-map.md` instead of duplicating them here;
- when historical status cannot be proven, use `NEEDS_RECONCILIATION`.
