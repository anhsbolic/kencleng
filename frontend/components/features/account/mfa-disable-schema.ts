import { z } from "zod";

/**
 * Single re-auth field — no length policy to enforce client-side (the
 * backend compares against the caller's own existing bcrypt hash, not a
 * new-password policy). Just required/non-empty — mirrors
 * `unlink-google-schema.ts`'s `unlinkGoogleSchema` exactly, same
 * re-auth-gated-destructive-action shape.
 */
export const mfaDisableSchema = z.object({
  password: z.string().min(1, "Password wajib diisi"),
});

export type MfaDisableFormValues = z.infer<typeof mfaDisableSchema>;
