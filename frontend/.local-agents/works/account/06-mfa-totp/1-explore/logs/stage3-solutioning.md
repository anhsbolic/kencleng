# Stage 3 — Solutioning: MFA TOTP (frontend surface)

> Builds on Stage 2's 5 area reports (`stage2-area1..5-*.md`). Confirmed
> accurate by Anhar before this stage started. This is trade-off/option
> comparison + rationale — not yet a file-by-file techplan (that's a
> separate synthesis step per `workflow/1-exploration/guidelines.md`).

## D1 — `MfaSection` composition

**Options:**
- A. One flat component, internal `useState` step machine
  (enroll/confirm/codes/disable) handling everything inline.
- B. Parent `MfaSection` (reads `useAccountMe()`, branches on
  `mfa_enabled` like `LoginMethodsSection` branches on
  `auth_providers`) delegating to `MfaEnrollFlow` (not-enrolled branch)
  or `MfaDisableForm` (enrolled branch), each its own file.

**Recommend B.** Matches the established one-component-per-action
convention (`SetPasswordForm`, `GoogleIdentityControl`/
`UnlinkGoogleForm`) and page.tsx's own stated philosophy ("independent
section components, not a monolithic form," D1 from Task 05) applied
one level down. Keeps each piece independently testable, same payoff
Task 05 got from splitting `GoogleIdentityControl` out.

## D2 — When does `enroll` actually fire?

**Decision:** Never auto-fire on mount. `MfaEnrollFlow` renders an
"Aktifkan MFA" button; the `useMfaEnroll` mutation only fires on click.
No resume-in-progress-enrollment UX is built — Stage 2 Area 1 already
found there's no reliable signal to resume from (`mfa_enabled: false`
can't distinguish never-started vs. pending-unconfirmed). If a user
abandons mid-flow and comes back, they see the initial button again;
clicking it again is safe per the spec's explicit overwrite-pending-
secret semantics. Avoids the anti-pattern of firing a mutation as a
side effect of rendering, which has no precedent anywhere in this
codebase.

## D3 — QR rendering library (new dependency) **[CONFIRMED by Anhar]**

**Options considered:**
- A. `qrcode.react`'s `QRCodeSVG` — SVG output, simple declarative
  `value`/`size`/`level` props, actively maintained, zero runtime deps.
- B. `react-qr-code` — also SVG-based, simpler API surface, less
  actively maintained.
- C. `qrcode` (imperative canvas/data-URI API) — would need a canvas
  polyfill in the jsdom test environment (this project's pinned test
  runtime, per `client.ts`'s own doc comment about `navigator.locks`
  absence in jsdom — canvas support is a similar class of gap), and is
  a heavier, non-declarative API to wrap.

**Decided: A — `qrcode.react`'s `QRCodeSVG`.** SVG output needs no
jsdom polyfill (matches how `lucide-react` icons already render as
inline SVG in this app, so there's already a working precedent for
SVG-in-jsdom), and the declarative props fit this codebase's component
style directly. This is the **first new dependency** this domain's
frontend track has needed.

**Testing approach:** wrap the library call in a thin `QrCode`
presentational component; component tests assert the wrapper received
the correct `otpauth_uri` value (e.g. by mocking the `qrcode.react`
module per `component-test-mocking-discipline.md`'s guidance), not by
inspecting rendered SVG path data.

## D4 — Backup codes reveal-once UX

**Decision:** After `enrollConfirm` succeeds, render all 10 codes as
plain text (not through `MaskedField` — that component is for PII
display/reveal-toggle, not a one-time-secret-reveal; different
purpose) plus an explicit "Saya sudah menyimpan kode ini" confirmation
action the user must take before the section proceeds to the enrolled
state. This directly answers Area 1's flagged risk (codes are gone for
good once this view unmounts, per spec) with at least an acknowledgment
gate — it cannot literally prevent a user from navigating away early
(no `beforeunload` interception proposed; that's over-engineering for
what this gate is trying to solve). Copy is new, frontend-owned (no
backend text exists for a client-only gate) — marked `// TBD` per this
codebase's existing placeholder-copy convention.

## D5 — Disable flow, `email_password` branch

**Decision:** Mirrors `UnlinkGoogleForm` almost exactly — single
password field, destructive submit button, `ApiError.detail` shown
verbatim on `401` (the schema's `MfaDisableRequest` 401 is generic/
undifferentiated, per Area 2 — there's only one message to show, not a
branch to distinguish). On `200`, invalidate `accountKeys.me()` only —
no session-clearing/redirect (unlike `SetPasswordForm`'s
`mode="change"` branch): nothing in the spec calls for MFA-disable to
revoke the current session.

## D6 — Disable flow, Google-only branch **[CONFIRMED by Anhar — Option B]**

