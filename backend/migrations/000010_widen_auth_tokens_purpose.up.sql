-- 000010 / account task #05 (feature: account linking)
--
-- Widen auth_tokens.purpose to admit 'email_verification_link': the
-- token issued by POST /account/security/set-password Branch 1 (adding
-- an email_password identity to a Google-only account). The distinct
-- purpose lets POST /auth/verify-email stay externally unchanged while
-- writing the user_logs audit entry truthfully when the redeemed token
-- came from the linking flow rather than registration
-- (techplan D7; docs/spec/1-account/features/05-account-linking.md).
--
-- Additive: no existing row violates either the old or the new
-- constraint; auth_tokens is small (per-registration issuance only), so
-- the brief lock during DROP + ADD CONSTRAINT is negligible.

ALTER TABLE auth_tokens DROP CONSTRAINT auth_tokens_purpose_check;
ALTER TABLE auth_tokens ADD CONSTRAINT auth_tokens_purpose_check
    CHECK (purpose IN ('email_verification', 'email_verification_link', 'password_reset'));
