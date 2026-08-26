# Stage 1 — Plan Announcement

> Announced before any implementation file was read; confirmed by Anhar
> before Stage 2 began.

**Feature**: Forgot & Reset Password (`POST /auth/forgot-password`,
`POST /auth/reset-password`) — Tier 1, key invariants INV-account-05
(atomic session revoke) and INV-account-08 (single-use tokens).

## Areas to explore, in order, and why

1. **Domain ground-truth docs** (`docs/spec/1-account/invariants.md`,
   `threat-model.md`, `tasks.md`) — first because they're the ground
   truth the feature spec references (INV-05, INV-08); everything else is
   measured against them.
2. **API contract** (`api/openapi/*.yaml` + generated bundle) — second,
   since handlers must match it and it defines the generic-202
   anti-enumeration shape.
3. **Existing token infrastructure** (`entity.go`, `repository.go`,
   `repository_db.go`, migrations 000003/000004) — bottom-up: data model
   first.
4. **Service layer patterns** (`service.go`) — closest sibling flows
   (Register's anti-enumeration branches, VerifyEmail's redeem-in-tx,
   validatePassword).
5. **Notification touchpoint** (`internal/platform/notification`) —
   cross-domain email dependency of forgot-password.
6. **Transport plumbing** (`cmd/server/main.go` routing, `errors.go`
   problem mapping, rate-limit middleware) — top of the stack.
7. **Test conventions** (`*_test.go`, integration harness) — last, since
   test shapes depend on all seams introduced above.

Rationale for order: spec ground truth → contract → bottom-up through the
stack (data model → service logic → cross-domain email → HTTP layer →
tests), so each layer's gap analysis can reference layers already
confirmed above it.

Stage 2 findings per area: `01`–`07` area logs in this directory.
Stage 3 decisions: `stage3-solutioning.md`.
