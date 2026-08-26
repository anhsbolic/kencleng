import type { paths } from "@/lib/api/schema";

/**
 * The `intent` values `GET /auth/google/redirect` actually accepts —
 * derived from the generated schema's own query-param type, not
 * hand-written, so a value outside this union is a compile-time error
 * rather than a runtime `400` (techplan account/02-google-oauth-
 * login-register, R2). This is exactly the class of bug this
 * component exists to prevent: `RegisterForm` previously hand-wrote
 * `intent=register`, a value the backend has never accepted
 * (`validIntent()` only allows `login`/`link`/`reauth`).
 */
export type GoogleAuthIntent =
  paths["/auth/google/redirect"]["get"]["parameters"]["query"]["intent"];

export interface GoogleAuthButtonProps {
  /** Always "login" for this feature — "link"/"reauth" belong to a
   * different, session-authenticated flow (account linking / MFA
   * re-auth, out of this component's scope), typed here only so the
   * prop can't silently drift from the backend's own accepted set. */
  intent: GoogleAuthIntent;
  label: string;
}

/**
 * Shared Google OAuth entry point for `/login` and `/register` (R1-R3).
 * Renders a real `<a href>` navigation — never `apiFetch`/XHR — since
 * the endpoint issues a `302` that only a real browser navigation
 * follows correctly.
 */
export function GoogleAuthButton({ intent, label }: GoogleAuthButtonProps) {
  return (
    <a
      href={`/auth/google/redirect?intent=${intent}`}
      className="inline-flex h-11 w-full items-center justify-center gap-2 rounded-md border border-neutral-200 text-base font-semibold text-neutral-700 hover:bg-neutral-100"
    >
      {label}
    </a>
  );
}
