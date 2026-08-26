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
 * Thin wrapper around `apiFetch` for this file's three POST-only
 * unauthenticated actions: normalizes a fetch-level rejection (network
 * down, DNS failure — `fetch` itself throws, it doesn't resolve with a
 * response) into the same `ApiError` shape as an HTTP-level failure,
 * so every caller has exactly one error type to check `instanceof`
 * against, never a mix of `ApiError` and a raw `TypeError`.
 */
async function postAccountAction(path: string, body: unknown): Promise<Response> {
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
