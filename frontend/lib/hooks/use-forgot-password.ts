import { useMutation } from "@tanstack/react-query";
import { forgotPassword } from "@/lib/api/account";

/**
 * Wraps `POST /auth/forgot-password`. No cache to invalidate — nothing is
 * cached for an unauthenticated user at this point in the flow, same
 * reasoning as `useRegister`. No `onSuccess`/`onSettled` side-effect either:
 * this endpoint issues no session, so there's nothing for the hook to do
 * beyond resolving the result — `ForgotPasswordForm` decides what to render
 * (techplan account/04, D5's reasoning applies here too).
 */
export function useForgotPassword() {
  return useMutation({
    mutationFn: forgotPassword,
  });
}
