# Post-Review Remediation Manifest — 01-register-email-verification

> Ticket    : 01-register-email-verification
> Generated : 2026-08-20
> Source    : `../report.md` (four-pass code review)
> Axis      : Finding-cluster (each task fixes one review finding or a
>              tightly-coupled cluster of findings)
> Verdict   : Request changes — 7 blocking findings, 7 optional.
> Status    : Ready

---

## Context

The build (`../3-build`) shipped 24 new files + 7 modified across 5
tasks. The code review (`../report.md`) found 7 blocking findings
(S1, S2, S3, S4, S5, L1, L2) and 7 minor/optional ones. This manifest
decomposes the blocking work into focused remediation tasks plus one
bucket for the optional minors.

Tasks A–E are **blocking** — land before merge. Task F is
**optional** — can be a follow-up PR.

None of the remediation tasks touch Tier 0 fenced paths
(`platform/crypto/`, `platform/auth/`), so no new §3 fence lift is
required.

---

## Task files

| # | File | Title | Findings | Blocking | Hard deps |
|---|---|---|---|---|---|
| A | `task-a-verifyemail-atomicity.md` | `VerifyEmail` atomicity + silent-failure fix | S1, S2 | yes | none |
| B | `task-b-anti-enumeration-db-time.md` | Anti-enumeration DB-time uniformity (R3/R4) | S3 | yes | none (same file as A but disjoint methods) |
| C | `task-c-resend-error-logging.md` | `ResendVerification` handler error logging | S4 | yes | none |
| D | `task-d-migration-down-ordering.md` | Down-migration ordering fix | S5 | yes | none |
| E | `task-e-sensitive-logging-hardening.md` | Sensitive-error logging hardening | L1, L2 | yes | none |
| F | `task-f-minor-hardening.md` | Minor hardening (optional) | R15, S6, H1, Q1, Q2, Q3, E1 | no | A, B (for test touches) |

---

## Dependency graph

```
Task A (VerifyEmail atomicity)      ─┐
Task B (anti-enumeration DB-time)   ─┤  all touch service.go / repo
Task C (resend error logging)       ─┤  → coordinate in one session
Task D (migration down ordering)    │   to avoid rebase churn
Task E (sensitive logging)          ─┘   (D and E are file-disjoint)

Task F (minor) ── after A and B (test assertions depend on the new shapes)
```

**Critical path:** A + B + C + D + E (all blocking). A and B both edit
`service.go` (A on `VerifyEmail`, B on `Register` — disjoint methods);
do them in one session to avoid conflicts.

---

## Rule coverage summary

Which review findings each task resolves:

| Finding | Severity | Task | Verification |
|---|---|---|---|
| S1 silent success in VerifyEmail | blocking | A | new test: post-redeem re-fetch failure returns 500, not 200 |
| S2 non-atomic redeem+verify | blocking | A | integration test: redeem+verify in one tx; set-verified failure rolls back the redeem |
| S3 anti-enumeration DB-timing | blocking | B | timing test against real Postgres (or latency fake): R3/R4 within band of R1/R2 |
| S4 swallowed resend error | blocking | C | test: resend service error → 202 still returned + log line present |
| S5 down-migration broken | blocking | D | `make migrate-down && make migrate-up` from clean state succeeds |
| L1 breachcheck verbatim log | blocking | E | test: API-unreachable log line contains no URL / SHA-1 prefix |
| L2 notification verbatim log | blocking | E | test: send-failure log line contains no recipient/token |

---

## Open items carried forward

None new. The three original techplan open items (crypto blocker,
INV-account-08 spec, rate-limit defaults) were resolved during the
build (`3-build/report.md` §8). The crypto §3 lift and invariants §4
edit were correctly handled as per-session exceptions — not findings.

---

## Model routing

**Tier: Complex** (security-critical remediation on a Tier 1 feature).
Per the original manifest's routing: **GLM 5.2 (max)** for the
security-critical tasks (A, B, C, E); **DeepSeek V4 Pro** acceptable
for the mechanical ones (D). Mandatory dual-model review after each
task ships. If only one model is available, GLM 5.2 (max) is the safer
single choice across all tasks.

---

## Cross-reference

- Review report: `../report.md` (findings, locations, suggested fixes).
- Source techplan: `../../2-plan/techplan.md` (contract — A/B reference
  §4 R7/R8-R12 and §5 Decision 8).
- Build report: `../../3-build/report.md` (what was shipped; risk note).
- Fencing rules: `AGENTS.md` §3 (Tier 0 paths) and §4 (spec/test
  authority) apply — no remediation task edits `docs/spec/*`.
- Review checklist: `workflow/4-code-review/checklist.md` — re-run
  after remediation before considering the change done.
