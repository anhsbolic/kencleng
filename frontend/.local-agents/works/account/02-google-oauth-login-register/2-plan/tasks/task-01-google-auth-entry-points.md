# Task 1: Google Auth Entry Points (`GoogleAuthButton`, `/login`, error banner)

> Originating contract techplan: `../techplan.md` ("Tech Plan: Google
> OAuth Login/Register (Frontend)", account/02-google-oauth-login-
> register, Status: Draft). Cross-check high-level decisions there
> whenever this task file is ambiguous — this file redistributes, not
> replaces, that document's detail.
>
> Splitting axis: **Component/module boundary** (see `../manifest.md`
> for the full rationale). This task owns the OAuth-initiation UI
> surface — the Google button and `/login`'s error display. It has no
> import/code dependency on Task 2 (session bootstrap) and can be
> built, reviewed, and tested independently of it.

## Scope

Fix the already-broken Google button on `/register`, extract it into a
shared component, and build `/login`'s Google entry point + its
`?error={code}` banner. Covers rules **R1–R7** and Decision Log
entries **D1, D2, D4, D5** from the originating techplan.

**Out of scope for this task**: anything about how the in-memory
access token gets populated after a successful redirect (Task 2's
scope — `AuthBootstrapProvider`, `client.ts`'s refresh export/type
fix, `app/layout.tsx`'s provider wiring, the `mocks/handlers.ts`
refresh handler). `/login`'s credential form (backend task #3,
out of scope for the whole techplan, not just this task).

## Background (condensed from techplan §1)

`RegisterForm` already has a "Daftar dengan Google" entry point — a
real `<a href="/auth/google/redirect?intent=register">` — but the
already-merged backend only accepts `intent` ∈ `{login, link, reauth}`
(`validIntent()` in `backend/internal/domain/account/google_oauth.go:
117-120`); `intent=register` gets a `400`. `/login`
(`app/(auth)/login/page.tsx`) is still the raw Phase 0 placeholder —
no Google button, no error handling, and (per its own comment) no
credential form either, since that's a separate task's scope.

## What this task builds (from techplan §10)

- **File**: `components/features/account/google-auth-button.tsx`
  (new) — shared, typed entry-point component. Props: `intent: "login"
  | "link" | "reauth"` (typed against the same literal union
  `schema.d.ts` generates for `/auth/google/redirect`'s query param —
  this task only ever passes `"login"`), `label: string`. Renders a
  real `<a href={"/auth/google/redirect?intent=" + intent}>` — never
  `apiFetch`/XHR.
- **File**: `components/features/account/register-form.tsx` (modify)
  — replace the existing inline, broken anchor with
  `<GoogleAuthButton intent="login" label="Daftar dengan Google" />`.
- **File**: `components/features/account/google-callback-error.tsx`
  (new) — Client Component. Reads the `error` search param via
  `useSearchParams()` (needs a `<Suspense>` boundary at the call site,
  same requirement `VerifyEmailStatus`/`app/verify-email/page.tsx`
  already established in this codebase), maps it per R6, renders
  `<Banner variant="error">` as `AuthShellClient`'s documented first
  child, moves focus into it on render (R7). Renders `null` when no
  `error` param is present.
- **File**: `app/(auth)/login/page.tsx` (modify) — replace the Phase 0
  placeholder with: heading, an explicit "email/password login coming
  soon" static note (task #3's scope — do not build a form here),
  `<Suspense>`-wrapped `GoogleCallbackError`, and
  `<GoogleAuthButton intent="login" label="Masuk dengan Google" />`.
  Continues to render inside the existing `AuthShellClient` —
  unmodified, reused as-is.

## Rules & Validation owned by this task

(Numbering matches the originating techplan §4 — not renumbered per
task.)

- **R1** (correct intent value): Given the "Masuk dengan Google" or
  "Daftar dengan Google" button, When clicked, Then it navigates to
  `/auth/google/redirect?intent=login` — never `intent=register`, on
  either page. The `login` intent already covers net-new-account
  creation server-side (`google_oauth.go`'s `callbackLogin`, middle
  branch: no existing identity, email not claimed elsewhere → creates
  a new `User`) — no distinct frontend "register intent" exists or is
  needed.
- **R2** (shared, typed component): Given both pages need this button,
  When implemented, Then both consume the one `GoogleAuthButton`
  component, not independently duplicated markup. `intent` is typed
  against the schema's own literal union — a value outside
  `"login" | "link" | "reauth"` is a TypeScript compile error, not a
  runtime `400`. This is exactly the class of bug R1 fixes; typing it
  out prevents recurrence.
- **R3** (real navigation only): Given the button renders, Then it's a
  real `<a href>` (or equivalent), never `apiFetch`/XHR — the endpoint
  issues a `302`, which only a real navigation follows correctly (task
  #1's own R7 precedent, carried forward here as this component's
  contract).
