import { apiFetch, ApiError, readProblemDetail } from "./client";
import type { components } from "./schema";

export type User = components["schemas"]["User"];
export type RegisterRequest = components["schemas"]["RegisterRequest"];
export type VerifyEmailRequest = components["schemas"]["VerifyEmailRequest"];
export type ResendVerificationRequest =
  components["schemas"]["ResendVerificationRequest"];
export type ValidationErrorItem = { field: string; message: string };

/**
 * Discriminated-union result for `register()` — chosen (techplan
 * account/01-register-email-verification, Decision D4) over throwing
 * a typed validation error, so field-level (`422`) and request-level
 * (network/5xx/429) failures can never be conflated by a caller: only
 * a genuine request-level failure reaches `catch`, and only the
 * `kind: "validation"` branch ever exposes `errors`.
 */
export type RegisterResult =
  | { ok: true; message?: string }
  | { ok: false; kind: "validation"; errors: ValidationErrorItem[] };

/**
 * `GET /account/me` — the current authenticated user's profile,
 * including `roles` (see `lib/types/roles.ts` for how the frontend
 * derives its `GlobalRole` union from this). Mocked in
 * `mocks/handlers.ts` ahead of Account Task #7 actually shipping,
 * per the Mock-First Development Workflow — Dashboard Shell's nav
 * role-filtering is a real consumer today, not speculative.
 *
 * Throws a generic error on a non-OK response — callers (TanStack
 * Query hooks) surface that as a generic error state, never the raw
 * response body (`loading-empty-error-state-conventions.md`).
 */
export async function getMe(): Promise<User> {
  const res = await apiFetch("/account/me", { method: "GET" });
  if (!res.ok) {
    throw new Error("Failed to load account profile");
  }
  return res.json();
}

/**
 * Thin wrapper around `apiFetch` for this file's POST-only actions
 * (register/verify-email/resend, plus login/login-mfa/logout — added by
 * account/03): normalizes a fetch-level rejection (network down, DNS
 * failure — `fetch` itself throws, it doesn't resolve with a response)
 * into the same `ApiError` shape as an HTTP-level failure, so every
 * caller has exactly one error type to check `instanceof` against,
 * never a mix of `ApiError` and a raw `TypeError`. `body` is optional —
 * `logout` sends no request body, `JSON.stringify(undefined)` yields
 * `undefined`, which `fetch` treats as no body sent.
 */
