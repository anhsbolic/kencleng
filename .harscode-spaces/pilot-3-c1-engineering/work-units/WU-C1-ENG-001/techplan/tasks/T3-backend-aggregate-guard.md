# T3 — Backend Organization aggregate and guard

> Phase: Techplan decomposition; Author: `PARTICIPANT-C1-ENG-DECOMPOSER-001`; Created/Updated: 2026-10-09. Model: `gpt-6-luna`; Reasoning: medium; Session: `SESSION-C1-ENG-DECOMPOSER-001` (Run binding). Target revision: `524ef600c7f71246af6b71671d89c3c040fd9d44`; workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Purpose and outcome

Implement preparation, atomic Organization plus initial Owner establishment, durable guard and operator-controlled audit boundary, and scoped durable reads in the backend. This task owns G2 and G3 implementation surfaces within the exact recorded Human authorization.

## Authority

- Parent: `../techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`, status Approved.
- Governing items: Q1–Q7; R1, R3–R8, R11–R14; D1, D2, D4, D6, D7; RISK-1, RISK-2, RISK-4–RISK-7; §§8–9 contracts; §10 backend/DB/migration anchors; §11 verification; §12 G2/G3; §13 Active item 5.
- G2/G3 authorization is recorded in WU `events.md` and parent Resolved items 10–11. Keep within their exact scopes.

## Hard dependencies

**T1 condition:** T1's coordinated OpenAPI source and validated bundle define the preparation, confirm, reads, guard failures and response/error shapes. Durable evidence: source/bundle revision and validation/bundle result.

**T2 condition before confirmation/authenticated reads cross their boundary:** backend exposes a verified, server-owned authenticated person/session context with same-person binding, CSRF/Origin enforcement for mutations, and session validity rechecked under the session-row lock. Durable evidence: T2 code/interface pointers and authored focused checks for protocol negatives/person upsert/session-lock behavior. Mere T2 completion without this observable interface is insufficient.

## Scoped implementation

Use live anchors `backend/cmd/server/main.go`, `backend/internal/platform/db/db.go`, `backend/internal/domain/`, `backend/migrations/`, and current auth interface from T2. Additive records/constraints for preparation, Organization, attribution, guard singleton/audit; service/repository/HTTP composition; frozen validated display name and same-person preparation; one READ COMMITTED transaction locking session → guard → preparation; complete aggregate/result linkage, replay and indeterminate outcome semantics; restricted scoped list/detail; default-closed guard and fail-closed behavior. Preserve DB constraints, lock order, historical replay while held, and availability tradeoff in §§8–9. Runtime role and operator control must satisfy §9; no bypass or application user override.

## Verification and handoff

Author isolated real-PostgreSQL migration/constraint, transaction fault, concurrent confirm/guard race, scoped access and actual privilege checks, with the configured Podman route and distinct roles described by the parent. Record exact evidence and distinguish authored checks from executed results. Provide T4 the implemented guard control/provisioning interface, privilege limits, audit fields, all-holds/reopen conditions and rollback behavior as durable code/schema/config references.

## Explicit boundaries

No name-based uniqueness/matching, conflict adjudication, legal proof, review taxonomy, extra roles/Owners/transfers, cross-service aggregate, partial success, blind new attempt after uncertain commit, or destructive rollback. No frontend or shared API edits. If a required material meaning is absent from parent, stop and route reconciliation.
