import { z } from "zod";

/**
 * Client-side validation is UX-only — the backend is the source of truth
 * (`form-validation-boundary.md`). The message here is deliberately the
 * same string `ForgotPasswordForm`'s defensive `422` handler falls back to
 * (techplan account/04, R2/R4/D6) — both paths mean the same thing to the
 * user ("this doesn't look like an email"), so they share one copy source
 * instead of drifting into two slightly different messages for the same
 * failure.
 */
export const forgotPasswordSchema = z.object({
  email: z.string().email("Format email tidak valid"),
});

export type ForgotPasswordFormValues = z.infer<typeof forgotPasswordSchema>;
