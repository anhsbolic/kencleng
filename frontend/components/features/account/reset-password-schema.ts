import { z } from "zod";

/**
 * Client-side validation is UX-only — the backend is the source of truth
 * (`form-validation-boundary.md`). Password length policy: same
 * length-only, no character-class requirement, as registration — see
 * `docs/spec/1-account/features/04-forgot-reset-password.md` and
 * `register-schema.ts`'s own identical comment. The breach-list check is
 * server-only and deliberately NOT replicated here — a compliant-looking
 * password can still come back as a `422` from the backend, surfaced as a
 * frontend-owned banner (`ResetPasswordForm`'s `WEAK_PASSWORD_MESSAGE`),
 * never a field-level error, since the real backend carries no field data
 * for this branch (techplan account/04, D2/D6).
 */
export const resetPasswordSchema = z.object({
  new_password: z.string().min(8, "Password minimal 8 karakter"),
});

export type ResetPasswordFormValues = z.infer<typeof resetPasswordSchema>;
