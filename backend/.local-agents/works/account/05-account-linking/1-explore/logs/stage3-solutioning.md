# Stage 3 — Solutioning (05-account-linking)

> Session: 2026-08-26. Written after Stage 2 gap analysis was reviewed.
> Open questions Q1–Q3 were answered by Anhar the same day — all three
> accepted as recommended; see "Open questions" section at the bottom.
> Status: solutioning complete, ready to move to 2-plan/techplan.

## D1 — INV-account-05 primitive: user-wide refresh-token revocation

**Options:**
- A: New repo method `RevokeAllRefreshTokensForUser(ctx, tx, userID)` —
  single goqu `UPDATE refresh_tokens SET revoked_at = now() WHERE
  user_id = $1 AND revoked_at IS NULL`.
- B: Loop over the user's families calling existing
  `RevokeRefreshTokenFamily`.

**Decision: A.** One statement, atomic inside the caller's tx, covered
by `ix_refresh_tokens_user_id`. B adds a query + indirection for zero
benefit and misstates intent (families are an internal lineage concept;
the requirement is user-scoped). Statement matches INV-account-05
exactly ("every row with revoked_at IS NULL", *including* rotated-out
rows — deliberately no `replaced_by_id` guard, same reasoning as family
revocation).

**Note:** this makes feature 05 the first implementer of
INV-account-05; Fitur 04 will reuse it later (dependency direction
inverted from the spec's wording — flagged in Stage 2, resolved here
rather than re-derived).

## D2 — Unlink atomicity: making check-then-delete non-racy

**Options:**
- A: Single guarded statement: `DELETE ... WHERE provider_type='google'
  AND user_id=$1 AND EXISTS(verified other identity) RETURNING id` —
  atomic, but affected-rows can't distinguish the two 409 cases (spec
  mandates distinct messages).
- B: Inside one tx: `SELECT id, provider_type, verified_at FROM
  auth_identities WHERE user_id=$1 FOR UPDATE` → classify in Go (no
  google row / no other identity / other-unverified / ok) → conditional
  `DELETE`.

**Decision: B.** Three distinguishable failure/success outcomes plus the
password re-auth comparison make single-statement affected-rows
arithmetic unreadable. `FOR UPDATE` on the user's identity rows
serializes concurrent unlinks: loser blocks, then classifies post-commit
state (READ COMMITTED re-read at lock acquisition), so the count guard
cannot race past. Follows the codebase's "correctness lives in Postgres
behind guarded statements" philosophy while staying testable.

Edge: concurrent winner already deleted the google row → loser sees no
google identity → map to idempotent success (200) rather than 409 —
deleting an already-absent link achieves the caller's goal.
*(Assumption: this concurrent-second-request status is not pinned by
the spec — recorded as such.)*

## D3 — Multi-google-identity users

`callbackLink` permits attaching a second google identity (different
google email). Spec assumes singular "the google AuthIdentity row".

**Decision:** unlink deletes **all** google identities of the caller
(the action's meaning is "remove Google as a login method"), under the
same INV-account-02/12 guard. Alternative (reject until one remains)
rejected as user-hostile with no security gain.

## D4 — Set-password branch selection

**Decision:** reuse existing `FindIdentifierHashByUserAndProvider(userID,
providerEmailPassword)` (`found bool`) for branch dispatch — no new repo
method needed for selection itself. Server-side only, never a client
flag (per AC). Ordering discipline copied from `Register`: policy
validation (`validatePassword`) → bcrypt always runs (timing
equivalence) → branch.

## D5 — Branch 1 duplicate-email conflict

**Decision:** mirror registration exactly — pre-check via
`FindAuthIdentityByIdentifierHash(email_password, hash)`; claimed by
anyone → send new nudge type `NudgeSetPasswordConflict =
"set_password_conflict"` (added to `platform/notification` constants;
FakeSender/DevSender need no logic change) and return nil → generic
202. Unique-violation fallback (`isUniqueViolation` → rollback → nudge
→ nil) covers the race the pre-check misses, same as R16. `dummyWrite`
preserves DB-time shape on the conflict branch. New identity insert +
token issuance in one tx; verification email sent after commit.

## D6 — Branch 2 change-password mechanics

**Decision:** one tx: `compare(current_password)` against the existing
identity's `credential_secret` (via `s.compare` seam — burns comparable
CPU even on failure paths) → `UpdateCredentialSecret(userID, newHash)`
(new repo method, single-row UPDATE keyed `(user_id,
'email_password')`) → `RevokeAllRefreshTokensForUser` (D1) → commit →
200. Validation-before-mutation ordering per AC (policy check before
touching anything). Wrong password → 401 via existing
`ErrInvalidCredentials` mapping. No email submitted in this branch —
the request schema is conditionally shaped like `MfaDisableRequest`
(established pattern).

## D7 — Audit-at-verification wrinkle (Branch 1)

Spec requires the Branch 1 audit entry *when the identity becomes
verified*, but `/auth/verify-email` must stay "unchanged" and its tokens
are indistinguishable (`purpose='email_verification'`; `VerifyEmail`
ignores purpose after redeem).

**Options:**
- A: Relax `auth_tokens.purpose` CHECK via migration 000010 to admit a
  third value (e.g. `email_verification_link`); issue Branch 1 tokens
  with it; `VerifyEmail` writes the audit entry when the redeemed purpose
  indicates the link flow. Endpoint code path genuinely unchanged (redeem
  is purpose-blind today); audit becomes possible without heuristics.
