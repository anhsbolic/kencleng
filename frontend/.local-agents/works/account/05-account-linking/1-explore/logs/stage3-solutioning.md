# Stage 3 — Solutioning

Feature: `docs/spec/1-account/features/05-account-linking.md`
(frontend surface). Builds directly on all six Stage 2 area docs in
this same directory — decisions below cite which finding each one
follows. Two items below are flagged for Anhar's confirmation before
this becomes a techplan, following the same convention Task #4's
Stage 3 used for its shell-placement decision.

## Decision 1 (flag for confirmation) — Page composition: independent
sections, not a monolithic form

**Recommended: `/dashboard/security` is composed as a flat JSX list of
independent, self-contained section components inside the page file,
not a single monolithic form.** This task builds and renders exactly
one section (`LoginMethodsSection`, covering set-password + Google
link/unlink). Task #6 (MFA), whenever it starts, adds its own
`MfaSection` as a sibling line in the same page file — a small,
easily-mergeable addition, not a structural rework of anything this
task builds. No throwaway placeholder MFA section is built now (avoids
Stage 2 Area 3's "wasted rework" risk); the page looks
narrower-than-final until Task #6 lands, which is accepted as correct
for the domain's current actual state (MFA genuinely doesn't exist
yet) rather than something to visually fake completeness around.

```tsx
// app/(dashboard)/dashboard/security/page.tsx
export default function SecurityPage() {
  return (
    <div className="flex flex-col gap-6">
      <h1 className="text-2xl font-bold text-neutral-900">Keamanan</h1>
      <LoginMethodsSection />
      {/* Account Task #6 (MFA) adds <MfaSection /> here as a sibling
          when that task starts — see stage2-area3-dashboard-shell-
          page.md for why this page is split across two tasks. */}
    </div>
  );
}
```

This resolves Stage 2 Area 3's biggest open finding (one page-map row,
two independently-scheduled tasks, no stated composition plan) with the
lowest-friction option: whichever task runs second only ever adds a
sibling line, never edits the first task's component internals,
matching `nav-items.ts`'s own "data, not structure" extensibility
philosophy already established in this codebase for the same Shell.

## Decision 2 (flag for confirmation) — This task's page also wires up
`GoogleAuthButton intent="link"`

**Recommended: yes, include it.** `page-map.md`'s `/dashboard/security`
row says "link/unlink Google identity" as one action — the "link"
half's backend (`intent=link` on `GET /auth/google/redirect`) already
shipped under Task #2, but no page anywhere renders the trigger for it
(Stage 2 Area 4). Since this task is the one actually building this
page's content, and no other task claims this UI element, leaving it
unwired would mean `page-map.md`'s stated action is never implemented
by anyone. This is a deliberate scope inclusion beyond
`05-account-linking.md`'s own two listed endpoints — flagged explicitly
per `docs/spec/README.md` §6.4 rather than silently expanded, since the
feature spec itself doesn't mention it.

## Decision 3 — Branch-selection UI heuristic, and a spec ambiguity it
surfaces (flag for Anhar, not resolved here)

`LoginMethodsSection` decides which form to show using
`user.auth_providers` + `user.email_verified` (both already available
via `useAccountMe()`, no new query):

| `auth_providers` has `email_password`? | `email_verified`? | UI shown |
|---|---|---|
| No | — | `SetPasswordForm mode="add"` (Branch 1) |
| Yes | No | "Menunggu verifikasi email" notice, no form |
| Yes | Yes | `SetPasswordForm mode="change"` (Branch 2) |

