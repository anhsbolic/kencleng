# Task Manifest — 02-google-oauth-login-register

> Ticket    : 02-google-oauth-login-register
> Generated : 2026-08-24
> Axis      : Dependency/sequence chain (primary) + component/layer alignment
> Back-ref  : `../techplan.md` — "Tech Plan: Google OAuth Login/Register" (originating contract techplan — source of all decisions redistributed into the task files below)
> Status    : Snapshot at generation time. This manifest does NOT track
> progress (done/in-progress/blocked) — status tracking belongs to the
> PR/ticket domain, not this workspace.
> Input status: techplan treated as locked per human confirmation in this
> session (header field still reads "Draft"; §14 records six decisions made
> by human review, 2026-08-22).

---

## Splitting axis

**Dependency/sequence chain (primary)**, with **component/module boundary**
and **layer (vertical slice)** as the natural alignment on top of it.

**Rationale.** Techplan §9 lays out a real dependency chain — migrations →
entities → repositories → platform client → service → cookies/handlers →
wiring — and the component boundaries map cleanly onto it. This is the
least-assumption axis: the sequence is stated in the techplan itself, so
splitting along it carries no extra judgment about which parts are
"sensitive" (which risk/blast-radius would require). The security-critical
work concentrates where it gets focused review: token forgery / JWKS /
replay surface in Task 02, no-auto-merge + verified_at state machine +
concurrent-duplicate handling in Task 03 (with a mandatory `-race` run,
backend/AGENTS.md §3), conditional authz + inline JWT verification in
Task 04.

Deviation from ticket 01's layout: the domain *data* layer here is thin
(2 structs, 2 inserts — no atomic redemption guard like INV-account-08), so
it merged into the service task rather than standing alone. Confirmed by
human in this session (4 tasks, not 5).

---

## Prerequisites

None. Ticket 01's Tier 0 crypto prerequisite (`platform/crypto/`) already
shipped; this ticket consumes it read-only. `platform/auth/` and
`platform/crypto/` remain fenced for every task below (root AGENTS.md §3):
Task 03 consumes `auth.Keys` read-only; Task 04 verifies JWTs inline rather
than modifying `platform/auth/`.

## Task files

| # | File | Title | Hard deps |
|---|---|---|---|
| 01 | `task-01-deps-and-migrations.md` | jwt dependency + migrations 000004/000005 | none |
| 02 | `task-02-platform-googleoauth.md` | googleoauth platform package (AuthURL, ExchangeCode, VerifyIDToken, JWKS cache) | Task 01 |
| 03 | `task-03-domain-data-and-service.md` | account: entities, repo inserts, GoogleRedirect/GoogleCallback/IssueTokens + tests (-race) | Tasks 01 + 02 |
| 04 | `task-04-transport-and-wiring.md` | cookies, handlers, inline ES256 verify, reauth marker store, main.go | Task 03 (+ Task 01 dep) |

## Dependency graph

```
Task 01 (deps + migrations)
        │
        v
Task 02 (googleoauth platform client)
        │
        v
Task 03 (domain data layer + OAuth service + tests)
        │
        v
Task 04 (transport + wiring)
```

**Critical path:** `Task 01 → Task 02 → Task 03 → Task 04`

Strictly linear — unlike ticket 01 there is no parallel opportunity worth
coordinating: every task consumes its predecessor's output directly
(Task 02 imports the dep from 01; Task 03 holds the client from 02 and
writes to the tables from 01; Task 04 calls the methods from 03).

---

## Rule coverage summary

End-user rules R1–R26 are satisfied across the tasks as follows.
R-numbers reference techplan §4. "Primary" = the task that owns the rule's
core logic and its test.

