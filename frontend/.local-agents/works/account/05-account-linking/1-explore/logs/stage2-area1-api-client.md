# Stage 2 — Area 1: API client + schema layer

## Current state

- **`lib/api/schema.d.ts` (generated) already has complete types for
  both endpoints** — this directly contradicts the feature spec's own
  "References" section, which says both endpoints "need a schema
  update — flagged as a follow-up." Confirmed the generation is not
  stale: `api/openapi.yaml` (root, hand-authored source) already
  defines both paths at lines 381 (`/account/security/google/unlink`)
  and 439 (`/account/security/set-password`), including the two
  distinct `409` examples for unlink (`only_identity` /
  `unverified_identity`), each with a distinguishable `type` URI
  (`https://kencleng.dev/errors/only-identity` vs
  `.../unverified-remaining-identity`) as well as different `detail`
  text.
  - `"/account/security/google/unlink"` (schema.d.ts:557-610): `post`
    takes `UnlinkGoogleRequest` (`{ password: string }`), responses
    `200` (`UnlinkGoogleResponse { message?: string }`), `401`
    (`responses["Unauthorized"]`), `409` (`application/problem+json` →
    `Problem`, undifferentiated at the type level — the two cases share
    one response type, distinguished only by body content).
  - `"/account/security/set-password"` (schema.d.ts:648-710): `post`
    takes `SetPasswordRequest` (`{ email?: string; current_password?:
    string; password: string }` — all three optional except `password`,
    no client-supplied branch discriminant field, matching the spec's
    "branch selection is server-side" rule), responses `200` (inline
    `{ message?: string }`, Branch 2 only), `202`
    (`GenericAcceptedMessage`, Branch 1 only), `401`
    (`responses["Unauthorized"]`, Branch 2 wrong `current_password`
    only), `422` (`responses["ValidationError"]` →
    `ValidationProblem`).
  - `components["schemas"]["User"].auth_providers?:
    ("email_password"|"google")[]` and `.mfa_enabled?: boolean`
    (schema.d.ts:3810-3823) — already present, already fetched via
    `useAccountMe()` (`lib/hooks/use-account-me.ts`, wraps `getMe()`
    from `lib/api/account.ts`). This is the data the frontend needs to
    decide which set-password branch's form to render (Branch 1 vs 2)
    and whether to show the "unlink Google" action at all — no new
    query needed, just a new consumer of an existing one.
- **`lib/api/account.ts` has zero code for either endpoint** — confirmed
  by grep and a full read of the file (265 lines): exports `getMe`,
  `register`, `verifyEmail`, `resendVerification`, `login`, `loginMfa`,
  `logout`, `forgotPassword`, `resetPassword`, plus the shared
  `postAccountAction` helper, `ApiError`-adjacent re-exports, and
  `RegisterResult`/`ForgotPasswordResult` discriminated-union types.
  Nothing for `setPassword`/`unlinkGoogle`.
- **Existing precedent shapes in this file, both partially reusable but
  neither an exact template**:
  1. `register()`/`forgotPassword()`'s discriminated-union
     anti-enumeration pattern (`{ok:true, message?}` vs `{ok:false,
     kind:"validation", errors}`) — single-branch functions. Set-password
     Branch 1 alone would fit this exactly (generic `202` regardless of
     conflict, `422` for policy failure). But `set-password` is **one
     endpoint serving two branches** with entirely different response
     codes/shapes per branch (200 vs 202, 401 meaning "wrong
     current_password" only in Branch 2) — no existing function in this
     file models "one endpoint, two mutually-exclusive response
     contracts selected by server-side state the client already knows
     ahead of time via `auth_providers`."
  2. `resetPassword()`'s pattern (resolve on `res.ok`, otherwise throw
     `ApiError`, caller branches on `.status`) — matches Branch 2's shape
     (200 / 401 / 422) and matches `unlinkGoogle`'s 200/401 shape, but
     neither of those needs to distinguish two sub-cases of the *same*
     status code the way unlink's `409` does.
- **Gap in `lib/api/client.ts`'s `ApiError`/`readProblemDetail`**:
  `ApiError` only carries `status` and `detail` (`client.ts:176-186`);
  `readProblemDetail` (`client.ts:194-202`) only extracts the RFC 9457
  `.detail` string, discarding `.type`. Unlink's two `409` cases are
  machine-distinguishable only via `.type` (both share `status: 409`,
  and both have distinct human-readable `.detail` text — but keying
  frontend branch logic off a translated Indonesian string is not a
  precedent used anywhere else in this codebase and isn't robust to
  backend copy changes). Nothing currently reads `.type` anywhere in
  `lib/api/`.

## Requirement

