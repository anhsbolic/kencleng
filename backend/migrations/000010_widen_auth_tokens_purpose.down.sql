-- 000010 down / account task #05 (feature: account linking)
--
-- Re-map BEFORE restoring the 2-value CHECK: any email_verification_link
-- rows are re-pointed to email_verification, which is semantically safe
-- because token redemption is purpose-blind (RedeemToken guards on the
-- hash's validity, not the purpose value) — a re-pointed token still
-- verifies the same identity it always did.

UPDATE auth_tokens SET purpose = 'email_verification'
    WHERE purpose = 'email_verification_link';

ALTER TABLE auth_tokens DROP CONSTRAINT auth_tokens_purpose_check;
ALTER TABLE auth_tokens ADD CONSTRAINT auth_tokens_purpose_check
    CHECK (purpose IN ('email_verification', 'password_reset'));
