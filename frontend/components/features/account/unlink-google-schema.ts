import { z } from "zod";

/**
 * Single re-auth field — no length policy to enforce client-side (the
 * backend compares against the caller's own existing bcrypt hash, not
 * a new-password policy). Just required/non-empty.
 */
export const unlinkGoogleSchema = z.object({
  password: z.string().min(1, "Password wajib diisi"),
});

export type UnlinkGoogleFormValues = z.infer<typeof unlinkGoogleSchema>;
