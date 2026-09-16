# Kencleng — Development Tracker

> Status: Living project status
> Last reconciled: 2026-09-16
> Kencleng reference baseline: `main@cc6552e5ae36e354388f5d9d3230672929b11055`
> Harscode workflow-v2 candidate for the next frontend run: `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`
> Purpose: Keep cross-domain delivery state visible without turning dated progress into workflow policy.

## 1. What this tracker owns

This file owns current project delivery state across domains and cross-cutting frontend/backend work.

It does not replace:

- `docs/spec/<domain-dir>/tasks.md` for task definition;
- feature specs for acceptance criteria;
- UI/UX authorities for design truth;
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
REBOOT_PREPARATION
```

Evidence-backed milestones:

```text
CONTRACT_READY
BACKEND_VERIFIED
FRONTEND_MOCK_VERIFIED
INTEGRATED_VERIFIED
DOMAIN_FINALIZED
DELIVERED
```

Code existing on `main` is not sufficient evidence for a verified milestone.

## 3. Current domain snapshot

The frontend reboot intentionally retires the current frontend product implementation. Historical frontend commits remain evidence of what was previously built, but after the reboot they do **not** represent active production/frontend completion state.

| Domain | Contract/spec state | Backend | Frontend | Integration | Current notes |
|---|---|---|---|---|---|
| Account | `NEEDS_RECONCILIATION` | `IN_PROGRESS` | `NOT_STARTED` for the new frontend generation | `NEEDS_RECONCILIATION` | Backend/account history remains. Existing frontend account implementation is being intentionally retired during the frontend reboot and must not be counted as delivered in the new generation. |
| Notification | Draft specs present | `NOT_STARTED` as standalone domain delivery | `NOT_STARTED` | `NOT_STARTED` | Existing incidental frontend/backend infrastructure does not equal standalone domain delivery. |
| Organization | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Future frontend work begins only after reboot baseline and required domain preflight. |
| Campaign | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Retired landing/campaign presentation components do not count as new-generation Campaign delivery. |
| Donation | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Correctness-critical money/ledger work includes Tier-0 fenced areas. |
| Disbursement | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | State-machine core includes Tier-0 fenced areas. |

## 4. Historical account/backend reconciliation evidence

Historical mainline evidence remains useful for backend/account reconciliation and is not erased by the frontend reboot.

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

These entries prove historical work happened. They do not automatically prove current completion/verification, and old frontend portions will be retired from the active tree.

Backend/account status reconciliation remains a separate evidence task; do not hide it inside the frontend reboot.

## 5. Current UI/UX authority readiness

The upstream product-brand/UI design exploration is complete and human-approved.

Active authority:

- `docs/ui-ux/README.md` — routing map;
- `docs/ui-ux/brand-product-ui-brief.md` — Sunlit Editorial / Evidence-Led Optimism;
- `docs/ui-ux/product-design-principles.md` — stable design judgment;
- `docs/ui-ux/design-guidelines.md` — concrete visual system;
- `docs/ui-ux/patterns.md` — recurring interaction behavior;
- `docs/ui-ux/asset-governance.md` — asset lifecycle/truthfulness;
- `docs/ui-ux/page-map.md` — persona/surface intent;
- `docs/ui-ux/visual-references/selected-direction/` — approved direction evidence.

Removed legacy design/prototype generations are available only through Git history and are not current authority.

## 6. Frontend reboot status

Current state:

```text
FRONTEND REBOOT PREPARATION
```

Authority:

```text
docs/project/frontend-reboot-plan.md
```

Purpose:

- reconcile project/spec/frontend documentation with the new design generation;
- explicitly classify what engineering scaffold is retained;
- retire current product/UI implementation and old reusable-component contracts;
- retire old active `.local-agents/works/**` history while keeping Git history as archive;
- keep `.local-agents/` tracked for new learning-by-doing process evidence;
- verify a minimal clean frontend scaffold;
- freeze a new frontend baseline before any new Harscode development run.

The reboot itself is project preparation/maintenance, **not** Frontend Experience Foundation implementation and not a Harscode feature-development run.

## 7. Cross-cutting frontend readiness

| Capability / authority | Status | Notes |
|---|---|---|
| Canonical product-design direction | READY | Sunlit Editorial / Evidence-Led Optimism is approved. |
| Concrete visual system | READY | Newsreader + Instrument Sans; warm-paper/Sun/Berry; semantic colors separate; border/spacing-first; restrained radius/elevation; Phosphor utility baseline; provenance/progress grammar. |
| Selected direction visual evidence | READY | Public composition + Evidence Journal PNGs are committed. |
| Frontend architecture v3 | PREPARED | Clean-start architecture in `docs/project/kencleng-frontend-tech-stack.md`. |
| Component governance | PREPARED | Semantic-owner-first; new-generation reusable registry intentionally starts empty. |
| Next.js/TypeScript/Tailwind tooling | TO VERIFY AFTER RESET | Capability expected to be retained; reboot verification will prove scaffold health. |
| Vitest/RTL/MSW capability | TO VERIFY AFTER RESET | Retain tooling/capability, not old product tests/fixtures by inertia. |
| Playwright capability | TO VERIFY AFTER RESET | Retained as on-demand browser automation, not phase ritual. |
| Old frontend product implementation | RETIRING | Must not become precedent for the new generation. |
| Old `.local-agents/works/**` frontend history | RETIRING FROM ACTIVE TREE | Git history remains archive; new runs remain committed for learning. |
| Frontend Experience Foundation Task 01 | BLOCKED | Starts only after reboot ready gate is green and baseline frozen. |

## 8. Frontend readiness gate before new development

Follow the checklist in `docs/project/frontend-reboot-plan.md`.

High-level sequence:

```text
documentation reconciliation
→ deletion manifest
→ retire old frontend implementation/process tree
→ leave minimal clean scaffold
→ verify scaffold health
→ reconcile tracker/status
→ freeze Frontend Reboot Baseline
→ start real Harscode frontend development
```

The next real development task remains:

```text
Frontend Experience Foundation
→ representative `/` slice
→ first production expression of Sunlit Editorial
```

This is intentionally not a requirement to finish the landing page or recreate the retired `/` implementation.

## 9. Continuous Real-Task Validation posture

The previous frontend validation run remains historical evidence.

The next CRTV measurement should begin only after the new clean frontend baseline is frozen. Preparation/reset work before that baseline is not part of the feature validation benchmark.

For the next run, preserve apples-to-apples operator telemetry at the session level and evaluate both:

```text
efficiency
→ context/usage/turns/rescue prompts

correctness
→ authority adherence/outcome/review/testing findings
```

Do not optimize token usage at the expense of correctness.

## 10. Update discipline

Update this file when project-level state materially changes.

For each update:

- point to concrete evidence when available;
- avoid copying detailed acceptance criteria from feature specs;
- avoid copying Harscode phase reports;
- keep blockers/provisional dependencies visible;
- do not mark work complete because implementation merely exists;
- when a historical status cannot be proven, use `NEEDS_RECONCILIATION`.