- Per the feature spec: `setPassword()` must resolve/throw distinctly
  per branch — `202` generic-message resolve (Branch 1, always, per
  anti-enumeration), `200` resolve with revoked-sessions message
  (Branch 2 success), `401` throw (Branch 2 wrong `current_password`),
  `422` throw/return validation errors (both branches, policy/breach
  failure). `unlinkGoogle()` must resolve `200`, throw/distinguish `401`
  (wrong password) from `409` (two distinct sub-cases, each needing its
  own frontend-owned copy per the spec's explicit requirement that the
  two `409` messages "distinguish... so the user knows whether they need
  to *start* set-password or just *finish* verifying").
- `frontend/AGENTS.md` §3: types come from `lib/api/` (generated), never
  hand-rolled duplicates — both endpoints' generated types already
  satisfy this; only the fetch-function wiring is missing.

## Gap

- `setPassword(input: SetPasswordRequest)` and
  `unlinkGoogle(input: UnlinkGoogleRequest)` don't exist in
  `lib/api/account.ts` — need to be added. No generated-type gap (the
  schema already has everything); this is purely a "wire up the fetch
  function" gap, same category as Task #4's finding.
- `setPassword()`'s two-branch response contract has no existing
  same-file precedent to copy verbatim — closest analogues (`register`/
  `forgotPassword`'s single-branch discriminated union, `resetPassword`'s
  status-branch-only shape) each cover half of what's needed here. This
  is a genuine design point for Stage 3, not just a copy-paste.
- `unlinkGoogle()`'s two-`409`-cases-same-status-code need is not
  representable with the current `ApiError(status, detail)` shape alone
  — either `ApiError`/`readProblemDetail` needs to also carry `.type`
  (a `client.ts` change, which is shared infra other domains' future
  problem-detail-branching needs could also benefit from), or
  `unlinkGoogle()` does its own one-off problem-body parsing bypassing
  `readProblemDetail`. Not resolving which here — flagging as the
  concrete design point for Stage 3.

## Page-consolidation check

- N/A at this layer (API client has no page-level consolidation
  question) — the relevant consolidation question is MFA-vs-linking on
  `/dashboard/security`, deferred to Area 3 where the actual page lives.
- No mismatch found between `page-map.md`'s action list and
  `tasks.md`'s endpoint list for *this* task specifically: both list
  exactly `POST /account/security/google/unlink` and
  `POST /account/security/set-password`, both present and fully typed
  in `schema.d.ts`/`api/openapi.yaml`. No orphaned page-map action
  without a backing endpoint, no orphaned endpoint without a page-map
  mention.

## Sniffing

- **Inconsistency (spec doc vs. actual repo state)**: the feature spec's
  "References" section states `api/openapi.yaml` "both need a schema
  update — flagged as a follow-up, alongside the other pending
  `openapi.yaml` changes from earlier feature specs." This is stale —
  the root `api/openapi.yaml` already has both endpoints fully defined,
  including the two differentiated `409` examples, and `schema.d.ts` is
  already regenerated from it. Worth flagging to a human rather than
  silently trusting the spec doc's claim, since a reader might
  reasonably (and wrongly) conclude backend/OpenAPI work is still
  pending here.
- **Misleading signal**: `schema.d.ts` having full types could look like
  "the frontend API layer for this feature is already done" if someone
  only checks the generated file — `account.ts` (the actually-consumed
  layer) has nothing, same shape of false signal noted in Task #4's
  equivalent area.
- **Risk**: the `readProblemDetail`/`ApiError` gap (missing `.type`) is
  shared infra (`client.ts`), not scoped to this feature alone — a
  change there is low-blast-radius (additive, existing callers only
  read `.status`/`.detail`, unaffected) but is exactly the kind of
  "frontend has no business logic" boundary line worth being precise
  about in Stage 3: exposing `.type` is still pure plumbing/transport
  concern, not business logic, so it stays in-bounds for `frontend/`.
- **Edge case**: `SetPasswordRequest.email` is optional at the type
  level (Branch 2 doesn't send it) but required in practice for Branch
  1 — the generated type can't express "required iff Branch 1" via
  OpenAPI's flat schema, so client-side `zod` validation (per
  `form-validation-boundary.md` conventions used elsewhere in this repo)
  will need to enforce this conditionally based on which branch's form
  is rendered, not rely on the generated type to catch a missing
  `email` before the request goes out.
- **Edge case**: nothing in `User` (`auth_providers`) tells the frontend
  whether an existing `email_password` identity is *verified* — by
  design (per INV-account-12, that state is only surfaced via the
  distinct `409` unlink response). This means the "Google-only user
  mid-way through the 3-step flow" state (submitted Branch 1, hasn't
  verified yet) looks identical to "fully Google-only" from
  `auth_providers` alone — the UI cannot pre-emptively show "you're
  waiting on verification" copy without either attempting unlink first
  or checking `email_verified` (which *is* on `User` and reflects
  exactly this: "True if at least one `email_password` AuthIdentity has
  `verified_at` set"). Worth using `email_verified` in Stage 3 to avoid
  a needless failed-unlink round trip for this specific state, rather
  than only reacting to the 409.

Proceeding to Area 2 (hooks layer).
