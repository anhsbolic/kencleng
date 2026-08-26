# Stage 2 — Area 5: Mocks / test conventions

## Current state

- `mocks/handlers.ts` (168 lines total): confirmed via grep, zero
  `forgot`/`reset` handlers exist. Established convention for this
  domain (lines 100-167, read in full):
  - One `mock*Accepted`/`mock*Ok` fixture constant per endpoint's
    happy-path response, defined near the top with a comment
    explaining which techplan/spec task it belongs to and which
    branches individual tests override via `server.use(...)`.
  - Default handlers registered in the `handlers` array are always the
    happy path (e.g. `http.post("/auth/register", () =>
    HttpResponse.json(mockRegisterAccepted, { status: 202 }))`) —
    per-test failure branches are added via `server.use()` inside the
    specific test, never as a second default handler.
  - Message copy in the fixtures matches the actual copy shown in the
    components exactly (e.g. `mockRegisterAccepted.message` is the
    literal string `RegisterForm`'s `DEFAULT_SUCCESS_MESSAGE` constant
    also contains) — fixtures and component copy are kept in sync by
    convention, not by shared import.
- Test-file conventions, confirmed by reading `register-form.test.tsx`
  (189 lines) and `verify-email-status.test.tsx` (140 lines, read in
  full) end-to-end:
  - `verify-email-status.test.tsx` is the **exact template** for
    `/reset-password`'s eventual test file: `vi.mock("next/navigation",
    () => ({ useSearchParams: () => ({ get: (key) => ... } ) }))` with
    a mutable `mockToken` variable reassigned per test; a
    `withQueryClient` helper wrapping `render()`; `server.use(...)`
    overriding the endpoint's response per test case (410 → expired
    message + resend affordance, 404 → generic invalid message, 429 →
    verbatim backend detail, network failure via `HttpResponse.error()`
    → generic frontend-owned message); an explicit test for the
    missing-token case asserting the **same** generic message as the
    404/invalid case (this is the test that actually proves R11, cited
    in Area 4).
  - `register-form.test.tsx` (line 71) has the field-level-422 template:
    "maps each 422 field error verbatim via setError, no banner (R5)" —
    posts a `422` with a `Problem`-shaped body plus an `errors: [...]`
    array, then asserts the specific field's rendered error message,
    and implicitly (per the test name) that no banner rendered — this
    is the exact test shape `/reset-password`'s new-password
    field-error case needs.
  - Every test file consistently names its `it()` blocks with a
    trailing rule-number tag (e.g. "(R5)", "(R11)") tracing back to the
    originating techplan/spec rule — an established codebase convention
    for traceability, not just this file's style.

## Requirement

- `frontend/AGENTS.md` §4: component/unit tests via `vitest`.
  `component-test-mocking-discipline` skill governs MSW/mocking depth
  (not yet consulted — Stage 3 concern, flagging its existence here
  since this area is exactly its subject matter).
- Feature spec's acceptance criteria (all bullet points under both
  endpoints) are the source of truth for which branches need a test —
  every named branch (202 registered/Google-only/no-match all
  identical, 429, 200, 404, 410, 422, concurrent double-submit) should
  map to a rule-tagged test the way `verify-email-status.test.tsx`
  already does for its own endpoint's branches.

## Gap

- `mocks/handlers.ts` needs two new default happy-path handlers
  (`/auth/forgot-password` → 202 generic accepted, `/auth/reset-password`
  → 200 generic success), plus fixture constants following the
  existing naming/comment convention. Purely additive.
- No test files exist yet for `ForgotPasswordForm`/`ResetPasswordForm`
  (components don't exist per Area 4) — once built, they should follow
  `register-form.test.tsx`'s shape for the forgot-password side and
  `verify-email-status.test.tsx`'s shape (token mock + `server.use`
  branch overrides + rule-tagged `it()` names) for the reset-password
  side, per the precedents above.
- One acceptance criterion from the spec has **no existing precedent
  test shape anywhere in this domain to mirror**: `/auth/reset-password`'s
  concurrent double-submit case ("the same valid token submitted twice
  concurrently ... exactly one request succeeds ... the other gets
  `404`"). Every existing frontend test in this domain mocks one
  request/response pair per test; a concurrency race is fundamentally a
  *backend* invariant (guarded `UPDATE`, INV-account-08) that the
  frontend cannot itself create or verify — the frontend's only
  testable obligation here is "renders whatever the backend returns for
  a given single response," which the existing 404-handling test shape
  already covers once written. Flagging this as a scope question for
  Stage 3 (or confirmation that no new test pattern is needed here),
  not a gap in the mocking convention itself.

## Sniffing

- **Misleading signal**: none.
- **Miscontext**: none — the feature spec doesn't claim anything about
  frontend test coverage for the concurrency case specifically; the
  spec's `TestResetPassword_TokenSingleUse_Concurrent` reference is
  explicitly a *backend* test name (`docs/spec/account/threat-model.md`
  component 3 / feature spec's Threat breakdown table), not a frontend
  one — no mismatch, just worth being explicit that this doesn't
  translate into a new frontend test obligation.
- **Risk**: low — this area is purely additive and has strong, complete
  precedent for every real frontend-testable branch.
- **Edge case**: none beyond what Areas 2/4 already flagged (429-on-reset
  question, missing-vs-invalid-token question) — this area doesn't
  introduce new ones, just needs test coverage for however those get
  resolved in Stage 3.
- **Inconsistency**: none found.

Proceeding to Area 6 (visual/prototype precedent).
