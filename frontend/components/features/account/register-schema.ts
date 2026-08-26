import { z } from "zod";

/**
 * Client-side validation is UX-only — the backend is the source of
 * truth (`form-validation-boundary.md`). Password length policy: see
 * `docs/spec/1-account/features/01-register-email-verification.md`
 * (>=8 chars, NIST 800-63B style, length-only — no character-class
 * requirement). The breach-list check is server-only and is
 * deliberately NOT replicated here (rule R2) — a compliant-looking
 * password can still come back as a `422` from the backend, surfaced
 * via `RegisterResult`'s `kind: "validation"` branch.
 */
export const registerSchema = z.object({
  name: z.string().min(1, "Nama wajib diisi"),
  email: z.string().email("Format email tidak valid"),
  password: z.string().min(8, "Password minimal 8 karakter"),
});

export type RegisterFormValues = z.infer<typeof registerSchema>;
