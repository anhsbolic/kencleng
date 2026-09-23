-- 000011 / Slice 1 public Campaign delivery
--
-- This is deliberately a minimum, additive schema for the persisted public
-- projection. It does not introduce creation, curation, donation, or upload
-- workflows from deferred Campaign slices.

CREATE TABLE organizations (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL CHECK (btrim(name) <> '')
);

CREATE TABLE campaigns (
    id UUID PRIMARY KEY,
    organization_id UUID NOT NULL REFERENCES organizations(id),
    title TEXT NOT NULL CHECK (btrim(title) <> ''),
    purpose TEXT NOT NULL CHECK (btrim(purpose) <> ''),
    story TEXT NOT NULL CHECK (btrim(story) <> ''),
    status TEXT NOT NULL CHECK (status IN ('draft', 'published', 'unpublished', 'closed')),
    published_at TIMESTAMPTZ,
    fundraising_ends_at TIMESTAMPTZ NOT NULL,
    target_amount NUMERIC(19,2),
    collected_amount NUMERIC(19,2),
    media_state TEXT NOT NULL CHECK (media_state IN ('absent', 'unavailable', 'available')),
    CONSTRAINT campaigns_funding_pair_check
        CHECK ((target_amount IS NULL) = (collected_amount IS NULL)),
    CONSTRAINT campaigns_funding_non_negative_check
        CHECK ((target_amount IS NULL OR target_amount >= 0)
            AND (collected_amount IS NULL OR collected_amount >= 0)),
    CONSTRAINT campaigns_published_timestamp_check
        CHECK (status <> 'published' OR published_at IS NOT NULL)
);

CREATE INDEX campaigns_public_detail_idx
    ON campaigns (id, organization_id) WHERE status = 'published';

CREATE TABLE campaign_media (
    id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    object_key TEXT NOT NULL UNIQUE CHECK (btrim(object_key) <> ''),
    display_order INTEGER NOT NULL CHECK (display_order >= 0),
    content_type TEXT NOT NULL CHECK (content_type IN ('image/jpeg', 'image/png')),
    alt_text TEXT NOT NULL CHECK (btrim(alt_text) <> ''),
    caption TEXT,
    CONSTRAINT campaign_media_display_order_unique UNIQUE (campaign_id, display_order)
);

CREATE INDEX campaign_media_public_lookup_idx ON campaign_media (campaign_id, id);
