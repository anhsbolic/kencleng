# Task 1 — API layer, hooks, mocks & manual-entry parsing utility

> Derived from: `../techplan.md` ("Tech Plan: MFA TOTP (Frontend)",
> account/06-mfa-totp). This task file redistributes §8-13 detail
> relevant to its own scope, in full — it does not summarize. For the
> Summary, §1-7 rationale (Background, Scope, Requirements, Decision
> Log, Backward Compatibility, Edge Cases & Risks), and §14 Open Items,
> read the source techplan directly; those sections are the contract
> and are not duplicated here.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **none** — this task can start immediately.
> Feeds into: **Task 2** (Enroll flow UI) and **Task 3** (Disable flow
> UI) both import this task's hooks; neither can start until this task's
> real exports exist (do not stub fake hooks "to unblock" — both
> downstream tasks' own tests exercise the real hook + MSW mock path,
> `component-test-mocking-discipline.md`'s network-layer-mocking
> principle).
> Recommended model: **DeepSeek V4 Pro** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row
> ("Decomposed: GLM 5.2 (max) / DeepSeek V4 Pro per sub-task") and its
> own tie-breaker ("DeepSeek V4 Pro when it's rule-table-heavy/precision
> work without a diagram") — this task is pattern-following (mirrors
> `resetPassword()`/`setPassword()`/`unlinkGoogle()`'s already-merged
> wrapper shapes almost exactly) with no branching/diagram of its own.

## Scope

Build the 3 typed fetch wrapper functions (`mfaEnroll`, `mfaEnrollConfirm`,
`mfaDisable`), their 3 corresponding TanStack Query mutation hooks, the
pure `parseOtpauthSecret()` parsing utility (manual-entry fallback, D11),
and the 3 default MSW mock handlers + fixtures.

