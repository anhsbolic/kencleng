# Task 1: Data Layer & Shared Infrastructure

> Originating contract techplan: `../techplan.md` ("Tech Plan: Register &
> Email Verification (Frontend)", account/01-register-email-verification,
> Status: Draft). Cross-check high-level decisions there whenever this
> task file is ambiguous — this file redistributes only the sections
> relevant to this task's scope, it does not restate the whole plan.
>
> Splitting axis: Dependency/sequence chain (primary). This task has
> **no dependency on Task 2 or Task 3** — it is the prerequisite both of
> them build on. See `../manifest.md` for the full dependency graph.

## Scope

Build the frontend-side API wrapper functions, one mutation hook per
action, the MSW test handlers, and the one shared UI component
(`ResendVerificationControl`) that both Task 2 (`/register`) and
Task 3 (`/verify-email`) consume. **Nothing in this task touches a
route or page** — it is pure data-layer + one reusable leaf component.

**Out of scope for this task** (belongs to Task 2 / Task 3 / other
tickets — do not build here): the register form itself, the
verify-email page itself, the Google OAuth callback, `/login`/
`/forgot-password`/`/reset-password`.

## Background (condensed from techplan §1)

The backend for this feature is already shipped and stable
(`14834e5`) — `POST /auth/register`, `POST /auth/verify-email`,
`POST /auth/verify-email/resend` all exist and their contract is
already reflected in `lib/api/schema.d.ts` (generated,
`openapi-typescript`). This task adds the frontend-side wrapper layer
around that already-fixed contract — no backend change, no new wire
contract, this is purely additive frontend code.

## API contract being wrapped (already shipped — from `lib/api/schema.d.ts`, not authored by this task)

```typescript
// POST /auth/register
type RegisterRequest = { name: string; email: string; password: string };
// 202 -> GenericAcceptedMessage { message?: string }
// 422 -> ValidationProblem { ...Problem, errors?: { field: string; message: string }[] }
// 429 -> Problem

// POST /auth/verify-email
type VerifyEmailRequest = { token: string };
// 200 -> { message?: string }
// 404 -> Problem   (not found / already used / revoked-by-resend)
// 410 -> Problem   (expired)
// 429 -> Problem

// POST /auth/verify-email/resend
type ResendVerificationRequest = { email: string };
// 202 -> GenericAcceptedMessage
// 429 -> Problem
```

## What this task builds (from techplan §8/§10)

```typescript
// lib/api/account.ts — NEW additions (existing getMe() untouched)
type ValidationErrorItem = { field: string; message: string };
type RegisterResult =
  | { ok: true }
  | { ok: false; kind: "validation"; errors: ValidationErrorItem[] };

async function register(input: RegisterRequest): Promise<RegisterResult>;
// verifyEmail/resendVerification keep the existing getMe-style
// throw-on-!ok contract — neither has a field-level-error case to
// represent, so the simpler shape is correct here, not an oversight.
async function verifyEmail(input: VerifyEmailRequest): Promise<{ message?: string }>; // throws on !ok (404/410/429/network) — thrown error must carry the Problem body so callers can read status + detail
async function resendVerification(input: ResendVerificationRequest): Promise<{ message?: string }>; // throws on !ok
```

- **File**: `lib/api/account.ts` — add the three functions above.
  `register` returns the `RegisterResult` discriminated union (Decision
  D4 below); the other two throw on `!ok`.
- **File**: `lib/hooks/use-register.ts` (new) — `useRegister()`, a thin
  `useMutation({ mutationFn: register })` wrapper. No query-key factory
  needed — nothing is cached anywhere for an unauthenticated user at
  this point in the flow, so there is nothing to invalidate
  (`data-fetching-conventions.md` — confirmed explicitly, not silently
  skipped).
- **File**: `lib/hooks/use-verify-email.ts` (new) — `useVerifyEmail()`,
  same shape. **The single-fire guard (rule R12) is the caller's
  responsibility (Task 3), not this hook's** — this hook is a plain
  `useMutation` wrapper with no built-in dedupe.
- **File**: `lib/hooks/use-resend-verification.ts` (new) —
  `useResendVerification()`, same shape.
