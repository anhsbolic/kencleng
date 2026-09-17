# Kencleng — Development Tracker

> Status: Living project status
> Last reconciled: 2026-09-17
> Current Kencleng main baseline before Product Authority promotion: `main@15a3e02cc88d00e8dee70f8dfb07c36cb1e2fe5a`
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
| Whole-product Product Authority | `PROMOTION_READY` on PR #27 branch | `docs/product/product-overview.md` promoted in branch; human product review PASS. Becomes repository authority on merge. |
| Product Design / Brand Authority | `READY` | Canonical `docs/ui-ux/`; Sunlit Editorial / Evidence-Led Optimism approved. |
| MVP release scope | `APPROVED` | `docs/product/mvp-scope.md`; human approval 2026-09-17. |
| MVP delivery sequencing | `APPROVED` | `docs/product/mvp-delivery-slices.md`; human approval 2026-09-17. |
| Product Authority routing | `PROMOTION_READY` on PR #27 branch | Root/scoped AGENTS, spec routing, and orchestration aligned in the promotion branch. |
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
| Campaign | `NEEDS_RECONCILIATION` for Slice 1 | `NOT_STARTED` active generation | `NOT_STARTED` active product surface | First active MVP semantic area through Public Campaign Understanding. |
| Donation | Historical/draft reference | `NOT_STARTED` active generation | `NOT_STARTED` | Enters baseline in Slice 2; correctness-critical money areas retain Tier-0 fencing. |
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

## 7. Next development selection

The next product surface is now explicitly selected:

```text
Slice 1 — Public Campaign Understanding
```

This is the first real post-promotion CRTV candidate.

Its high-level outcome is defined in `docs/product/mvp-delivery-slices.md`: a skeptical-but-open visitor can inspect a persisted public Campaign and correctly understand campaign/steward/funding/source context without internal-data leakage or fake downstream actions.

Do **not** pre-reconcile the whole Campaign domain/OpenAPI before Exploration.

Post-promotion execution path:

```text
Product Authority + MVP Slice 1 reference
→ canonical Harscode Exploration kickoff
→ agent discovers relevant Design Authority + historical Campaign/Organization/API/code evidence
→ gap analysis / solutioning through Harscode
→ reconcile only Slice 1 delivery specs/contracts
→ FE/BE delivery
→ real integration evidence
```

The backend-first vs contract-parallel choice should emerge from Slice 1 Exploration/Techplan based on actual contract stability and dependency shape; it is not pre-decided here.

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

## 9. Current merge gate

PR #27 (`docs/product-authority-reframe`) is the Product Authority promotion vehicle.

Before real Slice 1 development starts:

```text
PR #27 promotion branch review
→ merge to main
→ confirm main promotion checkpoint
→ start Slice 1 CRTV from canonical Harscode kickoff
```

Do not start a competing Slice 1 contract rewrite on `main` while PR #27 remains unmerged.

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