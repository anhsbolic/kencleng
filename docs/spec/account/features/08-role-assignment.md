# Feature Spec — Role Assignment (Admin)

> File: `docs/spec/account/features/08-role-assignment.md`
> Status: draft — all open items resolved, ready for human review
> Risk tier: 1
> Domain: account

## Endpoint

- `GET /admin/users`
- `POST /admin/users/{userId}/roles`
- `DELETE /admin/users/{userId}/roles?role={admin|kurator}`

## Acceptance criteria

### `GET /admin/users`

- Given the caller holds `role=admin`, When called, Then a paginated,
  cursor-based list of users is returned (`UserListResponse`), **hard
  capped at `limit=20`** regardless of what's requested (resolved in
  `docs/spec/account/threat-model.md` component 6 — bounds the
  per-request decrypt cost of bulk-exposing PII).
- Given the caller does **not** hold `role=admin`, When called, Then
  `403` — checked explicitly at the handler level, not only via a
  query filter (per `AGENTS.md`'s golden rule on authz).

### `POST /admin/users/{userId}/roles`

- Given the caller is Admin, the target `userId` exists, and assigning
  the requested role doesn't violate INV-account-09 or INV-account-10,
  When called, Then a `user_roles` row is created (`granted_by` =
  caller's `user_id`), response `201` with the updated `User`.
- Given the target `userId` doesn't exist, When called, Then `404`.
- Given assigning `admin` to a user who currently holds `kurator`
  (INV-account-09) or is an `organization_representatives` row of any
  organization (INV-account-10), When called, Then `409`, no state
  change — the two cases get distinct messages (mirrors the two-
  distinct-409-messages pattern already established in
  `05-account-linking.md` for unlink).
- Given the target user **already holds** the requested role (a
  duplicate assignment, e.g. re-clicking "make Admin" on an
  already-Admin user), When called, Then `409` "user already has this
  role" — see Assumption A, resolved.
- Given two concurrent assignment requests targeting the same user
  with conflicting roles (`admin` and `kurator`), Then exactly one
  succeeds; the guarded uniqueness/exclusivity check must not let both
  through (INV-account-09's concurrency test).

### `DELETE /admin/users/{userId}/roles?role=...`

- Given the caller is Admin, the target user exists and currently
  holds the specified role, When called, Then the `user_roles` row is
  deleted, response `204`.
- Given the target user doesn't exist, or doesn't currently hold the
  specified role, When called, Then `404` (both cases collapse to the
  same "nothing to revoke" response — no distinguishing signal needed
  here, unlike the enumeration-sensitive endpoints elsewhere in this
  domain, since this is an Admin-only authenticated action against a
  known-to-exist-in-the-UI target).
- Given the role being revoked is `admin`, and the target is the
  **last remaining Admin in the entire system** (`COUNT(user_roles
  WHERE role='admin') = 1` and this is that row), When called, Then
  **rejected** (`409`) — INV-account-13, resolved 2026-08-05: the
  system must always retain at least one Admin, since there is no
  self-service path back to Admin (Fitur 5's only bootstrap mechanism
  is a manual seed script, not usable at runtime).
- Given the role being revoked is `kurator`, and the target has a
  `pending` row in `organization_curation_assignments`
  (`kurator_id = userId`), When called, Then **rejected** (`409`) —
  INV-account-14, resolved 2026-08-05: demote is blocked until that
  assignment's `decision` moves to `approved`/`rejected` (a read
  against a table `organization` owns, not a write into it — see
  INV-account-14's cross-domain note).

## Error cases

| Condition | Expected response |
|---|---|
| Non-Admin calling any of these endpoints | `403` |
| Target `userId` not found (assign) | `404` |
| Target user not found, or doesn't hold the specified role (revoke) | `404` |
| Assigning `admin` to an existing `kurator` | `409` |
| Assigning `admin` to an existing organization representative | `409` |
| Assigning a role the target already holds | `409` |
| Revoking the last remaining Admin's own `admin` role | `409` |
| Revoking `kurator` from a user with a `pending` curation assignment | `409` |

## Applicable invariants

- `docs/spec/account/invariants.md#inv-account-09` — Admin ⊥ Kurator.
- `docs/spec/account/invariants.md#inv-account-10` — Admin ⊥
  organization Representative (cross-domain, declared here — see that
  invariant's "Reference for `organization` domain" section).
- `docs/spec/account/invariants.md#inv-account-11` — `user_logs`
  append-only; this endpoint is one of the mandatory write-sites.
- `docs/spec/account/invariants.md#inv-account-13` — at least one
  Admin always exists — resolved 2026-08-05, see Assumption B below.
- `docs/spec/account/invariants.md#inv-account-14` — revoking
  `kurator` requires no pending curation assignments (cross-domain
  read against `organization_curation_assignments`) — resolved
  2026-08-05, see Assumption C below.

## Threat breakdown

Derived from `docs/spec/account/threat-model.md` component 6, plus
two threats found while drafting this spec (Assumptions A & B):

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| `/admin/users` bulk-exposing decrypted PII per request | Hard-capped `limit=20` (resolved in threat-model), Admin-only authz explicit at handler level | `TestAdminUsers_PaginationCapped20`, `TestAdminUsers_RequiresAdminRole` |
| Forged `userId` targeting an unintended user | Explicit target validation + authz at handler level, not query filter alone | `TestAssignRole_TargetValidation` |
| Role assign/revoke repudiation | Mandatory `user_logs` entry (INV-account-11), append-only | `TestAssignRole_AuditLogWritten`, `TestRevokeRole_AuditLogWritten` |
| Admin ⊥ Kurator / Admin ⊥ Representative bypass via race | Guarded exclusivity check, concurrency test (INV-account-09/10) | `TestAssignRole_ConcurrentConflictingRoles_OnlyOneWins` |
| **New — duplicate role assignment** treated as a silent no-op or a confusing 500 (unique constraint violation surfacing as an unhandled DB error) instead of a clean response | `409` "already has this role," explicitly handled rather than left to bubble up as a raw constraint violation | `TestAssignRole_DuplicateAssignment_Clean409` |
| **New — Admin self-lockout**: the last remaining Admin revokes their own `admin` role (accidentally or via a compromised session), leaving zero Admins with no runtime recovery path | **Resolved 2026-08-05** — `409` guard on revoking the last Admin (INV-account-13) | `TestRevokeRole_LastAdminGuard` |
| **New — demoted-Kurator orphaned assignment**: revoking `kurator` while the user has a `pending` curation assignment leaves `organization`'s curation queue pointing at a non-Kurator | **Resolved 2026-08-05** — `409` guard, blocked until the assignment is decided (INV-account-14) | `TestRevokeRole_PendingCurationGuard` |

## Risk tier & rationale

**Tier 1** — INV-account-09/10 concurrency-safe exclusivity checks
(project goal #2), mandatory audit log writes (matches the Tier 1
example "audit log writes" in `kencleng-agentic-workflow.md` §4
verbatim), and this is the domain's only pure elevation-of-privilege
surface. No Tier 0 sub-area (no JWT/TOTP/crypto core logic here).

## Assumptions / open questions

**A. Resolved while drafting — duplicate role assignment.** Not
addressed in `kencleng-phase0-detail.md`. The `UNIQUE (user_id, role)`
constraint on `user_roles` means an unhandled duplicate assignment
would otherwise surface as a raw DB constraint violation (a poor
API experience, and a Tier 1 correctness smell — unhandled DB errors
leaking to the client). Resolved: explicit `409` "user already has
this role" check before the insert. High-confidence derivation, not
flagged for further confirmation.

**B. Resolved — 2026-08-05.** Found while drafting: nothing in the
source docs prevents revoking the last remaining Admin's `admin` role,
which would leave the system with zero Admins and no runtime recovery
path (Fitur 5's bootstrap is a one-time manual seed script, not an
endpoint). **Decision: block this case** (`409`), formalized as
**INV-account-13** in `invariants.md`. Accepted trade-off: adds a
`COUNT(*)` check to every Admin-role revocation (small, cheap query) —
worth it given the alternative is a state unrecoverable via the API.

**C. Resolved — 2026-08-05.** Carried forward from
`kencleng-phase0-detail.md` Fitur 5's own unresolved open question:
"Kurator yang di-demote — assignment kurasi yang sedang `pending`
miliknya diapakan?" **Decision: block the demote** until that
assignment is decided (`approved`/`rejected`), formalized as
**INV-account-14**. Chosen over auto-reassignment specifically to keep
`account` from needing to know `organization`'s assignment logic — the
guard is a **read** against a table `organization` owns, not a write
into it, preserving the domain-boundary principle
`kencleng-agentic-workflow.md` already establishes elsewhere. This is
a genuine cross-domain forward reference (same shape as
INV-account-10, but reversed — see INV-account-14's cross-domain note
in `invariants.md`).

## Audit log entry?

**Yes** — explicitly in Fitur 9's scope ("role assign/revoke"). Write
a `user_logs` entry for both successful assignment and successful
revocation. Per the established pattern from `05-account-linking.md`
and `06-mfa-totp.md`, also triggers a user-facing notification to the
**target** user ("Role kamu di Kencleng telah diubah oleh Admin") —
cross-domain dependency on the not-yet-built `notification` domain,
same forward-reference pattern as those two specs.

## References

- `docs/project/kencleng-phase0-detail.md` Fitur 5
- `docs/project/kencleng-erd.md` §1 (`user_roles`)
- `docs/spec/account/invariants.md` INV-account-09, 10, 11, 13, 14
- `docs/spec/account/threat-model.md` component 6
- `api/openapi.yaml` — `GET /admin/users`,
  `POST /admin/users/{userId}/roles`,
  `DELETE /admin/users/{userId}/roles` (needs the `LimitParam`
  override for `/admin/users` already flagged in the threat model, not
  yet applied)