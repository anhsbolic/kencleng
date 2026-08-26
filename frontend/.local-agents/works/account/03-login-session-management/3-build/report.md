# Build Report — Login & Session Management (account #03, frontend)

> Task      : account domain task #3, frontend surface —
>             `docs/spec/1-account/features/03-login-session-management.md`
> Executed  : 2026-08-26, tasks 1–4 per `2-plan/tasks/manifest.md`, plus a
>             same-day follow-up (§ below) not covered by the original
>             techplan/tasks — see "Follow-up: landing-page login/
>             register modal"
> Techplan  : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Draft — see Open Items below)
> Status    : Build complete, all tests/typecheck/lint/build green. No Tier 0 fenced sub-area in this plan (frontend has no JWT/token-crypto code) — nothing blocks committing on that front.

---

## Execution summary

| Task | Scope | Result |
|---|---|---|
| 1 | API layer & hooks for login/MFA-challenge (`login`, `loginMfa` fetch functions, `useLogin`/`useLoginMfa` hooks, 2 mock handlers) | ✅ 20 tests (`account.test.ts`, `use-login.test.ts`) |
| 2 | Session infrastructure: cross-tab refresh coordination (Web Locks + `BroadcastChannel`) + session-guard redirect | ✅ 20 tests (`client.test.ts`, `auth-channel.test.ts`, `session-guard-provider.test.tsx`, `auth-bootstrap-provider.test.tsx`) |
| 3 | Login form + MFA challenge UI (`LoginForm`, both steps) | ✅ 13 tests (`login-form.test.tsx`, `login/page.test.tsx`) |
| 4 | Logout entry point (`useLogout`, "Keluar" button) | ✅ 8 tests (`dashboard-shell-client.test.tsx`) |

Executed sequentially in one session (Task 1 → Task 2 → Task 3 → Task 4), not in parallel across agents — the manifest's Task 1/Task 2 parallel-eligibility was a scheduling option, not exercised here since one agent did all four in dependency order anyway.

## Files changed

**New**: `lib/api/auth-channel.ts` + `.test.ts` · `lib/api/client.test.ts` · `lib/hooks/{use-login,use-login-mfa,use-logout}.ts` + `use-login.test.ts` · `components/features/account/{login-form,login-schema}.tsx/.ts` + `login-form.test.tsx` · `components/providers/session-guard-provider.tsx` + `.test.tsx` · `mocks/fake-broadcast-channel.ts` (shared test double, see Deviations below).

