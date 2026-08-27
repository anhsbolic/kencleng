# Task 01: Foundation artifacts — TOTP dependency + openapi source amendment + bundle regen

> Back-reference : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar) — sections 2 (In scope bullet 7), 5 (D10), 9 (steps 1 & 8), 7 (risk row "Stale bundle"), 14 (Resolved #5)
> Depends on    : nothing — first task in the chain
> Model         : DeepSeek V4 Flash (purely mechanical: no business logic; verification is objective). Escalate to DeepSeek V4 Pro the moment the STOP condition below fires.

## Objective

Land the two zero-business-logic artifacts everything else builds on: the `pquerna/otp` dependency (tasks 03–04 import it) and the approved openapi source amendment with its mechanical bundle regeneration (D10 / Resolved #5). No Go code changes beyond `go.mod`/`go.sum`.

## Files

| File | Change |
|---|---|
| `backend/go.mod`, `backend/go.sum` | `require github.com/pquerna/otp <latest stable>` (pin); `go mod tidy` |
| `api/openapi/account.yaml` | `/account/security/mfa/enroll` gains a `"409"` response entry |
| `api/openapi.yaml` | regenerated mechanically (`cd api && npm run bundle`) |

## Implementation constraints

- **openapi amendment shape**: mirror the sibling error-response pattern already present elsewhere in `account.yaml` (e.g. the role-endpoints' `409` at lines ~590-594 reference Problem Details content). Response description should carry enough context for contract readers: enrollment rejected because MFA is already enabled; problem type `https://kencleng.dev/errors/mfa-already-enabled`. Do NOT touch any other path, schema, or component in the sources.
- Index/common/account.yaml hygiene: only `account.yaml` carries the delta (techplan §11 NOT-changed table).
- Bundle regeneration travels WITH the source edit in the same change — never one without the other (stale-bundle drift risk row).
- Commit both artifact groups (go.mod/change and api/) — this workspace does not auto-commit; leave staged for the human unless instructed otherwise.

## STOP condition (techplan D9-B lineage)

If `npm run bundle` drops components (e.g. `securitySchemes`) or otherwise diverges from an honest merge of the sources — STOP and report. That is a bundler defect, not something to fix by hand-editing the generated bundle.

## Verification

- `cd api && npm run bundle` exits 0; `git diff api/openapi.yaml` shows the 409 appearing and NO unrelated churn
- `go build ./...` clean with the new dependency present
- `govulncheck ./...` reports no new finding against `pquerna/otp`

## Rules enabled (not yet proven here)

None directly — this task enables R1–R11's imports (TOTP library) and keeps the handler↔spec contract verifiable at the code-review gate (D10 consequence). All behavioral proof lives in tasks 02–06.