**Rules owned by this task** (full text, copied from techplan §4 — this
task is responsible for the *wrapper-function/hook-layer* behavior; the
*component-layer* behavior for the same rule numbers is Task 2's or
Task 3's responsibility, using the hooks/utility this task produces):

- **R6** (enroll `409`, defensive) — **wrapper-layer portion**: Given
  `enroll` returns `409` (undocumented in `schema.d.ts`, per the feature
  spec's already-enabled guard), When received, Then `mfaEnroll()` throws
  `ApiError(409, detail)`. The *component*-layer response (rendering the
  generic error banner) is Task 2's responsibility.
- **R7** (enroll network/5xx) — **wrapper-layer portion**: Given a
  network failure or unexpected `5xx` on enroll, When received, Then
  `mfaEnroll()` throws `ApiError` via the usual `postAccountAction` path
  — no discriminated-result branch, matching `resetPassword()`'s shape.
- **R9** (confirm success) — **hook-layer portion only**: Given confirm
  resolves `200` with `{ backup_codes }` (exactly 10 strings), When
  received, Then `useMfaEnrollConfirm`'s `onSuccess` invalidates
  `accountKeys.me()`; the resolved `backup_codes` are available via the
  mutation's own `.data`/resolved value for the caller to read — no
  separate plumbing needed for that part. The *component*-layer response
  (lifting the codes to `MfaSection` via a callback, replacing the QR UI)
  is Task 2's responsibility.
- **R10** (confirm `422`) — **wrapper-layer portion**: Given confirm
  returns `422`, When received, Then `mfaEnrollConfirm()` throws
  `ApiError` (per D8 — **not** a discriminated-result shape like
  `register`/`setPassword`, despite `schema.d.ts` typing this response as
  `ValidationError` — the feature spec is explicit there's no useful
  field-level distinction to preserve here). The *component*-layer
  response (fixed message, keeping the form mounted) is Task 2's.
- **R11** (confirm network/5xx) — **wrapper-layer portion**: same shape
  as R7, for `mfaEnrollConfirm()`.
- **R15** (disable, `email_password` success) — **hook-layer portion**:
  Given the correct current password, When `mfaDisable({ password })`
  resolves `200`, Then `useMfaDisable`'s `onSuccess` invalidates
  `accountKeys.me()`.
- **R16** (disable, `email_password` `401`) — **wrapper-layer portion**:
  Given the wrong password, When submitted, Then `mfaDisable()` throws
  `ApiError(401, detail)` — `.detail` passed through verbatim from the
  backend's `Unauthorized` response (undifferentiated in the schema, per
  Stage 2 Area 2's finding — there is only one message string to expect
  here, unlike `login`/`setPassword`'s dedicated credential-error text).
- **R18** (disable, Google-only success) — **hook-layer portion**: same
  invalidation shape as R15, for the no-body call variant.
- **R19** (disable, Google-only `401`) — **wrapper-layer portion**: same
  throw shape as R16 — the *component*-layer response (error banner +
  `GoogleAuthButton intent="reauth"` prompt, retry availability) is
  Task 3's responsibility.
- **R23** (mocks) — **this task's full scope**: Given no `server.use()`
  override, When any of the 3 endpoints is called in a test/dev-mode
  context, Then the default MSW handler returns the documented
  happy-path fixture (`mockMfaEnrollResponse`, `mockMfaEnrollConfirmResponse`
  with exactly 10 codes, `mockMfaDisableOk`).
- **R24** (manual-entry secret fallback) — **parsing-utility portion
  only**: Given an `otpauth://` URI, When `parseOtpauthSecret(uri)` is
  called, Then it extracts and returns the `secret` query parameter, or
  returns `null` if absent/unparseable (never throws — `TBD — verify`
  the real backend-generated URI always carries a `secret` param before
  treating its absence as impossible rather than just unhandled). The
  *component*-layer response (rendering it as selectable text next to
  the QR, hiding gracefully on `null`) is Task 2's responsibility.

## Interface Contract (relevant subset of techplan §8)

**API contract consumed** (already shipped, generated from
`api/openapi.yaml` into `lib/api/schema.d.ts` — not authored by this
task):

```typescript
// POST /account/security/mfa/enroll
type MfaEnrollResponse = { otpauth_uri: string };
// 200 only typed in schema.d.ts; 409 required by feature spec but NOT
// typed in schema.d.ts (Stage 2 Area 2 finding — spec/schema
// inconsistency, handled defensively anyway per R6).

// POST /account/security/mfa/enroll/confirm
type MfaEnrollConfirmRequest = { totp_code: string };
type MfaEnrollConfirmResponse = { backup_codes: string[] }; // exactly 10
// 200 | 422 (ValidationError shape, treated as plain throw per D8)

// POST /account/security/mfa/disable
type MfaDisableRequest = { password?: string }; // omitted for Google-only
// 200 { message?: string } | 401 (Unauthorized, undifferentiated)
```

**This task's exports:**

```typescript
// lib/api/account.ts
export type MfaEnrollResponse = components["schemas"]["MfaEnrollResponse"];
export type MfaEnrollConfirmRequest = components["schemas"]["MfaEnrollConfirmRequest"];
export type MfaEnrollConfirmResponse = components["schemas"]["MfaEnrollConfirmResponse"];
export type MfaDisableRequest = components["schemas"]["MfaDisableRequest"];

export async function mfaEnroll(): Promise<MfaEnrollResponse>;
export async function mfaEnrollConfirm(
  input: MfaEnrollConfirmRequest
): Promise<MfaEnrollConfirmResponse>; // throws ApiError on 409/422/network/5xx
export async function mfaDisable(
  input: MfaDisableRequest
): Promise<{ message?: string }>; // throws ApiError on 401/network/5xx

// lib/hooks/use-mfa-enroll.ts
export function useMfaEnroll(): UseMutationResult<MfaEnrollResponse, ApiError, void>;

// lib/hooks/use-mfa-enroll-confirm.ts
export function useMfaEnrollConfirm(): UseMutationResult<MfaEnrollConfirmResponse, ApiError, MfaEnrollConfirmRequest>;

// lib/hooks/use-mfa-disable.ts
export function useMfaDisable(): UseMutationResult<{ message?: string }, ApiError, MfaDisableRequest>;

// lib/otpauth.ts
export function parseOtpauthSecret(otpauthUri: string): string | null;
```

**This task's consumers must know:** Task 2's `MfaEnrollFlow` calls
`useMfaEnroll()`/`useMfaEnrollConfirm()` and `parseOtpauthSecret()`; Task
3's `MfaDisableForm` calls `useMfaDisable()`. Both downstream tasks catch
`ApiError` for their own banner rendering — this task's wrapper functions
never render UI, only resolve/throw.

**Business logic flow (this task's slice, verbatim from §8):**

```
mfaEnroll()            -> 200: resolve MfaEnrollResponse
                        -> 409/network/5xx: throw ApiError (R6/R7)
mfaEnrollConfirm(input) -> 200: resolve MfaEnrollConfirmResponse, useMfaEnrollConfirm's
                                onSuccess invalidates accountKeys.me() (R9-hook)
                        -> 422/network/5xx: throw ApiError (R10/R11)
mfaDisable(input)       -> 200: resolve {message}, useMfaDisable's onSuccess
                                invalidates accountKeys.me() (R15/R18-hook)
                        -> 401/network/5xx: throw ApiError (R16/R19)
parseOtpauthSecret(uri) -> secret string | null, never throws (R24)
```

## Architecture (relevant note from §9)

None of the 3 hooks need `useSetPassword`'s extracted-standalone-function
treatment (see the source techplan's account/05-account-linking
precedent) — no divergent-branch logic to unit-test in isolation; each
hook is a bare `useMutation` with at most a one-line `onSuccess`
invalidation, following `useUnlinkGoogle`'s simplest existing shape
directly.

## Implementation Details (verbatim from §10)

**File**: `lib/api/account.ts`
- Add `export type MfaEnrollResponse`, `MfaEnrollConfirmRequest`,
  `MfaEnrollConfirmResponse`, `MfaDisableRequest` (all
  `components["schemas"][...]` re-exports, not hand-written).
- Add `mfaEnroll()`, `mfaEnrollConfirm(input)`, `mfaDisable(input)` —
  each a thin `postAccountAction` wrapper, `409`/`422`/`401` handled per
  R6/R10/R16/R19 (this task's wrapper-layer portions).

**File**: `lib/hooks/use-mfa-enroll.ts` (new)
- `useMfaEnroll()` — bare `useMutation({ mutationFn: mfaEnroll })`, no
  `onSuccess` (nothing on `User` changes until confirm succeeds).

**File**: `lib/hooks/use-mfa-enroll-confirm.ts` (new)
- `useMfaEnrollConfirm()` — `useMutation({ mutationFn: mfaEnrollConfirm,
  onSuccess: () => queryClient.invalidateQueries({ queryKey:
  accountKeys.me() }) })`.

**File**: `lib/hooks/use-mfa-disable.ts` (new)
- `useMfaDisable()` — same invalidation shape as
  `use-mfa-enroll-confirm.ts`.

**File**: `lib/otpauth.ts` (new)
- `parseOtpauthSecret(otpauthUri: string): string | null` —
  `new URL(otpauthUri).searchParams.get("secret")`, wrapped so a
  malformed URI never throws (returns `null` instead), matching
  `readProblemDetail()`'s existing "best-effort, never throws" convention
  in `lib/api/client.ts`.

**File**: `mocks/handlers.ts`
- Add `mockMfaEnrollResponse: MfaEnrollResponse` (a well-formed fake
  `otpauth://` URI — no real secret needs to be cryptographically valid
  for a mock, just parseable by `parseOtpauthSecret()` for Task 2's
  tests), `mockMfaEnrollConfirmResponse: MfaEnrollConfirmResponse`
  (**exactly 10** backup-code strings — get the count right, per §13's
  flagged common mistake), `mockMfaDisableOk = { message: "MFA berhasil
  dinonaktifkan." }` (verbatim from `schema.d.ts`'s own `@example`, no
  placeholder-copy uncertainty here) + 3 `http.post(...)` default
  handlers, following this file's existing per-endpoint convention
  exactly (named fixture constant, comment naming the source task, one
  default handler — individual tests in Task 2/3 override
  non-default branches via `server.use(...)`).

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `lib/api/account.ts` | Modify | Add `mfaEnroll`/`mfaEnrollConfirm`/`mfaDisable` + types (shared file — no other task modifies it) |
| `lib/hooks/use-mfa-enroll.ts` | Add | Mutation hook, no invalidation |
| `lib/hooks/use-mfa-enroll-confirm.ts` | Add | Mutation hook + cache invalidation |
| `lib/hooks/use-mfa-disable.ts` | Add | Mutation hook + cache invalidation |
| `lib/otpauth.ts` | Add | `parseOtpauthSecret()` — manual-entry fallback parsing (D11) |
| `mocks/handlers.ts` | Modify | Add 3 fixtures + 3 default handlers (shared file — no other task modifies it) |
| `lib/otpauth.test.ts` | Add | Unit tests for `parseOtpauthSecret()` (valid URI, missing `secret`, malformed URI) |
| `lib/api/account.test.ts` | Modify | Add `describe` blocks for the 3 new wrapper functions (same file already houses `register`/`login`/`loginMfa` tests) |

**Reason untouched** (relevant row from §11): `lib/api/client.ts` —
existing `postAccountAction`/`apiFetch` reused as-is, no new auth/CSRF
handling needed. `lib/api/schema.d.ts` — generated file, not
hand-edited, already has all 3 endpoints' types.

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R6 (wrapper-layer): a mocked `409` on enroll causes `mfaEnroll()`
  to throw `ApiError(409, detail)`
- [ ] R7 (wrapper-layer): a mocked network failure/`5xx` on enroll causes
  `mfaEnroll()` to throw `ApiError`
- [ ] R9 (hook-layer): a mocked `200` on confirm, resolved by
  `useMfaEnrollConfirm()`, invalidates the `account.me` query and
  resolves with `backup_codes` (10 items) intact
- [ ] R10 (wrapper-layer): a mocked `422` on confirm causes
  `mfaEnrollConfirm()` to throw `ApiError` (not a discriminated result)
- [ ] R11 (wrapper-layer): a mocked network failure/`5xx` on confirm
  causes `mfaEnrollConfirm()` to throw `ApiError`
- [ ] R15 (hook-layer): a mocked `200` on disable (with `password`),
  resolved by `useMfaDisable()`, invalidates the `account.me` query
- [ ] R16 (wrapper-layer): a mocked `401` on disable causes
  `mfaDisable()` to throw `ApiError` with `.detail` set from the
  response body
- [ ] R18 (hook-layer): a mocked `200` on disable (no body), resolved by
  `useMfaDisable()`, invalidates the `account.me` query
- [ ] R19 (wrapper-layer): a mocked `401` on disable (no-body call)
  causes `mfaDisable()` to throw `ApiError`
- [ ] R23: default MSW handlers return the documented fixtures with no
  `server.use()` override; `mockMfaEnrollConfirmResponse.backup_codes`
  has exactly 10 entries
- [ ] R24: `parseOtpauthSecret()` returns the correct secret for a
  well-formed URI, `null` for a URI missing `secret`, and `null` (never
  throws) for a malformed/non-URI string

**Count-check** (this task's slice): 12 checklist items above, covering
R6, R7, R9 (hook-layer), R10, R11, R15 (hook-layer), R16, R18
(hook-layer), R19, R23, R24 (parsing-layer) — the *component*-layer
halves of R6/R7/R9/R10/R11/R16/R19/R24 are Task 2's or Task 3's checklist
responsibility (see those task files), not double-counted here.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Treating `enrollConfirm`'s `422` like `register`'s discriminated-validation shape | `mfaEnrollConfirm()` would need to return `{ ok: false, errors: [...] }` instead of throwing, contradicting D8 and breaking Task 2's expected `try/catch` usage | Per D8, throw `ApiError` on `422` — don't model a `SetPasswordResult`-style union for this endpoint |
| Backup-codes fixture with fewer/more than 10 entries in `mocks/handlers.ts` | A "shows 10 codes" test (in this task or in Task 2/4) silently passes/fails against the wrong invariant | Confirm the fixture array has exactly 10 entries (matches the real spec requirement, not an arbitrary mock convenience) |
| Assuming every real `otpauth_uri` has a `secret` query param and letting `parseOtpauthSecret()` throw if it's missing/malformed | A malformed or unexpected URI shape (e.g. a future backend change) would crash the whole `MfaEnrollFlow` in Task 2, not just the manual-entry line | `parseOtpauthSecret()` must catch/return `null` defensively (R24) — never let it throw |
</content>
