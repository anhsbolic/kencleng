# Task 04: UnlinkGoogle service (guarded classification + concurrency harness)

> Back-reference : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved) — sections 4 (R9–R13), 5 (D1–D3, evaluation-order micro-decision), 8 (UnlinkGoogle flow block — keep verbatim as the logic source)
> Depends on    : task-02 (repository methods incl. locked finder); task-03 (shares `security.go` — serial to avoid same-file conflicts, and sentinels/consts co-locate there)
> Model         : GLM 5.2 (max) (concurrency-invariant design is the highest-risk reasoning in the slice; compensating control: mandatory Complex-tier dual-model code review per model-routing)

## Objective

Implement unlink with atomic guard classification under row locks, password re-auth as the last gate, hard delete of ALL the caller's google identities, and the atomic audit entry. Then prove the guard cannot race past with a ≥100-goroutine stress harness asserting end-state invariants.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/security.go` | Extended — `UnlinkGoogle`, sentinels `ErrOnlyIdentity`, `ErrRemainingUnverified` |
| `backend/internal/domain/account/security_test.go` | Extended — tests below |

## Flow (authoritative — techplan §8)

```
UnlinkGoogle(ctx, userID, password):
  tx {
     rows := SELECT id,provider_type,verified_at,credential_secret
             FROM auth_identities WHERE user_id=$1 FOR UPDATE       # serialize unlinks
     google := rows[provider=google]
     if google empty: commit; return 200            # idempotent no-op (incl. race loser)
     others := rows - google
     if others empty            -> ErrOnlyIdentity            # 409 only-identity
     if none(others.verified)   -> ErrRemainingUnverified     # 409 unverified-remaining
     compare(password, verifiedOther.secret) != nil -> ErrInvalidCredentials # 401
     DELETE FROM auth_identities WHERE id IN google.ids     # hard delete
     InsertUserLog({user_id, action_type:"account_linking"})
  } commit -> 200 UnlinkGoogleResponse
```

Key decisions inherited: D2 (`FOR UPDATE` + in-Go classification — NOT a single guarded DELETE; three outcomes need distinct messages), D3 (delete **all** google identities; multi-google users are reachable via `intent=link`), evaluation-order micro-decision (**guards before re-auth** — a Google-only caller has no password yet; the idempotent no-op returns without password; password is the last gate). Edge note from techplan §7: if >1 `email_password` identity ever existed (out-of-band writes only), compare against the verified one deterministically (`ORDER BY created_at LIMIT 1` if ever ambiguous).

## Rules to prove

- **R9** `TestUnlinkGoogle_Success_HardDeletesAndAudits` — google rows → 0 for `(user_id,'google')`; exact `action_type='account_linking'` asserted in same-tx `user_logs` row (tasks.md audit KPI)
- **R10** `TestUnlinkGoogle_OnlyIdentity_Rejected409` — sentinel maps later to type `https://kencleng.dev/errors/only-identity`
- **R11** `TestUnlinkGoogle_RejectsUnverifiedRemainingIdentity` — distinct sentinel for `https://kencleng.dev/errors/unverified-remaining-identity`
- **R12** ordering matrix: `TestUnlinkGoogle_RequiresReauth`, `TestUnlinkGoogle_WrongPassword_Rejected` (401, zero state change), `TestUnlinkGoogle_IdempotentNoGoogleRow_Returns200`
- **R13** `TestUnlinkGoogle_ConcurrentRequests_GuardHolds` — ≥100 goroutines mixing unlink/verify interleavings against one user under `-race`. Assert END-STATE INVARIANTS, not "didn't crash": at most one deletion success; after every successful unlink the user retains ≥1 identity AND ≥1 verified remainder (INV-account-02/12 hold at every observed end-state); losers ∈ {idempotent 200} ∪ {correct 409} — never a spurious success-after-guard-failure. Per tasks.md KPI this is the slice's concurrency stress test.
- **R16** log-scan coverage extended to unlink paths (no email/password/hash leakage)

## Common mistakes (techplan §13 subset)

- Treating the concurrent loser as 409 → classify post-lock; absent google row = goal already achieved → 200
- Asserting only "no panic" in race tests → invariant assertions mandatory
- Building the DELETE with `fmt.Sprintf`'d ids → goqu `Where(goqu.C("id").In(ids))`
- Reordering guards behind the password check → breaks reachability of both 409s for passwordless Google-only callers
- Deadlock paranoia → none warranted: waits are one-directional vs `SetUserVerified` (techplan §7 Low row); do not add retry loops

## Out of scope here

HTTP status mapping (task-05); testcontainers proofs of R9/R13 DB truths (task-06).

## Verification

`go test -race ./internal/domain/account/...` green including the stress harness.
