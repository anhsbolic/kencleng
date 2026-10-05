# WU-S2-008 — Donation Retry Credential Source Reconciliation

## Definition

- Type: `RECONCILIATION`
- Parent Outcome: `S2-GUEST-DONATION-TRUTHFUL-STATE`
- Coordination owner: Orchestration Operator
- Planned first phase role: Explorer
- Communication language: Bahasa Indonesia

## Outcome

Reconcile the settled same-key/equivalent-request retry credential behavior into the owning Donation spec and authored API sources, affected generated/internal counterparts, and exact owner-acceptance evidence. This Work Unit does not implement backend behavior or close the separate O3/O4/O5 controls.

## Scope

- Current requirement and source gap: WU-S2-003 `techplan.candidate.md`, §13 Active item 9, exact current hash `e895a1da8b90e7f88c449651a9a46add59e1a1d610cce3cc7b739d0c12c30315`.
- Candidate authored sources: `docs/spec/5-donation/features/01-submit-donation-settlement.md`, `docs/spec/5-donation/features/02-donation-status-check.md`, `docs/spec/5-donation/invariants.md` (INV-donation-10 and INV-donation-05), and `api/openapi/donation.yaml` (Donation POST retry/response wording).
- Discover and reconcile affected counterparts from live source impact: at minimum consider `api/openapi.yaml`, `frontend/lib/api/generated/openapi.ts`, and affected internal frontend Donation consumer/flow fixtures or tests. Record exact affected and unaffected files with evidence; do not assume the list is exhaustive before Exploration.
- Preserve the Human-settled policy recorded in parent `events.md` (2026-10-04): exact retries return the original Donation and issue a new credential each time; every credential expires 24 hours after its own issuance; persist only a one-way verifier, never bearer material; multiple credentials may remain valid concurrently.
- Required acceptance: independent source Review and applicable counterpart verification, followed by Anhar's explicit acceptance of the exact hashes/bytes for every changed owning source and affected counterpart. Record the exact receipt and handoff to WU-S2-003.

## Boundaries

- This is a source/spec/API and internal/generated-counterpart reconciliation. WU-S2-003 owns backend delivery and implementation.
- Do not treat WU-S2-006's seven accepted source hashes or WU-S2-007's five accepted source/counterpart hashes as acceptance of the retry issuance rule. Those receipts remain limited to their recorded bytes and semantics.
- Do not change Product/MVP authority, migration/schema, backend/frontend production behavior, tests except contract-facing counterpart fixtures/assertions authorized by the eventual reviewed Techplan, or protected Tier-0 paths during Exploration/Techplan.
- The Human decision above is settled. Do not reopen or extend it. Numeric token strength/generation, key controls, comparison, expiry enforcement/cleanup, idempotency retention, D1 pairing, and runtime/security risk acceptance remain separate WU-S2-003 gates.
- No acceptance, implementation, migration, runtime readiness, or Slice completion is implied by planning or source reconciliation alone.

## Completion condition

Every affected owning source and internal/generated counterpart has an exact current hash, applicable independent Review/verification evidence, and explicit owner acceptance by Anhar against those exact revisions; a handoff to WU-S2-003 identifies the accepted bytes and remaining independent gates. WU-S2-008 then becomes `DONE`; no delivery milestone is earned.

## Current State

- Execution status: `ACTIVE`
- Scheduling state: `ACTIVE / PARKED` — Human HOLD of Slice 2 development, 2026-10-05. `EXP-S2-008-001` remains prepared and undispatched; it must not be dispatched during HOLD.
- Immediate route: No development action during HOLD. On explicit Human resume, re-ground the prepared canonical Exploration Invocation and determine phase applicability from current evidence before dispatch.
- Dependencies: HARD on WU-S2-002 accepted Donation baseline and settled Human direction; WU-S2-007 is a read-only exact-current API/counterpart baseline. WU-S2-006 is historical evidence only and does not accept this behavior.
- Downstream: WU-S2-003 reliance on Q19/R4/R8 and candidate approval remain HARD-gated on WU-S2-008 exact source/counterpart acceptance.
- Separate pending gates: mechanical §13 Active item 5 pointer correction (only if still relevant after OI9), exact converged candidate report and Human approval, fresh positive migration-design Review before schema Build, Open Item 7, D1 Human pairing/Tier-0 authorization, O3/O4/O5, migration `000012` application, PostgreSQL/runtime, and whole-spine Testing.