- **File**: `mocks/handlers.ts` — add three `http.post(...)` handlers
  for `/auth/register`, `/auth/verify-email`, `/auth/verify-email/resend`,
  default-happy-path responses per the contract above. Individual tests
  in Task 2/Task 3 override via `server.use(...)` for the 422/404/410/
  429/network-error cases — matches the existing `mockUser`/roles
  override convention already used for `GET /account/me` in this same
  file.
- **File**: `components/features/account/resend-verification-control.tsx`
  (new) — the shared "Kirim ulang" affordance (email input pre-filled
  where available + button), wrapping `useResendVerification()`.
  Consumed by both Task 2's `RegisterForm` and Task 3's
  `VerifyEmailStatus` (Decision D2) — build it generic over its
  trigger context (don't bake in register-specific or verify-email-
  specific copy/layout).

## Rules & Validation owned by this task

(Full numbering matches techplan §4 — R-numbers are not renumbered
per task, so cross-task references stay unambiguous.)

- **R6** (universal fallback): Given a response outside an endpoint's
  documented status set (network failure, unexpected `5xx`), When it
  occurs on any of the three actions, Then the wrapper function/hook
  surfaces it as a thrown error (verify/resend) or a value the caller
  can't mistake for `ok:true`/`kind:"validation"` (register) — **this
  task is responsible for making sure that distinction is impossible to
  get wrong at the type level**; the actual banner rendering happens in
  Task 2/Task 3, but the contract shape that makes it easy to get right
  is built here (Decision D4). Raw response body is never exposed
  beyond what's needed to read `Problem.detail`/`message` — no full
  response object leaks to a component that would render it unfiltered.
- **R8** (resend control contract): Given `ResendVerificationControl`
  receives an `email` prop and the user activates it, When clicked,
  Then it calls `resendVerification({ email })` — verified at this
  component's own test level; Task 2/Task 3 only need to verify they
  pass the right `email` in, not re-verify the call itself.
