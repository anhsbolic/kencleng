# Task List — account

> File: `docs/spec/account/tasks.md`
> Status: draft — KPI/metrics proposed by the assisting agent, pending
> Anhar's review/approval
> Last updated: 2026-08-05

## Delivery KPI / metrics

Proposed once here, applies to every task below unless a task's row
explicitly overrides it with something stricter. This is a first
attempt at a concept (`kencleng-agentic-workflow.md` §12 step 5)
that's never been made concrete before this domain — treat these
thresholds as a starting point to tune once real gate output exists,
not a fixed standard to defend.

| Metric | Applies to | Threshold |
|---|---|---|
| Automated gates (§7) | Every task | 100% pass, 0 skipped/`t.Skip`'d, `make verify` exits 0 |
| Named invariant test | Every task referencing an `INV-account-NN` | At least one test per referenced invariant, named so it's traceable back to that `INV-` ID (e.g. `TestInvariant_Account03_RefreshRotationSingleUse`) — a claim of "covered" with no such test is treated as unverified, per `AGENTS.md` §5 |
| Concurrency stress test | Tasks touching a race-sensitive invariant (INV-account-01, 02, 03, 04, 08, 09, 10) | `go test -race` clean, plus a stress harness with **≥100 concurrent goroutines** hitting the same contested row/constraint, 0 invariant violations across the run |
| Test coverage | Every task | ≥80% of new/changed lines in the task's package(s) |
| Security layer A | Every task | `gosec`/`gitleaks`/`govulncheck` — 0 findings, or an explicit accepted-risk note if a finding is a deliberate false positive |
| Security layer B (adversarial) | Tier 0/1 tasks | Fresh-context adversarial pass completed; every finding triaged (fixed, or explicitly accepted with reasoning) before merge |
| Tier 0 sub-area authorship | Tasks #3 and #6 (see below) | The specific Tier 0 file(s) are human-authored or human-rewritten, not agent-generated wholesale — a boolean checklist item, not a numeric metric |
| Audit log test | Tasks whose scope includes a Fitur 9 action (role assign/revoke, MFA enable/disable, account linking, self-PII reveal) | A test asserting the exact `user_logs.action_type` row is written, not just that the primary operation succeeded |

## Tasks

20 `account`-tagged OpenAPI endpoints grouped into 8 vertical slices —
grouped by shared flow/tables rather than 1:1 per endpoint, consistent
with `kencleng-agentic-workflow.md` §11 "one vertical slice per
session" (a slice can span more than one endpoint when they're one
cohesive flow).

