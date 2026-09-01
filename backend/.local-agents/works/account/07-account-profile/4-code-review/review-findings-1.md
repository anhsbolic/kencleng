# Review Findings — account #07 (initial build), round 1

> Reviewed: the diff behind `3-build/report.md` — `internal/domain/account/profile.go`,
> `internal/domain/account/service_test.go` (R7 additions), `internal/transport/http/account_profile.go`
> + `_test.go`, `internal/transport/http/auth_login.go` + `_test.go` (D10 changes), `cmd/server/main.go`
> (route registration). Passes run in order: Safety → Quality → Stack-Specific → Consistency.

## 1. Safety

No findings.

Verified explicitly (not skipped):
- **Nil dereference** — the `(nil, nil)` gone-user path is handled before mapping
  (`account_profile.go:88-93`), and `toUserResponse` is itself nil-safe
  (`account_profile.go:45-47`), so the handler cannot panic even if the ordering changed. R4 pins it.
- **Nil slice serialization** — `Roles`/`AuthProviders` normalized to `[]string{}` before the
  JSON boundary (`account_profile.go:48-55`); R1 asserts they decode as `[]any`, not `null`.
- **Concurrency** — the only shared mutable state added is `fakeRepo.getViewCalls`
  (`service_test.go:352-355`), mutated under the existing `f.mu` lock. No goroutines added.
- **Error propagation** — `GetProfile` errors flow through `MapServiceError` (`account_profile.go:83-87`);
  unknown errors become generic 500 with server-side logging (`errors.go:130-136`), detail not leaked (R5).
- **Context propagation** — `r.Context()` passed through handler → service → repo unchanged.
- **Resources** — none acquired in this slice (no tx, no external call, no file handle).
- **Techplan-drift cross-check** (guidelines §1): the areas Safety would flag (decrypted-PII
  response surface, JSON contract shape) are both explicitly recorded in techplan §12's Test
  Focus Pointer as N/A with reasons, and the slice is a Tier 2 pure read with no concurrency —
  the pointer table is consistent with the code. No drift finding.

## 2. Quality

