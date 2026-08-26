import { useMutation } from "@tanstack/react-query";
import { resendVerification } from "@/lib/api/account";

/** Wraps `POST /auth/verify-email/resend`. No cache to invalidate. */
export function useResendVerification() {
  return useMutation({
    mutationFn: resendVerification,
  });
}
