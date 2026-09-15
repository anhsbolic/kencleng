# Kencleng — Development Tracker

> File: `docs/project/kencleng-development-tracker.md`
>
> Status: Living project status
>
> Last reconciled: 2026-09-15 against Kencleng main `ec789070` and Harscode `workflow-v2` validation baseline `fce5721d`
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

The account task/spec documents intentionally remain unchanged in this cleanup sequence: changing domain contract/status files deserves its own evidence-backed reconciliation rather than being hidden inside cross-cutting frontend work.

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
| Authority cleanup after `3f5de5e` | Present | Merged in PR #1 / `bdba6b0`; compact authority routing, Tier-0 fencing, living tracker, and stale-guidance cleanup are on main. |
| Button Secondary v2 runtime adoption | Present | Merged in PR #2 / `749807a`; `Button.secondary` is neutral/outlined and covered by a targeted primitive regression test. |
| Browser automation capability | Proven | PR #3 / `b07a256`; Playwright is wired as an explicit Chromium real-browser capability, `npm run verify` passed 40 files / 226 tests, and the deterministic `/login` browser smoke passed 1/1 from a real checkout. It intentionally remains outside the fast baseline. |
| Codex harness optimization | Present in Harscode | Harscode proposal `0027` is Accepted and merged on Harscode main at `5f3c9bc`; `harness-optimization/codex/` now defines the translation-only Codex layer. |
| Continuous Real-Task Validation #1 | NEXT | Frontend Experience Foundation / brand calibration using a representative `/` slice. The goal is to validate enough public-facing brand, shell, CTA hierarchy, typography/spacing/color, and responsive behavior to calibrate subsequent frontend work — not to finish the landing page or resolve Campaign product semantics. |
| Continuous Real-Task Validation #2 | PLANNED | Organization Registration at `/dashboard/organization/new` remains the next representative feature vertical slice unless Validation #1 produces evidence that changes the sequencing decision. |

## 6. Readiness before frontend Continuous Real-Task Validation

The staged prerequisite sequence is complete:

```text
post-migration authority cleanup         ✓ merged
→ Button.secondary shared-component      ✓ merged
→ Playwright browser verification        ✓ proven and merged
→ Harscode Codex harness translation     ✓ accepted and merged
→ workflow-v2 audit/remediation gate     ✓ verified
→ frontend experience foundation run     NEXT
```

The frontend is ready to start Continuous Real-Task Validation **as real project work, not as proof that the workflow or project guidance is already optimal**.

The first run is intentionally foundation-oriented:

```text
representative surface: /
primary intent:          public brand / experience calibration
scope shape:             enough shell + hero + representative content/CTA surface
not required:            complete landing page, final Campaign semantics, or live Campaign backend integration
```

The second planned run is Organization Registration, which exercises a different concern set: contract-driven form behavior, multipart uploads, API errors, mocks, authenticated dashboard composition, and representative browser behavior.

During validation, classify problems by the layer that actually failed:

```text
project guidance / authority
Harscode workflow or best-practice
Codex harness translation
feature implementation bug
```

Do not respond to a feature bug by expanding global agent policy, and do not respond to a real reusable harness gap by patching only the one feature's local instructions.

The first validation task should be representative enough to exercise current frontend product/design/component authorities and rendered verification while avoiding a Tier-0/Tier-1 security or money-critical surface as the first workflow experiment.

Readiness evidence now includes:

- compact hierarchical Kencleng `AGENTS.md` routing is on `main`;
- current design/component authorities are discoverable;
- shared-component blast-radius governance was exercised in B1;
- the real-browser Playwright capability was executed successfully in B2;
- Harscode Codex proposal `0027` was human-approved and merged without creating a second project policy source;
- workflow-v2 remediation findings were independently re-verified before the validation baseline was selected.

## 7. Update discipline

Update this file when a project-level state materially changes.

For each update:

- point to concrete evidence when available;
- avoid copying detailed acceptance criteria from feature specs;
- avoid copying Harscode phase reports;
- keep blockers and provisional dependencies visible;
- do not mark work complete because implementation merely exists.

If a historical status cannot be established confidently, use `NEEDS_RECONCILIATION` rather than guessing.
