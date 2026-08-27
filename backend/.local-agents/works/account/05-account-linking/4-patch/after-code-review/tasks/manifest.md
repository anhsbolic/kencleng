# Manifest — 05-account-linking patch tasks (after code review)

> Generated : 2026-08-27
> Source    : `4-code-review/report.md` (Verdict: Request changes) + `4-patch/after-code-review/plan.md`
> Note      : snapshot at generation time — this file does NOT track progress status; status belongs to the PR/ticket domain.

## Task list

| # | File | Title | Blocking? |
|---|---|---|---|
| 1 | `task-01-error-dispatch-fix.md` | Extend `MapServiceError` with the two 409 sentinels; route both security handlers' non-validation errors through it; add 500-dispatch transport tests | **Yes** (S1 / BP1 / C1) |
| 2 | `task-02-optional-cleanup.md` | Optional follow-ups: typed `SetPassword` result, shared password-validation helper, down-migration comment tightening | No (Q1 / Q2 / Q3-BP2) |

## Splitting axis

**Risk/sequence**. Task 01 is the only blocking fix and must land
before merge; task 02 is a clearly-marked optional bucket that can be
deferred or cherry-picked per item. They are independent — task 02
does not depend on task 01 landing (though if both land in one PR,
run task 01 first so the dispatch change is reviewable in isolation).

## Dependency graph

```
01 (blocking)        02 (optional, independent)
```

## Back-reference

- Code review report: `4-code-review/report.md` — findings S1/BP1/C1 (blocking), Q1/Q2/Q3-BP2 (optional).
- Originating techplan: `2-plan/techplan.md` — task-05 Files table originally specified the `MapServiceError` extension that this patch restores.
- Repo conventions: `backend/AGENTS.md` (golden rules, error handling), `internal/transport/http/errors.go:76-109` (`MapServiceError`), `auth_register.go:60` / `auth_password_reset.go:89` (usage precedent).

## Review note

Task 01's diff is small and surgical — the reviewer should confirm that
(a) `MapServiceError` now covers `ErrOnlyIdentity`, `ErrRemainingUnverified`,
and `ErrInvalidCredentials` (the latter already mapped at `errors.go:98-101`,
should stay there), (b) both handlers call `MapServiceError(w, err)` for
every non-`ErrValidation` error, (c) the new 500-dispatch test injects a
non-sentinel wrapped error and asserts the `internal` problem type, and
(d) no inline 409 strings remain in the handler files.
