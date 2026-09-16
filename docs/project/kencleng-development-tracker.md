# Kencleng — Development Tracker

> Status: Living project status
> Last reconciled: 2026-09-17
> Current Kencleng checkpoint: `main@fca5a8f3178b53e5eec006064d8bcf2b078771b3`
> Harscode operational baseline: `main@b64fa11082a094d0e1b6e9488c20eac1c7f9777b`
> Purpose: Keep cross-domain delivery state visible without turning dated progress into workflow policy.

## 1. What this tracker owns

This file owns current project delivery state across domains and cross-cutting frontend/backend work.

It does not replace:

- `docs/spec/<domain-dir>/tasks.md` for task definition;
- feature specs for acceptance criteria;
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
DOMAIN_FINALIZED
DELIVERED
```

Code existing in history is not sufficient evidence for a verified milestone.

## 3. Current domain snapshot

The frontend reboot intentionally retired the previous frontend product implementation. Historical frontend commits remain evidence of what was previously built, but they do **not** represent active completion state for the new frontend generation.

| Domain | Contract/spec state | Backend | Frontend | Integration | Current notes |
|---|---|---|---|---|---|
| Account | `NEEDS_RECONCILIATION` | `IN_PROGRESS` | `NOT_STARTED` for the new frontend generation | `NEEDS_RECONCILIATION` | Backend/account history remains. Retired frontend account implementation is archived in Git and is not active delivery. |
| Notification | Draft specs present | `NOT_STARTED` as standalone domain delivery | `NOT_STARTED` | `NOT_STARTED` | Historical incidental infrastructure does not equal standalone domain delivery. |
| Organization | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Future frontend work starts from the new frontend foundation and still requires normal domain preflight. |
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
- `docs/ui-ux/product-design-principles.md` — stable design judgment and frontend design-readiness/autonomy boundary;
- `docs/ui-ux/design-guidelines.md` — concrete visual system;
- `docs/ui-ux/patterns.md` — recurring interaction behavior;
- `docs/ui-ux/asset-governance.md` — asset lifecycle, truthfulness, recommendation/generation flow;
- `docs/ui-ux/page-map.md` — persona/surface intent;
- `docs/ui-ux/visual-references/selected-direction/` — approved direction evidence.

Removed legacy design/prototype generations are available only through Git history and are not current authority.

High-fidelity UI is not a universal prerequisite for frontend implementation. Frontend work proceeds from sufficient canonical product/design authority; material ambiguity is resolved through low-fidelity alternatives/recommendation, while missing product/domain truth is routed back to its owning authority.

## 6. Frontend reboot and foundation status

Frontend reboot:

```text
COMPLETE
```

Reboot authority:

```text
docs/project/frontend-reboot-plan.md
docs/project/frontend-reboot-deletion-manifest.md
```

Reboot implementation checkpoint:

```text
main@6b5cdd3d4710a64be0d00ce4479bad118fd6cf48
```

The reboot:

- reconciled project/spec/frontend documentation with the new design generation;
- classified retained engineering scaffold versus retired implementation;
- removed previous product/UI implementation and old reusable-component contracts from the active tree;
- removed old active `.local-agents/works/**` while preserving Git history as archive;
- kept `.local-agents/` tracked for future learning-by-doing process evidence;
- reduced the frontend to a minimal technical scaffold;
- synchronized the dependency lockfile;
- received operator-reported successful verification of dependency install, `npm run verify`, `npm run build`, and minimal boot/render inspection;
- was squash-merged through PR #20.

The reboot was project preparation/maintenance, not Frontend Experience Foundation implementation and not a Harscode feature-development run.

Frontend Experience Foundation Task 01:

```text
DELIVERED
```

Evidence:

- PR #24 / `71093b687cd7135495bc6ed62d520a621f96f586` implemented the representative `/` calibration surface and received human rendered acceptance PASS;
- PR #25 / `fca5a8f3178b53e5eec006064d8bcf2b078771b3` closed Validation 02.

## 7. Cross-cutting frontend readiness

| Capability / authority | Status | Notes |
|---|---|---|
| Canonical product-design direction | READY | Sunlit Editorial / Evidence-Led Optimism approved. |
| Concrete visual system | READY | Newsreader + Instrument Sans; warm-paper/Sun/Berry; semantic colors separate; border/spacing-first; restrained radius/elevation; Phosphor utility baseline; provenance/progress grammar. |
| Frontend design autonomy | READY | High-fidelity design files are optional; clear intent may be implemented directly, material ambiguity uses low-fi alternatives, and product-truth gaps must be resolved at authority. |
| Asset recommendation/generation governance | READY | Canonical reuse first; agent may recommend, brief, prompt, or generate candidates when justified; synthetic assets may not masquerade as evidence. |
| Selected direction visual evidence | READY | Public composition + Evidence Journal PNGs committed. |
| Frontend architecture v3 | READY | Clean-start architecture in `docs/project/kencleng-frontend-tech-stack.md`. |
| Contract-parallel coordination map | READY | `docs/project/kencleng-integration-map.md` owns structural surface→capability→contract→backend-owner mapping without duplicating delivery status. |
| Component governance | READY | Semantic-owner-first; new-generation reusable registry intentionally starts empty. |
| Next.js/TypeScript/Tailwind tooling | OPERATOR VERIFIED | Minimal scaffold verification reported successful on 2026-09-16. |
| Vitest/RTL/MSW capability | OPERATOR VERIFIED | Test configuration initialized through the requested verification run; old product tests/fixtures remain retired. |
| Playwright capability | RETAINED | On-demand browser automation capability retained; no old product browser regression is carried forward by inertia. |
| Previous frontend product implementation | RETIRED | Git history is archive; active tree no longer treats it as precedent. |
| Previous `.local-agents/works/**` history | RETIRED FROM ACTIVE TREE | Git history remains archive; new runs will be committed fresh. |
| Frontend Experience Foundation Task 01 | DELIVERED | PR #24 implemented and verified the foundation; PR #25 closed Validation 02. |

## 8. Next development selection

The project is no longer waiting for frontend foundation work.

The next frontend/backend sequence should select a **real product surface/capability** and determine whether backend-first or contract-parallel delivery is appropriate.

Preferred CRTV candidate characteristics:

- meaningful page/flow rather than another foundation-only task;
- stable-enough product/API contract or a small clearly resolvable contract gap;
- useful frontend presentation behavior and states;
- suitable for MSW-backed frontend progress before the real backend is available when contract-parallel is chosen;
- not unnecessarily combining first-time contract-parallel validation with the highest-risk money/security core.

Before implementation of the selected cross-stack surface:

```text
select page/flow
→ add only the actionable mapping to docs/project/kencleng-integration-map.md
→ confirm product/design readiness
→ confirm API contract readiness
→ choose backend-first or contract-parallel sequencing
→ run the normal Harscode lifecycle for each coherent task
→ perform real integration verification when both sides are available
```

The exact next surface is intentionally not guessed in this tracker.

## 9. Continuous Real-Task Validation posture

Validation 02 is complete. The Harscode workflow-v2 candidate was subsequently promoted to the operational default on Harscode `main`.

Current Harscode authority:

```text
main@b64fa11082a094d0e1b6e9488c20eac1c7f9777b
Operational Default / Level 2
```

The historical `workflow-v2` branch is experimental lineage, not authority for new Kencleng work.

CRTV now continues through ordinary Kencleng development. Harscode process changes should be driven by evidence from real tasks rather than speculative workflow iteration.

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
- keep structural cross-stack mappings in `kencleng-integration-map.md` instead of duplicating them here;
- when historical status cannot be proven, use `NEEDS_RECONCILIATION`.
