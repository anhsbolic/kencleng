# Kencleng — Development Tracker

> Status: Living project status
> Last reconciled: 2026-09-16
> Current main before reboot merge: `cc6552e5ae36e354388f5d9d3230672929b11055`
> Frontend reboot branch: `frontend-reboot-preparation`
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
READY_TO_START
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

Code existing in history is not sufficient evidence for a verified milestone.

## 3. Current domain snapshot

The frontend reboot intentionally retires the previous frontend product implementation. Historical frontend commits remain evidence of what was previously built, but they do **not** represent active completion state for the new frontend generation.

| Domain | Contract/spec state | Backend | Frontend | Integration | Current notes |
|---|---|---|---|---|---|
| Account | `NEEDS_RECONCILIATION` | `IN_PROGRESS` | `NOT_STARTED` for the new frontend generation | `NEEDS_RECONCILIATION` | Backend/account history remains. Retired frontend account implementation is archived in Git and is not active delivery. |
| Notification | Draft specs present | `NOT_STARTED` as standalone domain delivery | `NOT_STARTED` | `NOT_STARTED` | Historical incidental infrastructure does not equal standalone domain delivery. |
| Organization | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Future frontend work starts from the new reboot baseline and still requires normal domain preflight. |
| Campaign | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Retired landing/campaign components do not count as new-generation Campaign delivery. |
| Donation | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Correctness-critical money/ledger work includes Tier-0 fenced areas. |
| Disbursement | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | State-machine core includes Tier-0 fenced areas. |

## 4. Historical account/backend reconciliation evidence

Historical backend/account evidence remains useful and is not erased by the frontend reboot.

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

These entries prove historical work happened. They do not automatically prove current completion/verification. Backend/account status reconciliation remains a separate evidence task.

## 5. Current UI/UX authority readiness

The upstream Product Brand + UI Design Exploration is complete and human-approved.

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
READY FOR MERGE
```

Authority:

```text
docs/project/frontend-reboot-plan.md
docs/project/frontend-reboot-deletion-manifest.md
```

The reboot has:

- reconciled project/spec/frontend documentation with the new design generation;
- classified retained engineering scaffold versus retired implementation;
- removed previous product/UI implementation and old reusable-component contracts from the active branch tree;
- removed old active `.local-agents/works/**` while preserving Git history as archive;
- kept `.local-agents/` tracked for future learning-by-doing process evidence;
- reduced the frontend to a minimal technical scaffold;
- synchronized the dependency lockfile;
- received operator-reported successful verification of dependency install, `npm run verify`, `npm run build`, and minimal boot/render inspection.

ChatGPT verified the resulting Git tree and lockfile synchronization from GitHub but did not execute the local verification commands itself.

The reboot is project preparation/maintenance, **not** Frontend Experience Foundation implementation and not a Harscode feature-development run.

## 7. Cross-cutting frontend readiness

| Capability / authority | Status | Notes |
|---|---|---|
| Canonical product-design direction | READY | Sunlit Editorial / Evidence-Led Optimism approved. |
| Concrete visual system | READY | Newsreader + Instrument Sans; warm-paper/Sun/Berry; semantic colors separate; border/spacing-first; restrained radius/elevation; Phosphor utility baseline; provenance/progress grammar. |
| Selected direction visual evidence | READY | Public composition + Evidence Journal PNGs committed. |
| Frontend architecture v3 | READY | Clean-start architecture in `docs/project/kencleng-frontend-tech-stack.md`. |
| Component governance | READY | Semantic-owner-first; new-generation reusable registry intentionally starts empty. |
| Next.js/TypeScript/Tailwind tooling | OPERATOR VERIFIED | Minimal scaffold verification reported successful on 2026-09-16. |
| Vitest/RTL/MSW capability | OPERATOR VERIFIED | Test configuration initialized through the requested verification run; old product tests/fixtures remain retired. |
| Playwright capability | RETAINED | On-demand browser automation capability retained; no old product browser regression is carried forward by inertia. |
| Previous frontend product implementation | RETIRED | Git history is archive; active branch tree no longer treats it as precedent. |
| Previous `.local-agents/works/**` history | RETIRED FROM ACTIVE TREE | Git history remains archive; new runs will be committed fresh. |
| Frontend Experience Foundation Task 01 | READY_TO_START AFTER MERGE | First real frontend development task from the clean baseline. |

## 8. Frontend readiness gate before new development

The detailed gate is green in `docs/project/frontend-reboot-plan.md`.

Remaining repository transition:

```text
merge frontend-reboot-preparation
→ record exact merged main commit as Frontend Reboot Baseline
→ begin CRTV telemetry before first Exploration session
→ start Frontend Experience Foundation through Harscode
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

The next CRTV measurement begins only from the merged clean frontend baseline. Preparation/reset work before that baseline is not part of the feature validation benchmark.

Session-level telemetry should preserve both:

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
- distinguish operator-reported verification from checks actually executed by an agent/tool;
- avoid copying detailed acceptance criteria from feature specs;
- avoid copying Harscode phase reports;
- keep blockers/provisional dependencies visible;
- do not mark work complete because implementation merely exists;
- when historical status cannot be proven, use `NEEDS_RECONCILIATION`.
