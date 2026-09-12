# Kencleng — Development Tracker

> File: `docs/project/kencleng-development-tracker.md`
>
> Status: Living project status
>
> Last reconciled: 2026-09-12 against main checkpoint `3f5de5e`
>
> Purpose: Keep cross-domain backend/frontend/integration status visible without putting dated progress into workflow policy.

## 1. What this tracker owns

This file owns **current project delivery state** across domains and cross-cutting frontend/backend work.

It does not replace:

- `docs/spec/<domain-dir>/tasks.md` for domain-scoped task definition;
- feature specs for acceptance criteria;
- Harscode phase artifacts for one task/session;
- `docs/kencleng-agentic-workflow.md` for project orchestration policy.

When this tracker disagrees with an older domain task status, reconcile the evidence instead of silently choosing one. Git history, merged code, tests, and current specs are evidence; commit-message wording alone is not proof of verification.

## 2. Status vocabulary

Ordinary progress labels:

```text
NOT_STARTED
IN_PROGRESS
BLOCKED
NEEDS_RECONCILIATION
```

Evidence-backed project milestones:

```text
CONTRACT_READY
BACKEND_VERIFIED
FRONTEND_MOCK_VERIFIED
INTEGRATED_VERIFIED
DOMAIN_FINALIZED
DELIVERED
```

Do not promote a row to an evidence milestone without the evidence that milestone requires.

In particular:

- code existing on `main` does not automatically mean `BACKEND_VERIFIED`;
- frontend code/tests existing does not automatically mean `FRONTEND_MOCK_VERIFIED` unless the scoped mock/rendered verification is known;
- a commit touching FE + BE does not automatically mean `INTEGRATED_VERIFIED`;
- domain completion requires an explicit finalization/integration sweep.

Design readiness (`READY` / `PARTIAL` / `OPEN`) is classified per material UI work item, not permanently per domain.

## 3. Current domain snapshot

| Domain | Contract/spec state | Backend | Frontend | Integration | Current notes / evidence |
|---|---|---|---|---|---|
| Account | `NEEDS_RECONCILIATION` | `IN_PROGRESS` | `IN_PROGRESS` | `NEEDS_RECONCILIATION` | Main history contains implementation commits for backend tasks 01–07 and exploration for task 08; frontend commits cover landing plus account tasks 01–06. The older account task tracker predates several of those commits. Do not claim `INTEGRATED_VERIFIED` until real-stack integration evidence is reconciled. |
| Notification | Draft specs present | `NOT_STARTED` as standalone domain delivery | `NOT_STARTED` | `NOT_STARTED` | Notification contracts exist and later domains depend on them. Account may contain call sites/infrastructure, but that does not equal standalone notification-domain verification. |
| Organization | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Start only after domain preflight and dependency review. |
| Campaign | Draft specs present | `NOT_STARTED` | `NOT_STARTED` as domain delivery | `NOT_STARTED` | Frontend campaign/reference components may exist from scaffold/design work; folder presence is not counted as delivered campaign functionality. |
| Donation | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | Correctness-critical money/ledger work includes Tier-0 fenced areas. |
| Disbursement | Draft specs present | `NOT_STARTED` | `NOT_STARTED` | `NOT_STARTED` | State-machine core includes Tier-0 fenced areas. |

## 4. Account reconciliation evidence

Mainline history after the older `docs/spec/1-account/tasks.md` status snapshot includes at least:

| Scope | Main commit evidence |
|---|---|
| Account backend task 01 | `14834e5` — finalize register/email verification |
| Account backend task 02 | `efc1111` → `ce61841` — Google OAuth exploration/build/review/testing sequence |
| Account backend task 03 + frontend tasks 01–02 | `16a4bf9` |
| Account backend task 04 + frontend tasks 03–04 | `ea7d5bc` |
| Account backend/frontend task 05 | `6f036c6` |
| Account backend/frontend task 06 | `50b9a18` |
| Account backend task 07 | `6a846bd` |
| Account backend task 08 | `0798c5d` — exploration only at the `3f5de5e` checkpoint |

These entries prove that the old domain status table is stale. They do **not** by themselves prove every Harscode verification/finalization requirement was satisfied; that is why the domain remains `IN_PROGRESS` / `NEEDS_RECONCILIATION` here rather than being promoted automatically.

The account task/spec documents intentionally remain unchanged in this cleanup PR: changing domain contract/status files deserves its own evidence-backed reconciliation rather than being hidden inside instruction compaction.

**Remaining reconciliation work:** before relying on account `tasks.md` as current status, perform a dedicated evidence sweep of tasks 03–08 and update only status/path/reference metadata that can be proven. Do not bundle that work into unrelated frontend runtime changes.

## 5. Cross-cutting frontend readiness

| Capability | Status | Notes |
|---|---|---|
| Frontend scaffold / Next.js runtime | Present | Current frontend package/scripts are authoritative. |
| Product Design Principles | Present | `docs/ui-ux/product-design-principles.md` |
| UX Pattern System | Present | `docs/ui-ux/patterns.md` |
| Visual Design Guidelines v2 | Present | `docs/ui-ux/design-guidelines.md` |
| Brand & Visual Asset System | Present | `docs/ui-ux/brand-and-visual-assets.md` |
| Living Component System | Present | `frontend/components/README.md` |
| Prototype authority/usage v2 | Present | `docs/ui-ux/prototype-reference.md`, `design-reference-usage.md` |
| Frontend architecture v2 | Present | `docs/project/kencleng-frontend-tech-stack.md` |
| Authority cleanup after `3f5de5e` | In review | Current PR compacts always-loaded instructions, restores fencing, creates this tracker, and removes stale competing guidance. |
| Browser automation capability | Planned | Playwright selected, not yet wired; do not claim Playwright runs. |
| Button Secondary v2 runtime adoption | Planned | Documentation targets neutral/outlined Secondary; runtime migration is a separate shared-component change. |
| Codex harness optimization | Planned outside Kencleng | Harscode handoff will be staged only after Kencleng project truth/runtime feedback loop are clean. |

## 6. Current blockers before Codex frontend dogfood

The intended sequence is:

```text
post-migration authority cleanup
→ Button.secondary shared-component dogfood
→ Playwright browser-verification capability
→ Harscode Codex harness translation
→ first representative Codex frontend feature
```

Once the current authority-cleanup PR is merged, the next implementation work is the Button shared-component migration followed by Playwright wiring.

Before declaring the Codex frontend workflow operationally ready, verify that:

- always-loaded instructions are compact and non-conflicting;
- current design/component authorities are discoverable;
- shared-component change governance has been exercised on a real change;
- rendered/browser verification is executable;
- Codex translation remains harness translation rather than project policy duplication.

## 7. Update discipline

Update this file when a project-level state materially changes.

For each update:

- point to concrete evidence when available;
- avoid copying detailed acceptance criteria from feature specs;
- avoid copying Harscode phase reports;
- keep blockers and provisional dependencies visible;
- do not mark work complete because implementation merely exists.

If a historical status cannot be established confidently, use `NEEDS_RECONCILIATION` rather than guessing.