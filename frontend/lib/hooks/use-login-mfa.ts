import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { loginMfa } from "@/lib/api/account";
import { applyLoginSuccess } from "./use-login";

/**
 * Wraps `POST /auth/login/mfa` — completes the MFA challenge step (R7).
 * Reuses `applyLoginSuccess` from `use-login.ts` rather than
 * re-implementing the store/cache/redirect sequence a second time (see
 * that file's doc comment).
 */
export function useLoginMfa() {
  const queryClient = useQueryClient();
  const router = useRouter();

  return useMutation({
    mutationFn: loginMfa,
    onSuccess: (result) => applyLoginSuccess(result, queryClient, router),
  });
}