| Rule | Primary | Supporting | Notes |
|---|---|---|---|
| R1 (login redirect, no auth) | 04 | 03 | 03 generates state/nonce/cookie payload; 04 skips auth + writes cookie + 302 |
| R2 (link/reauth redirect w/o session → 401 pre-redirect) | 04 | — | `TestGoogleRedirect_LinkReauthRequireAuth`; explicit authz check at handler boundary |
| R3 (session user_id into cookie) | 04 | 03 | cookie payload shape from 03's flow |
| R4 (state mismatch → state_mismatch, no Google call) | 03 | 04 | constant-time compare in service; zero-call assertion; 302 written by handler |
| R5 (nonce mismatch → nonce_mismatch) | 02 | 03, 04 | ErrNonceMismatch sentinel in VerifyIDToken; mapping in service; redirect by handler |
| R6 (Google timeout → google_unavailable, not raw 500) | 02 | 03 | 10s timeout client in package; result mapping in service |
| R7 (login + existing google identity → tokens) | 03 | 04 | IssueTokens path; Set-Cookie half in 04 |
| R8 (login + new user → User+AuthIdentity+tokens) | 03 | 04 | single tx; verified_at=now |
| R9 (no-auto-merge on email_password match) | 03 | — | `TestGoogleCallback_NoAutoMerge_Login`; top-severity threat |
| R10 (link conflict → google_link_conflict) | 03 | — | `TestGoogleCallback_NoAutoMerge_Link` |
| R11 (link success → attach + audit log) | 03 | 01 | user_logs row action_type=account_linking (Fitur 9); table created in 01 |
| R12 (reauth → 5-min marker, no side effects) | 03 | 04 | no-op guarantees tested in 03; sync.Map store set/checked in 04 |
| R13 (fixed redirect_uri from env) | 02 | 04 | only configured URI reaches token endpoint; env wiring in main.go |
| R14 (google identities verified_at=now) | 03 | — | asserted across all creation branches |
| R15 (concurrent duplicate registration clean failure) | 03 | 01 | unique index on auth_identities pre-exists (migration 000002, verified in 01's checklist); race test + clean error in 03 |
| R16 (no secrets/tokens in logs) | 03 | 02, 04 | log-capture assertions in both 03 and 04; sanitized categories in 02 |
| R17 (Google client errors sanitized before logging) | 02 | — | category strings only ("timeout", "http error") |
| R18 (invalid intent → 400) | 04 | 03 | validation logic in GoogleRedirect; boundary write in handler |
| R19 (missing code/state → state_mismatch, no call) | 04 | 03 | param presence check first; zero-call assertion via stub |
| R20 (expired/missing cookie → state_mismatch) | 04 | 03 | cookie read/clear in handler |
| R21 (JWKS refresh-on-miss) | 02 | — | cache TTL 15 min + one refetch + retry |
| R22 (explicit http.Client timeout, NewRequestWithContext) | 02 | 04 | never http.DefaultClient |
| R23 (constant-time compare) | 03 | 02 | state compare in service; nonce compare in VerifyIDToken |
| R24 (cookie attributes HttpOnly/Secure/Lax/MaxAge=600/Path) | 04 | — | exact attribute assertions in handler tests |
| R25 (inline ES256 verify; platform/auth untouched) | 04 | — | public key injected as handler dependency; fenced path untouched |
| R26 (bad sig/iss/aud/exp → google_token_invalid; only nonce → nonce_mismatch) | 02 | 03 | strict verification in package; error-code mapping in service |

---

## Open items carried forward from techplan §14

Active items that must be surfaced by the executing agent in their PR risk
note (or resolved by a human before the relevant work ships):

1. **Caddy routing for `/auth/*`** (techplan Active #1) — dev uses direct
   `:8090` GOOGLE_REDIRECT_URI (already in .env.example); non-dev requires a
   Caddyfile fix, which is a **root-level session**, outside backend/
   boundary. Affects Task 04's manual smoke testing only, not code. Do not
   edit the Caddyfile from these tasks.
2. **Error-code set v2 + `account_linking` action_type literal — frontend
   sign-off pending** (techplan Active #2) — backend pins both as specified
   in techplan §8/§14; divergence found later is a flag-for-human item, not
   a local edit. Relevant to Tasks 03 (audit literal) and 04 (?error={code}
   contract).
3. **Deferred-by-design scope** (not gaps, recorded so reviewers don't flag
   them): refresh_tokens rotation indexes/constraints → task #3;
   user_logs REVOKE constraint + full vocabulary → task #08; reauth-marker
   consumption → task #06; unlink/set-password → task #05.

---

## Model routing

**Tier: Complex** (26 rules ≥ 15 threshold + touches auth — both trigger
Complex per the routing table).

**Stage: Coding / build (decomposed)** — the applicable row is "Decomposed:
GLM 5.2 (max) / DeepSeek V4 Pro per sub-task." Per-task recommendations
apply the table's GLM↔DeepSeek tiebreaker: GLM when the work leans on
diagrams/state-transitions/multi-step reasoning; DeepSeek V4 Pro when it's
rule-table-heavy/precision work without a diagram. The Testing Complex row
(DeepSeek V4 Pro once rule-ID count ≥ 15) governs test-suite authoring.

Source: `/home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/model-routing.md`
(DRAFT — personal reference, not yet workspace guidance; re-check
pricing/benchmarks before trusting beyond a few weeks old).

| Task | Recommended model | Rationale |
|---|---|---|
| 01 — deps + migrations | **DeepSeek V4 Pro** | Pure precision work (SQL DDL, go.mod). No diagrams, no multi-step reasoning. DeepSeek ties Claude on SWE-bench Verified — closest benchmark to "resolve a real issue in a real codebase." |
| 02 — googleoauth platform package | **GLM 5.2 (max)** | Security-relevant throughout: RS256-only enforcement, iss/aud/exp checks, JWKS rotation, replay-vs-forgery error split. GLM has demonstrated signal on security-pattern detection (matches ticket 01's platform-package routing). |
| 03 — domain data + service | **GLM 5.2 (max)** for the build; **DeepSeek V4 Pro** for the test suite | Build half: intent-branched state machine + no-auto-merge + tx orchestration — exactly the diagrams/state-transitions work routed to GLM, and the highest-risk task of the ticket. Test half: 26-rule ID-to-test mapping needs the Testing-row reliability floor (DeepSeek V4 Pro, count ≥ 15). |
| 04 — transport + wiring | **GLM 5.2 (max)** for handlers; **DeepSeek V4 Pro** for wiring + tests | Handlers carry conditional-auth multi-step reasoning (R2/R18/R25 inline verify) — GLM-leaning; main.go wiring + mechanical handler tests are precision work — DeepSeek-leaning. |

**Mandatory downstream step — Code review (Complex tier, non-negotiable
dual-model row):**
GLM 5.2 (max) + DeepSeek V4 Pro parallel (mandatory), diff manually — no
exception at this tier. After any task ships, run both models on the diff
and diff their outputs. Per the routing doc, DeepSeek is a *supplement* to
GLM here, not a substitute (known weakness on cybersecurity-specific evals)
— don't run DeepSeek alone for review.

**Single-model fallback:** if only one model is available, **GLM 5.2 (max)**
is the safer choice across all four tasks — Tasks 02/03/04 are all
security-critical, which outweighs DeepSeek's precision edge on Task 01.

**Caveats (from the routing doc itself):**
- The routing doc is explicitly DRAFT — treat recommendations as hypotheses,
  not settled policy.
- All benchmark figures are vendor-reported unless noted; independent
  reviews have found vendor/independent gaps for at least one listed model.
- Pricing moves fast — re-check before assuming cost math still holds.

---

## Cross-reference

- Source techplan: `../techplan.md` (contract/derived authoring process
  precedes this decomposition — see
  `harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`;
  decomposition process itself:
  `harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md`).
- Review checklist: task files produced here are still subject to
  `harscode-workspace/workflow/4-code-review/checklist.md`, same as any
  other derived-section content.
- Fencing rules: root AGENTS.md §3 (Tier 0 paths — platform/auth/,
  platform/crypto/) and §4 (spec/test authority) apply to every task in
  this manifest. Neither fenced path is modified by any task above.
