# Implementation Report — account/02-google-oauth-login-register (frontend)

> Ticket    : account/02-google-oauth-login-register (frontend surface)
> Feature   : Google OAuth Login/Register — `/login` + `/register` Google
>             entry points, `?error={code}` handling, session hydration
> Date      : 2026-08-26
> Spec ref  : `docs/spec/1-account/features/02-google-oauth-login-register.md`
> Techplan  : `.local-agents/works/account/02-google-oauth-login-register/2-plan/techplan.md`
> Tasks     : `.local-agents/works/account/02-google-oauth-login-register/2-plan/tasks/` (2 task files + manifest)

---

## 1. Summary

Backend task #2 (`GET /auth/google/redirect`, `GET /auth/google/callback`)
shipped and merged (`efc1111`→`ce61841`) before this session started.
The frontend side was partially built and actively broken: task #1
had already wired a "Daftar dengan Google" button on `/register`, but
its `intent=register` query value is not one the merged backend
accepts (`400`) — and `/login` was still the raw Phase 0 placeholder
with no Google entry point, no error handling, and no bridge from the
OAuth callback's cookie-delivered tokens into the SPA's in-memory
session state at all.

This session built both decomposed tasks. Per the manifest's
component/module-boundary axis (no hard dependency between them), Task
1 (Google Auth Entry Points) and Task 2 (Session Bootstrap/Hydration)
were each implemented, tested, and verified independently — order
between them was arbitrary, both are done. Every rule in the
originating techplan's §4 (R1–R13) has at least one named test proving
it. No Open Item was resolved during this build (all four remain
Active — see §8); no deviation from the originating techplan's
Interface Contract was needed (see §6, empty by design).

---

## 2. Files changed

### New files (7)

| File | LoC | Task | Description |
|---|---|---|---|
| `components/features/account/google-auth-button.tsx` | 40 | 1 | Shared entry-point component — `intent` typed against `schema.d.ts`'s own literal union (R1-R3) |
| `components/features/account/google-auth-button.test.tsx` | 42 | 1 | R1-R3 coverage |
| `components/features/account/google-callback-error.tsx` | 59 | 1 | `/login`'s `?error={code}` banner — `google_email_conflict` distinguished, four codes collapsed to one fallback (R5-R7) |
| `components/features/account/google-callback-error.test.tsx` | 56 | 1 | R5-R7 coverage |
| `app/(auth)/login/page.test.tsx` | 22 | 1 | R4 composition coverage |
| `components/providers/auth-bootstrap-provider.tsx` | 57 | 2 | Silent-refresh-on-boot hydration provider (R8-R11, R13) |
| `components/providers/auth-bootstrap-provider.test.tsx` | 122 | 2 | R8-R11 coverage |

### Modified files (6)

| File | Task | Description |
|---|---|---|
| `components/features/account/register-form.tsx` | 1 | Replaced the broken inline `intent=register` anchor with `<GoogleAuthButton intent="login" .../>` (D1, D5) |
| `components/features/account/register-form.test.tsx` | 1 | Updated the R7 test's expected `href` from `intent=register` to `intent=login` |
| `app/(auth)/login/page.tsx` | 1 | Replaced the Phase 0 placeholder with heading + "coming soon" note + `<Suspense>`-wrapped error banner + Google button (R4) |
| `app/layout.tsx` | 2 | Mounted `AuthBootstrapProvider` inside `QueryProvider`/`MockingProvider` (R13) |
| `lib/api/client.ts` | 2 | Exported `tryRefreshOnce`; replaced the hand-written `RefreshResponse` type with the generated `components["schemas"]["RefreshResponse"]` (R12) |
| `mocks/handlers.ts` | 2 | Added `POST /auth/refresh`, defaulting to `401` (guest is the common case — this fires on every app load, not just post-OAuth) |

### Pre-existing changes (NOT this feature — out of scope, flagged)