async function postAccountAction(path: string, body?: unknown): Promise<Response> {
  try {
    return await apiFetch(path, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
  } catch {
    throw new ApiError(0);
  }
}

/**
 * `POST /auth/register` — always `202` regardless of which internal
 * branch actually fired server-side (new user / resend-verification
 * nudge / already-verified nudge / Google-only-conflict nudge), by
 * design (anti-enumeration; see the feature spec's Assumption A/B).
 * The only branch this function itself distinguishes is `422`
 * (password fails length policy or is breach-listed) — everything
 * else (network failure, `429`, unexpected `5xx`) throws `ApiError`,
 * never silently returned as if it were a validation result.
 */
export async function register(input: RegisterRequest): Promise<RegisterResult> {
  const res = await postAccountAction("/auth/register", input);

  if (res.status === 202) {
    const body: { message?: string } = await res.json();
    return { ok: true, message: body.message };
  }

  if (res.status === 422) {
    const body: { errors?: ValidationErrorItem[] } = await res.json();
    return { ok: false, kind: "validation", errors: body.errors ?? [] };
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}

/**
 * `POST /auth/verify-email` — unauthenticated; the token itself is the
 * credential. `200` resolves with the backend's own confirmation
 * message; every other status (`404` not found/used/revoked, `410`
 * expired, `429`, network/unexpected `5xx`) throws `ApiError` so the
 * caller can branch on `.status` (see
 * `components/features/account/verify-email-status.tsx`).
 */
export async function verifyEmail(
  input: VerifyEmailRequest
): Promise<{ message?: string }> {
  const res = await postAccountAction("/auth/verify-email", input);

  if (res.ok) {
    return res.json();
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}

/**
 * `POST /auth/verify-email/resend` — always `202` generic regardless
 * of whether the email matched anything (anti-enumeration by design,
 * consistent with `register`'s pattern even though this endpoint isn't
 * itself enumeration-critical). `429` and other failures throw
 * `ApiError`.
 */
export async function resendVerification(
  input: ResendVerificationRequest
): Promise<{ message?: string }> {
  const res = await postAccountAction("/auth/verify-email/resend", input);

  if (res.ok) {
    return res.json();
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}

export type LoginRequest = components["schemas"]["LoginRequest"];
export type LoginMfaRequest = components["schemas"]["LoginMfaRequest"];
export type LoginResponse = components["schemas"]["LoginResponse"];
export type LoginMfaRequiredResponse =
  components["schemas"]["LoginMfaRequiredResponse"];

/**
 * `POST /auth/login`'s success shape — `status: "ok"` carries the
 * session (cookie already set by the time this resolves) or
 * `status: "mfa_required"` carries only a short-lived
 * `mfa_pending_token` (no cookie yet). The backend's own `status` field
 * is the discriminant — deliberately not wrapped in a locally-invented
 * `ok`/`kind` union, since the wire shape already discriminates cleanly
 * (techplan account/03-login-session-management, task-01).
 */
export type LoginResult = LoginResponse | LoginMfaRequiredResponse;

/**
 * `POST /auth/login` — email_password credential check. `200` resolves
 * to either branch of `LoginResult`. `401` (wrong credentials) and
 * `429` (lockout) both throw `ApiError` carrying the **identical**
 * generic detail text the backend sends for either case (spec 03,
 * `errors.go`'s `problemDetailGenericCredential`) — callers never need
 * to branch on `.status` for copy, only for control flow.
 */
export async function login(input: LoginRequest): Promise<LoginResult> {
  const res = await postAccountAction("/auth/login", input);

  if (res.status === 200) {
    return res.json();
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}

/**
 * `POST /auth/login/mfa` — completes login via `totp_code`/
 * `backup_code` against a valid `mfa_pending_token`. Always resolves the
 * `LoginResponse` ("ok") shape on `200` — this endpoint never returns
 * `mfa_required` again. `401` (wrong code, or expired/invalid token) and
 * `429` (MFA-stage lockout) both throw `ApiError`, same generic-detail
 * treatment as `login`.
 */
export async function loginMfa(input: LoginMfaRequest): Promise<LoginResponse> {
  const res = await postAccountAction("/auth/login/mfa", input);

  if (res.status === 200) {
    return res.json();
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}

/**
 * `POST /auth/logout` — idempotent from the client's perspective (spec
 * 03: no refresh cookie present is not an error, `204`). No
 * discriminated result needed: the caller (`useLogout`) clears local
 * state unconditionally regardless of network outcome, so this function
 * only needs to resolve on success — a genuine failure (network error,
 * unexpected `5xx`) still throws `ApiError` via the usual path, but
 * `useLogout`'s `onSettled` never branches on it.
 */
export async function logout(): Promise<void> {
  const res = await postAccountAction("/auth/logout", undefined);

  if (!res.ok) {
    throw new ApiError(res.status, await readProblemDetail(res));
  }
}

export type ForgotPasswordRequest = components["schemas"]["ForgotPasswordRequest"];
export type ResetPasswordRequest = components["schemas"]["ResetPasswordRequest"];

/**
 * Discriminated-union result for `forgotPassword()` — mirrors
 * `RegisterResult`'s shape (techplan account/04, D3). The `202` branch is
 * always identical regardless of which internal case fired
 * (registered/unregistered/Google-only — anti-enumeration by design, R3).
 * The `422` branch is a defensive addition: confirmed directly against
 * `auth_password_reset.go` that the backend rejects a malformed email with
 * a real, spec-undocumented `422` (`{field:"email", ...}`) — client-side
 * `zod` should already prevent reaching it in normal use (R4).
 */
export type ForgotPasswordResult =
  | { ok: true; message?: string }
  | { ok: false; kind: "validation"; errors: ValidationErrorItem[] };

/**
 * `POST /auth/forgot-password` — see `ForgotPasswordResult`'s doc comment
 * for the `202`/`422` split. Everything else (`429`, network, unexpected
 * `5xx`) throws `ApiError`, never silently returned as a result branch.
 */
export async function forgotPassword(
  input: ForgotPasswordRequest
): Promise<ForgotPasswordResult> {
  const res = await postAccountAction("/auth/forgot-password", input);

  if (res.status === 202) {
    const body: { message?: string } = await res.json();
    return { ok: true, message: body.message };
  }

  if (res.status === 422) {
    const body: { errors?: ValidationErrorItem[] } = await res.json();
    return { ok: false, kind: "validation", errors: body.errors ?? [] };
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}

/**
 * `POST /auth/reset-password` — resolves on `200`, throws `ApiError` for
 * every other status (`404`/`410`/`422`/`429`/network/unexpected `5xx`),
 * matching `verifyEmail()`'s exact shape (techplan account/04, D2) rather
 * than a discriminated validation result: confirmed directly against
 * `service.go`'s `validatePassword` that the real `422` never carries
 * `errors[]` (a bare `account.ErrValidation` sentinel mapped to a fieldless
 * `Problem`), so there is no field-level data to model here. The caller
 * (`ResetPasswordForm`) branches on `.status` and renders frontend-owned
 * copy per branch — never the backend's raw `detail`, confirmed to be
 * unlocalized English placeholder text for every branch this endpoint can
 * return (D6).
 */
export async function resetPassword(
  input: ResetPasswordRequest
): Promise<{ message?: string }> {
  const res = await postAccountAction("/auth/reset-password", input);

  if (res.ok) {
    return res.json();
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}

export type SetPasswordRequest = components["schemas"]["SetPasswordRequest"];
export type UnlinkGoogleRequest = components["schemas"]["UnlinkGoogleRequest"];

/**
 * Discriminated-union result for `setPassword()` — techplan account/05-
 * account-linking, D4. `branch` distinguishes which server-side case
 * fired (`"added"` = Branch 1, generic `202`; `"changed"` = Branch 2,
 * `200`) so callers (`useSetPassword`'s `onSuccess`) don't need to
 * re-derive it from status codes themselves — branch selection is
 * entirely server-side (never a client-supplied flag), confirmed
 * verified-agnostic against the real backend (`security.go`). `422` is
 * a return branch, never thrown — same rule as `register`/
 * `resetPassword`'s `kind: "validation"` shape.
 */
export type SetPasswordResult =
  | { ok: true; branch: "added"; message?: string }
  | { ok: true; branch: "changed"; message?: string }
  | { ok: false; kind: "validation"; errors: ValidationErrorItem[] };

/**
 * `POST /account/security/set-password`. `401` (Branch 2 wrong
 * `current_password`) and network/5xx throw `ApiError` — its `.detail`
 * is the backend's own confirmed-correct Indonesian text
 * (`problemDetailGenericCredential`, shared with `login`'s own 401),
 * shown verbatim by callers (D4) — no frontend override needed there,
 * unlike the `422` branch.
 */
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

  throw new ApiError(res.status, await readProblemDetail(res));
}

export type UnlinkGoogleResult = { ok: true; message?: string };

/**
 * `POST /account/security/google/unlink` — resolves on `200`, throws
 * `ApiError` for `401` (wrong password) and `409` (two distinct cases,
 * INV-account-02/INV-account-12) alike. Both `409` cases already carry
 * distinct, correct, final Indonesian `.detail` text from the backend
 * (confirmed directly against `account_security.go`) — callers show it
 * verbatim, no `.type` field parsing needed (techplan account/05-
 * account-linking D4).
 */
export async function unlinkGoogle(input: UnlinkGoogleRequest): Promise<UnlinkGoogleResult> {
  const res = await postAccountAction("/account/security/google/unlink", input);

  if (res.ok) {
    const body: { message?: string } = await res.json();
    return { ok: true, message: body.message };
  }

  throw new ApiError(res.status, await readProblemDetail(res));
}
