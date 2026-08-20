# Task Manifest — 01-register-email-verification

> Ticket    : 01-register-email-verification
> Generated : 2026-08-19
> Axis      : Dependency/sequence chain (primary) + component/layer alignment
> Back-ref  : `../2-plan/techplan.md` (originating contract techplan — source of all decisions redistributed into the task files below)
> Status    : Snapshot at generation time. This manifest does NOT track
> progress (done/in-progress/blocked) — status tracking belongs to the
> PR/ticket domain, not this workspace.

---

## Splitting axis

**Dependency/sequence chain (primary)**, with **component/module boundary**
and **layer (vertical slice)** as the natural alignment on top of it.

**Rationale.** Techplan §9 already lays out a real dependency chain —
entity → repository interface → repository impl → service → handlers →
wiring — and the component boundaries (migrations → platform → domain →
transport → wiring) map cleanly onto it. This is the least-assumption
axis: the dependency chain is stated in the techplan itself, so
splitting along it carries no extra judgment about which parts are
"sensitive" (which risk/blast-radius would require). The
security-critical work (constant-time R7, concurrent token redemption
R12/INV-account-08, concurrent duplicate registration R16/INV-account-01)
is isolated in Task 04 so it gets focused review and a `-race` run,
matching AGENTS.md §3's Tier 1 requirements.

---

## Prerequisite (HUMAN — Tier 0 fenced, NOT an agent task)

**`platform/crypto/` encrypt/decrypt/HMAC functions.**

`backend/internal/platform/crypto/` is file-path-fenced per AGENTS.md §3.
An agent must not author these functions. They are a hard prerequisite
for Task 03 (the repository cannot encrypt `primary_email` / `identifier`
for insertion without them). A human session must produce:

- `crypto.Encrypt(plaintext []byte, key *crypto.Key) ([]byte, error)` — AES-GCM
- `crypto.Decrypt(ciphertext []byte, key *crypto.Key) ([]byte, error)`
- `crypto.HMAC(data []byte, key *crypto.Key) string` — HMAC-SHA256 hex

This prerequisite is **not** tracked as a task file (it cannot be
executed by an agent). It is recorded here so the dependency graph is
honest: Task 03 cannot start until this is done. See techplan §5
Decision 1, §7 risk row 1, §14 Open Item #1.

---

## Task files

| # | File | Title | Hard deps |
|---|---|---|---|
| 01 | `task-01-deps-and-migrations.md` | Dependencies + 3 migrations | none |
| 02 | `task-02-platform-packages.md` | secrets, breachcheck, notification | Task 01 (for `golang.org/x/crypto` direct) |
| 03 | `task-03-domain-data-layer.md` | account: entity + repository (interface + impl) | **Tier 0 crypto prerequisite**, Task 01 |
| 04 | `task-04-domain-service-and-tests.md` | account: service + tests (R1-R19, -race) | Task 02, Task 03 |
| 05 | `task-05-transport-and-wiring.md` | handlers, middleware, errors, main.go | Task 04, Task 01 (for `golang.org/x/time/rate`) |

---

## Dependency graph

```
[Prereq: Tier 0 crypto (HUMAN, fenced — NOT an agent task)]
        │
        v
Task 01 (deps + migrations) ──────────────────────────┐
        │                                              │
        ├──→ Task 02 (platform packages)              │
        │            │                                 │
        │            ├──→ Task 03 (domain data layer) ← needs crypto + Task 01
        │            │            │                    │
        │            │            v                    │
        │            ├──→ Task 04 (domain service + tests)
        │            │            │                    │
        │            │            v                    │
        │            └──→ Task 05 (transport + wiring) ← also needs Task 01 (rate dep)
        │                                             │
        └─────────────────────────────────────────────┘
```

**Critical path:**
`Tier 0 crypto (human) + Task 01 → Task 03 → Task 04 → Task 05`

**Parallel opportunity:**
Task 02 can run concurrently with Task 01 and Task 03's early work
(entity + repository interface are schema-shape-only and don't need
the platform packages or crypto to be drafted; only `repository_db.go`
needs crypto). In practice, sequence them in the order listed unless
parallelism is explicitly coordinated by a human.

---

## Rule coverage summary

End-user rules R1-R19 are satisfied across the tasks as follows.
R-numbers reference techplan §4.

| Rule | Primary task | Notes |
|---|---|---|
| R1 (register new) | 04 | service orchestrates; 03 provides inserts; 05 writes 202 |
| R2 (register unverified existing) | 04 | resend branch inside register |
| R3 (register verified existing) | 04 | password-reset nudge |
| R4 (register Google-only conflict) | 04 | Google-only nudge |
| R5 (password validation order) | 04 | service checks before branch lookup; 05 boundary-checks length |
| R6 (breach check fail-open) | 02 (client) + 04 (use) | 02 provides fail-open; 04 consumes |
| R7 (constant-time anti-enumeration) | 04 | always-bcrypt + DB-write-shaped work; 05 writes identical 202 |
| R8 (verify valid token) | 04 | 03 provides RedeemToken + SetVerifiedAt; 05 writes 200 |
| R9 (verify expired) | 04 | ErrTokenExpired; 05 writes 410 |
| R10 (verify not found / already used) | 04 | ErrTokenNotFound; 05 writes 404 |
| R11 (verify revoked) | 03 (3-clause guard) + 04 (disambiguate) | storage-level guard lives in 03's RedeemToken |
| R12 (verify concurrent double-submit) | 03 (atomic UPDATE) + 04 (test) | storage-level single-use in 03; orchestration+test in 04 |
| R13 (resend unverified match) | 04 | 03 provides RevokeTokens + InsertAuthToken |
| R14 (resend no match / verified / google-only) | 04 | no token, no email, identical 202 |
| R15 (rate limit) | 05 | middleware + 429; `TestResend_RateLimited` |
| R16 (concurrent duplicate registration) | 01 (unique index) + 03 (error mapping) + 04 (test) | schema-level uniqueness in 01; clean rollback in 03/04 |
| R17 (Google-only generic response) | 04 | branch equivalence test |
| R18 (password policy 422) | 04 (service) + 05 (handler field errors) | 05 surfaces field-level detail |
| R19 (breach check fail-open test) | 04 | `TestRegister_BreachCheck_FailOpen` |