**Spec ambiguity surfaced while designing this** (not found in Stage 2 —
found here, while working out exactly what "has an email_password
identity" means for branch selection): the feature spec's own branch-
selection rule says set-password branches "based on whether the
authenticated `user_id` **currently has** an `email_password`
`AuthIdentity`" — not qualified by verification status. But Branch 2's
own acceptance criteria justify skipping re-verification specifically
"since the email was already verified." **These two statements are in
tension** for the exact mid-flow state this task's 3-step design
creates: a user who submitted Branch 1 but hasn't clicked the
verification link yet now *does* have an `email_password`
`AuthIdentity` (just an unverified one). If the backend's actual branch
check is "any identity, verified or not," a second `set-password` call
during this window would route into Branch 2 and silently update the
unverified identity's credential "immediately, no verification needed"
— contradicting Branch 2's own stated justification for skipping
verification. This is a backend-spec question, not something the
frontend can resolve — flagging per `docs/spec/README.md` §6.3 (an
implementing agent must not edit the spec to make its own code pass)
rather than guessing further.

**Why this doesn't block the frontend design above**: the table's
heuristic is a **proactive UX nicety only** — the actual branch a
request lands in is determined entirely server-side and reported back
via the response's status code (`200` vs `202`), which is what
`useSetPassword`'s `onSuccess` actually branches on (Decision 4, below)
— not the frontend's own pre-submit guess. If the backend resolves the
ambiguity in either direction, this frontend code needs no change: the
heuristic table only ever affects *which fields are shown before
submit*, and even if it's occasionally wrong for the narrow mid-flow
window above, the post-submit handling stays correct because it trusts
the response, not the guess.

## Decision 4 — `client.ts`: expose `.type`, additively

`ApiError` gains an optional `type` field; a new `readProblem()` helper
returns both `.detail` and `.type` from one parse, and `readProblemDetail`
becomes a thin wrapper over it (existing callers/behavior unchanged):

```ts
// client.ts
export interface ProblemDetail {
  detail?: string;
  type?: string;
}

export async function readProblem(res: Response): Promise<ProblemDetail> {
  try {
    const body: unknown = await res.json();
    const detail = (body as { detail?: unknown } | null)?.detail;
    const type = (body as { type?: unknown } | null)?.type;
    return {
      detail: typeof detail === "string" ? detail : undefined,
      type: typeof type === "string" ? type : undefined,
    };
  } catch {
    return {};
  }
}

export async function readProblemDetail(res: Response): Promise<string | undefined> {
  return (await readProblem(res)).detail;
}

export class ApiError extends Error {
  status: number;
  detail?: string;
  type?: string;
  constructor(status: number, detail?: string, type?: string) {
    super(detail ?? `Request failed with status ${status}`);
    this.name = "ApiError";
    this.status = status;
    this.detail = detail;
    this.type = type;
  }
}
```

Purely additive — every existing `new ApiError(status, detail)` call
site keeps compiling unchanged (`type` is optional, last positional
param). This is shared infra (Stage 2 Area 1 flagged it as
low-blast-radius, in-bounds for `frontend/` as transport plumbing, not
business logic).

## File plan

### New files

| File | Purpose |
|---|---|
| `lib/hooks/use-set-password.ts` | `useSetPassword()` — see Decision 5. |
| `lib/hooks/use-unlink-google.ts` | `useUnlinkGoogle()` — see Decision 5. |
| `components/features/account/set-password-schema.ts` | `addPasswordSchema` (`email`+`password`) and `changePasswordSchema` (`current_password`+`password`), both reusing `register-schema.ts`'s length-only password rule (Stage 2 Area 4). |
| `components/features/account/set-password-form.tsx` | `mode: "add" \| "change"` prop, renders the branch-appropriate fields, calls `useSetPassword()`. See Component design. |
| `components/features/account/unlink-google-form.tsx` | Password-confirmation field + `Button variant="destructive"`, calls `useUnlinkGoogle()`, maps the two `409` `.type` values to distinct frontend-owned copy. |
| `components/features/account/google-identity-control.tsx` | Picks between `<GoogleAuthButton intent="link" />`, a proactive-block notice, or `<UnlinkGoogleForm />` per Decision 3's table (extended with the `hasGoogle` axis). |
| `components/features/account/login-methods-section.tsx` | The page section: heading + `SetPasswordForm`/pending-notice + `GoogleIdentityControl`, reads `useAccountMe()` once and passes derived booleans down. |
| Four matching `*.test.tsx` files | See Testing plan. |

### Modified files

| File | Change |
|---|---|
| `lib/api/client.ts` | Decision 4 — additive `readProblem()`, `ApiError.type`. |
| `lib/api/account.ts` | Add `SetPasswordRequest`/`UnlinkGoogleRequest` type aliases, `SetPasswordResult`/`UnlinkGoogleResult` discriminated types, `setPassword()`, `unlinkGoogle()`. Purely additive. |
| `app/(dashboard)/dashboard/security/page.tsx` | Placeholder replaced per Decision 1. |
| `components/features/account/verify-email-status.tsx` | Small additive change — see Decision 6. |
| `mocks/handlers.ts` | Add `mockGoogleOnlyUser` fixture (`auth_providers: ["google"]`) usable via `server.use()` overrides (matching this file's existing per-test override convention), plus default handlers for both new endpoints. |

No changes needed to `lib/api/schema.d.ts` (already fully generated,
Stage 2 Area 1), `app/(dashboard)/layout.tsx`,
`_components/dashboard-shell-client.tsx`, or `_components/nav-items.ts`
(the "Keamanan" entry already points at the right, unchanged path).

## API layer design (`lib/api/account.ts`)

```ts
export type SetPasswordRequest = components["schemas"]["SetPasswordRequest"];
export type UnlinkGoogleRequest = components["schemas"]["UnlinkGoogleRequest"];

/**
 * POST /account/security/set-password. `202`/`200` are both real
 * successes distinguished by which branch fired server-side (never a
 * client-supplied flag, per the feature spec) — `branch` lets the
 * caller (useSetPassword's onSuccess) react differently without
 * re-deriving which case happened from status codes itself. `422`
 * follows register()/resetPassword()'s established rule: validation
 * failures are always a return branch, never thrown. `401` (Branch 2
 * wrong current_password) and everything else throw ApiError.
 */
export type SetPasswordResult =
  | { ok: true; branch: "added"; message?: string }
  | { ok: true; branch: "changed"; message?: string }
  | { ok: false; kind: "validation"; errors: ValidationErrorItem[] };

export async function setPassword(input: SetPasswordRequest): Promise<SetPasswordResult> {
  const res = await postAccountAction("/account/security/set-password", input);

  if (res.status === 202) {
    const body: { message?: string } = await res.json();
    return { ok: true, branch: "added", message: body.message };
  }
  if (res.status === 200) {
    const body: { message?: string } = await res.json();
    return { ok: true, branch: "changed", message: body.message };
  }
  if (res.status === 422) {
    const body: { errors?: ValidationErrorItem[] } = await res.json();
    return { ok: false, kind: "validation", errors: body.errors ?? [] };
  }

  throw new ApiError(res.status, await readProblemDetail(res)); // 401 (Branch 2 wrong password), network, 5xx
}

export type UnlinkGoogleResult = { ok: true; message?: string };

/**
 * POST /account/security/google/unlink. The two `409` cases share one
 * status code — the caller distinguishes them via the thrown
 * ApiError's `.type` (Decision 4), not `.detail` text-matching.
 */
export async function unlinkGoogle(input: UnlinkGoogleRequest): Promise<UnlinkGoogleResult> {
  const res = await postAccountAction("/account/security/google/unlink", input);

  if (res.ok) {
    const body: { message?: string } = await res.json();
    return { ok: true, message: body.message };
  }

  const problem = await readProblem(res);
  throw new ApiError(res.status, problem.detail, problem.type); // 401, 409 (two types), network, 5xx
}
```

## Hooks design

```ts
// use-set-password.ts
export function useSetPassword() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: setPassword,
    onSuccess: (result) => {
      if (!result.ok || result.branch !== "changed") return;
      // Branch 2 success revokes ALL of the user's refresh tokens
      // server-side (INV-account-05, no carve-out for this session's
      // own token — Stage 2 Area 2). Clearing the token here — the
      // same action useLogout takes in its own onSettled — is enough:
      // SessionGuardProvider already redirects to /login on any real→
      // null accessToken transition, "regardless of what caused it."
      // No new redirect call site needed.
      useAuthStore.getState().clearAccessToken();
      queryClient.clear();
      postAuthChannelMessage({ type: "logged-out" });
    },
  });
}

// use-unlink-google.ts
export function useUnlinkGoogle() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: unlinkGoogle,
    onSuccess: () => {
      // auth_providers loses "google" — the section must re-render
      // without the unlink action/notice.
      queryClient.invalidateQueries({ queryKey: accountKeys.me() });
    },
  });
}
```

Branch 1 (`added`) success deliberately has no `onSuccess` side-effect —
matches `useForgotPassword`/`useResetPassword`'s "the form decides what
to render" precedent (Stage 2 Area 2); no session or cache change is
implied by creating an unverified identity.

## Decision 6 — `VerifyEmailStatus`: small additive change for the
authenticated caller

Two changes, both additive, both harmless for the existing logged-out
registration caller:

1. On success, invalidate `accountKeys.me()` — so navigating back to
   `/dashboard/security` after verifying reflects the now-verified
   identity without a stale cache (Stage 2 Area 2's gap).
2. The terminal CTA link is conditioned on whether an access token
   currently exists (`useAuthStore`): logged-in → `/dashboard/security`
   ("Kembali ke Keamanan"); logged-out (today's only real caller,
   registration) → `/login` ("Masuk sekarang", unchanged).

```tsx
// verify-email-status.tsx, inside the component
const queryClient = useQueryClient();
const accessToken = useAuthStore((state) => state.accessToken);

useEffect(() => {
  if (verifyEmail.isSuccess) {
    queryClient.invalidateQueries({ queryKey: accountKeys.me() });
  }
}, [verifyEmail.isSuccess, queryClient]);

// in the "verified" outcome branch:
<Link href={accessToken ? "/dashboard/security" : "/login"} ...>
  {accessToken ? "Kembali ke Keamanan" : "Masuk sekarang"}
</Link>
```

No new page/component — reusing the existing `/verify-email` route
stays literally true to the spec's "reusing `POST /auth/verify-email`
unchanged" framing at the endpoint level, while fixing the
component-level gap Stage 2 Area 2 found (the page was built for a
different persona than this feature's actual step-2 caller).

## Component design

### `LoginMethodsSection`

```
user = useAccountMe().data
if (!user) → loading state (skeleton row, Pattern 3 idle-loading convention)

hasEmailPassword = user.auth_providers?.includes("email_password") ?? false
hasGoogle        = user.auth_providers?.includes("google") ?? false
verified         = user.email_verified ?? false

render:
  <section> (card, radius-lg, per design-guidelines.md)
    <h2>Metode Masuk</h2>
    {hasEmailPassword && !verified
      ? <PendingVerificationNotice />   // "Menunggu verifikasi email kamu."
      : <SetPasswordForm mode={hasEmailPassword ? "change" : "add"} />}
    <GoogleIdentityControl hasGoogle={hasGoogle}
                           canUnlink={hasEmailPassword && verified} />
  </section>
```

### `SetPasswordForm`

- `mode="add"`: fields `email` + `password`, schema `addPasswordSchema`.
  On submit success (`branch: "added"`): inline success banner ("Cek
  email kamu untuk verifikasi.") + form reset, matching
  `ForgotPasswordForm`'s idle→success swap convention (Stage 2 Area 4).
  On `401`: unreachable for this mode (Branch 1 never returns 401,
  Area 1) — not handled as a distinct branch.
- `mode="change"`: fields `current_password` + `password`, schema
  `changePasswordSchema`. On submit success (`branch: "changed"`):
  `useSetPassword`'s `onSuccess` already clears the session — this
  component renders nothing further, since `SessionGuardProvider`'s
  redirect fires before any success view would be meaningfully seen
  (matches the "terminal action, redirect" half of `patterns.md` §B's
  Success convention, not the toast half). On `401`: frontend-owned
  banner ("Password saat ini salah." — placeholder pending product
  copy, same `// TBD` treatment as every other hand-written string in
  this domain).
- Both modes: `422` maps `result.errors` to `form.setError("password",
  ...)`, reusing `WEAK_PASSWORD_MESSAGE`-shape copy from
  `reset-password-form.tsx`'s existing constant (not reinvented).
- Password field uses the same show/hide toggle sub-pattern as
  `LoginForm` (Stage 2 Area 4/6) — a candidate for extraction into a
  shared `PasswordInput` once a third instance needs it (see
  Assumptions).

### `GoogleIdentityControl`

```
if (!hasGoogle) → <GoogleAuthButton intent="link" label="Hubungkan ke Google" />
if (hasGoogle && !canUnlink) → proactive Banner, frontend-owned copy
  matching whichever of the backend's two 409 messages currently
  applies (computed client-side from the same booleans, avoiding a
  guaranteed-409 round trip — Stage 2 Area 1's flagged UX improvement):
  - no email_password at all → "only-identity"-style copy
  - email_password present but unverified → "unverified-identity"-style copy
if (hasGoogle && canUnlink) → <UnlinkGoogleForm />
```

### `UnlinkGoogleForm`

- One `password` field (re-auth confirmation) + `Button
  variant="destructive"` ("Lepas Tautan Google").
- On `401`: banner "Password salah." (placeholder copy).
- On `409`: maps `error.type` to the two spec-required distinct
  messages — this is the actual enforcement point; the proactive block
  above is UX-only and can be stale (e.g. another tab already changed
  state), so this mapping must exist regardless of how rarely it's
  expected to fire.
- On success: `useUnlinkGoogle`'s `onSuccess` invalidation causes
  `GoogleIdentityControl` to re-render into the "no Google" branch —
  no local success view needed in this component itself.

## Assumptions / open questions (per `docs/spec/README.md` §6.4)

1. **Backend-spec ambiguity flagged in Decision 3** — whether
   `auth_providers` includes an identity before `verified_at` is set,
   and whether Branch 2's server-side check itself requires
   verification or merely existence. **Needs confirmation from Anhar/
   backend before merge** — suggest checking
   `backend/internal/domain/account/service.go`'s branch-selection and
   `auth_providers` construction directly, same style as Task #4's
   Assumption #1. Frontend design is robust to either answer (see
   Decision 3's last paragraph), but the exact UI heuristic table may
   need a one-line adjustment once confirmed.
2. **`SetPasswordRequest.email` conditional-required-ness** (Stage 2
   Area 1) is enforced only by `addPasswordSchema` client-side, not by
   the generated type — accepted, matches `MfaDisableRequest`'s
   already-established shape for the same kind of branch-conditional
   request.
3. **`422`'s `field` value assumed to be the literal string
   `"password"`**, by direct analogy with `register`'s `field:
   "password"` — same class of assumption as Task #4's Assumption #1,
   same recommended verification method (grep the backend's validation
   error construction), not re-derived further here.
4. **Branch 2's post-success UX (immediate session cutover via
   `clearAccessToken`) is a frontend UX decision, not a spec
   requirement** — the spec only mandates backend-side revocation.
   Chosen for consistency with `useLogout`'s existing pattern and to
   close the "hijacked session's access token still works for ~15 more
   minutes" window (Stage 2 Area 2's flagged risk) as tightly as the
   existing session-management infrastructure allows. Flag for Anhar
   if a softer UX (stay on page, show a "sesi lain telah keluar"
   notice, let this tab's access token expire naturally) is preferred
   instead.
5. **No client-side guard against resubmitting Branch 1 while already
   mid-verification** — deliberately out of scope; the spec's
   acceptance criteria don't address this edge case, and Decision 3's
   heuristic already keeps the common path correct.
6. **`PendingVerificationNotice` has no resend-verification affordance**
   wired in this design — `ResendVerificationControl` already exists
   (Stage 1) and could plausibly be reused here, but the feature spec
   doesn't request it for this flow specifically (unlike registration's
   own verify-email page). Flagging as a possible small enhancement,
   not adding it speculatively.
7. **Password show/hide toggle extraction** (Stage 2 Area 4's
   observation) — this task is the second place needing it
   (`LoginForm` was the first) but only within forms this task itself
   builds (`SetPasswordForm`, `UnlinkGoogleForm` each need one
   instance). Recommend extracting a shared `PasswordInput` as part of
   this task's own scope (three total instances across two components
   built in this same task, plus `LoginForm`'s existing one, clears the
   "second domain needs it" bar from `phase0-shared-infra.md`'s
   Incremental Growth Rule) rather than deferring — but not deciding
   the exact API here; a small enough call that implementation can
   settle it directly.

## Risk note (full, per `AGENTS.md` §8 exception)

- **Highest-risk single piece**: the backend-spec ambiguity in Decision
  3 (whether an unverified identity already counts as "has an
  email_password identity" for branch selection). This isn't a
  frontend bug risk — the frontend design is deliberately built to
  trust the response over its own guess — but it's a real open question
  about backend behavior during the exact mid-flow window this feature
  introduces (submitted Branch 1, hasn't verified yet), and it should
  be confirmed before this task is considered fully specified, not
  just before this task's frontend is considered done.
- **Second risk**: Assumption #3 (the `"password"` field-name guess for
  `422` mapping) — same shape and same mitigation as Task #4's
  equivalent, silent-failure risk (message just doesn't attach to the
  field) rather than a crash, must be confirmed against real backend
  behavior before merge.
- **Third risk**: the `UnlinkGoogleForm`'s `409`-`.type` mapping is new
  code exercising a codepath (`ApiError.type`) nothing else in the
  codebase uses yet — low usage surface means a mistake in the mapping
  (e.g. swapped message-per-type) wouldn't be caught by any other
  component's tests. Mitigated by writing both `409` cases as explicit,
  separately-asserted test cases (see Testing plan), not inferring
  coverage from a single generic "409 shows an error" test.
- **Fourth risk**: Decision 1's page-composition approach (independent
  sections as sibling JSX) is a low-risk, low-commitment structural
  choice — worth confirming with Anhar since it's the first time this
  cross-task-shared-page pattern has come up in this domain, and
  whatever's decided here becomes the precedent Task #6 is expected to
  follow. Getting it wrong doesn't break anything already built;
  Task #6 would just need to restructure `SecurityPage` slightly if a
  different composition approach is later preferred, which is
  contained, reversible cost.
- **Residual/accepted risk**: Decision 2's inclusion of
  `intent="link"` and the "second identity every user must go through
  `/verify-email` again mid-session" flow (Decision 6) both extend
  slightly past `05-account-linking.md`'s literal endpoint list, in the
  direction the page-map's own action list requires. Both are flagged
  explicitly for Anhar per `docs/spec/README.md` §6.3/§6.4 rather than
  silently expanded or silently deferred.

## Testing plan

| Test file | Cases |
|---|---|
| `set-password-form.test.tsx` | `mode="add"`: client-side email/password validation; happy path (202 → success banner, form reset); 422 → field error on password, form stays. `mode="change"`: happy path (200 → no local success view, asserts `clearAccessToken`/`queryClient.clear` were triggered via the hook, not this component re-implementing it); 401 → banner "Password saat ini salah.", form stays; 422 → field error on password. |
| `unlink-google-form.test.tsx` | Happy path (200 → `useUnlinkGoogle`'s invalidate called); 401 → banner, form stays; 409 `only-identity` type → distinct banner A; 409 `unverified-remaining-identity` type → distinct banner B (both asserted as different text, not just "shows an error"); network failure → generic banner. |
| `google-identity-control.test.tsx` | No-Google case renders `GoogleAuthButton intent="link"`; Google-present-but-blocked case (both sub-cases: no email_password, and unverified email_password) renders the correct proactive notice text per sub-case; Google-present-and-unlinkable renders `UnlinkGoogleForm`. |
| `login-methods-section.test.tsx` | Composes the right children per `auth_providers`/`email_verified` combination (four states from Decision 3's table) using `server.use()` overrides on `GET /account/me`, mirroring `mockGoogleOnlyUser`. |
| `verify-email-status.test.tsx` (existing file, extended) | New cases: authenticated caller (accessToken present) sees "Kembali ke Keamanan" linking to `/dashboard/security` instead of "Masuk sekarang"; `accountKeys.me()` invalidation is called on success regardless of auth state. |

No dedicated hook test for `useUnlinkGoogle` (bare invalidate,
mirrors `useLogin`'s precedent of only hooks with the actual branching
logic under test getting their own file) — `useSetPassword`'s
`onSuccess` session-cutover logic is real branching logic, so it
inherits a slightly larger share of its coverage via
`set-password-form.test.tsx`'s `mode="change"` happy-path assertion
above rather than a separate hook test file, consistent with
`use-login.test.ts` being the only hook test file that currently exists
in this domain (for the same reason — `applyLoginSuccess` was the one
hook with real logic).
