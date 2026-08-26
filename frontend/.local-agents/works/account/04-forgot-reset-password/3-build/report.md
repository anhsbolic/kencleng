# Build Report — Forgot & Reset Password (account #04, frontend)

> Task      : account domain task #4, frontend surface —
>             `docs/spec/1-account/features/04-forgot-reset-password.md`
> Executed  : 2026-08-26, single session, no task decomposition (see
>             `2-plan/`'s decomposition-evaluation note — the techplan was
>             judged small/linear enough for one pass)
> Techplan  : `.local-agents/works/account/04-forgot-reset-password/2-plan/techplan.md` (Status: Approved by Anhar)
> Status    : Build complete, all tests/typecheck/lint/build green. No Tier 0 fenced sub-area in this plan (frontend has no JWT/token-crypto code) — nothing blocks committing on that front.

---

## Execution summary

Built in dependency order per the techplan's §9: schema regen first, then the API layer, then hooks, then the two form components + pages, then mocks, then tests.

| Step | Scope | Result |
|---|---|---|
| 1 | `lib/api/schema.d.ts` regeneration (R17) | ✅ picked up `/auth/reset-password`'s `429`; see Deviations #1 for an unrelated bundled change |
| 2 | `lib/api/account.ts`: `forgotPassword()` (D3), `resetPassword()` (D2) + types | ✅ |
| 3 | `lib/hooks/use-forgot-password.ts`, `use-reset-password.ts` | ✅ bare mutation wrappers, no side-effect (D5) |
| 4 | `ForgotPasswordForm` + `forgot-password-schema.ts` + `/forgot-password` page rebuild | ✅ 6 tests (R1-R6) |
| 5 | `ResetPasswordForm` + `reset-password-schema.ts` + `/reset-password` page rebuild | ✅ 9 tests (R7-R15) |
| 6 | `mocks/handlers.ts` — 2 default handlers (R16) | ✅ |

## Files changed

**New**: `lib/hooks/{use-forgot-password,use-reset-password}.ts` · `components/features/account/{forgot-password-form,forgot-password-schema}.tsx/.ts` + `forgot-password-form.test.tsx` · `components/features/account/{reset-password-form,reset-password-schema}.tsx/.ts` + `reset-password-form.test.tsx`.

**Edited**: `lib/api/account.ts` (+`forgotPassword`/`resetPassword` + `ForgotPasswordRequest`/`ResetPasswordRequest`/`ForgotPasswordResult` types) · `lib/api/schema.d.ts` (regenerated) · `mocks/handlers.ts` (+2 default handlers) · `app/(auth)/forgot-password/page.tsx` (renders `ForgotPasswordForm`) · `app/(auth)/reset-password/page.tsx` (renders `Suspense`-wrapped `ResetPasswordForm`, stays in the Auth Shell per D1).

**Untouched by design**: `app/(auth)/layout.tsx`, `_components/auth-shell-client.tsx` (D1 — both pages reuse the existing shell/banner-first convention as-is) · `components/features/account/login-form.tsx` (its `/forgot-password` link was already correct) · `components/ui/{input,button,banner,label}.tsx` (reused as-is) · `app/verify-email/page.tsx`, `verify-email-status.tsx` (D6's copy-source deviation deliberately not retrofitted here — Open Item, techplan §14 Resolved #8) · `api/openapi/account.yaml`, `api/openapi.yaml` (already correct; only the frontend's generated file was behind) · anything under `backend/` (read-only cross-checks only).

## Verification results

- **Unit/component suite**: `npx vitest run` — **156/156 tests, 30/30 files, all green** (up from the pre-build baseline of 141/141 — 15 net new tests: 6 in `forgot-password-form.test.tsx`, 9 in `reset-password-form.test.tsx`).
- **Typecheck**: `npx tsc --noEmit` — clean, 0 errors, both immediately after the schema regen (step 1) and again at the end.
- **Lint**: `npm run lint` (ESLint) — clean, 0 errors/warnings.
- **Production build**: `npm run build` (Next.js/Turbopack) — compiles successfully; `/forgot-password` and `/reset-password` both appear in the route list, statically pre-rendered (`○`), no SSR crash — confirms the `<Suspense fallback={null}>` boundary around `ResetPasswordForm`'s `useSearchParams()` read works under a real build, not just jsdom.
- **R13's correctness-critical assertion** (the techplan's one High-severity risk, §7) verified directly: the test drives a real MSW-mocked `422` response, asserts the form stays mounted/interactive, then submits a second time and asserts both requests carried the identical `token` value — not just "the form didn't disappear," but that the actual resubmit payload proves the link wasn't invalidated client-side.

## Deviations from the plan (implementation-level judgment calls, not scope changes)

None of these change what the techplan decided (D1-D7) — refinements made while writing the actual code:

1. **`schema.d.ts`'s regeneration picked up an unrelated change**, exactly as the techplan's §7 risk row anticipated: `/auth/login`'s `429` response type split from `components["responses"]["TooManyRequests"]` into a new `LockedOutGenericCredentials` component (a pending, already-committed-to-source change in `api/openapi/account.yaml` unrelated to this task). Confirmed safe before proceeding: grepped all hand-written code for both component names — neither is referenced anywhere outside `schema.d.ts` itself (every caller checks `res.status` numerically and reads the body at runtime via `readProblemDetail`/`res.json()`, never by generated-type component name), and `tsc --noEmit` stayed clean. Diff-reviewed per the plan's own mitigation step; not reverted, since it's a strict quality improvement to the login endpoint's type documentation and genuinely harmless to every existing caller.
2. **`ResetPasswordForm`'s focus-management reuses `LoginForm`'s banner-ref pattern for all four of its non-idle states** (missing-token/invalid, expired, success, and the non-terminal request-error banner) rather than the techplan's Implementation Details sketch, which didn't specify a mechanism. Considered a second heading-based approach (mirroring `RegisterForm`'s own success-heading-focus convention) but rejected it: `VerifyEmailStatus`'s own precedent for this same "outcome resolves, move focus" situation keeps one stable heading and moves focus via the banner/heading depending on the state, and reusing one single `bannerRef` for every branch (rather than a heading per branch) avoids introducing a second heading string per outcome that would just duplicate the page's own `<h1>Reset Password</h1>`.
3. **`forgotPassword()`'s `422` handling loops over `result.errors` checking `field === "email"`** rather than assuming there's exactly one error at index 0 — defensive, and matches `RegisterForm`'s existing loop-based convention for its own (larger) set of possible fields, even though this endpoint only has one possible field in practice.

## Open items carried forward (unresolved by this build, same as techplan §14)

None — the techplan's Active list was already empty at Approved status (all items resolved during planning, techplan §14 Resolved #6-#8: copy adopted, regen sequencing decided, backend follow-up scoped as non-blocking). No new open item surfaced during the build itself; everything encountered (the unrelated schema diff, the focus-management mechanism choice) was either already anticipated in the techplan or resolved as a small, self-contained implementation judgment call above.
