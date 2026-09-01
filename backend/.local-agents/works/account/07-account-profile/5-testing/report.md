# Testing Report — account #07 (GET /account/me)

> Phase: 5-testing — after implementation + code review (review-findings-1.md, verdict
> "Approve with minor comments", no code changes required)
> Gate commands: `backend/README.md` § Testing / § Verification (`make verify` stages)

## Step 0 — Sweep (spot-check, don't redo)

**Named-test spot checks** — every rule the build report claims coverage for was re-run, not
rewritten. All 8 pass on re-run (`-run` targeted, `-count=1`):

| Claim | Test | Result |
|---|---|---|
| R7 | `TestGetProfile_PassesThroughToRepository` | PASS |
| R1–R6 | `TestAccountMe_*` (5 tests + 3 R3 subtests) | PASS |
| R8 | `TestLoginHandler_UserShapeSnakeCase`, `TestLoginMfaHandler_UserShapeSnakeCase` | PASS |

**Build report's own gap statement** ("no `-race`, perf/load, or security-class test was run
in this iteration") — treated as this phase's priority list, addressed under Test Focus Pointer
and the gate runs below.

**Test Focus Pointer (techplan §12)** — both rows marked N/A. Cross-checked against the raw
exploration logs' Risk-lens findings: areas 1–4 record "Risk: None" (pure read, pass-through,
mechanical wiring); area 5's one real finding (JSON camelCase drift) is exactly the N/A row's
"contract-correctness class, covered by R1/R8 unit tests" disposition. **Verdict: no relevant
pointer rows → no separate Test Execution Plan required.** The only non-"None" underlying risk
(the torn four-query read, `repository_db.go:679`) is recorded as an *accepted* Low risk in
techplan §7 — documented, not silent, so no techplan-drift flag. Race coverage for the new code
is still obtained: the scoped-per-guideline race run below includes both touched packages and
the new mutex-guarded `fakeRepo.getViewCalls`. No bcrypt/KDF or perf-sensitive primitive is
exercised by this slice (Tier 2 pure read — no threshold to set).

## Step 1 — Rule coverage (techplan §4 R1–R8)

| Rule | Proven by | Verified this phase |
|---|---|---|
| R1 success shape | `TestAccountMe_Success_UserShapeSnakeCase` | re-run PASS |
| R2 auth required (spec-named) | `TestAccountMe_RequiresAuth` + existing `TestRequireSession_*` | re-run PASS (full suites) |
| R3 session-scoped (spec-named) | `TestAccountMe_NoIDParameter_SessionScoped` | re-run PASS |
| R4 gone user → 401 | `TestAccountMe_UserGone_SessionInvalidated` | re-run PASS |
| R5 internal error → generic 500 | `TestAccountMe_InternalError_Generic500` | re-run PASS |
| R6 no PII in logs | `TestAccountMe_LogsFreeOfPII` | re-run PASS |
| R7 pass-through contract | `TestGetProfile_PassesThroughToRepository` | re-run PASS |
| R8 login/mfa user shape (D10) | both `*UserShapeSnakeCase` tests | re-run PASS |

Count-check: R1–R8 each map to ≥1 confirmed test. Nothing needed re-derivation.

### What only this phase could do — real-interface exercise (added)

No test exercised the actual **route registration** (`cmd/server/main.go:199-202` — the mux
pattern + `RateLimit` ∘ `RequireSession` composition): handler tests invoke
`AccountMeHandler` directly, so a wiring mistake (wrong pattern, missing middleware) would pass
every existing test. Added `internal/transport/http/account_profile_wire_test.go`, which
rebuilds the exact registration shape behind a real `httptest.Server` with real ES256
verification (`GoogleTokenVerifier` + `newES256Signer`) and issues real HTTP requests:

- `TestAccountMeWire_Success_BearerToken` — valid bearer over the wire: 200, exact shape,
  service received the session userID, `Content-Type: application/json`
- `TestAccountMeWire_NoToken_Middleware401` — 401 at the middleware, service never called
- `TestAccountMeWire_InvalidTokens_401` — garbage / wrong-key / expired (15 min past, beyond
  the 1 min leeway) / non-Bearer scheme → all 401
- `TestAccountMeWire_GoneUser_401ByteIdenticalToNoToken` — R4/D5's "same generic body as a
  missing token" proven **byte-for-byte on the wire**, across the two different emitters
  (middleware vs handler)
- `TestAccountMeWire_WrongMethod_405` — POST to the GET-only pattern → 405, service not called

### Four categories