| File | Note |
|---|---|
| `backend/internal/domain/account/*`, `backend/internal/platform/auth/*`, `backend/migrations/000006–000009*` | Pre-existing, backend-side, unrelated to this frontend session — visible in `git status` from a separate in-progress backend track (`03-login-session-management`), not touched here (directory-boundary rule, root `AGENTS.md` §7) |
| `docs/kencleng-agentic-workflow.md` | Pre-existing modification, unrelated to this feature |
| `lib/api/account.ts` | Already modified before this session started (task #1's work); untouched by this session — this feature needed no new function in `account.ts` (the two Google endpoints aren't JSON endpoints, and `/auth/refresh` is wrapped in `client.ts`, not `account.ts`) |

---

## 3. Routes delivered / changed

| Route | Shell | Pattern | Notes |
|---|---|---|---|
| `/login` | `AuthShellClient` (unmodified, reused as-is) | Form (Auth sub-variant) | Replaces the Phase 0 placeholder — Google button + error banner only; credential form is backend task #3's scope (D2) |
| `/register` | `AuthShellClient` (unmodified, reused as-is) | Form (Auth sub-variant) | Existing route from task #1; its Google button fixed (D1) and extracted into the shared component (D5) |

No new route was added — confirmed against the originating techplan
§9: success lands on the already-existing bare `/`, hydrated globally
by `AuthBootstrapProvider` in the root layout, not a dedicated landing
page.

Confirmed via `next build`: `/login` prerenders as static (`○ /login`
in the build output) — confirms the `<Suspense>` boundary around
`GoogleCallbackError`'s `useSearchParams()` call is correctly placed.

---

## 4. Rule coverage (R1–R13)

| Rule | Named test(s) | Status |
|---|---|---|
| R1 (correct intent value) | `google-auth-button.test.tsx` — "targets intent=login for the login label (R1)", "targets intent=login for the register label — never intent=register (R1)"; `register-form.test.tsx` — "renders 'Daftar dengan Google' as a real navigation link with intent=login... (R7; account/02 D1/R1)" | ✅ |
| R2 (shared, typed component) | `google-auth-button.test.tsx` — "both /login and /register call sites share this one component (R2)" (structural — see test comment for the enforcement mechanism) | ✅ |
| R3 (real navigation only) | `google-auth-button.test.tsx` — "renders a real navigation link, not a button, to the redirect-initiation endpoint (R3)" | ✅ |
| R4 (`/login` scope boundary) | `app/(auth)/login/page.test.tsx` — "shows a heading, the 'coming soon' note, and the Google button — no credential fields (R4)" | ✅ |
| R5 (error banner presence) | `google-callback-error.test.tsx` — "renders nothing when no error param is present (R5)" | ✅ |
| R6 (error code → copy mapping) | `google-callback-error.test.tsx` — "shows the distinct email-conflict message...", 4× "shows the shared generic retry message for {code}...", "falls back to the same generic message for an unrecognized code..." | ✅ |
| R7 (banner placement + focus) | `google-callback-error.test.tsx` — "renders as an alert (role) and moves focus into it on render (R7)" | ✅ |
| R8 (bootstrap hydration trigger) | `auth-bootstrap-provider.test.tsx` — "calls refresh exactly once on mount when accessToken is null (R8)", "does not call refresh if accessToken is already set before mount (R8)" | ✅ |
| R9 (hydration success) | `auth-bootstrap-provider.test.tsx` — "populates useAuthStore and invalidates the account.me query on success (R9)" | ✅ |
| R10 (hydration failure is silent) | `auth-bootstrap-provider.test.tsx` — "leaves accessToken null and renders no error/toast on refresh failure (R10)" | ✅ |
| R11 (at most one attempt) | `auth-bootstrap-provider.test.tsx` — "does not trigger a second refresh call on re-render within the same mount (R11)" | ✅ |
| R12 (generated type, not hand-written) | Type-level — `lib/api/client.ts`'s `RefreshResponse` is now `components["schemas"]["RefreshResponse"]`; verified by `tsc --noEmit` passing clean, not a dedicated runtime test (no behavioral change, only the type source) | ✅ |
| R13 (provider placement) | Structural — every `auth-bootstrap-provider.test.tsx` test renders the provider inside `QueryClientProvider` and exercises `useQueryClient()` (R9's invalidation); a wrong placement would throw on render rather than pass. Also confirmed live in `app/layout.tsx`'s actual provider stack. | ✅ |

**Count-check**: 13 rules in the techplan's §4, 13 rows above —
matched (per `rules.md` §4's mandatory check, carried over from the
techplan into this report).

---

## 5. Verification results

| Gate | Command | Result |
|---|---|---|
| Type-check | `npx tsc --noEmit` | ✅ clean |
| Unit/component tests (full suite) | `npm run test` | ✅ 85/85 passed, 22 test files — 18 new tests (4 new test files), 0 regressions in the pre-existing 67 |
| Lint | `npm run lint` | ✅ clean, no warnings |
| Build | `npm run build` | ✅ compiles, TypeScript passes, `/login` prerenders as static (confirms the `Suspense` boundary) |
| Contract check (mock fixtures vs. generated types) | Manual — `client.ts`'s refresh-response parsing now types against `lib/api/schema.d.ts`'s generated `components["schemas"]["RefreshResponse"]`, no hand-written parallel type (R12) | ✅ |
| Accessibility (manual review, no automated `jest-axe` gate configured in this project) | Focus management (R7) verified via a `toHaveFocus()`-equivalent assertion on the banner's focusable wrapper; `Banner`'s existing `role="alert"` reused as-is (pre-existing, already tested primitive) | ✅ (automated a11y gate not part of this project's `npm run verify` — see Risk note, consistent with task #1's own report) |

---

## 6. Process deviations (flagged for audit trail)

None. Every file built matches the originating techplan's §10
Implementation Details and each task file's own "What this task
builds" section exactly — no signature, prop shape, or file-path
deviation was needed during implementation. This is worth stating
explicitly (not just omitting the section) per root `AGENTS.md`'s
honesty-reporting principle: an empty deviations section means
verified-empty, not skipped.

---

## 7. Risk note

(Per root `AGENTS.md` §5's required PR structure.)

- **Assumptions made:**
  - None beyond what the originating techplan's §1 Background already
    confirmed via direct backend code reads (the `login` intent's
    net-new-account branch, the callback's hardcoded redirect targets,
    the refresh-token compatibility chain) — this build session made
    no new assumption of its own; it implemented exactly what the
    techplan and task files specified.
  - `GoogleAuthButton`'s `intent` prop is typed by deriving from
    `paths["/auth/google/redirect"]["get"]["parameters"]["query"]["intent"]`
    in the generated schema, rather than a hand-written union — this
    keeps the type tied to a single source of truth per
    `frontend/AGENTS.md` §3, and was the natural implementation of
    R2's "typed against the schema's own literal union" requirement
    (the techplan didn't prescribe the exact derivation mechanism,
    only the requirement).

- **Edge cases intentionally NOT handled (and why):**
  - `/login`'s email/password credential form — explicitly out of
    scope (D2; backend task #3, `03-login-session-management.md`, a
    separate, currently in-progress track).
  - `link`/`reauth` intents and the `/account/security` surface —
    explicitly out of scope (backend tasks #5/#6). The
    `/account/security` vs. `page-map.md`'s `/dashboard/security`
    mismatch remains unresolved (Open Item #2 — see §8), not this
    session's to fix.
  - Multi-tab session sync (`BroadcastChannel` or equivalent) — not
    built (Open Item #1 — see §8). A user completing Google login in
    one tab will not be reflected in another open tab until that tab's
    own reload/hydration.
  - Error-banner copy (R6) and `/login`'s "coming soon" note use
    placeholder Indonesian text, marked `TBD` via code comment,
    pending product sign-off (Open Item #4).

- **Concurrency assumptions:** `AuthBootstrapProvider`'s "at most one
  attempt" guard (R11) is a `useRef` boolean, intentionally
  React-render-cycle-scoped, not a server-side idempotency guarantee —
  the backend's own refresh-token rotation/reuse-detection
  (`INV-account-03`/`INV-account-04`, task #3's Tier-0 fenced scope)
  is the actual source of truth for correctness under real concurrent
  refresh attempts (e.g. two tabs racing); this guard only prevents a
  *client-side* double-fire (e.g. React's dev-mode effect
  double-invoke) from wasting one legitimate refresh attempt. Verified
  under a forced `rerender()` in `auth-bootstrap-provider.test.tsx`,
  not under true concurrent network races.

- **What is not tested, and why:**
  - No automated accessibility gate (`jest-axe` or equivalent)
    configured in this project — same as task #1's report notes.
  - Live end-to-end Google OAuth flow (button click → Google consent →
    real callback → real cookies → real hydration) is unverified in
    this session — component tests assert the button's `href` and the
    hydration provider's behavior against MSW-mocked `/auth/refresh`,
    not a real browser round-trip through Google. This is the same
    class of limitation task #1's report flagged for the register
    button, now extended to the full round-trip including hydration.
  - Per the techplan's own D7/Open Item context: even a fully correct
    live round-trip today may not authenticate cleanly against
    whatever endpoint-protection middleware backend task #3 ultimately
    ships, independent of anything in this frontend session — flagged
    in the techplan, not re-litigated here.

---

## 8. Open items status (techplan §14)

| # | Item | Status | Resolution |
|---|---|---|---|
| 1 | Multi-tab session sync (`BroadcastChannel`) | ⚠️ still open | Not built — scope decision needed from whoever reviews the techplan: build now vs. defer until `logout` (task #3) exists |
| 2 | `/account/security` vs. `/dashboard/security` mismatch | ⚠️ still open | Not this task's scope (link/reauth untouched); needs an owner before backend tasks #5/#6 start |
| 3 | `api/openapi.yaml`'s misleading `intent=login`/`register` prose | ⚠️ still open | Suggested doc fix not applied — shared cross-track file, not this session's to edit unilaterally |
| 4 | Exact copy for R6's error messages / `/login`'s "coming soon" note | ⚠️ still open | Placeholder Indonesian copy shipped, marked `TBD` in `google-callback-error.tsx` and `app/(auth)/login/page.tsx`, pending product sign-off |

No Open Item was resolved during this build — all four remain exactly
as the techplan left them.

---

## 9. How to run

```bash
# 1. Install deps (if not already)
npm install

# 2. Dev server (MSW mocks the account/campaign/notification endpoints —
#    no live backend needed to exercise /login or /register)
npm run dev

# 3. Full test suite
npm run test          # or: npx vitest run

# 4. This feature's tests only
npx vitest run \
  components/features/account/google-auth-button.test.tsx \
  components/features/account/google-callback-error.test.tsx \
  "app/(auth)/login/page.test.tsx" \
  components/providers/auth-bootstrap-provider.test.tsx \
  components/features/account/register-form.test.tsx

# 5. Type-check + lint + build (full gate)
npx tsc --noEmit
npm run lint
npm run build

# 6. Manually exercise the flow
#    /login    — see the Google button + (if you append ?error=google_email_conflict
#                 or any of the other documented codes) the mapped error banner
#    /register — "Daftar dengan Google" now targets intent=login (fixed)
#    Session hydration (AuthBootstrapProvider) fires silently on every
#    app load in dev; override `POST /auth/refresh` via `server.use(...)`
#    in a test, or point `NEXT_PUBLIC_API_MOCKING=true` dev mode's mock
#    handler, to exercise the success/failure branches.
```
