# Stage 2 — Area 4: Feature components precedent
(`components/features/account/`)

## Current state

No `set-password-form.tsx` / `unlink-google-*.tsx` (or equivalent)
exists yet. Read every existing form/action component in this
directory in full; the precedents actually relevant to this task:

- **`LoginForm`** (login-form.tsx, 284 lines) — the strongest structural
  precedent. Demonstrates:
  - **Multi-step-in-one-component** (`step: "password" | "mfa"` state) —
    directly analogous to this task's "which section renders" question
    (Branch 1 vs Branch 2 set-password form, plus a separate unlink
    action, all potentially in one component/page).
  - **Password show/hide toggle** (`showPassword` state, `Eye`/`EyeOff`
    from `lucide-react`, a `Ghost` `Button` with `aria-label` switching
    between "Tampilkan password"/"Sembunyikan password") — directly
    reusable for every password field this task needs (`current_password`,
    new `password`, unlink's `password` confirmation).
  - **Request-level banner as first child + focus-on-error convention**
    (`bannerRef.current?.focus()` in a `useEffect` keyed on the error
    state) — the exact fix for the `/login` prototype's Known Issue #1
    (field-level vs banner-level conflation), already the established
    house style across `LoginForm`, `ResetPasswordForm`,
    `ForgotPasswordForm`, `GoogleCallbackError`, `VerifyEmailStatus`.
  - **`ApiError.detail` used directly as banner copy in one specific
    case** (`onSubmitPassword`'s catch: `error.detail ?? GENERIC_ERROR_MESSAGE`)
    — this is the *one* place in the codebase that renders a backend
    `detail` string as-is rather than mapping to frontend-owned copy,
    justified there because `errors.go`'s generic-credential detail is
    already the intended user-facing string. **Not** a safe default to
    copy blindly for this task: `patterns.md` §B's "never render raw
    backend text" rule is the norm everywhere else (`ResetPasswordForm`,
    `ForgotPasswordForm`, `VerifyEmailStatus`, `GoogleCallbackError` all
    map status/code to frontend-owned strings instead) — worth being
    deliberate in Stage 3 about which convention this task's two
    endpoints should follow, since unlink's two `409`s in particular
    need frontend-owned, distinct copy per the spec, not raw `.detail`
    passthrough.
- **`GoogleAuthButton`** (google-auth-button.tsx, 40 lines) — its
  `intent` prop is typed directly from the generated schema's
  `paths["/auth/google/redirect"]["get"]["parameters"]["query"]["intent"]`
  union, which already includes `"link"` alongside `"login"`/`"reauth"`.
  The component's own doc comment explicitly anticipates this: *"link"/
  "reauth" belong to a different, session-authenticated flow (account
  linking / MFA re-auth, out of this component's scope)* — written by
  Task #2's session, evidently already aware this task would need it.
  **No page anywhere currently renders `<GoogleAuthButton intent="link" .../>`**
  — Task #2 built the generic capability but had nowhere to use the
  "link" branch yet.
- **`register-schema.ts`** — the canonical password-length-only zod
  rule (`z.string().min(8, "Password minimal 8 karakter")`), reused
  verbatim by `reset-password-schema.ts`. This task's `password` field
  (both branches) should reuse the same rule, not reinvent it.
- **`google-callback-error.tsx`** — smallest example of the
  code-to-frontend-copy mapping convention (`messageForCode`), useful
  precedent for mapping unlink's two `409` `type`/`detail` values to
  the two spec-required distinct messages.

## Requirement

- Per the spec, the page needs: a set-password section whose *shape*
  depends on `auth_providers` (Branch 1 form: `email` + `password`;
  Branch 2 form: `current_password` + `password`), and an unlink
  section (visible only when `google` is in `auth_providers`) requiring
  a `password` confirmation field and rendering one of two distinct
  `409` messages when blocked.
- Per `page-map.md`, a Google-only user should also see an affordance
  to add Google in the first place is *not* applicable here (they
  already have Google) — but the inverse case (an `email_password`-only
  user with no Google linked) is implied by the same page-map row's
  "link/unlink Google identity" phrasing, and that direction's backend
  is Task #2's already-shipped `intent=link` endpoint.

## Gap

- Both form components need to be built from scratch; `LoginForm`'s
  multi-step pattern and password-toggle sub-pattern are the closest
  reusable shapes, not a literal template for either branch's exact
  field set.
- **Scope question, not yet resolved by either task's spec**: does
  building `/dashboard/security` under Task #5 include wiring up
  `<GoogleAuthButton intent="link" .../>` for the "add Google" direction
  (an `email_password`-only user linking Google), even though the
  underlying endpoint is Task #2's? The account-linking spec's own
  structural note only says the *behavior* "isn't duplicated here" — it
  doesn't say the *page* doesn't render the trigger. Since Task #5 is
  the one actually building this page's content, leaving this
  unaddressed would mean no task ever wires up the one remaining
  `GoogleAuthButton` intent, and `page-map.md`'s "link... Google
  identity" action would go unimplemented by any task. Flagging for
  Stage 3, not resolving here.
- Password-field validation schema for this task's forms should reuse
  `register-schema.ts`'s length-only rule (already proven reusable once
  by `reset-password-schema.ts`) rather than reintroduce a third copy.
- No component in this codebase yet demonstrates "same status code
  (409), branch copy on a secondary field" the way unlink needs —
  `google-callback-error.tsx`'s `messageForCode` is the closest shape
  (switch on a string value to pick frontend copy) but keys off a query
  param, not a thrown `ApiError`'s parsed body field. Given Area 1's
  finding that `ApiError`/`readProblemDetail` don't currently expose
  `.type`, this component-level gap and that client-layer gap are the
  same underlying issue viewed from two layers.

## Page-consolidation check

- No new finding beyond Area 3's — component-level precedent confirms
  the *pattern* to reuse (`LoginForm`'s multi-step shape) but doesn't
  resolve the cross-task (#5/#6) page-ownership question raised there.

## Sniffing

- **Misleading signal**: `GoogleAuthButton` already generically
  supporting `intent="link"` could look like "linking-to-Google is
  already wired up somewhere" on a shallow read of that one file — a
  full repo search confirms no page uses it with that intent yet.
- **Inconsistency**: `LoginForm`'s one instance of rendering
  `ApiError.detail` directly (rather than mapping to frontend-owned
  copy) is inconsistent with the otherwise-universal "map, don't
  passthrough" convention used by every other form/status component in
  this directory — not a bug in `LoginForm` itself (justified there),
  but not a pattern to reach for by default when building this task's
  components, especially given the spec's explicit requirement for two
  *distinct* unlink messages that don't come free from any single
  backend `.detail` string pairing cleanly with "which case is this."
- **Risk**: reusing `LoginForm`'s password show/hide toggle across
  potentially 2-3 password fields on one page (current_password, new
  password, unlink confirmation password) multiplies the amount of
  `showX`/`showY` boolean state if copied naively per-field — worth
  considering in Stage 3 whether this warrants extracting a small
  `PasswordInput` wrapper instead of three independent copies of the
  same toggle logic (a "second domain needs it" promotion candidate per
  `phase0-shared-infra.md`'s Incremental Growth Rule, since this would
  be the second time the exact same toggle shape is needed within the
  same domain, arguably even the same page).

Proceeding to Area 5 (`components/ui/` primitives inventory).
