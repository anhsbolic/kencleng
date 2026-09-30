# Review Findings — `RV-S2-002-007`

> Phase: Code Review  
> Author: Codex Reviewer  
> Created: 2026-09-30  
> Model / reasoning: `gpt-6-luna` / `high` (dispatch configuration; runtime selection not independently exposed)  
> Session: Fresh independent Reviewer Session; Session ID not exposed  
> Target revision: `6891341a050982e14174ab5af132a200f24e71d9` plus current Task 01 source diff and Run artifacts  
> Workflow revision: `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`

## Scope reviewed

Actual Task 01 source diff only:

- `docs/spec/5-donation/invariants.md`
- `docs/spec/5-donation/threat-model.md`
- `docs/spec/5-donation/tasks.md`
- `docs/spec/5-donation/features/01-submit-donation-settlement.md`
- `docs/spec/5-donation/features/02-donation-status-check.md`
- `docs/spec/5-donation/features/03-public-donor-list.md` through `06-guest-email-reveal.md`
- `docs/spec/4-campaign/invariants.md`
- `docs/spec/4-campaign/features/09-closure.md`

Compared against Approved `TP-S2-002-011`, the Human-accepted Task 01, the task manifest, current repo instructions, and canonical Review guidance. Build report was context only; the changed source files and diff were inspected directly.

## 1. Safety

**Finding F-01 — Blocking.**

- **Location:** `docs/spec/4-campaign/features/09-closure.md:22-26`
- **Problem:** The Summary says the threshold, deadline, and force-close triggers all share the same `WHERE status = 'published'` idempotency guard and attributes that mechanism to `INV-campaign-13`. The reconciled invariant at `docs/spec/4-campaign/invariants.md:210-223` explicitly leaves the transaction/locking mechanism unselected, and the approved Techplan D1/R7 likewise forbids selecting one here. The new introductory cross-reference also limits this edit to the narrow D1 boundary; it does not reconcile the broader closure mechanism.
- **Impact:** Readers may treat a specific SQL guard as an approved concurrency design and assume it proves the winning close reason/order. That misstates the Tier-1 contract and could constrain or misdirect later Campaign delivery and verification.
- **Suggested resolution:** In a new Build/Patch Run, remove the shared-guard mechanism assertion from the Summary. Keep the trigger list only as historical/current feature context, and point D1 readers to `INV-campaign-13` while stating that the applicable ordering mechanism remains unselected and belongs to the authorized Campaign closure delivery work.
- **Classification:** Blocking; the active D1 cross-reference currently coexists with a contradictory implementation-mechanism claim.

No other safety findings. The Donation invariants preserve atomic exact-once success/funding, backend-owned terminal state, request idempotency separately from settlement replay, and the accepted-pending-after-close rule. No mechanism was inferred from these requirements.

## 2. Quality

No findings. The active Donation criteria and task split are readable and retain the required Open Item gates. Features 03–06 are clearly marked deferred, and the historical behavior is visibly separated from active Slice 2 acceptance.

## 3. Stack-Specific Best Practices

Reviewed the routed matching guidance:

- `best-practices/go/decimal-and-money.md`
- `best-practices/restapi/idempotency-and-versioning.md`
- `best-practices/postgresql/financial-invariant-enforcement.md`

No findings. The spec requires exact decimal/no `float64`, keeps request retry idempotency distinct from settlement replay, and leaves database transaction/locking design for its authorized owner and later runtime evidence. This documentation-only diff does not establish runtime compliance with those practices.

## 4. Consistency

**Finding F-01 — Blocking; same contradiction as Safety pass.**

- **Location:** `docs/spec/4-campaign/features/09-closure.md:22-26`
- **Target-repo authority:** `docs/spec/4-campaign/invariants.md#inv-campaign-13` says the D1 reference does not select a transaction/locking mechanism. The accepted Task 01 says not to select a mechanism or reconcile the broader Slice 3 closure lifecycle. `docs/spec/README.md` §§1–2 makes reconciled domain specs subordinate to Product/MVP authority and requires ambiguity to remain explicit.
- **Problem / impact / suggested resolution:** As described in F-01 under Safety; the Summary retains a lower-level historical mechanism claim that contradicts the approved D1 scope.
- **Classification:** Blocking.

No additional consistency findings. The Donation artifacts use domain-first paths, remain `draft` pending owner/Human review, preserve O1/O11 and O2–O5/O8 boundaries, propagate the exact Q11 copy, and keep deferred features outside active acceptance. No residual Security/PII risk is marked accepted.

## Verification executed during Review

None. Review used read-only inspection of the scoped diff and authority files. No runtime reproduction, automated test, or broad suite was needed or run; this document review provides no runtime/testing evidence.

## Verdict

**Request changes.** F-01 is blocking. No other patch loop is recommended.

## Phase handoff

- **Completed:** Independent four-pass review of the actual Task 01 diff; verdict is Request changes.
- **Artifacts:** `review-findings.md`, `patch-plan.md`, and `launch-record.md` in `RV-S2-002-007`.
- **Human decision:** None needed to route the scoped correction to Build. Applicable Campaign/Donation, Security/PII, API, and Human/domain-owner review remains required before affected specs may become `agreed`.
- **Open / deferred:** Blocking F-01; O1–O5, conditional O8, and O11 remain as recorded in the Approved spine. Runtime Testing remains downstream.
- **Recommended next step:** Start a new Build/Patch Run with `patch-plan.md`, correcting only the stale Campaign Summary claim. Then route the corrected scope through applicable owner/Human review; do not claim `CONTRACT_READY` from this Review.
- **Session transition:** Start a fresh Build/Patch Participant Session in a new Run, re-grounded on this finding and patch plan, as required for orchestrated re-entry.
- **Context pointers:** `TP-S2-002-011/techplan.md`; accepted `TPD-S2-002-001/tasks/01-donation-domain-spec-reconciliation.md`; F-01 at `docs/spec/4-campaign/features/09-closure.md:22-26`.
