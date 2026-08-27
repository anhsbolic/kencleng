# Stage 2 — Area 4: Mocks (`mocks/handlers.ts`)

## Current state

`mocks/handlers.ts` (205 lines, read in full) has no handlers for any of
the 3 MFA endpoints — confirmed directly, not just by the earlier grep.
`GET /account/me` returns a single default fixture, `mockUser`, which
already has `mfa_enabled: false` set explicitly (added ahead of this
task, presumably when the `User` schema field was generated) — so the
default logged-in-donatur fixture is already MFA-not-enrolled, matching
the "just the default case" convention this file's own comments
describe (Dashboard Shell nav checkpoint precedent).

Established per-endpoint pattern, consistent across every prior task
(register, forgot/reset-password, set-password/unlink-google, login):

1. A `mock<Thing><Branch>` constant holding the exact success-branch
   response body, typed against the real `lib/api/account.ts` type
   where one exists (e.g. `const mockLoginOk: LoginResponse = ...`).
2. A comment block above the constant(s) naming the source task
   (`account/0N-...`) and stating the copy is the backend's own real
   text (not invented), with a pointer to the techplan decision that
   confirmed it (e.g. "see the techplan's D4").
3. One `http.post(...)` entry per endpoint in the `handlers` array,
   returning the default/happy-path fixture.
4. A stated convention that individual tests override non-default
   branches (`422`/`401`/`409`/network-error/a different `/account/me`
   identity shape) via `server.use(...)` inline in the test file itself
   — no alternate-branch fixtures live in this shared file.

## Requirement

Default (happy-path) handlers for all 3 MFA endpoints, so `MfaSection`
and its children have something to render against in dev mode and in
component tests before individual tests override specific branches.

## Gap

- `POST /account/security/mfa/enroll` — needs a default `200` handler
  returning a fixture `MfaEnrollResponse` (an `otpauth://` URI string;
  no real secret/issuer decision has been made yet — Stage 3 territory,
  but the mock just needs *a* well-formed URI, not a real one).
- `POST /account/security/mfa/enroll/confirm` — needs a default `200`
  handler returning a fixture `MfaEnrollConfirmResponse` (10 backup
  code strings, per the spec's "exactly 10 backup codes" requirement —
  the fixture array length itself is a small but real correctness
  detail worth getting right, since a test asserting "10 codes shown"
  would otherwise silently test against a wrong-length fixture).
- `POST /account/security/mfa/disable` — needs a default `200` handler
  returning the message fixture (real backend text already known:
  `"MFA berhasil dinonaktifkan."`, quoted directly in `schema.d.ts`'s
  own `@example`, so — unlike some earlier tasks — there's no
  "TBD"/placeholder-copy uncertainty here at all for the success
  branch).
- No fixture yet for an **MFA-enabled** `mockUser` variant
  (`mfa_enabled: true`) — needed for any test exercising the "already
  enrolled" state of `MfaSection` (e.g. the disable flow, or the `409`
  re-enroll-while-enabled guard). Individual tests can override this
  inline the same way Task 05's tests override `auth_providers` for a
  Google-only user (per this file's own stated convention) — not
  necessarily a new shared constant, but worth deciding explicitly in
  Stage 3 rather than each test re-deriving the override ad hoc.

## Page-consolidation check

N/A for this area.

## Sniffing

- **Edge case**: the backup-codes fixture array length matters for
  test correctness (spec says exactly 10) — an easy place to
  accidentally under/over-populate a fixture and have a test silently
  validate the wrong invariant. Flagging so Stage 3/implementation
  doesn't treat "any array of strings" as good enough.
- **Consistency**: this file's convention of sourcing "verbatim backend
  text" for success messages holds here too — `schema.d.ts`'s own
  `@example` already gives the exact disable-success string, so unlike
  `SetPasswordForm`'s `GENERIC_ERROR_MESSAGE` (marked `// TBD`
  everywhere), there's no placeholder-copy ambiguity for at least this
  one string.
- **Miscontext risk (carried from Area 2)**: since `schema.d.ts` doesn't
  type a `409` response for `enroll`, there's no existing convention in
  this file for mocking an *undocumented-in-schema* error branch — every
  prior override-branch precedent (422/401/429) had a typed schema
  response to mock against. Whoever adds the `409` test-override case
  will be typing against the feature spec's prose, not the generated
  schema — worth being explicit about that provenance rather than
  treating it as equally "generated-schema-backed" as the others.
- No new inconsistency or risk beyond what Areas 1–2 already surfaced.
</content>