**Edited**: `lib/api/account.ts` (+`login`/`loginMfa`/`logout`, `postAccountAction` doc comment) + `account.test.ts` · `lib/api/client.ts` (+`coordinatedRefresh`, `apiFetch`'s 401 handler now calls it) · `mocks/handlers.ts` (+3 handlers) · `components/providers/auth-bootstrap-provider.tsx` (+channel listener effect, boot hydration now calls `coordinatedRefresh`) + `.test.tsx` · `app/layout.tsx` (mount `SessionGuardProvider`) · `app/(auth)/login/page.tsx` (renders `LoginForm`) + `page.test.tsx` (rewritten — see Deviations) · `app/(dashboard)/_components/dashboard-shell-client.tsx` (+`LogoutButton`) + `.test.tsx`.

**Untouched by design**: `components/ui/input.tsx` (D6 — password show/hide is a local composition, not a shared-primitive change) · `lib/stores/auth-store.ts` (shape already correct, only new callers) · `app/(auth)/{forgot-password,reset-password}/page.tsx` (task #4) · `app/(dashboard)/dashboard/security/page.tsx` (MFA enrollment, tasks #5/#6) · `lib/api/schema.d.ts` (generated, already complete) · anything under `backend/` (read-only cross-checks only, per root `AGENTS.md` §7).

## Verification results

- **Unit/component suite**: `npx vitest run` — **127/127 tests, 27/27 files, all green** (up from the pre-build baseline of 111/111 — 16 net new tests across the touched areas, several existing files also gained cases: `account.test.ts` +9, `auth-bootstrap-provider.test.tsx` +3, `dashboard-shell-client.test.tsx` +4).
- **Typecheck**: `npx tsc --noEmit` — clean, 0 errors, at every checkpoint (after each of the 4 tasks, not just at the end).
- **Lint**: `npm run lint` (ESLint) — clean, 0 errors/warnings.
- **Production build**: `npm run build` (Next.js/Turbopack) — compiles successfully, all 12 routes (including `/login` and every `/dashboard/*` page) statically pre-render with no SSR crash — this specifically exercises the module-scope-construction risk flagged in the techplan's §7 (`BroadcastChannel`/`navigator.locks` referenced at module top-level would crash SSR); the lazy feature-detection design held up under a real build, not just under jsdom.
- **jsdom API-availability claim** (techplan §14 Resolved #6) re-confirmed live during this build: `client.test.ts`'s R12 test asserts `"locks" in navigator === false` directly against the real test environment, not a mock — still true on this run.

## Deviations from the plan (implementation-level judgment calls, not scope changes)

None of these change what the techplan/task files decided (D1-D7) — they're refinements made while writing the actual code, each individually small:

1. **`LoginResult` uses the backend's own `status` field as the discriminant**, not an `ok`-wrapped union as sketched illustratively in task-01's Interface Contract section. Simpler, avoids inventing a parallel field the wire shape doesn't have. `login()`/`loginMfa()`'s actual signatures are documented in `lib/api/account.ts` itself now.
2. **`login-form.test.tsx` created as a new file**, not anticipated as a separate file when the tasks were decomposed. Matches `register-form.test.tsx`'s existing precedent (component owns its own detailed test file; the page-level test stays a thin composition check) — `login/page.test.tsx` was slimmed to a composition-only smoke test instead of carrying full R1-R10 coverage itself, avoiding duplicating that coverage across two files.
3. **`mocks/fake-broadcast-channel.ts`** extracted as a shared test double, reused by `auth-channel.test.ts`, `auth-bootstrap-provider.test.tsx`, and `dashboard-shell-client.test.tsx` — avoids three copies of the same ~30-line fake class.
4. **`SessionGuardProvider` uses Zustand's own `(state, prevState)` subscription pair** instead of a manually-tracked `useRef` baseline as sketched in task-02's Implementation Details — simpler, and sidesteps a mount-ordering question the task file raised (whether the baseline needs to be seeded strictly after `AuthBootstrapProvider`'s effect) that turned out to be moot once traced through: `useAuthStore`'s initial state is always `null` before any hydration effect runs, so there was never a real race there — noted directly in the component's own doc comment rather than left as a latent concern.
5. **A "Kembali ke halaman login" link added to the MFA step**, not explicitly required by any rule (R1-R19). R8's own text ("the user must restart from the password step") implied an escape hatch should exist for the case where the `mfa_pending_token` has genuinely expired (indistinguishable from a wrong-code `401` per spec 03's anti-enumeration-adjacent design) — without it, a user in that state would be stuck retrying the same failing action indefinitely with no way back. Small, self-contained addition.
6. **`account.ts`'s `postAccountAction` doc comment and signature updated** (`body` is now optional) to accommodate `logout()`'s no-body request — a one-line, backward-compatible widening of an existing internal helper, not a new decision.

## Open items carried forward (unresolved by this build, same as techplan §14 Active)

1. **Redirect target after successful login** (`/dashboard/profile`) — implemented as recommended, still pending Anhar's explicit confirmation.
2. **Copy sign-off** — MFA-step labels ("Kode OTP", "Gunakan kode cadangan", etc.) and the generic error/placeholder strings are functional but not yet product-approved final copy.
3. **`page-map.md` Cross-Cutting UI Elements table has no logout row** — doc suggestion, not made as part of this build (out of scope, per the techplan's own §2).

No new open item surfaced during the build itself — everything encountered was already anticipated in the techplan's Decision Log or Edge Cases table, or resolved via the small implementation-level judgment calls listed above.

---

## Follow-up: landing-page login/register modal

Requested directly by Anhar after the build above shipped: "*register and login form are on different url path. i need them on landing page (public page), as modal.*" Not part of the original techplan/task files — no contract update was written for it (see the note at the end of this section), scoped and executed in the same session as a direct follow-up.

**Constraint surfaced before implementing**: the Google OAuth callback (task #02) has the backend hardcoded to redirect to `/login?error={code}` on failure — `/login`/`/register` can't be removed as routes. Presented Anhar a choice between (a) a client-state modal (routes untouched, modal has no shareable URL) and (b) Next.js intercepting routes (shareable modal URLs, new routing pattern this codebase hasn't used). **Anhar chose (a).**

**What shipped**:
- `lib/stores/auth-modal-store.ts` (new) — Zustand `mode: 'login' | 'register' | null`.
- `components/features/account/auth-modal.tsx` (new) — overlay + panel (desktop centered / mobile full-screen, CSS-only breakpoint switch matching `AuthShellClient`), focus-trapped on both breakpoints (unlike `AuthShellClient`, which only traps on desktop), Escape/backdrop-click/close-button all dismiss. Mounted once in `app/(public)/layout.tsx` as a sibling of `{children}` — not a separate route — so the actual landing page stays mounted and visible behind it.
- `components/features/account/auth-modal-triggers.tsx` (new) — small Client Component for the desktop nav's Masuk/Daftar buttons (the nav itself is a Server Component).
- `LoginForm`/`RegisterForm` — added optional `onSwitchToRegister`/`onSwitchToLogin` props: real `<Link>` navigation by default (unchanged standalone-route behavior), a mode-switching button when the prop is passed (modal context). Same components serve both `/login`/`/register` directly and the modal — no duplication.
- `app/(public)/layout.tsx`, `app/(public)/_components/public-shell-client.tsx` — desktop and mobile "Masuk"/"Daftar" triggers switched from `<Link href="/login|register">` to modal-opening buttons.
- `/login` and `/register` themselves: **unchanged** — still real routes, still work exactly as before if visited directly (confirmed by re-running the production build: both still appear as pre-rendered routes).

**Verification**:
- Full suite: **141/141 tests** (14 net new: `auth-modal.test.tsx` 8, 2 new cases each in `login-form.test.tsx`/`register-form.test.tsx` for the new prop, 2 new cases in `public-shell-client.test.tsx` for the button-vs-link switch).
- Typecheck, lint, production build: all clean (caught and fixed one self-inflicted JSDoc-merge syntax error in `auth-modal.tsx` before this counted as green).
- **Verified live, not just via jsdom**: no `chromium-cli`/Playwright available in this environment by default — installed Playwright + downloaded the Chromium/headless-shell binaries (network access confirmed available), launched `npm run dev` against the MSW mocks, and drove it with a small script: landing page → click Masuk → modal opens over the still-rendered page → click Daftar inside the modal → form switches, **`page.url()` confirmed still `http://localhost:3000/`** throughout, zero unexpected console/page errors (the one console entry was the already-expected silent 401 from `AuthBootstrapProvider`'s boot-time check). Three screenshots sent to Anhar directly. Dev server and Playwright's temp install cleaned up after.

**Process note**: this follow-up went straight from request → implementation → live verification in one pass, without a matching techplan/task-file update — reasonable for a same-day, contained UI refinement, but means `2-plan/techplan.md` and `2-plan/tasks/` now describe the pre-modal architecture only. If this domain's frontend track continues, worth a note (not necessarily a full new techplan) capturing the modal decision (client-state vs. intercepting routes, and why) somewhere more durable than this report, so it doesn't get lost the next time someone reads the techplan looking for how `/login` actually behaves.
