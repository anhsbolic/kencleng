# Implementation Report — account/01-register-email-verification (frontend)

> Ticket    : account/01-register-email-verification (frontend surface)
> Feature   : Register & Email Verification — `/register` + `/verify-email`
> Date      : 2026-08-26
> Spec ref  : `docs/spec/1-account/features/01-register-email-verification.md`
> Techplan  : `.local-agents/works/account/01-register-email-verification/2-plan/techplan.md`
> Tasks     : `.local-agents/works/account/01-register-email-verification/2-plan/tasks/` (3 task files + manifest)

---

## 1. Summary

The backend for this feature shipped weeks earlier (`14834e5`) with zero
frontend changes. This session built both frontend surfaces the feature
spec + `page-map.md` require: the `/register` form (name/email/password,
Google-register entry point, uniform anti-enumeration success state)
and a new `/verify-email` route (token-in-URL, four outcome states, a
shared resend affordance) — the latter didn't exist as a route at all
before this session, not even a placeholder.

All three decomposed tasks were executed in their dependency order
(Task 1 → Task 2 → Task 3, per the manifest's graph — Task 1 is a hard
prerequisite for the other two, which have no dependency on each other).
Every rule in the originating techplan's §4 (R1–R18) has at least one
named test proving it. One Active Open Item (`<Suspense>` boundary
requirement, #6) was resolved during Task 3's build and moved to the
techplan's Resolved list, with the resolution verified against a real
`next build`, not just asserted.

---

## 2. Files changed

### New files (12)

| File | LoC | Task | Description |
|---|---|---|---|
| `lib/hooks/use-register.ts` | 15 | 1 | `useRegister()` — thin `useMutation` wrapper, no cache to invalidate |
| `lib/hooks/use-verify-email.ts` | 14 | 1 | `useVerifyEmail()` — same shape |
| `lib/hooks/use-resend-verification.ts` | 9 | 1 | `useResendVerification()` — same shape |
| `components/features/account/resend-verification-control.tsx` | 79 | 1 | Shared "Kirim ulang" affordance (D2) — editable email field, pre-filled via `defaultEmail` where known |
| `components/features/account/resend-verification-control.test.tsx` | 72 | 1 | R8/R9/R10 coverage |
| `lib/api/account.test.ts` | 160 | 1 | `register`/`verifyEmail`/`resendVerification` behavior against MSW |
| `components/features/account/register-schema.ts` | 19 | 2 | `zod` schema — `name`/`email`/`password` (R1/R2) |
| `components/features/account/register-form.tsx` | 176 | 2 | Register form: fields, states, Google entry point, success view |
| `components/features/account/register-form.test.tsx` | 172 | 2 | R1–R7, R10, R17, R18 coverage |
| `app/verify-email/page.tsx` | 31 | 3 | New top-level route (D1), `<Suspense>`-wrapped |
| `components/features/account/verify-email-status.tsx` | 128 | 3 | Outcome view — R11–R16 |
| `components/features/account/verify-email-status.test.tsx` | 140 | 3 | R6, R10–R16 coverage |

### Modified files (4)

| File | Task | Description |
|---|---|---|
| `lib/api/client.ts` | 1 | Added `ApiError` (status + `detail`) and `readProblemDetail` — shared, endpoint-agnostic error-shape infra, not account-specific, so future domains can reuse it rather than re-inventing per-domain error types |
| `lib/api/account.ts` | 1 | Added `register`, `verifyEmail`, `resendVerification`, plus a local `postAccountAction` helper that normalizes a `fetch`-level rejection (network down) into the same `ApiError` shape as an HTTP-level failure |
| `mocks/handlers.ts` | 1 | Added 3 default happy-path handlers for `/auth/register`, `/auth/verify-email`, `/auth/verify-email/resend` |
| `app/(auth)/register/page.tsx` | 2 | Replaced the Phase 0 placeholder with a thin wrapper around `RegisterForm` |

### Pre-existing changes (NOT this feature — out of scope, flagged)

| File | Note |
|---|---|
| `backend/internal/domain/account/*`, `backend/migrations/000006–000009*` | Pre-existing, backend-side, unrelated to this frontend session — visible in `git status` from a separate in-progress backend track (`03-login-session-management`), not touched here (directory-boundary rule, root `AGENTS.md` §7) |
| `docs/kencleng-agentic-workflow.md` | Pre-existing modification, unrelated to this feature |

---

## 3. Routes delivered

| Route | Shell | Pattern | Notes |
|---|---|---|---|
| `/register` | `AuthShellClient` (unmodified, reused as-is) | Form | Replaces the Phase 0 placeholder |
| `/verify-email` | None — top-level, Status/Tracking-style minimal shell (Decision D1) | Status/Tracking (unauthenticated) | New route; deliberately **not** nested under `app/(auth)/` |

Confirmed via `next build`: both routes compile and prerender as
static (`○ /register`, `○ /verify-email` in the build output).

---

## 4. Rule coverage (R1–R18)

| Rule | Named test(s) | Status |
|---|---|---|
| R1 (register form fields) | `register-form.test.tsx` — "shows field-level validation errors on submit without touching any field (R1)" | ✅ |
| R2 (no client breach-check) | `register-form.test.tsx` — "accepts a password >= 8 chars locally with no breach-list check (R2)" | ✅ |
| R3 (submit button state) | `register-form.test.tsx` — "submit button has type=submit and disables while pending (R3)" | ✅ |
| R4 (register 202 → success) | `register-form.test.tsx` — "replaces the form with a fixed success view on 202, verbatim backend message (R4)"; `account.test.ts` — "resolves ok:true with the backend's own message on 202 (R4)" | ✅ |
| R5 (register 422 → field errors) | `register-form.test.tsx` — "maps each 422 field error verbatim via setError, no banner (R5)"; `account.test.ts` — "resolves ok:false with per-field messages on 422, never throws (R5)" | ✅ |
| R6 (universal fallback) | `register-form.test.tsx` — "shows a generic banner on a network failure (R6)"; `verify-email-status.test.tsx` — "shows a generic banner on a network failure (R6)"; `account.test.ts` — "throws a plain ApiError on a network failure (R6)" | ✅ |
| R7 (Google entry point) | `register-form.test.tsx` — "renders 'Daftar dengan Google' as a real navigation link, not a button (R7)" | ✅ |
| R8 (resend control contract) | `resend-verification-control.test.tsx` — "pre-fills the email field from defaultEmail and calls resend with it (R8)" | ✅ |
| R9 (resend outcome uniform) | `resend-verification-control.test.tsx` — "shows the same generic confirmation on 202 regardless of match (R9)"; `account.test.ts` equivalent | ✅ |
| R10 (429 handling) | Verified independently at all three consuming sites: `register-form.test.tsx`, `verify-email-status.test.tsx`, `resend-verification-control.test.tsx`, `account.test.ts` | ✅ |
| R11 (missing token) | `verify-email-status.test.tsx` — "shows a generic invalid-link message when the token is missing (R11)" | ✅ |
| R12 (single-fire guard) | `verify-email-status.test.tsx` — "fires verifyEmail exactly once per token, even under a forced re-render (R12)" | ✅ |
| R13 (verify 200) | `verify-email-status.test.tsx` — "shows the verified message plus a link to /login on 200 (R13)" | ✅ |
| R14 (verify 410) | `verify-email-status.test.tsx` — "shows the expired message plus a resend control on 410 (R14)" | ✅ |
| R15 (verify 404) | `verify-email-status.test.tsx` — "shows a generic invalid-link message on 404 (R15)" | ✅ (copy is a placeholder — Open Item #1 still active) |
| R16 (focus on verify-email resolution) | `verify-email-status.test.tsx` — "moves focus into the result heading once the loading state resolves (R16)" | ✅ |
| R17 (focus on register success) | `register-form.test.tsx` — "moves focus into the success heading once the form is replaced (R17)" | ✅ |
| R18 (no enumeration-defeating client check) | `register-form.test.tsx` — "never fires a request on email blur/change beyond the explicit submit (R18)" | ✅ |

**Count-check**: 18 rules in the techplan's §4, 18 rows above — matched (per `rules.md` §4's mandatory check, carried over from the techplan into this report).

---

## 5. Verification results

| Gate | Command | Result |
|---|---|---|
| Unit/component tests (full suite) | `npx vitest run` | ✅ 67/67 passed, 18 test files — 31 new tests, 0 regressions in the pre-existing 36 |
| Lint | `npm run lint` | ✅ clean, no warnings |
| Build | `npm run build` | ✅ compiles, TypeScript passes, both `/register` and `/verify-email` prerender as static |
| Contract check (MSW fixtures vs. generated types) | Manual — all three new `lib/api/account.ts` functions typed against `lib/api/schema.d.ts`'s generated `components["schemas"]`, no hand-written parallel types | ✅ |
| Accessibility (manual review, no automated `jest-axe` gate configured in this project) | `Input`/`Label` `htmlFor`/`aria-*` wiring reused as-is (pre-existing, already tested); focus-management (R16/R17) verified via `toHaveFocus()` assertions | ✅ (automated a11y gate not part of this project's `npm run verify` — see Risk note) |

---

## 6. Process deviations (flagged for audit trail)

### 6.1 `register()`'s success type gained a `message` field

The techplan's Interface Contract (§8) sketched `RegisterResult`'s
success branch as bare `{ ok: true }`. Rule R4 requires the success
view to display the response's own `GenericAcceptedMessage.message`
verbatim — unsatisfiable without carrying the message through. Changed
to `{ ok: true; message?: string }` during Task 1's build. No other
part of the plan depended on the narrower shape.

### 6.2 Network-level failures normalized into `ApiError`

Found while writing Task 1's R6 test: a `fetch`-level rejection
(`HttpResponse.error()` in MSW, modeling a real network failure) throws
a raw `TypeError`, not an HTTP-status response — so it never reaches
the `!res.ok`/`res.status` branches at all. Without normalization,
callers would need to check two different error shapes (`ApiError` for
HTTP failures, raw `Error`/`TypeError` for network failures) to
implement one R6 rule. Added `postAccountAction`, a local helper in
`lib/api/account.ts` that wraps the three POST calls and re-throws any
`fetch`-level rejection as `ApiError(0)` — one consistent type for
every caller's `instanceof ApiError` check.

### 6.3 `ResendVerificationControl`'s prop became `defaultEmail?: string`

The techplan (Task 1) specified the control's prop as `{ email: string
}`. Building Task 3 surfaced the reason this was too narrow:
`/verify-email`'s expired-token view has no email at all to offer —
only a token — while `/register`'s success view does. Changed the
shared component to hold its own editable email field, pre-filled via
an optional `defaultEmail` where the caller has one. This is a strict
superset of the original design (register's call site still gets a
pre-filled value; verify-email's call site gets a blank, user-fillable
field) and keeps both consumers using the exact same component rather
than forking behavior.

### 6.4 Open Item #6 resolved mid-build

`app/verify-email/page.tsx` wraps `VerifyEmailStatus` in `<Suspense>`
(required by `useSearchParams()` in the Next.js App Router). Verified,
not just asserted: `next build` (16.2.12) compiles cleanly and
`/verify-email` prerenders as static (`○ /verify-email`), no
Suspense-boundary error/warning. Moved from the techplan's Active to
Resolved Open Items list, and the Summary's Open Items line was
regenerated to match (`rules.md` §8, `guardrails.md` §11) — a
duplication slip made while editing that doc (the item briefly
appeared in both lists at once) was caught and corrected in the same
pass, not left for a later cleanup.

---

## 7. Risk note

(Per root `AGENTS.md` §5's required PR structure.)

- **Assumptions made:**
  - `RegisterRequest.name` is a required field (confirmed directly
    from `lib/api/schema.d.ts`, not inferred from `page-map.md`'s
    terser one-line description) — the register form collects it.
  - Every literal string rendered for a *documented* backend outcome
    (register's `202`, verify-email's `200`/`410`, resend's `202`) is
    read live from the response at runtime, never hardcoded into a
    component — the strings quoted in the techplan/this report are
    what the MSW mocks return in tests, not baked-in copy. Verified by
    `account.test.ts`'s exact-string assertions against the mocked
    responses.
  - The Google button (R7) is a plain `<a href>` navigation, not a
    `fetch`/XHR call — verified by asserting the rendered `href`
    attribute directly (`register-form.test.tsx`), not by mocking
    `window.location`.

- **Edge cases intentionally NOT handled (and why):**
  - `/verify-email`'s `404` outcome and the frontend-owned fallback
    banner (R6/R15) both render placeholder copy marked `TBD` in code
    comments — Open Items #1 and #5 on the techplan, pending product
    sign-off. Functionally correct (right banner variant, right
    trigger condition), just not final copy.
  - No resend affordance is offered on the `404` outcome (only on
    `410`) — Open Item #2, an explicit unresolved UX judgment call, not
    an oversight.
  - The "Daftar dengan Google" button's navigation target
    (`/auth/google/redirect?intent=register`) is not itself tested
    end-to-end against a real Google OAuth flow — Task #2
    (`02-google-oauth-login-register.md`) owns the callback side; this
    session only verifies the link's `href` is correct.

- **Concurrency assumptions:** `/verify-email`'s single-fire guard
  (R12, a `useRef` boolean) is intentionally React-render-cycle-scoped,
  not a server-side idempotency guarantee — the backend's own
  single-use token guard (`INV-account-08`'s 3-clause `UPDATE ... WHERE
  used_at IS NULL AND revoked_at IS NULL AND expires_at > now()`,
  already shipped and tested on the backend) is the actual source of
  truth; this guard only prevents a *client-side* double-fire (e.g.
  React's dev-mode effect double-invoke) from wasting the one
  legitimate attempt a real user gets. Verified under a forced
  `rerender()` in `verify-email-status.test.tsx`, not under true
  concurrent network races (that's the backend's own, already-covered
  concern, not this session's).

- **What is not tested, and why:**
  - No automated accessibility gate (`jest-axe` or equivalent) is
    configured in this project's `npm run verify` — focus management
    (R16/R17) is asserted directly via `toHaveFocus()`, but broader
    contrast/ARIA-tree correctness relies on the pre-existing,
    already-tested `Input`/`Label`/`Banner`/`Button` primitives, not a
    fresh audit of this session's new markup.
  - `Suspense` fallback content (`VerifyEmailLoadingFallback`) has no
    dedicated test — it's a two-line static string, verified only
    indirectly by the successful `next build`. Low risk given its
    triviality.
  - Cross-browser/real-network behavior of the Google redirect
    (`302` handling) is unverified in this session — component tests
    can only assert the `href`, not a real browser's navigation
    behavior.

---

## 8. Open items status (techplan §14)

| # | Item | Status | Resolution |
|---|---|---|---|
| 1 | 404 copy for `/verify-email` | ⚠️ still open | Placeholder copy shipped (`Link verifikasi tidak valid atau sudah digunakan.`), marked `TBD` in `verify-email-status.tsx` |
| 2 | Resend affordance on 404's "revoked" sub-case | ⚠️ still open | Not added — UX judgment call left to the human, per the techplan |
| 3 | Google button/navigation scope split (D3) vs. Task #2 | ⚠️ still open | Implemented per D3 (button + navigation now, callback later), but cross-task ratification with `tasks.md`/Task #2's owner not yet confirmed |
| 4 | D5's security reasoning (410 vs 404 distinction) | ⚠️ still open | Implemented per D5, but the security sanity-check itself is still pending human review |
| 5 | Fallback banner copy (R6) | ⚠️ still open | Placeholder copy shipped (`Terjadi kesalahan. Silakan coba lagi.`), marked `TBD` in both `register-form.tsx` and `verify-email-status.tsx` |
| 6 | `<Suspense>` requirement for `useSearchParams()` | ✅ resolved | See §6.4 above — verified via `next build`, moved to techplan's Resolved list |

---

## 9. How to run

```bash
# 1. Install deps (if not already)
npm install

# 2. Dev server (MSW mocks the account/campaign/notification endpoints —
#    no live backend needed to exercise /register or /verify-email)
npm run dev

# 3. Full test suite
npm run test          # or: npx vitest run

# 4. This feature's tests only
npx vitest run lib/api/account.test.ts \
  components/features/account/resend-verification-control.test.tsx \
  components/features/account/register-form.test.tsx \
  components/features/account/verify-email-status.test.tsx

# 5. Lint + build (full gate)
npm run lint
npm run build

# 6. Manually exercise the flow
#    /register            — fill the form, submit, see the success view
#    /verify-email?token=x — any non-empty token resolves via the
#                            mocked 200 in dev (mocks/handlers.ts);
#                            override with server.use(...) in a test to
#                            exercise the 410/404/429/error branches
```