| # | Task | Feature spec file | Endpoints | Tier | Rationale | Parallel group |
|---|---|---|---|---|---|---|
| 1 | Register & Email Verification | `01-register-email-verification.md` | `POST /auth/register`, `POST /auth/verify-email`, `POST /auth/verify-email/resend` | **1** | INV-account-01 (concurrent-uniqueness race test); email-enumeration generic-response logic (resolved 2026-08-05) needs a constant-time correctness test | Serial group S1 |
| 2 | Google OAuth Login/Register | `02-google-oauth-login-register.md` | `GET /auth/google/redirect`, `GET /auth/google/callback` | **1** | State/nonce CSRF-critical; no-auto-merge anti-takeover rule needs human review even though most of the flow is standard OAuth plumbing | Serial group S1 |
| 3 | Login & Session Management (incl. lockout, Fitur 2C) | `03-login-session-management.md` | `POST /auth/login`, `POST /auth/login/mfa`, `POST /auth/refresh`, `POST /auth/logout` | **1**, with a **Tier 0 fenced sub-area**: refresh-token rotation/reuse-detection logic (INV-account-03, INV-account-04) and JWT issuance/verification — these specific files must be human-authored/paired, not agent-generated, and marked "no-agent-write" in `AGENTS.md` | Matches the Tier 0 examples in `kencleng-agentic-workflow.md` §13.2 verbatim ("JWT/TOTP, refresh-token reuse detection") | Serial group S1 |
| 4 | Forgot & Reset Password | `04-forgot-reset-password.md` | `POST /auth/forgot-password`, `POST /auth/reset-password` | **1** | INV-account-05 (revoke-all-sessions atomicity), INV-account-08 (single-use token race), anti-enumeration | Serial group S1 |
| 5 | Account Linking | `05-account-linking.md` | `POST /account/security/google/unlink`, `POST /account/security/set-password` | **1** | INV-account-02 (min-1-identity guard, concurrency-sensitive), identifier-conflict handling on set-password | Serial group S1 |
| 6 | MFA TOTP | `06-mfa-totp.md` | `POST /account/security/mfa/enroll`, `POST /account/security/mfa/enroll/confirm`, `POST /account/security/mfa/disable` | **1**, with a **Tier 0 fenced sub-area**: TOTP secret generation/encryption (`secret_encrypted`) and verification — same reasoning as task #3 | Group B (independent tables: `mfa_totp_secrets`, `mfa_backup_codes`) |
| 7 | Account Profile (read) | `07-account-profile.md` | `GET /account/me` | **2** | Pure read, no mutation, no invariant involved, standard OpenAPI-derived DTO | Can run anytime, no grouping needed |
| 8 | Role Assignment (Admin) | `08-role-assignment.md` | `GET /admin/users`, `POST /admin/users/{userId}/roles`, `DELETE /admin/users/{userId}/roles` | **1** | INV-account-09, 10 (exclusivity, concurrency + elevation-of-privilege sensitivity); INV-account-13 (last-Admin guard), INV-account-14 (cross-domain read against `organization`'s pending curation assignments); mandatory `user_logs` write | Group C (independent table: `user_roles`) |

No task in this domain lands at Tier 0 in full or Tier 3 — expected,
`account` is inherently security-critical across the board.

## Parallel / serial grouping

- **Serial group S1** (tasks #1-#5): all touch the same core tables
  (`auth_identities`, `refresh_tokens`, `auth_tokens`) and have a
  natural dependency order — Register → Login → Forgot Password →
  Google OAuth → Linking (Linking depends on Google OAuth and
  set-password both already existing). Running these in parallel risks
  exactly the kind of migration-numbering collision flagged as an
  unresolved risk in `kencleng-agentic-workflow.md` §12
  ("Known risk with parallel task execution").
- **Group B** (task #6, MFA): independent tables
  (`mfa_totp_secrets`, `mfa_backup_codes`), can run in parallel with
  S1 or Group C.
- **Group C** (task #8, Role Assignment): independent table
  (`user_roles`), can run in parallel with S1 or Group B.
- **Task #7** (`GET /account/me`): trivial, no grouping constraint,
  fits in wherever there's capacity.
- **Migration ownership note** (added 2026-08-26, per task #3's
  approved techplan D1-C — see
  `.local-agents/works/account/03-login-session-management/2-plan/techplan.md`):
  migrations `000006`–`000009` (`login_attempts`, `mfa_totp_secrets`,
  `mfa_backup_codes`, `user_roles`) are created by **task #3** as
  schema-pre-settle. Tasks #6 and #8 own table *logic* only — do not
  create these tables' migrations again.

Suggested execution order given the above:
`S1 (serial: #1→#2→#3→#4→#5)` with `#6` and `#8` run in parallel
alongside S1 once their independent tables' migrations are settled;
`#7` slotted in wherever convenient.

## Status tracker

| # | Status | Notes |
|---|---|---|
| 1 | merged | Register & email verification shipped on main (`14834e5` finalize; review/patch commits precede it) |
| 2 | merged | Google OAuth login/register shipped on main (`efc1111` → `ce61841` explore→build→review→testing) |
| 3 | in progress | Explore logs + techplan done (`.local-agents/works/account/03-login-session-management/`); open items resolved with Anhar 2026-08-26; build not started |
| 4 | not started | |
| 5 | not started | |
| 6 | not started | Migrations for `mfa_totp_secrets`/`mfa_backup_codes` pre-created by task #3 — logic only |
| 7 | not started | |
| 8 | not started | `user_roles` migration pre-created by task #3 — logic only |

(Statuses updated 2026-08-26 from git history ground truth — the tracker
had been left at all-"not started" despite tasks #1–#2 shipping.)

## References

- Tier definitions: `kencleng-agentic-workflow.md` §4
- Tier 0 file-path fencing: `kencleng-agentic-workflow.md` §11,
  `AGENTS.md`
- Invariants referenced: `docs/spec/account/invariants.md`
- Threat model referenced: `docs/spec/account/threat-model.md`
- OpenAPI source: `api/openapi.yaml`, tag `account` (20 endpoints)