- **R9** (resend outcome uniform): Given any `202` response from
  `resendVerification`, When received, Then `ResendVerificationControl`
  shows the same generic confirmation text
  (`GenericAcceptedMessage.message`, e.g. *"Kalau email terdaftar,
  instruksi sudah dikirim."*) regardless of whether the email actually
  matched anything server-side — the component has no way to know or
  show otherwise, by construction.
- **R10** (429 handling, resend): Given a `429` on resend, When
  received, Then `ResendVerificationControl` shows the response's own
  `Problem.detail` text verbatim (this is a documented, backend-
  authored user-facing string — distinct from R6's fallback, which is
  for genuinely undocumented failures). The same `detail`-rendering
  pattern this component establishes should be mirrored by Task 2/
  Task 3 for their own R10 obligations (register submit, verify-email
  fetch) — implemented independently in each, since register/verify-
  email aren't wrapped by this shared component.

## Decision Log entries relevant to this task

**D4 — `register()`'s error-handling shape**

| Option | Why rejected/accepted |
|---|---|
| A. Throw a custom `ValidationError` subclass on 422, plain `Error` otherwise | Rejected — still routes both field-level and request-level failures through one `throw`/`catch` path; relies on every caller remembering an `instanceof` check. |
| B. Discriminated union return (`{ok:true} \| {ok:false, kind:"validation", errors}`), `throw` reserved for genuine request-level failure (**chosen**) | Encodes `patterns.md`'s "never conflate field-level and request-level errors" at the type level. Matches `form-validation-boundary.md`'s own worked example almost exactly. |
| C. Return the raw `Response`, parse in the hook | Rejected — pushes response-shape knowledge out of `lib/api/`, the layer that should own it once. |

## Backward Compatibility (from techplan §6, applies as-is)

- **Database**: N/A — no persistence layer in `frontend/`.
- **API**: No API changes — this task only adds new *consumers* of an
  already-shipped contract.
- **Existing clients/data**: none affected — confirmed via `grep` that
  zero existing frontend code calls any of these three endpoints today.

## Edge Cases & Risks relevant to this task

| Risk | Likelihood | Severity | Mitigation |
|---|---|---|---|
| `register()` copies the existing `getMe`/`getCampaigns` throw-only pattern instead of the D4 discriminated union | Medium — the existing pattern is the easiest thing to copy-paste | Medium — degrades UX downstream (422 shows as a generic banner instead of field errors) but doesn't break anything | Implement per D4 exactly as specified above; code review checks the return shape against the type, not just "it compiles" |
| A new MSW handler is missed for one of the 3 endpoints | Low | Low — `onUnhandledRequest: "error"` (already configured in `vitest.setup.ts`) fails any test hitting the gap loudly, not silently | Add all three handlers in this task, before Task 2/Task 3 need them |

## Files Changed / NOT Changed (this task's subset)

| File | Change Type | Description |
|---|---|---|
| `lib/api/account.ts` | Modify | Add `register`, `verifyEmail`, `resendVerification` |
| `mocks/handlers.ts` | Modify | Add 3 new `/auth/*` MSW handlers |
| `lib/hooks/use-register.ts` | Add | New mutation hook |
| `lib/hooks/use-verify-email.ts` | Add | New mutation hook |
| `lib/hooks/use-resend-verification.ts` | Add | New mutation hook |
| `components/features/account/resend-verification-control.tsx` | Add | Shared resend affordance, consumed by Task 2 + Task 3 |
| Corresponding `*.test.tsx`/`*.test.ts` for each file above | Add | Per Testing Checklist below |

| File | Reason untouched (this task) |
|---|---|
| `app/(auth)/register/page.tsx`, `app/verify-email/page.tsx` | Task 2 / Task 3's scope, not this task's |
| `lib/stores/auth-store.ts` | Confirmed no interaction needed anywhere in this feature — register/verify never issue a token |
| `components/ui/banner.tsx`, `input.tsx`, `button.tsx` | Already support everything needed; reused as-is |
| `lib/api/client.ts` (`apiFetch`) | Already endpoint-agnostic; no change needed |

## Testing Checklist (this task's subset)

- [ ] R6: `register()` returns `{ok:false, kind:"validation", errors}` on a mocked 422 (never throws for this case); `verifyEmail`/`resendVerification` throw on a mocked network error / unexpected 5xx, and the thrown error carries enough of the `Problem` body for a caller to read `detail`
- [ ] R8: `ResendVerificationControl`, given an `email` prop, calls `resendVerification({email})` on click (unit test, mocked hook or MSW)
- [ ] R9: `ResendVerificationControl` renders the same generic confirmation text on a mocked 202 regardless of the mock's simulated match/no-match
- [ ] R10: `ResendVerificationControl` renders the mocked 429 response's `Problem.detail` text verbatim

**Note on count-check**: this task's subset intentionally does not
cover all 18 rules from the parent techplan — see `../manifest.md` for
the full R1–R18 coverage map across all three tasks combined, which is
where the mandatory count-check against the parent's §4 is actually
verified (`rules.md` §4 applies at the manifest/whole-plan level for a
decomposed techplan, not per task file).

## Testing Examples & Common Mistakes (this task's subset)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Copying `getMe`/`getCampaigns`'s throw-only pattern for `register()` | 422 field errors collapse into one generic banner instead of per-field messages, discovered only downstream in Task 2 | Use the discriminated-union return (D4) — verify the *type*, not just that the function runs |
| Hardcoding the `202`/`429` copy into `ResendVerificationControl` | Copy silently drifts from the backend's actual (and potentially later-changed) response text | Always render the response's own `message`/`detail` field; only the MSW mock fixtures in tests should contain literal example strings |

## Open Items relevant to this task

- **Exact copy for the frontend-owned fallback banner (R6)** — no
  backend string exists to reuse on the fallback path, by design (the
  body is deliberately never inspected there). This task only needs to
  make sure the *contract shape* correctly distinguishes "fallback" from
  "documented outcome" — the actual banner copy is rendered by Task 2/
  Task 3/wherever the fallback is shown, and is listed as an Active Open
  Item on the parent techplan (`../techplan.md` §14, item 5), needing
  product sign-off before that copy is finalized.