- **R4** (`/login` scope boundary): Given `/login` loads, When
  rendered, Then it shows: a heading, an explicit "email/password
  login coming soon" placeholder note (task #3's scope, not built
  here), the `GoogleAuthButton`, and the error-banner slot (R5) — no
  credential input fields.
- **R5** (error banner presence): Given `/login` loads with
  `?error={code}` present, Then render an error banner; given the
  param is absent, Then render nothing.
- **R6** (error code → copy mapping): Given `code === "google_email_
  conflict"`, Then show distinguishable copy indicating the email is
  already registered via password login (pointing at the existing
  login path, not encouraging a retry of the same failing action).
  Given `code` is one of `state_mismatch`/`nonce_mismatch`/
  `google_token_invalid`/`google_unavailable`, Then show one shared
  generic retry message. Given `code` is anything else/unrecognized,
  Then show the same generic fallback. The raw `code` value is never
  rendered to the user.
- **R7** (banner placement + focus): Given the error banner renders,
  Then it is `AuthShellClient`'s documented first child
  (`<Banner variant="error">`), and focus moves into it (or its
  containing heading) on render — matching the focus-management
  convention already established by `RegisterForm`/`VerifyEmailStatus`.

## Decision Log entries relevant to this task

**D1 — Fix the `intent=register` bug**

| Option | Why rejected/accepted |
|---|---|
| A. `intent=login` (**chosen**) | Confirmed from the merged backend: `login` intent's three branches already include "no existing identity, email free → create `User`" — exactly what "register via Google" needs. No missing backend capability, only a wrong query value. |
| B. Add a fourth `register` value to the backend's enum | Rejected — out of this frontend task's authority; backend task #2 is merged, changing its accepted-value contract needs its own cross-track review, and A makes it unnecessary. |

**D2 — `/login` page scope**

| Option | Why rejected/accepted |
|---|---|
| A. Build only what this task owns: Google button, error banner, explicit "coming soon" note (**chosen**) | Mirrors task #1's own precedent. Keeps root `AGENTS.md`'s honesty principle intact — the page visibly says what's missing rather than looking finished or looking unchanged. |
| B. Build the entire page including a stub credential form | Rejected — would write UI against `POST /auth/login`, an endpoint this task has no contract for; risks duplicating/conflicting with task #3's actual build. |
| C. Leave `/login` as the Phase 0 placeholder entirely | Rejected — leaves `/login` incomplete against its `page-map.md` row for no technical reason, since the Google button is fully buildable today. |

**D4 — Error banner copy per code**

| Option | Why rejected/accepted |
|---|---|
| A. One generic message for every code | Rejected — `google_email_conflict` (the no-auto-merge case) is the one branch where a fully generic message actively misleads a legitimate user into retrying the exact same failing action; spec 02 treats it as the top-severity anti-takeover threat in the whole feature. |
| B. Distinguish `google_email_conflict`; collapse the other four into one shared fallback (**chosen**) | Matches this codebase's own precedent (`/verify-email`'s 410-vs-404 distinction, task #1 D5) — keep backend-distinguished, actionable outcomes distinguishable. None of the five codes are enumeration-sensitive, so nothing is leaked by differentiating. |

**D5 — Shared component vs. duplicated markup**