This is the area's central unresolved gap (Stage 2 Area 1's main
finding): neither `docs/spec/account/features/06-mfa-totp.md` nor
`api/openapi.yaml` pin down the redirect target or query-param contract
for `GET /auth/google/callback` when `intent=reauth` — the schema only
says (verbatim) "always ends in a redirect to a frontend route (success
or error state indicated via query param)," with no route name or
param name specified anywhere. Task 05 never built handling for any
callback-redirect query param either (it relies entirely on refetching
`/account/me` afterward, which works for `link` because linking is
inspectable from account state — the `reauth` marker is not part of
`User` at all, so that trick doesn't carry over).

**Options:**
- A. **Two-click flow keyed off a redirect-return query param.**
  `MfaDisableForm`'s Google-only branch shows `<GoogleAuthButton
  intent="reauth" .../>`. On return to `/dashboard/security`, read a
  query param (e.g. `?reauth=ok`/`?reauth=error`) via
  `useSearchParams()` to decide whether to show an armed "Nonaktifkan
  MFA" button or an error banner. **Requires a backend contract that
  isn't confirmed anywhere in scope for this frontend track** — the
  redirect target/param name would need to be pinned down (by Anhar or
  a backend-track decision) before this could be built correctly.
- B. **Optimistic single button, rely on the documented `401`.** Always
  render one "Nonaktifkan MFA" button for Google-only users. Clicking
  it calls `mfaDisable({})` (no body) directly. If the marker is
  missing/expired, the already-documented `401` fires, and the error
  banner adds a `<GoogleAuthButton intent="reauth" .../>` prompt;
  after reauthenticating and returning, the user clicks disable again
  — same button, now succeeds. Needs **no unconfirmed contract at
  all** — only the `401` response already typed in `schema.d.ts`.

**Decided: B.** It uses only what's already confirmed in the generated
schema, degrades gracefully (worst case: one extra click-and-retry
round trip after reauth), and matches `UnlinkGoogleForm`'s own existing
philosophy of "always render the action, let the backend's response
drive the copy" rather than trying to pre-empt the backend's state
client-side. Recorded here as an explicit assumption (per
`docs/spec/README.md`'s rule that ambiguity must be recorded, not
silently resolved) — A remains available later if the backend redirect
contract gets pinned down and a smoother single-click flow becomes
worth building.

## D7 — `enroll`'s undocumented `409`

**Decision:** Handle it defensively in the wrapper function
(`mfaEnroll()` checks `res.status === 409` explicitly, throws
`ApiError`) even though `schema.d.ts` only types `200` for this
endpoint (Area 2's flagged schema/spec inconsistency). In normal
single-tab use this becomes unreachable in practice, since
`MfaEnrollFlow` (the "Aktifkan MFA" trigger) only ever renders in the
not-enrolled branch — same client-side guard-mirroring reasoning
`LoginMethodsSection`/`GoogleIdentityControl` already use for their own
blocked-state derivations. The `409` path exists only as a genuine
defensive fallback (e.g. a second tab, a stale cache) — not a reachable
UI path this task needs to build a dedicated view for beyond a generic
error banner.

## D8 — `enrollConfirm`'s `422` shape

**Decision:** Despite `schema.d.ts` typing it as `ValidationError`
(`{errors: [{field, message}]}`, same shape as `register`/
`setPassword`'s discriminated-result branches), **do not** model this
as a discriminated result. The feature spec is explicit: "treated the
same as an invalid code — no distinguishing response needed." Follow
`resetPassword()`'s precedent instead (Task 04, D2: confirmed the real
`422` carries no useful field-level data worth branching on) — throw
`ApiError` on `422`, and the confirm form shows one fixed, frontend-
owned message ("Kode tidak valid, coba lagi." — `// TBD`) regardless of
whatever `errors[].field`/`.message` the backend actually sends. Simpler
than the register/setPassword shape, and matches the spec's own stated
intent more directly than mirroring the schema's generic validation
shape would.

## D9 — Mocks

Add, following the existing per-endpoint convention exactly
(`mock<Thing>` constant + comment naming the source task + one default
`http.post` handler):

- `mockMfaEnrollResponse: MfaEnrollResponse` — a well-formed but fake
  `otpauth://` URI (no real secret ever needs to be valid for a mock).
- `mockMfaEnrollConfirmResponse: MfaEnrollConfirmResponse` — **exactly
  10** backup-code strings (Area 4's flagged correctness detail — get
  the count right so a "10 codes shown" test isn't silently validating
  a wrong-length fixture).
- `mockMfaDisableOk = { message: "MFA berhasil dinonaktifkan." }` —
  already known verbatim from `schema.d.ts`'s own `@example`, no
  placeholder-copy uncertainty here.
- No new shared "MFA-enabled user" fixture constant needed as a rule —
  individual tests override `mfa_enabled: true` via `server.use()` on
  `/account/me`, same per-test inline convention `LoginMethodsSection`'s
  own test file already uses for `auth_providers`/`email_verified`.

## D10 — "Regenerate backup codes" wording (`page-map.md` miscontext)

**Decision:** Do not build a one-click "regenerate" action — none
exists on the backend (Area 1's miscontext finding: the spec is
explicit the only path is disable → re-enroll). `MfaDisableForm`'s
enrolled-state view instead carries a short explanatory line near the
disable action (e.g. "Untuk mendapatkan kode cadangan baru,
nonaktifkan MFA lalu aktifkan kembali.") so the UI doesn't imply a
capability that doesn't exist. This is a copy/framing choice within
this task's own scope, not a new endpoint — `page-map.md`'s wording
itself is a docs-accuracy question outside this frontend code change,
worth mentioning to Anhar separately but not blocking this task.

## Confirmation status

1. **D6** — **CONFIRMED by Anhar**: Option B (optimistic single button
   + `401`-driven reauth prompt) for the Google-only disable flow.
2. **D3** — **CONFIRMED by Anhar**: `qrcode.react` as the QR rendering
   dependency (first new package for this domain's frontend track).
3. New frontend-owned copy (D4's "Saya sudah menyimpan kode ini" gate,
   D10's disable-then-re-enroll explanatory line) stays `// TBD`,
   same placeholder-copy convention used everywhere else in this
   codebase pending final copy — not a blocking decision.

All decisions (D1–D10) are now settled. Ready for implementation
techplan synthesis.
</content>