- B: Write Branch-1 audit at identity-creation time instead (violates
  spec's explicit "not at initial creation").
- C: Heuristic (created_at deltas between token and identity) — fragile,
  rejected.

**Decision: A — RESOLVED 2026-08-26, accepted by Anhar.** Migration
000010 relaxes the `auth_tokens.purpose` CHECK to admit a third value;
`/auth/verify-email` stays behaviorally unchanged (redemption is
purpose-blind) while the audit becomes truthful. Rationale recorded:
Option B directly contradicts the spec text ("not at initial creation")
and Option C breaks under clock skew/concurrency; the durable purpose
distinction also pre-empts task #08's action_type vocabulary needing to
separate "identity activated via registration" vs "via linking".
Honest fallback had Q1 been declined: defer the Branch-1 audit with an
explicit assumption note — never write it at creation time.

## D8 — Authenticated session extraction for `/account/security/*`

**Options:**
- A: Inline per-handler extraction (copy `sessionToken` + verifier into
  each handler).
- B: Small `RequireSession`-style middleware over an `accountMux`, built
  on the existing `sessionToken()` helper + a generalized verifier
  (kept out of `platform/auth/` — Tier 0 fenced).

**Decision: B.** Two endpoints now (+ `/account/me` in task #7, MFA in
#6, reset-password handlers in #4 later) justify one middleware; 401
Problem written once. Explicit authz boundary stays visible per root
golden rule. Naming/location settled during implementation so tasks
#4/#6/#7 can share it without a breaking refactor.

## D9 — Contract repairs discovered during exploration

1. `api/openapi.yaml`: add the missing
   `components.securitySchemes.bearerAuth` definition (referenced
   globally but undefined — tooling-breaking defect, safe mechanical fix).
2. Stale spec reference ("both endpoints need a schema update") — do
   NOT edit the feature spec (root AGENTS.md §4); record the staleness
   in the risk note instead.

## Execution shape (files to touch)

| Layer | File(s) | Change |
|---|---|---|
| migrations | `000010_*.sql` (Q1 resolved = A) | relax `auth_tokens.purpose` CHECK to admit `email_verification_link` |
| repo port | `internal/domain/account/repository.go` | `DeleteAuthIdentitiesByProvider` (or guarded variant), `FindAuthIdentitiesByUser`, `UpdateCredentialSecret`, `RevokeAllRefreshTokensForUser` |
| repo adapter | `internal/domain/account/repository_db.go` | implementations (goqu only, parameterized) |
| service | new `security.go` in domain pkg | `SetPassword`, `UnlinkGoogle`, sentinels (`ErrOnlyIdentity`, `ErrRemainingUnverified`) |
| notification | `internal/platform/notification/sender.go` | `NudgeSetPasswordConflict` const |
| transport | new `account_security.go` | both handlers + session middleware |
| wiring | `cmd/server/main.go` | `accountMux` + routes |
| tests | beside each | table-driven units + race tests named per spec + integration |

No changes to: Tier 0 paths (`platform/auth/`, crypto), `docs/spec/*`,
existing handlers beyond route mounting.

## Test plan sketch (maps to threat breakdown)

- `TestUnlinkGoogle_ConcurrentRequests_GuardHolds` — ≥100-goroutine
  stress per tasks.md KPI; guard must hold under `-race`.
- `TestSetPassword_ConcurrentDuplicateEmail_Race`,
  `TestSetPassword_GenericResponse_AllBranches`.
- `TestSetPassword_PasswordPolicy`, `TestSetPassword_BreachCheck_FailOpen`.
- Token single-use: reuse `TestVerifyEmail_TokenSingleUse_Concurrent`
  pattern (same endpoint).
- `TestUnlinkGoogle_RequiresReauth`, `TestUnlinkGoogle_WrongPassword_Rejected`.
- `TestUnlinkGoogle_RejectsUnverifiedRemainingIdentity`.
- `TestSetPassword_Branch2_RequiresCurrentPassword`,
  `TestSetPassword_Branch2_AllSessionsRevoked` (INV-account-05 — first
  ever implementation).
- Invariant traceability tests named per tasks.md KPI
  (`INV-account-01/02/05/08/12`).
- Audit test asserting exact `user_logs.action_type` rows (KPI row:
  account linking actions).

## Open questions for Anhar — ALL RESOLVED 2026-08-26

- **Q1 (D7):** RESOLVED — Option A accepted (migration-based third
  token purpose; see D7).
- **Q2 (D2/D3):** RESOLVED — both assumptions confirmed:
  - Idempotent success (200) on concurrently-already-deleted unlink,
    matching existing idempotent patterns (`RevokeRefreshTokenByHash`,
    logout, `callbackLink` no-op) and the caller's achieved intent.
  - Delete-all-google-identities semantics — the endpoint's meaning is
    "remove Google as a login method"; multi-google users are rare but
    reachable via `callbackLink`, so singularity must not be assumed.
- **Q3:** RESOLVED — proceed despite tasks.md's S1 serial order (#3, #4
  unfinished), with the inversion explicitly recorded in the risk note:
  (a) the order inversion itself, (b) `RevokeAllRefreshTokensForUser`
  is written first here by necessity and Fitur 04 must reuse it rather
  than re-derive it, (c) nothing in D1–D8 depends on #4's handler shape;
  #3's remaining work doesn't conflict with this task's files.
