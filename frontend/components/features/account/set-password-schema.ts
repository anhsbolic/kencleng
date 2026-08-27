import { z } from "zod";

/**
 * Client-side validation is UX-only — the backend is the source of
 * truth (`form-validation-boundary.md`). Password length policy: see
 * `docs/spec/1-account/features/05-account-linking.md` — same
 * length-only (>=8 chars), NIST 800-63B-style rule as registration,
 * reused verbatim from `register-schema.ts` rather than reinvented.
 * The breach-list check is server-only and deliberately NOT replicated
 * here — a compliant-looking password can still come back as a `422`
 * from the backend, surfaced via `SetPasswordResult`'s `kind:
 * "validation"` branch.
 */
export const addPasswordSchema = z.object({
  email: z.string().email("Format email tidak valid"),
  password: z.string().min(8, "Password minimal 8 karakter"),
});

export type AddPasswordFormValues = z.infer<typeof addPasswordSchema>;

export const changePasswordSchema = z.object({
  current_password: z.string().min(1, "Password saat ini wajib diisi"),
  password: z.string().min(8, "Password minimal 8 karakter"),
});

export type ChangePasswordFormValues = z.infer<typeof changePasswordSchema>;