1. **Finding (minor, optional):** the gone-user 401 branch is silent — no log line distinguishes
   "session referenced a deleted user" from a token the middleware rejected.
   **Location:** `internal/transport/http/account_profile.go:88-93`.
   **Why it matters:** this is a rare decision point; when someone debugs "I have a valid token
   but get 401", server logs currently show nothing at all. A single line logging `userID`
   (a UUID — not PII, and not the view/response, so R6's no-email guarantee is unaffected)
   would make that diagnosable.
   **Suggested fix:** `log.Printf("transport: /account/me session user not found: %s", userID)`
   in the nil-view branch. Optional — Tier 2 read, zero clients, acceptable to defer.

2. **Finding (minor, optional):** the 401 `WriteProblem` triple
   (`errors/unauthorized` / "Unauthorized" / "Authentication required.") now appears twice in
   `AccountMeHandler` (lines 77-80 and 89-92), plus once each in `RequireSession`
   (`account_security.go:44-47, 51-54`).
   **Location:** `internal/transport/http/account_profile.go:77-80, 89-92`.
   **Why it matters:** a future wording/type change must be applied in 5 places.
   **Suggested fix:** a small `writeUnauthorized(w)` helper in `errors.go`. **Not blocking:** the
   inline triple *is* the existing convention in this package (established by
   `account_security.go`), so this slice is consistent; extraction would touch pre-existing
   files and belongs in its own cleanup slice, not this one (root AGENTS §7 scope discipline).

## 3. Stack-Specific Best Practices

Keyword match against `best-practices/index.md`: this diff's technology is Go + REST API.
Files opened and applied: `go/nil-and-zero-values.md`, `go/authorization-and-idor.md`,
`restapi/openapi-spec-first-drift.md`, `go/secrets-and-sensitive-logging.md`. (No Kafka/Pub-Sub/
Redis/GraphQL/React surface is touched; `restapi/pagination-and-status-codes.md` was already
applied at techplan time via D5 and re-checked here — status classes are correct.)

1. **Finding (minor, optional)** — from `restapi/openapi-spec-first-drift.md` (checklist item 3:
   happy-path *and* edge shape validated against the spec): `loginOKResponse.User` keeps
   `json:"user,omitempty"` (`auth_login.go:41`), but openapi `LoginResponse` lists `user` under
   `required`. On the only reachable path this is fine (login always assembles the view before
   `issueSessionTokens`, `login.go:144/241`), so `omitempty` is effectively dead. But if a future
   refactor ever reaches the `"ok"` writer with a nil `res.User`, the required `user` key
   silently vanishes from the wire instead of failing loudly.
   **Location:** `internal/transport/http/auth_login.go:37-42`.
   **Why it matters:** silent spec drift on an impossible-today path — exactly the class the
   drift checklist warns about, and cheap to make loud now.
   **Suggested fix:** either drop `omitempty` and make the nil-User-on-"ok" case impossible to
   write silently (e.g. a debug log or a test asserting `user` is always present on ok bodies),
   or leave as-is with a comment noting the assumption. Optional — no current code path hits it.

Checklist results from the opened files (all satisfied, no findings):
- `go/nil-and-zero-values.md` §2 — nil→`[]` normalization is owned at the transport boundary
  (D7) and asserted by type (`[]any`) in R1, not just presence. Satisfied.
- `go/authorization-and-idor.md` — no resource ID accepted from input; the resource is keyed
  entirely by `UserIDFromContext` (`account_profile.go:75`), and R3
  (`TestAccountMe_NoIDParameter_SessionScoped`) is the dedicated IDOR test proving foreign
  identifiers in query/path/body are inert. Satisfied.
- `go/secrets-and-sensitive-logging.md` — the handler never logs the view or response; R6
  proves the decrypted email never reaches captured logs. The `MapServiceError` default logs
  the raw error, but this repo's `GetLoginUserView` errors wrap only SQL-build/scan/decrypt
  failures whose text contains no PII or parameter values. Satisfied.

## 4. Consistency

No findings. Checked against `backend/README.md` + `backend/AGENTS.md` (§2 style, §4 local
commands) and the surrounding transport/domain code:

- **Seam convention** — `profileService` one-method interface in the handler file matches the
  per-file seam pattern (`loginSessionService` `auth_login.go:17`, `securityService`
  `account_security.go:66`). ✓
- **Error handling convention** — reuse of `MapServiceError` (no new sentinel added, per
  techplan §11 "Files NOT changed") and the byte-identical `WriteProblem` triple already
  established at `account_security.go:44-47`. ✓
- **Doc comments** — every exported item (`GetProfile`, `AccountMeHandler`) has one; the
  unexported `userResponse`/`toUserResponse` are documented too, matching the file's neighbors. ✓
- **Naming** — `stubProfileService`/`fullView`/`assertUserShape`/`getAccountMe` follow the
  existing test-helper naming style (`stubLoginService`, `postJSON`, `findCookie`); `fullView`
  is shared across the two test files exactly as the shape-divergence guard intends. ✓
- **Route wiring** — the `main.go:196-202` registration mirrors the `/account/security/` mount
  (`RateLimit` ∘ `RequireSession`), with an explanatory comment in house style. ✓
- **Validation convention** — no request input exists to validate; nothing skipped that applies.

One pre-flagged item verified rather than re-raised: the report's faithful-reconciliation note
(techplan §10 sketches `toUserResponse` returning a value while also specifying a
`*userResponse` field assigned from it — contradictory). The implemented `*userResponse` return
is the only reading that compiles against both call sites and preserves the `omitempty`
omission semantics; the wire contract is unchanged. Correctly documented in the report, no
action needed.

## Verdict

**Approve with minor comments.**

- Blocking: none.
- Optional (all deferrable, none required for this slice):
  1. §2.1 — add a log line (userID only) to the gone-user 401 branch for debuggability.
  2. §2.2 — consolidate the 401 `WriteProblem` triple into a helper (separate cleanup slice;
     would touch pre-existing files).
  3. §3.1 — make a nil-`User`-on-"ok" login response loud (comment, log, or drop `omitempty`).

No patch-plan file written: no finding requires a code change in this slice.
