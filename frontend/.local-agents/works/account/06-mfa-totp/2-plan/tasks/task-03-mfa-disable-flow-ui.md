# Task 3 — MFA disable flow UI (password / Google re-auth)

> Derived from: `../techplan.md` ("Tech Plan: MFA TOTP (Frontend)",
> account/06-mfa-totp). This task file redistributes §8-13 detail
> relevant to its own scope, in full — it does not summarize. For the
> Summary, §1-7 rationale, and §14 Open Items, read the source techplan
> directly.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **Task 1** (API layer, hooks, mocks & manual-entry
> parsing utility) — this task imports and calls `useMfaDisable()`,
> which must exist first. Do not stub a fake hook "to unblock" — this
> task's own tests exercise the real hook + MSW mock path
> (`component-test-mocking-discipline.md`'s network-layer-mocking
> principle — see Task 1's mocks).
> Parallel-eligible with: **Task 2** (Enroll flow UI) — no shared files,
> no shared state, both consume Task 1 independently.
> Feeds into: **Task 4** (`MfaSection` composition + page wiring)
> imports `MfaDisableForm` from this task.
> Recommended model: **DeepSeek V4 Pro** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row
> ("Decomposed: GLM 5.2 (max) / DeepSeek V4 Pro per sub-task") and its
> own tie-breaker ("DeepSeek V4 Pro when it's rule-table-heavy/precision
> work without a diagram") — the `email_password` branch is a
> near-direct mirror of the already-merged `UnlinkGoogleForm` (same
> re-auth-gated-destructive-action shape), and the Google-only branch is
> a linear button→401→retry sequence with no diagram of its own — this
> is precedent-following, precision work, not novel state-machine
> design.

## Scope

Build `MfaDisableForm` (the enrolled-state UI: branches on
`auth_providers` into a password-confirm form for `email_password` users
or a single re-auth-aware button for Google-only users) and its `zod`
password schema.

**Rules owned by this task** (full text, copied from techplan §4 — this
task owns the *component-layer* behavior for every rule below, including
R15/R16/R18/R19 whose *wrapper-function/hook-layer* behavior was already
built in Task 1):

- **R14** (disable, `email_password` branch shown): Given
  `user.auth_providers` includes `"email_password"`, When
  `MfaDisableForm` renders, Then show a password field + destructive
  "Nonaktifkan MFA" button — mirrors `UnlinkGoogleForm`'s exact shape
  (single field, destructive button, `bannerRef`-focus convention).
- **R15** (disable, `email_password` success) — **component-layer
  portion**: Given the correct current password, When submitted, Then
  call Task 1's `useMfaDisable()` with `{ password }`; on `200` (Task
  1's hook already invalidated `accountKeys.me()`), Then render nothing
  further locally — no local success view needed (matches
  `UnlinkGoogleForm`'s own precedent: the parent, `MfaSection` in Task
  4, re-renders into the not-enrolled branch once `mfa_enabled` flips to
  `false`).
- **R16** (disable, `email_password` `401`): Given the wrong password,
  When submitted, Then show `ApiError.detail` verbatim in a banner
  (Task 1's wrapper passes it through; generic fallback if absent, per
  the schema's undifferentiated `401`); form stays interactive.
- **R17** (disable, Google-only branch shown): Given
  `user.auth_providers` does NOT include `"email_password"`, When
  `MfaDisableForm` renders, Then show a single "Nonaktifkan MFA" button,
  no password field.
- **R18** (disable, Google-only success) — **component-layer portion**:
  Given the re-auth marker is already valid (a prior `intent=reauth`
  round trip succeeded — D6, Option B: this component does not
  pre-detect the marker's validity, it just attempts the call), When the
  button is clicked, Then call Task 1's `useMfaDisable()` with no body;
  on `200`, render nothing further locally (same as R15's note).
- **R19** (disable, Google-only `401`): Given the marker is
  missing/expired, When the button is clicked, Then show an error
  banner plus a `<GoogleAuthButton intent="reauth" .../>` prompt; the
  disable button remains available to retry after the user returns from
  re-authenticating (D6's chosen optimistic-single-button approach —
  this is the entire re-auth UX for the Google-only branch, no
  query-param-driven redirect handling is built, see D6 in the source
  techplan for why).
- **R20** (no regenerate action): Given the enrolled state
  (`MfaDisableForm`), When rendered, Then no one-click "regenerate
  backup codes" action exists; show a short explanatory line instead
  ("Untuk mendapatkan kode cadangan baru, nonaktifkan MFA lalu aktifkan
  kembali.") — `page-map.md`'s "regenerate" wording does not correspond
  to any real backend endpoint (D10, Stage 2 Area 1 miscontext finding).
- **R21** (a11y — banner focus, this component's own banners): Given any
  error banner in this component renders (R16/R19), When it renders,
  Then focus moves into it (`bannerRef` + `useEffect`, matching
  `UnlinkGoogleForm`/`SetPasswordForm`'s existing convention).

## Interface Contract (relevant subset of techplan §8)

**This task consumes from Task 1:**

```typescript
useMfaDisable(): UseMutationResult<{ message?: string }, ApiError, MfaDisableRequest>;
```

**This task's exports:**

```typescript
// components/features/account/mfa-disable-form.tsx (new)
function MfaDisableForm(props: {
  hasEmailPassword: boolean; // mirrors GoogleIdentityControl's prop-driven-by-parent convention
}): JSX.Element;
```

**Business logic flow (this task's slice, verbatim from §8):**

```
MfaDisableForm({ hasEmailPassword })
  hasEmailPassword === true:
    password field + "Nonaktifkan MFA" button
      submit -> useMfaDisable().mutate({ password })
        -> 200  => no local view, parent re-renders once mfa_enabled flips (R15)
        -> 401  => banner, .detail verbatim, form stays interactive (R16)
  hasEmailPassword === false:
    single "Nonaktifkan MFA" button
      click -> useMfaDisable().mutate({})
        -> 200  => no local view, parent re-renders once mfa_enabled flips (R18)
        -> 401  => banner + GoogleAuthButton(intent="reauth") prompt,
                    button stays available to retry (R19)
  always: explanatory "disable then re-enroll" line, no regenerate action (R20)
```

## Architecture (relevant note from §9)

`mfa-disable-form.tsx` branches on `auth_providers` (R14/R17), passed in
as a prop from the parent (`MfaSection`, Task 4) — same
prop-driven-by-parent convention as `GoogleIdentityControl`, not internal
state. `mfa-disable-schema.ts` mirrors `unlink-google-schema.ts` almost
exactly: a single required `password` field, no length policy to enforce
client-side (the backend compares against the caller's existing bcrypt
hash, not a new-password policy).

## Implementation Details (verbatim from §10)

**File**: `components/features/account/mfa-disable-form.tsx` (new)
- Props: `hasEmailPassword: boolean`. Renders the password form
  (R14-R16) or the single button (R17-R19), plus the always-shown
  explanatory line (R20).

**File**: `components/features/account/mfa-disable-schema.ts` (new)
- `zod` schema for `password`: required, non-empty — same shape as
  `unlink-google-schema.ts`'s `unlinkGoogleSchema`.

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `components/features/account/mfa-disable-form.tsx` | Add | Password / Google-only disable branches |
| `components/features/account/mfa-disable-schema.ts` | Add | `zod` schema for `password` (`email_password` branch) |
| `components/features/account/mfa-disable-form.test.tsx` | Add | Component tests |

**Reason untouched** (relevant row from §11): `components/features/account/mfa-enroll-flow.tsx`, `qr-code.tsx` and their schema — Task 2's independent scope, no shared file. `components/features/account/unlink-google-form.tsx` — precedent read for shape only, not modified.

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R14: `email_password` branch shows password field + destructive
  button
- [ ] R15 (component-layer): correct password → `200` (via MSW) → no
  local success view rendered
- [ ] R16: wrong password → `401` (via MSW) → `.detail` shown verbatim,
  form stays interactive
- [ ] R17: Google-only branch shows single button, no password field
- [ ] R18 (component-layer): Google-only success → `200` (via MSW) → no
  local success view rendered
- [ ] R19: Google-only `401` (via MSW) → error banner + reauth link,
  button stays available for retry
- [ ] R20: enrolled view shows no regenerate action, shows the
  disable→re-enroll explanatory line
- [ ] R21: focus moves into any of this component's banners on render

**Count-check** (this task's slice): 8 checklist items above, covering
R14, R15 (component-layer), R16, R17, R18 (component-layer), R19, R20,
R21 — the *wrapper-function/hook-layer* halves of R15/R16/R18/R19 live
in Task 1's own checklist, not duplicated here.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Building a one-click "regenerate backup codes" button to match `page-map.md`'s literal wording | Invents a backend capability that doesn't exist (D10) — no such endpoint | Show the disable→re-enroll explanatory line instead (R20) |
| Rendering a local success message/view after disable succeeds | Duplicates state the parent (`MfaSection`, Task 4) already owns — the parent re-renders into the not-enrolled branch once `mfa_enabled` flips via cache invalidation | Render nothing further locally on success (R15/R18), same as `UnlinkGoogleForm`'s existing precedent |
| Pre-detecting whether the Google re-auth marker is valid before enabling the disable button | No mechanism exists for this (D6's whole point) — would require inventing an unconfirmed backend contract | Always render the button; let the documented `401` drive the reauth-prompt UI (R19) |
</content>
