-- 000010 down / account task #05 (feature: account linking)
--
-- Re-map BEFORE restoring the 2-value CHECK. Safe when the migration
-- rolls back alongside the feature code: rolled-back VerifyEmail has no
-- link-audit logic and its 2-value purpose guard accepts the remapped
-- 'email_verification' token, so every still-unredeemed link token keeps
-- verifying the same identity it always did. A standalone rollback
-- (migration down while the feature code stays) is NOT equivalent: the
-- new code's guard accepts both purposes but writes the R14 audit only
-- for 'email_verification_link', so a remapped token would redeem as a
-- registration token — the identity still verifies, but the audit entry
-- is silently skipped. Do not roll back this migration alone.

UPDATE auth_tokens SET purpose = 'email_verification'
    WHERE purpose = 'email_verification_link';

ALTER TABLE auth_tokens DROP CONSTRAINT auth_tokens_purpose_check;
ALTER TABLE auth_tokens ADD CONSTRAINT auth_tokens_purpose_check
    CHECK (purpose IN ('email_verification', 'password_reset'));
