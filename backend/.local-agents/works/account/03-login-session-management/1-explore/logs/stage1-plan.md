# Stage 1 — Plan Announcement

> Task: account #03 — Login & Session Management
> Feature spec: `docs/spec/1-account/features/03-login-session-management.md`
> Date: 2026-08-26
> Status: confirmed by Anhar, proceeded to Stage 2 unchanged

## Inputs read before announcing

- Workflow docs: `harscode-workspace/workflow/1-exploration/{guidelines,sniffing-checklist,examples}.md`
- `docs/spec/README.md` (four doc types, authority order)
- Feature spec itself (4 endpoints, lockout folded in, Tier 1 + Tier 0 fenced sub-area)
- Convention files only (AGENTS.md root + backend) + structure-only listing of
  `backend/` — no implementation files read yet at this point.

## Announced exploration areas & order

1. **Requirement anchors first** — `api/openapi.yaml` (`/auth/login`,
   `/auth/login/mfa`, `/auth/refresh`, `/auth/logout` + schemas/cookies/error
   codes), `docs/spec/1-account/invariants.md` (INV-03/04/06/07), threat-model
   component 2.
   *Why:* every gap statement needs a precise requirement baseline; the spec's
   resolved assumptions (A/B: HS256 mfa_pending token w/ separate secret;
   C: `login_attempts.stage`; D: multi-tab deferred to frontend) define done.
2. **`internal/domain/account/`** — what exists from features 01/02 that login
   builds on (identity lookup, bcrypt, token entities, repo patterns).
3. **`internal/platform/auth/` + `ratelimit/` + `secrets/`** — Tier 0 fenced
   area state; JWT/signing capability; config wiring; limiter state.
4. **`internal/transport/http/` + `cmd/server/main.go`** — middleware, cookie
   infra, error mapping, handler patterns, route registration/wiring.
5. **`migrations/` + DB layer** — actual DDL vs required tables. *Why last:*
   whether a migration is even needed depends on areas 2–4 findings.

## Pre-flagged suspicions (to verify during Stage 2)

- `/auth/login/mfa` depends on TOTP/backup-code infra owned by feature 06 —
  sequencing boundary unclear.
- Feature spec internal tension: "Risk tier" § says both tokens share the same
  ES256 key while Assumption A resolves to a separate HS256 secret.

Both confirmed in Stage 2 (see gap-area reports).