| Option | Why rejected/accepted |
|---|---|
| A. Duplicate the anchor markup a second time for `/login` | Rejected — directly reproduces the class of bug found live in `/register`'s existing code (an independently hand-copied `intent` value). |
| B. Extract a shared, typed `GoogleAuthButton` (**chosen**) | Cheap (one anchor's `href`/label differ, no real behavior divergence to abstract prematurely); makes `intent` a typed prop against the schema's own literal union, turning the exact bug class in D1 into a compile-time error for any future caller. |

## Interface Contract (subset relevant to this task)

```typescript
// components/features/account/google-auth-button.tsx
type GoogleAuthButtonProps = {
  intent: "login" | "link" | "reauth"; // R2 — typed against schema.d.ts's own literal union; this task only ever passes "login"
  label: string; // e.g. "Masuk dengan Google" / "Daftar dengan Google"
};
function GoogleAuthButton(props: GoogleAuthButtonProps): JSX.Element; // real <a href>, R3

// components/features/account/google-callback-error.tsx
function GoogleCallbackError(): JSX.Element | null; // reads ?error={code} via useSearchParams (R5/R6/R7)
```

**Business logic flow (this task's slice):**
```
GoogleAuthButton (mounted on both /login and /register)
  -> real navigation to /auth/google/redirect?intent=login  (R1, R3)

GoogleCallbackError (mounted on /login, inside <Suspense>)
  -> read `error` search param
  -> absent                    => render nothing (R5)
  -> "google_email_conflict"   => distinct copy (R6)
  -> any other known/unknown   => shared generic fallback (R6)
  -> on render                 => focus moves into the banner (R7)
```

## Backward Compatibility

- **Database**: N/A — no persistence layer in `frontend/`.
- **API**: No API changes; consumes the already-merged `GET
  /auth/google/redirect`/`GET /auth/google/callback` contract.
- **Existing clients/data**: `RegisterForm`'s Google button is not yet
  committed to git (`components/features/account/` is untracked in
  `git status`) — fixing its `intent` value is not a hotfix to shipped
  code, it's a fix before first commit. No deployed client affected.

## Edge Cases & Risks relevant to this task

| Risk | Likelihood | Severity | Mitigation |
|---|---|---|---|
| `/login`'s Google button hand-copied instead of using `GoogleAuthButton`, reintroducing an invalid `intent` value | Low once R2's typed prop exists, but only if a future change bypasses the shared component | **High** — breaks the entry point exactly like the bug found in `/register` today | R1/R2: typed `intent` prop makes an invalid value a compile-time error |
| Error banner accidentally renders the raw `code` string instead of mapped copy for an unrecognized/future code | Low | Low — a cryptic string briefly shown to the user, not a security issue since the code vocabulary carries no sensitive detail | R6: explicit fallback-to-generic rule, dedicated test for an unmapped code |
| Focus left on a removed/hidden element when the error banner replaces the loading state | Medium — invisible in a visual-only review pass | Medium — screen-reader/keyboard users lose their place | R7 + dedicated a11y test |
| `google_email_conflict` copy read as blaming the user or as an invitation to just retry | Low once R6/D4's distinct copy exists | Low — UX polish issue, not a security one | R6, copy exact-text pending Open Item (see below) |

## Files Changed / NOT Changed (this task's subset)

| File | Change Type | Description |
|---|---|---|
| `components/features/account/google-auth-button.tsx` | Add | Shared, typed Google entry-point component (R1-R3) |
| `components/features/account/register-form.tsx` | Modify | Replace inline anchor with `GoogleAuthButton` (D1, D5) |
| `components/features/account/google-callback-error.tsx` | Add | `?error={code}` banner (R5-R7) |
| `app/(auth)/login/page.tsx` | Modify | Replace Phase 0 placeholder (R4) |
| Corresponding `*.test.tsx` for the four files above | Add | Per Testing Checklist below |

| File | Reason untouched (this task) |
|---|---|
| `components/providers/auth-bootstrap-provider.tsx`, `app/layout.tsx`, `lib/api/client.ts`, `mocks/handlers.ts` | Task 2's scope |
| `app/(auth)/layout.tsx`, `_components/auth-shell-client.tsx` | Unmodified — `/login` continues using the existing shell as-is (D2) |
| `lib/api/account.ts` | No new typed request/response function needed — the two Google endpoints aren't JSON endpoints |
| `app/(auth)/forgot-password/page.tsx`, `reset-password/page.tsx` | Out of scope — task #4 |

## Testing Checklist (this task's subset)

- [ ] R1: both Google buttons render with `href` containing exactly `intent=login`, never `intent=register`
- [ ] R2: `GoogleAuthButton`'s `intent` prop is typed against the schema's `"login" | "link" | "reauth"` union (type-level check, not just runtime)
- [ ] R3: the button renders as a real `<a>` element, not wired to any `apiFetch`/mutation call
- [ ] R4: `/login` renders a heading, the "coming soon" note, and the Google button — no credential input fields present
- [ ] R5: `/login` with no `error` param renders no banner; with `error` present, renders one
- [ ] R6: `error=google_email_conflict` shows the distinct message; each of the other four documented codes and one unmapped/unknown code all show the same shared generic fallback; the raw code string is never present in rendered output
- [ ] R7: the error banner is the first child inside `AuthShellClient`'s panel, and focus moves into it on render

## Testing Examples & Common Mistakes (this task's subset)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Hand-copying the Google anchor into `/login` instead of using `GoogleAuthButton` | Reintroduces the exact `intent=register`-style bug class, undetected until a live `400` | Always route through the shared, typed component (R2/D5) |
| Rendering the raw `error` query-param value directly in the banner | A cryptic backend-internal string (e.g. `state_mismatch`) shown verbatim to the user | R6 — always map through the known-code table, fallback to the shared generic string for anything unmapped |
| Building `/login`'s email/password form "while I'm in here" | Duplicates/conflicts with task #3's actual build of the same page | D2/R4 — this task's `/login` scope stops at the Google button + error banner + placeholder note |

## Open Items relevant to this task

- **`api/openapi.yaml`'s misleading `intent=login`/`register` prose**
  (originating techplan §14, Active #3) — likely origin of the D1 bug
  this task fixes. Suggested one-line doc fix, but a shared cross-track
  document this task doesn't have sole authority to edit unilaterally.
- **Exact copy for R6's error messages** (originating techplan §14,
  Active #4) — both the `google_email_conflict`-specific message and
  the shared four-code generic fallback are placeholder Indonesian
  text pending product sign-off, same treatment as two prior open
  items already recorded in this codebase's `RegisterForm`/
  `VerifyEmailStatus` work. Placeholder text is acceptable to start
  building with — do not invent final product copy.
