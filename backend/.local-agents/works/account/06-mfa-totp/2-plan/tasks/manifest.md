# Manifest — 06-mfa-totp task decomposition

> Generated : 2026-08-27
> Source    : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar)
> Note      : snapshot at generation time — this file does NOT track progress status; status belongs to the PR/ticket domain.

## Task list

| # | File | Title |
|---|---|---|
| 1 | `task-01-dependency-and-contract-artifacts.md` | `pquerna/otp` dependency + openapi source 409 amendment (D10) + mechanical bundle regen |
| 2 | `task-02-repository-foundation.md` | Six new MFA repository operations + `MFABackupCode` entity; guarded SQL per §8 contract |
| 3 | `task-03-service-layer.md` | `MfaEnroll` / `MfaEnrollConfirm` / `MfaDisable`, sentinels, audit constants, backup-code material helpers |
| 4 | `task-04-tier0-verifier-draft-for-pairing.md` | Tier 0 core: `totpVerifier` replacing stub — DRAFT-FOR-PAIRING, hard STOP at the pairing checkpoint (D12) |
| 5 | `task-05-transport-and-wiring.md` | Three handlers, `ConsumeReauthMarker`, sentinel→Problem mappings, `main.go` wiring + routes |
| 6 | `task-06-integration-and-race-gate.md` | testcontainers DB-truth suite, ≥100-goroutine `-race` harnesses (R4/R8/R9), full `make verify` gate |

## Splitting axis

**Dependency/sequence chain** (layer-aligned). Rationale: every stage waits on the previous at compile time (repo ports before services before handlers), mirroring task #05's decomposition shape for a same-tier slice in the same package. A component/module axis is impossible (one domain package); risk separation IS expressed but as boundaries inside the chain — most notably task-04's independent existence exists precisely to make the D12 Tier 0 pairing checkpoint an enforceable stop rather than prose. No parallelism is claimed or intended.

## Dependency graph

```
01 ──► 02 ──► 03 ──► 04 ──► 05 ──► 06
                          ▲
        (pairing sign-off │ after 04's suites go green;
         gates MERGE, not scaffold progress)
```

Fully serial by construction:
- 02 needs nothing from code (pure additive layer) but starts the Go chain
- 03 needs 01 (otp import) + 02 (ports) and owns the shared material helpers
- 04 needs 02 (redemption port) + 03 (normalize/hash helpers); its suites double as the pairing harness
- 05 needs 03 (service seam) + 04-scaffold (constructor name); final merge ordering still respects the pairing gate
- 06 runs against the assembled slice

## Model routing per task

Overall tier: **Complex** (16 rules ≥15 threshold; TOTP-crypto/auth surface → mandatory dual-model code review). Per-task picks from `best-practices/model-routing.md` Coding/Complex row ("per sub-task"):

| Task | Model | Why |
|---|---|---|
| 01 | DeepSeek V4 Flash | Purely mechanical (yaml add + bundler + dep pin); objective verification. Escalation trigger documented in-file (bundler-drops-components STOP condition). |
| 02 | DeepSeek V4 Pro | Rule-table-heavy precision: conflict-arm predicate encoding, parameterization golden rule, nullable scans |
| 03 | GLM 5.2 (max) | Multi-step branching + guarded-tx ordering + timing-parity reasoning (GLM's strength; reward-hacking caveat covered by mandatory dual-model review) |
| 04 | GLM 5.2 (max) | Highest-risk reasoning in the slice — but output is human-paired regardless (D12), which is itself the compensating control |
| 05 | DeepSeek V4 Pro | Contract-shape precision: status codes, problem-type URIs, marker atomicity semantics |
| 06 | DeepSeek V4 Pro | Testing/Complex row verbatim: rule-ID count-check reliability once rules ≥15 |

Cross-task non-negotiables from model-routing: **code review must run GLM 5.2 (max) + DeepSeek V4 Pro in parallel with manual diffing** (no exception at this tier), and any Claude fallback follows that doc's mapping table.

## Back-reference

Originating contract techplan: `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` — "Tech Plan: MFA TOTP Enrollment, Confirmation & Disable (account #06)". Every task file carries this back-reference; executors cross-check decisions there (D1–D12, R1–R16, and the §8 SQL/flow blocks are authoritative over any paraphrase in task files).

## Review note

These derived task files remain subject to `workflow/4-code-review/checklist.md` like all derived-section content.
