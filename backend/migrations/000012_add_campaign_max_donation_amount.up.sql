-- Slice 2 per-Campaign individual donation ceiling, in whole Rupiah.
ALTER TABLE campaigns
    ADD COLUMN max_donation_amount NUMERIC(10,0) NOT NULL DEFAULT 1000000000,
    ADD CONSTRAINT campaigns_max_donation_amount_check
        CHECK (max_donation_amount BETWEEN 5000 AND 1000000000);
