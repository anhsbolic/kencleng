# Build Report — Account Linking (account #05, frontend)

> Task      : account domain task #5, frontend surface —
>             `docs/spec/1-account/features/05-account-linking.md`
> Executed  : 2026-08-27, single session, no task decomposition (see
>             `2-3-techplan-decomposition-prompt.md`'s Step 0 evaluation
>             — the techplan was judged small/linear enough for one
>             pass, comparable in scope to `account/04`'s own
>             un-decomposed build)
> Techplan  : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved by Anhar)
> Status    : Build complete, all tests/typecheck/lint/build green. No Tier 0 fenced sub-area in this plan (frontend has no JWT/token-crypto code) — nothing blocks committing on that front.

---

## Execution summary

Built in dependency order per the techplan's §9: API layer first, then
hooks, then the shared `PasswordInput` extraction, then the form/
control components, then the page, then the `VerifyEmailStatus`
additive fix, then mocks, then tests.

| Step | Scope | Result |
|---|---|---|
| 1 | `lib/api/account.ts`: `setPassword()`, `unlinkGoogle()` + types | ✅ |
| 2 | `lib/hooks/use-set-password.ts` (+ extracted `applySetPasswordSuccess`), `use-unlink-google.ts` | ✅ |
| 3 | `components/shared/password-input.tsx` | ✅ 3 tests |
| 4 | `SetPasswordForm` + `set-password-schema.ts` | ✅ 11 tests (R4-R12) |
| 5 | `UnlinkGoogleForm` + `unlink-google-schema.ts` | ✅ 7 tests (R16-R18 + client validation) |
| 6 | `GoogleIdentityControl` | ✅ 4 tests (R13-R16) |
| 7 | `LoginMethodsSection` | ✅ 5 tests (R1-R3) |
| 8 | `app/(dashboard)/dashboard/security/page.tsx` rebuild | ✅ |
| 9 | `VerifyEmailStatus` additive fix (R19/D6) | ✅ 2 new tests, 8 pre-existing tests unaffected |
| 10 | `mocks/handlers.ts` — 2 default handlers (R20) | ✅ |

## Files changed

**New**: `components/shared/{password-input,password-input.test}.tsx` ·
`lib/hooks/{use-set-password,use-set-password.test,use-unlink-google}.ts`
· `components/features/account/{set-password-schema,unlink-google-schema}.ts`
· `components/features/account/{set-password-form,set-password-form.test}.tsx`
· `components/features/account/{unlink-google-form,unlink-google-form.test}.tsx`
· `components/features/account/{google-identity-control,google-identity-control.test}.tsx`
· `components/features/account/{login-methods-section,login-methods-section.test}.tsx`.

