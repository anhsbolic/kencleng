# Task 1 — API layer & hooks for login / MFA-challenge

> Derived from: `../techplan.md` ("Tech Plan: Login & Session Management
> (Frontend)", account/03-login-session-management). This task file
> redistributes §8-13 detail relevant to its own scope, in full — it does
> not summarize. For the Summary, §1-7 rationale (Background, Scope,
> Requirements, Decision Log, Backward Compatibility, Edge Cases & Risks),
> and §14 Open Items, read the source techplan directly; those sections
> are the contract and are not duplicated here.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **none** — this task can start immediately.
> Feeds into: Task 3 (Login form + MFA challenge UI) depends on this
> task's two hooks.
> Shared file note: `lib/api/account.ts` and `mocks/handlers.ts` are also
> touched by Task 4 (Logout) — this task owns the `login`/`loginMfa`
> additions to those files; Task 4 owns the `logout` addition. No
> overlapping lines, but sequence this task before Task 4 if applying
> both to avoid an avoidable merge conflict on the same file.
> Recommended model: **DeepSeek V4 Pro** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row
> ("Decomposed: GLM 5.2 (max) / DeepSeek V4 Pro per sub-task") and its
> own tie-breaker ("DeepSeek V4 Pro when it's rule-table-heavy/precision
> work without a diagram") — this task is pattern-following (mirrors
> `useRegister`/`register()`'s already-merged shape) with no branching
> logic or diagram of its own.

## Scope

Build the typed fetch functions and TanStack Query mutation hooks for
`POST /auth/login` and `POST /auth/login/mfa`, including the shared
success-handling side effects (store write + cache write). Add the two
corresponding default mock handlers.

**Rules owned by this task** (full text, copied from techplan §4 — this
task is responsible for the *hook-layer* behavior; the *component-layer*
behavior for the same rule numbers is Task 3's responsibility, using the
hooks this task produces):

- **R3** (password-step submit → success): Given valid credentials with
  no MFA enrolled, When `POST /auth/login` returns `200 status=ok`, Then
  `useAuthStore.setAccessToken(access_token)`,
  `queryClient.setQueryData(accountKeys.me(), user)` (D7), and redirect
  to `/dashboard/profile` (Open Item #1 in the source techplan — the
  redirect target is not yet locked; implement against `/dashboard/
  profile` and treat it as a one-line change if overridden later).
- **R4** (password-step submit → MFA required) — **hook-layer portion
  only**: Given valid credentials with MFA enrolled, When
  `POST /auth/login` returns `200 status=mfa_required`, Then the hook
  returns the `mfa_required` branch of its discriminated result
  unmodified — it must NOT treat this as an error (no `ApiError` thrown),
  and must NOT call `setAccessToken`/`setQueryData` for this branch (no
  cookie is set yet per spec 03). The *component*-layer response to this
  branch (transitioning `step` to `'mfa'`, storing `mfa_pending_token`)
  is Task 3's responsibility — this task only needs to pass the branch
  through correctly.
- **R7** (MFA-step submit → success): Given a valid `mfa_pending_token`
  and correct code, When `POST /auth/login/mfa` returns `200`, Then
  identical handling to R3 — same success handler, reused (not a second
  implementation).
- **R19** (mocks) — **this task's slice**: `mocks/handlers.ts` gains
  default handlers for `POST /auth/login` (`200 status=ok`) and
  `POST /auth/login/mfa` (`200`) — individual tests override via
  `server.use(...)` for every other branch (401/429/mfa_required),
  matching the existing `/auth/refresh` convention already in that file.

## Interface Contract (relevant subset of techplan §8)

**API contract consumed** (already built/tested — cross-checked directly
against `internal/transport/http/auth_login.go` and `internal/domain/
account/login.go`, not just the generated types):

```typescript
// POST /auth/login
// body: LoginRequest { email: string; password: string }
// 200 -> LoginResponse { status: "ok"; access_token: string; access_token_expires_at?: string; user: User }
//     -> LoginMfaRequiredResponse { status: "mfa_required"; mfa_pending_token: string }
//     (refresh token set as HttpOnly+Secure+SameSite=Strict cookie "kencleng_refresh" ONLY on the "ok" branch)
// 401 -> Problem { type, title, status: 401, detail: "Email atau password salah." }  (wrong credentials)
// 429 -> TooManyRequests  (same detail string as above — lockout)

// POST /auth/login/mfa
// body: LoginMfaRequest { mfa_pending_token: string; totp_code?: string; backup_code?: string }
// 200 -> LoginResponse (identical shape/handling to /auth/login's "ok" branch)
// 401 -> Problem (invalid code, or expired/malformed mfa_pending_token)
// 429 -> TooManyRequests (MFA-stage lockout, same generic detail)
```

**This task's exports:**
```typescript
// lib/api/account.ts
export function login(input: LoginRequest): Promise<LoginResult>;
// LoginResult = discriminated union, mirrors RegisterResult's already-merged shape:
//   { ok: true; status: "ok"; user: User; access_token: string; access_token_expires_at?: string }
// | { ok: true; status: "mfa_required"; mfa_pending_token: string }
// | never for 401/429 — those throw ApiError(status, detail), same as every other lib/api/account.ts function
export function loginMfa(input: LoginMfaRequest): Promise<LoginResult>;
// same LoginResult shape as login() — only ever resolves the "ok" branch (an mfa_required-shaped
// response is not a valid outcome of this endpoint per the OpenAPI contract), throws ApiError on 401/429

// lib/hooks/use-login.ts
export function useLogin(): UseMutationResult<LoginResult, ApiError, LoginRequest>;

// lib/hooks/use-login-mfa.ts
export function useLoginMfa(): UseMutationResult<LoginResult, ApiError, LoginMfaRequest>;
```

**This task's consumers must know:** Task 3's `LoginForm` component calls
`useLogin()`/`useLoginMfa()`, inspects the resolved `LoginResult`'s
`status` field to decide whether to transition to the MFA step or treat
it as success (the success side effects already ran inside this task's
`onSuccess`, Task 3 does not need to repeat them), and catches `ApiError`
for the 401/429 banner cases.

**Business logic flow (this task's slice, verbatim from §8):**
```
success handler (shared by both steps, R3/R7):
  setAccessToken(access_token)
  setQueryData(accountKeys.me(), user)   (D7)
  router.push('/dashboard/profile')       (Open Item #1)
```

## Architecture (relevant note from §9)

No new TanStack Query *query* hook is needed for login/MFA — both are
one-shot mutations (`useMutation`), consistent with `useRegister`'s
existing precedent (`lib/hooks/use-register.ts`).

## Implementation Details (verbatim from §10)

**File**: `lib/api/account.ts`
- Change: add `login`, `loginMfa` (R3, R4-hook-layer, R7), reusing the
  existing `postAccountAction` helper (already handles network-error
  normalization into `ApiError(0)` — no new low-level plumbing needed).
  Discriminated-result shape mirrors `register()`'s existing
  `RegisterResult` pattern in the same file — read that function first
  for the exact style to match (status-code branching, `ApiError` throw
  points).

**File**: `lib/hooks/use-login.ts` (new)
- `useMutation` wrapping `login`; `onSuccess` performs the shared success
  handler (D7, §8) when the result's `status` is `"ok"` — the
  `mfa_required` branch is handled entirely by `LoginForm`'s own state
  (Task 3), not this hook. Needs `useQueryClient()` (for `setQueryData`)
  and `useAuthStore.getState().setAccessToken` and `useRouter()` (for the
  redirect) — or `router.push` may instead be left to the calling
  component if a hook-level router dependency is undesirable; either is
  acceptable as long as R3's full behavior (store write + cache write +
  redirect) fires exactly once per successful "ok" result. Document
  whichever placement is chosen, since Task 3's component will call this
  hook and needs to know whether it or the hook performs the redirect.

**File**: `lib/hooks/use-login-mfa.ts` (new)
- Same shape as `use-login.ts`, wrapping `loginMfa`. Reuses the identical
  success-handler logic — do not reimplement it a second time; extract a
  small shared helper (e.g. `applyLoginSuccess(user, access_token, ...)`)
  called from both hooks' `onSuccess` if that keeps the two files from
  duplicating the store/cache/redirect sequence verbatim.

**File**: `mocks/handlers.ts`
- Change: add default handlers for `POST /auth/login` (`200
  status=ok`) and `POST /auth/login/mfa` (`200`) (R19). Follow the
  file's own stated convention (top-of-file comment: "One handler per
  endpoint, added as the page/component that needs it gets built... not
  speculatively ahead of demonstrated need") and the existing
  `/auth/refresh` handler's override pattern (`server.use(...)` per
  test for non-default branches).

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `lib/api/account.ts` | Modify | Add `login`, `loginMfa` (shared file — Task 4 also modifies, adding `logout`; no line overlap) |
| `lib/hooks/use-login.ts` | Add | Login mutation + success handler (D7) |
| `lib/hooks/use-login-mfa.ts` | Add | MFA-step mutation |
| `mocks/handlers.ts` | Modify | Add two default handlers (shared file — Task 4 also modifies, adding the `/auth/logout` handler; no line overlap) |
| `lib/api/account.test.ts` (or equivalent, per this task's test coverage below) | Add/Modify | Per Testing Checklist below |

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R3 (hook-layer): a mocked `200 status=ok` response, resolved by
  `useLogin()`, calls `setAccessToken`, writes `user` into the
  `account.me` cache via `setQueryData`, and triggers the redirect —
  test at the hook level (e.g. `renderHook` + MSW), not requiring
  `LoginForm` to exist.
- [ ] R4 (hook-layer): a mocked `200 status=mfa_required` response,
  resolved by `useLogin()`, does **not** call `setAccessToken` or
  `setQueryData`, and the hook's resolved value carries the
  `mfa_required` branch with the `mfa_pending_token` intact.
- [ ] R7: a mocked MFA-step `200`, resolved by `useLoginMfa()`, performs
  identical success handling to R3 (reuse the same assertion helper/test
  case shape as R3, don't write a second, drifted copy).
- [ ] R19: `mocks/handlers.ts`'s two new default handlers (`/auth/login`,
  `/auth/login/mfa`) resolve with the documented default shapes;
  confirm existing/new tests can override each via `server.use(...)`.

**Count-check** (this task's slice): 4 checklist items above, covering
R3 (hook-layer), R4 (hook-layer), R7, R19 — the *component*-layer halves
of R3/R4 are Task 3's checklist responsibility (see that task file), not
double-counted here.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Duplicating the success-handler logic (store write + cache write + redirect) separately in `use-login.ts` and `use-login-mfa.ts` instead of sharing one implementation | The two paths drift apart over time (e.g. one gets updated for a redirect-target change, the other doesn't) | Extract one shared helper, call it from both hooks' `onSuccess` |
| Treating a `mfa_required` response as an error and throwing `ApiError` from `login()` | Breaks Task 3's ability to distinguish "needs MFA" from "actually failed" — the whole point of the discriminated `LoginResult` | `mfa_required` is a valid, non-throwing branch of the resolved value, exactly like `RegisterResult`'s `kind: "validation"` branch is not a thrown error |
