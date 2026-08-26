# Task 3 — Login form + MFA challenge UI

> Derived from: `../techplan.md` ("Tech Plan: Login & Session Management
> (Frontend)", account/03-login-session-management). This task file
> redistributes §8-13 detail relevant to its own scope, in full — it does
> not summarize. For the Summary, §1-7 rationale, and §14 Open Items,
> read the source techplan directly.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **Task 1** (API layer & hooks for login/MFA) — this task
> imports and calls `useLogin()`/`useLoginMfa()`, which must exist first.
> Do not start this task by stubbing fake hooks "to unblock" — wait for
> Task 1's real exports, since this task's own tests exercise the real
> hook + MSW mock path (`component-test-mocking-discipline.md`'s
> network-layer-mocking principle — see `task-01`'s mocks).
> Recommended model: **GLM 5.2 (max)** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row and its
> tie-breaker ("GLM when the work leans on diagrams, state-transitions,
> or multi-step reasoning") — this task has genuine multi-branch state
> (password step ↔ MFA step, 4 distinct response outcomes per step) and
> the source techplan's own §7 flags it as a High-severity risk
> specifically because it has **zero existing visual or code precedent
> anywhere in this codebase** — the highest first-time-implementation
> risk in the whole plan.

## Scope

Build `/login`'s credential form: email/password fields, password
show/hide toggle, the MFA challenge step (its own render branch inside
the same component, not a route), and error handling for both steps —
replacing `/login/page.tsx`'s current placeholder content and rewriting
its now-contradicted test file.

**Rules owned by this task** (full text, copied from techplan §4 — this
task owns the *component-layer* behavior for every rule below, including
R3/R4/R7 whose *hook-layer* behavior was already built in Task 1):

- **R1** (login form fields): Given `/login` loads, When rendered, Then
  it shows email + password fields (password schema: non-empty, min 8
  chars — matches the backend's length-only policy already documented
  for registration, `RegisterForm`'s own precedent in `register-
  schema.ts`), a "Lupa password?" link to `/forgot-password`, a submit
  button, and — only in the password step (R9) — a divider +
  `GoogleAuthButton intent="login" label="Masuk dengan Google"`.
- **R2** (password show/hide): Given the password field, When the toggle
  is clicked, Then the field's `type` switches between `password`/
  `text`, and the toggle's accessible label switches between
  "Tampilkan password"/"Sembunyikan password" — composed locally (D6),
  not a change to the shared `Input` primitive (`components/ui/
  input.tsx` is explicitly untouched by this whole plan — see Files NOT
  Changed below).
- **R3** (password-step submit → success) — **component-layer portion**:
  Given `useLogin()` resolves with `status: "ok"`, Then this component
  does not need to perform the store/cache/redirect side effects itself
  (Task 1's hook's `onSuccess` already did) — this component's own job
  is simply to call the mutation and let its resolved/error state drive
  rendering; no separate success-handling code belongs here.
- **R4** (password-step submit → MFA required) — **component-layer
  portion**: Given `useLogin()` resolves with `status: "mfa_required"`,
  Then this component's `step` state becomes `'mfa'`, and
  `mfa_pending_token` (from the resolved value) is stored in local
  `useState` (D2) — no store/cache mutation, no cookie set yet.
- **R5** (password-step submit → failure): Given wrong credentials or a
  locked-out identifier, When `useLogin()` rejects with an `ApiError`
  (`401` or `429`), Then render `<Banner variant="error">
  {error.detail}</Banner>` as the form's first child (**never** attached
  to the email/password input's own `error` prop — this is the
  confirmed, not-fixed Known Issue #1 from the Tier 1 prototype
  [`design-reference/login-register.html`], and must not be
  reproduced), falling back to a generic frontend-owned message only if
  `.detail` is absent (network/5xx).
- **R6** (MFA step fields): Given `step === 'mfa'`, When rendered, Then
  show one field pair — `totp_code` (primary) and a "Gunakan kode
  cadangan" toggle revealing `backup_code` instead — with client-side
  validation requiring exactly one of the two to be non-empty (the
  generated `LoginMfaRequest` type only documents this as a comment;
  this task is where the frontend actually enforces it as UX, via a
  `zod` `.refine()`). No Google button in this step (R9).
- **R7** (MFA-step submit → success) — **component-layer portion**: same
  as R3's component-layer note — `useLoginMfa()`'s `onSuccess` (Task 1)
  already performed the side effects; this component just needs to
  render nothing special on this branch (the redirect Task 1's hook
  triggers will unmount this component).
- **R8** (MFA-step submit → failure): Given an invalid code or MFA-stage
  lockout, When `useLoginMfa()` rejects (`401` or `429`), Then the same
  banner-first treatment as R5, and the user **stays on the MFA step**
  (not bounced back to re-enter the password) — the `mfa_pending_token`
  is still valid (5-minute TTL) unless it has expired, in which case R8
  still shows the generic failure banner (spec 03: an expired/malformed
  token also returns `401`) and the user must restart from the password
  step since there is no token left to retry with (there is no
  automatic detection of "token expired specifically" vs. "code was
  wrong" — both collapse to the same `401`/banner, per spec 03's own
  anti-enumeration-adjacent design; do not try to distinguish them).
- **R9** (Google button gating): Given `step === 'mfa'`, Then
  `GoogleAuthButton` and its divider are not rendered — only relevant
  during the password step.
- **R10** (`mfa_pending_token` lifetime): Given a page refresh/navigation
  away mid-MFA-step, When the component remounts, Then it starts back at
  `step === 'password'` with no memory of the prior attempt
  (component-local state, not persisted — D2's accepted trade-off; do
  not add any persistence mechanism here, this is intentional).

## Interface Contract (relevant subset of techplan §8)

**API contract this task's UI reflects** (already built by Task 1 — this
task consumes the hooks, does not call `fetch`/`apiFetch` itself):
```typescript
// POST /auth/login  — via Task 1's useLogin()
// POST /auth/login/mfa — via Task 1's useLoginMfa()
// (see task-01-api-layer-login-mfa.md for the full wire-shape detail)
```

**This task's exports:**
```typescript
// components/features/account/login-form.tsx (new)
function LoginForm(): JSX.Element; // R1-R10
```

**This task consumes from Task 1:**
```typescript
useLogin(): UseMutationResult<LoginResult, ApiError, LoginRequest>;
useLoginMfa(): UseMutationResult<LoginResult, ApiError, LoginMfaRequest>;
// LoginResult's `status` field discriminates "ok" vs "mfa_required" — see task-01 for the exact shape
```

**Business logic flow (this task's slice, verbatim from §8):**
```
LoginForm (mounted on /login, step state 'password' | 'mfa')
  password step:
    submit -> useLogin().mutate(values)
      -> resolves status=ok         => (Task 1's hook already handled side effects — this component renders nothing extra)
      -> resolves status=mfa_required => step='mfa', store mfa_pending_token locally (R4)
      -> rejects (401 | 429)          => banner, error.detail verbatim (R5)
  mfa step:
    submit -> useLoginMfa().mutate({ mfa_pending_token, totp_code | backup_code })
      -> resolves                     => (same as password-step success)
      -> rejects (401 | 429)          => banner, stay on mfa step (R8)
```

## Architecture (relevant note from §9)

`LoginForm` is the single new feature component for this page, following
`RegisterForm`'s exact shape (`components/features/account/register-
form.tsx` — read this file first as the composition precedent named
explicitly by `/login/page.tsx`'s own existing code comment):
`react-hook-form` + `zodResolver`, a `Banner variant="error"` as first
child on request-level failure, per-field errors via `Input`'s own
`error` prop for the two genuinely field-level cases (empty-required,
min-length — client-side UX validation only, per `form-validation-
boundary.md`), never for the backend's generic credential failure
(R5/R8). The MFA step is a second render branch inside the same
component (`step === 'mfa'`), not a route change — no new page-map.md
entry, consistent with the source techplan's finding that this is a
sub-state of the existing `/login` Form pattern, not a separate flow.

## Implementation Details (verbatim from §10)

**File**: `components/features/account/login-form.tsx` (new)
- New Client Component. `step` state (`'password' | 'mfa'`),
  `mfaPendingToken` state, `useLogin`/`useLoginMfa` mutations (from Task
  1), `zodResolver`-backed forms per step (R1-R10). Embeds a local
  password-show/hide composition (D6, see below) and, password-step-only,
  `GoogleAuthButton` (R9) — reuse `components/features/account/
  google-auth-button.tsx` as-is (already built by task #02), do not
  duplicate its markup.
  - **Password show/hide (D6)**: a small, local composition — do **not**
    modify `components/ui/input.tsx`. Reuse `Input`'s existing visual
    tokens directly (border/radius/focus-ring classes) plus the existing
    `Button variant="ghost"` for the toggle button, rendering only the
    Eye/EyeOff Lucide icon (no label text) — matches `MaskedField`'s
    already-established spec for the same visual pattern
    (`patterns.md` §C / `design-guidelines.md`'s `MaskedField` section:
    "Ghost button variant, Eye/EyeOff Lucide icon only, to keep it
    compact").

**File**: `components/features/account/login-schema.ts` (new)
- `loginSchema` (email + password, min 8) and `loginMfaSchema`
  (`totp_code`/`backup_code`, `.refine()`-d so exactly one is non-empty —
  R6) — comments pointing at spec 03
  (`docs/spec/1-account/features/03-login-session-management.md`) as the
  authoritative source, matching `register-schema.ts`'s own convention
  (comment style: point at the spec, don't invent the rule ad hoc, per
  `form-validation-boundary.md`).

**File**: `app/(auth)/login/page.tsx`
- Change: replace the static "coming soon" note + `GoogleAuthButton`
  with `<LoginForm />`; `GoogleCallbackError` (built by task #02) stays
  exactly as-is, still rendered first, still inside its own `<Suspense>`
  boundary.

**File**: `app/(auth)/login/page.test.tsx`
- Change: replace the current negative assertions (which currently
  assert email/password fields do **not** exist — this task directly
  contradicts that, intentionally) with coverage of R1-R10 (see Testing
  Checklist below) — this file is rewritten, not incrementally extended.

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `components/features/account/login-form.tsx` | Add | Password + MFA step form (R1-R10) |
| `components/features/account/login-schema.ts` | Add | `zod` schemas for both steps |
| `app/(auth)/login/page.tsx` | Modify | Render `LoginForm` in place of the placeholder |
| `app/(auth)/login/page.test.tsx` | Modify (rewrite) | Replace negative assertions with R1-R10 coverage |

**Reason untouched** (relevant row from §11): `components/ui/input.tsx`
— D6, password show/hide is a local composition, not a shared-primitive
change. `app/(auth)/layout.tsx`, `_components/auth-shell-client.tsx` —
unmodified, `/login` continues using the existing shell/banner-first
convention as-is.

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R1: `/login` renders email + password fields, "Lupa password?" link, submit button, divider + Google button (password step only)
- [ ] R2: password field toggles `type="password"`/`type="text"` on click; accessible label switches correctly
- [ ] R3 (component-layer): submitting valid credentials against Task 1's real `useLogin()` hook (via MSW, not a mocked hook) results in no residual error banner and no premature re-render glitch — the store/cache/redirect assertions themselves belong to Task 1's own test suite, not duplicated here
- [ ] R4 (component-layer): a mocked `200 status=mfa_required` response (via MSW) transitions this component's `step` to `'mfa'` and renders the MFA fields
- [ ] R5: a mocked `401` and a mocked `429` (via MSW) both render the identical banner text from `error.detail`, as the form's first child, never attached to the email/password input's own `error` prop
- [ ] R6: the MFA-step schema rejects submission with both `totp_code` and `backup_code` empty, and with both filled simultaneously (exactly one required)
- [ ] R7 (component-layer): a mocked MFA-step `200` (via MSW) does not leave the component in an error or stuck-loading state
- [ ] R8: a mocked MFA-step `401`/`429` (via MSW) renders the banner and leaves `step` at `'mfa'`, not reverted to `'password'`
- [ ] R9: `GoogleAuthButton` is absent from the DOM while `step === 'mfa'`
- [ ] R10: remounting `LoginForm` (simulating a refresh) always starts at `step === 'password'`, regardless of prior state

**Count-check** (this task's slice): 10 checklist items above, covering
R1-R10 — R3/R4/R7's *hook-layer* assertions (store write, cache write,
redirect) live in Task 1's own checklist, not duplicated here; this
task's R3/R4/R7 items verify the component correctly renders/transitions
given each hook outcome, without re-asserting side effects it doesn't
own.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Attaching the generic credential error to the email input's own `error` prop | Reproduces the confirmed, not-fixed Known Issue #1 from the Tier 1 prototype (`design-reference/login-register.html`) — leaks which field was "more wrong" | R5/R8 — always the banner, never the input's `error` prop, for this specific failure class |
| Mocking `useLogin`/`useLoginMfa` directly in this component's tests instead of using MSW against Task 1's real hooks | Test can no longer catch a broken query key, wrong endpoint, or serialization bug in Task 1's actual implementation (`component-test-mocking-discipline.md`'s core warning) | Mock at the network layer (MSW), exercise the real hook |
| Re-implementing the store/cache/redirect success handling inside this component "just to be safe" | Duplicates Task 1's `onSuccess` logic, risking the two copies drifting apart | Trust Task 1's hook to have already performed R3/R7's side effects — this component's job on success is just to let the resulting navigation/unmount happen |