---

## Open items carried forward from techplan §14

These remain unresolved at the task level and must be surfaced by the
executing agent in their PR risk note (or resolved by a human before
the relevant task ships):

1. **Tier 0 crypto prerequisite** — see "Prerequisite" above. Hard
   blocker for Task 03.
2. **INV-account-08 verification description inconsistency** — the
   invariant's Verification field omits `revoked_at IS NULL`. Task 03
   and Task 04 implement the 3-clause Statement version and treat the
   2-clause Verification field as a documented spec error. The agent
   must NOT edit `docs/spec/domains/account/invariants.md` (AGENTS.md
   §4) — flag it for a human to fix. Task 03's
   `TestVerifyEmail_RevokedToken_Rejected` and Task 04's test of the
   same name are the regression guards.
3. **Rate limit RPS/burst values for `/auth/*`** — Task 05's
   middleware is configurable; default values are TBD. The wiring
   reads from env and fails fast if unset rather than silently
   disabling the limiter. Human resolves the concrete defaults.

---

## Model routing

**Tier: Complex** (19 rules ≥ 15 threshold + touches auth/PII — both
trigger Complex per the routing table).

**Stage: Coding / build (decomposed)** — the applicable row is
"Decomposed: GLM 5.2 (max) / DeepSeek V4 Pro per sub-task." Per-task
recommendations below apply the table's GLM↔DeepSeek tiebreaker:
GLM when the work leans on diagrams/state-transitions/multi-step
reasoning; DeepSeek V4 Pro when it's rule-table-heavy/precision work
without a diagram.

Source: `../../harscode-workspace/best-practices/model-routing.md`
(DRAFT — personal reference, not yet workspace guidance; re-check
pricing/benchmarks before trusting beyond a few weeks old).

| Task | Recommended model | Rationale |
|---|---|---|
| 01 — deps + migrations | **DeepSeek V4 Pro** | Pure precision work (SQL DDL, `go.mod`). No diagrams, no multi-step reasoning. DeepSeek ties Claude on SWE-bench Verified — closest to "resolve a real issue in a real codebase." |
| 02 — platform packages | **GLM 5.2 (max)** | Security-relevant concerns (explicit timeout, fail-open, no PII in logs). GLM has demonstrated signal on security-pattern detection. DeepSeek V4 Pro acceptable but GLM edges it here. |
| 03 — domain data layer | **GLM 5.2 (max)** | Multi-step reasoning around the INV-account-08 3-clause guard + entity write-path design (who encrypts). The kind of "diagrams/state-transitions" work the table routes to GLM. |
| 04 — service + tests | **GLM 5.2 (max)** for `service.go`; **DeepSeek V4 Pro** for the test suite | Highest-risk task. Build half leans GLM (constant-time R7, 4-branch state logic, concurrent double-submit R12). Test half leans DeepSeek V4 Pro per the Testing Complex row: "rule-ID count-check needs more reliability than Flash once rule count ≥15" — 19 rules to map 1:1 to tests. |
| 05 — transport + wiring | **DeepSeek V4 Pro** | Mostly mechanical (handlers, middleware, wiring) with security checklists (no internal leakage, idle-key eviction). Precision work. GLM acceptable. |

**Mandatory downstream step — Code review (Complex tier, non-negotiable
dual-model row):**
GLM 5.2 (max) + DeepSeek V4 Pro parallel (mandatory), diff manually —
no exception at this tier. After any task ships, run both models on
the diff and diff their outputs. Per the routing table, DeepSeek is a
*supplement* to GLM here, not a substitute (DeepSeek has a known
weakness on cybersecurity-specific evals) — don't run DeepSeek alone
for review.

**Single-model fallback:** if only one model is available, **GLM 5.2
(max)** is the safer single-model choice across all five tasks because
Task 04's security-criticality outweighs DeepSeek's precision edge on
the mechanical tasks.

**Caveats (from the routing doc itself):**
- The routing doc is explicitly **DRAFT — personal reference, not yet
  workspace guidance**. Treat recommendations as hypotheses, not
  settled policy.
- All benchmark figures cited when building that table are
  vendor-reported unless otherwise noted — treat as hypotheses to
  validate against real stories.
- Pricing moves fast — re-check before assuming today's cost math
  holds.

---

## Cross-reference

- Source techplan: `../2-plan/techplan.md` (contract/derived authoring
  process precedes this decomposition — see
  `../../harscode-workspace/workflow/techplan-synthesis-prompt.md`).
- Review checklist: task files produced here are still subject to
  `../..//harscode-workspace/workflow/code-review/checklist.md`, same
  as any other derived-section content.
- Fencing rules: AGENTS.md §3 (Tier 0 paths) and §4 (spec/test
  authority) apply to every task in this manifest.