- **Happy path**: R1 + wire success (both confirmed above).
- **Negative**: R2/R5 + wire invalid-token matrix + wrong method (added this phase).
- **Edge**: R3 (foreign identifiers), R4 (gone user), and — **gap found & closed this phase** —
  the nil→`[]` normalization branch in `toUserResponse` (`account_profile.go:48-55`) had zero
  coverage: `fullView` and `seedView` both set non-nil slices, and the repository initializes
  both non-nil, so D7's core guarantee was only tested for the already-normalized case. Added
  `TestAccountMeWire_NilSlices_NormalizedToEmptyArrays` (nil `Roles`/`AuthProviders` in the
  view → `[]`, never `null`, on the wire). This matches the known pattern "Nil/empty dependency
  data not handled" from `workflow/5-testing/examples.md` — existing entry, no new one needed.
- **Backward compatibility**: D10's login wire-shape change — R8 proves both success writers
  emit the new snake_case shape, and the pre-existing login tests (asserting `user` presence,
  cookie flags, mfa-required shape) all still pass in the full suite. Zero clients existed at
  disposition (techplan §6, Anhar 2026-09-01) — no old-client surface to preserve. Existing
  data: `GetLoginUserView` integration tests unchanged and untouched by this slice.

### Error verification (precise, not "an error happened")

- Gone user → 401 with problem type `errors/unauthorized`, **byte-identical** to the
  middleware's no-token body (wire test) — not 404/500, no account-state leak.
- Internal failure → 500 `problems/internal` with generic detail; R5 asserts the wrapped error
  text ("db exploded") does not reach the client, and the server-side log line was observed
  working during the re-run (logged by `MapServiceError` default).
- No token / invalid token / expired / wrong key → 401 `errors/unauthorized` at the middleware
  layer before the handler (wire tests assert status + problem type + no service call).
- Wrong method → mux 405 (route-level, distinct from auth 401s).

## Final verification — repo's own gate commands

| Stage | Result |
|---|---|
| `go build ./...`, `go vet ./...`, `gofmt -l internal/` | clean |
| `go test ./...` (unit) | ok — all packages |
| `go test -race ./...` (full repo, includes both touched packages) | ok |
| `go test -tags=contract ./...` | ok |
| `staticcheck ./...` | **failed at HEAD before this phase** (2 pre-existing unused test symbols) — fixed, see Findings |
| `gosec ./...` | **13 pre-existing findings, all present at HEAD**, none in this slice's files — flagged, not fixed |
| `gitleaks detect` | **could not run — `gitleaks` binary not installed** in this environment |
| `govulncheck ./...` | 24 vulnerabilities affecting code (2 modules + stdlib) — pre-existing, repo-wide dependency posture, unrelated to this slice |

## Findings

1. **Coverage gap (closed in this phase):** D7's nil→`[]` guarantee was unproven for its actual
   trigger condition (nil input); `TestAccountMeWire_NilSlices_NormalizedToEmptyArrays` closes it.
2. **Real-interface gap (closed in this phase):** route/middleware wiring had zero coverage;
   `account_profile_wire_test.go` (7 tests) closes it.
3. **Pre-existing gate failure (partially fixed, rest flagged — needs a human/root-level
   cleanup slice):** `make verify` could not pass at HEAD independent of this task.
   - staticcheck U1000 ×2 (dead test symbols, pre-existing at HEAD — verified via
     `git stash` + re-run): removed `fakeRepo.seedMFADisabled`
     (`service_test.go:546-555`) and `stubSecurityService.reauthResult`
     (`account_security_test.go:43`). Pure dead-code deletion — no assertion loosened
     (root AGENTS §4 not violated).
   - gosec 13 findings (G124 cookie attrs, G401/G505 sha1, G304 key files, G112
     ReadHeaderTimeout, G710/G706 taint, G101) — all pre-date this slice; they span
     `platform/auth` (Tier 0 fenced) and `cookie.go`/`googleoauth`. Per root AGENTS §3/§7 these
     are **not** fixable from this session — flagging for a dedicated security-hygiene slice.
   - `gitleaks` not installed — the `security` stage cannot run here at all; root-level tooling gap.
4. **Techplan §8 minor inaccuracy (no action):** it calls the read surface "migrations
   000001-000008 territory", but `000009_user_roles` / `000010_widen_auth_tokens_purpose`
   landed 2026-08-26/27, before the techplan (2026-09-01). No collision either way: this slice
   adds no migration, and nothing landed since the techplan touched the schema.

## Verdict

**Testing complete — slice verified.** R1–R8 all confirmed via re-run; the two gaps only this
phase could see (route wiring, nil-slice edge) are closed with new passing tests; full unit,
race, and contract gates are green. The remaining `make verify` failure is entirely
pre-existing (gosec findings in Tier 0/touch-later files + missing `gitleaks` binary) and is
flagged for a separate cleanup slice, not silently patched.
