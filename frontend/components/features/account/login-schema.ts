import { z } from "zod";

/**
 * Client-side validation is UX-only — the backend is the source of
 * truth (`form-validation-boundary.md`). Password length policy: same
 * length-only, no character-class requirement, as registration — see
 * `docs/spec/1-account/features/03-login-session-management.md` and
 * `register-schema.ts`'s own identical comment. A compliant-looking
 * password can still come back as a `401` (wrong credentials, R5) —
 * this schema never tries to guess correctness, only shape.
 */
export const loginSchema = z.object({
  email: z.string().email("Format email tidak valid"),
  password: z.string().min(8, "Password minimal 8 karakter"),
});

export type LoginFormValues = z.infer<typeof loginSchema>;

/**
 * MFA-challenge step (`POST /auth/login/mfa`) — the generated
 * `LoginMfaRequest` type only documents "one of `totp_code`/
 * `backup_code` must be present" as a comment (spec 03); this schema is
 * where the frontend actually enforces it as UX (R6). Both fields stay
 * optional at the type level since `LoginForm` only ever renders one at
 * a time (a toggle switches which is visible) — the `superRefine` below
 * requires at least one to be non-empty regardless of which one that
 * is, and attaches the same message to both paths so whichever field is
 * currently rendered shows it.
 */
export const loginMfaSchema = z
  .object({
    totp_code: z.string().optional(),
    backup_code: z.string().optional(),
  })
  .superRefine((data, ctx) => {
    const hasTotp = Boolean(data.totp_code?.trim());
    const hasBackup = Boolean(data.backup_code?.trim());

    if (!hasTotp && !hasBackup) {
      const message = "Masukkan kode OTP atau kode cadangan";
      ctx.addIssue({ code: z.ZodIssueCode.custom, message, path: ["totp_code"] });
      ctx.addIssue({ code: z.ZodIssueCode.custom, message, path: ["backup_code"] });
    }
  });

export type LoginMfaFormValues = z.infer<typeof loginMfaSchema>;