**Edited**: `lib/api/account.ts` (+`setPassword`/`unlinkGoogle` +
`SetPasswordRequest`/`UnlinkGoogleRequest`/`SetPasswordResult`/
`UnlinkGoogleResult` types) · `mocks/handlers.ts` (+2 default handlers)
· `app/(dashboard)/dashboard/security/page.tsx` (placeholder replaced
with `LoginMethodsSection` + Task #6 extension-point comment) ·
`components/features/account/verify-email-status.tsx` (+cache
invalidation on success, +conditional terminal CTA) ·
`components/features/account/verify-email-status.test.tsx` (+2 new
cases for R19, existing 8 cases unmodified).

**Untouched by design**: `lib/api/client.ts` (D4 — verbatim
`.detail`/`.message` sufficient, no `.type` discriminant needed) ·
`lib/api/schema.d.ts` (already complete and current, confirmed not
stale during planning) · `app/(dashboard)/layout.tsx`,
`_components/{dashboard-shell-client,nav-items}.tsx` (the "Keamanan"
nav item already existed) · `components/ui/*` (reused as-is, no
primitive gap) · `components/features/account/login-form.tsx` (not
retrofitted to the new `PasswordInput` extraction — out of scope) ·
`backend/**` (read-only cross-checks during planning only).

## Verification results

- **Unit/component suite**: `npx vitest run` — **191/191 tests, 36/36
  files, all green** (up from the pre-build baseline of 156/156 — 35
  net new tests: 3 `password-input.test.tsx`, 3
  `use-set-password.test.ts`, 11 `set-password-form.test.tsx`, 7
  `unlink-google-form.test.tsx`, 4 `google-identity-control.test.tsx`,
  5 `login-methods-section.test.tsx`, 2 new cases added to the
  pre-existing `verify-email-status.test.tsx`).
- **Typecheck**: `npx tsc --noEmit` — clean, 0 errors.
- **Lint**: `npm run lint` (ESLint) — clean, 0 errors/warnings.
- **Production build**: `npm run build` (Next.js/Turbopack) — compiles
  successfully; `/dashboard/security` appears in the route list,
  statically pre-rendered (`○`), no SSR crash.
- **R18's distinct-message assertion** verified directly with three
  separate mocked responses (`401`, `409 only-identity`, `409
  unverified-remaining-identity`) each asserted against its own exact
  text — not just "shows an error," per the techplan's §13 flagged
  common mistake.
- **R2's "form stays interactive, not hidden" requirement** verified
  directly: the pending-verification-banner test asserts both the
  banner text and the `current_password`/`Ganti Password` fields are
  simultaneously present, not a hidden-form assertion.

## Deviations from the plan (implementation-level judgment calls, not scope changes)

None of these change what the techplan decided (D1-D6, R1-R20) —
refinements made while writing the actual code:

1. **`useSetPassword`'s `onSuccess` logic was extracted into a
   standalone `applySetPasswordSuccess(result, queryClient)` function**,
   mirroring `use-login.ts`'s existing `applyLoginSuccess` pattern
   exactly, rather than left inline inside `useMutation({ onSuccess })`
   as the techplan's Implementation Details sketch showed. This makes
   the branching logic (added/changed/validation) directly
   unit-testable via `use-set-password.test.ts` without a `renderHook`
   harness — consistent with this codebase's existing convention that
   only hooks with real branching logic get their own extracted,
   testable function and test file (`useLogout`, with no branching, has
   no such file either).
2. **The pending-verification banner (R2) renders *above*
   `SetPasswordForm`**, matching R2's own prose ("renders above the
   Branch 2 form") rather than §8's business-logic-flow pseudocode
   sketch, which listed it after the form. Treated the R-numbered
   acceptance criterion as the more authoritative statement of the two
   — a genuine, small inconsistency inside the techplan itself, not
   something introduced during the build. Worth a one-line fix to §8's
   pseudocode if the techplan is revised again, but not a reason to
   deviate from R2's explicit wording.
3. **`mockGoogleOnlyUser` was not added as a shared, exported fixture in
   `mocks/handlers.ts`** as the techplan's R20 wording literally
   suggested. Checked the established convention first
   (`dashboard-shell-client.test.tsx`'s `mockMe(roles)` helper) and
   found every existing test file that needs a non-default `/account/me`
   shape constructs it inline via `server.use(...)`, not by importing a
   shared fixture — no fixture in this file is exported today. Followed
   that existing pattern instead (each new test file that needs a
   Google-only/mixed-identity user builds the literal object itself),
   rather than introducing a new "shared exported fixture" convention
   this codebase doesn't otherwise use. `mocks/handlers.ts` still gained
   exactly what R20's other half required: two new default handlers for
   the endpoints this task adds.
4. **A `key={mode}` was added to `<SetPasswordForm>` in
   `LoginMethodsSection`** (not specified in the techplan) — forces a
   clean remount if the account's identity shape changes underneath an
   already-mounted form (e.g. right after verifying mid-session, when
   `mode` flips from `"add"` to `"change"`), avoiding stale local
   success/error state carrying across the transition. Low-risk,
   additive, and consistent with standard React practice for this exact
   "identity-changing prop" situation.
5. **`SetPasswordForm`'s `mode="add"` catch block uses a bindingless
   `catch { ... }`** rather than `catch (error) { ... }`, since Branch
   1's only documented failure path reaching the `catch` is network/5xx
   (no `401` branch exists for that mode, confirmed via `security.go`)
   — nothing in the caught value needs inspecting, so no unused binding
   is introduced.

## Open items carried forward (unresolved by this build, same as techplan §14)

None — the techplan's Active list was already empty at Approved status
(all ten items resolved during planning, §14 Resolved #1-10). No new
open item surfaced during the build itself; both things noticed while
writing the actual code (the §2/§8 pending-verification-banner ordering
inconsistency, and the `mockGoogleOnlyUser` fixture-placement question)
were resolved as small, self-contained implementation judgment calls
above, not left open.
