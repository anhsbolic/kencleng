ALTER TABLE campaigns
    DROP CONSTRAINT campaigns_max_donation_amount_check,
    DROP COLUMN max_donation_amount;
