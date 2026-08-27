import { z } from "zod";

/**
 * Single `totp_code` field — required, non-empty. No digit-count/format
 * regex: not stated anywhere in the feature spec
 * (`docs/spec/1-account/features/06-mfa-totp.md`), so none is invented
 * here (`form-validation-boundary.md`'s checklist item on not inventing
 * unstated client-side policy). The backend is the real validator.
 */
export const mfaEnrollConfirmSchema = z.object({
  totp_code: z.string().min(1, "Kode wajib diisi"),
});

export type MfaEnrollConfirmFormValues = z.infer<typeof mfaEnrollConfirmSchema>;
