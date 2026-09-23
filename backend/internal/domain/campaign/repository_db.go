package campaign

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // register PostgreSQL dialect
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pgDialect = goqu.Dialect("postgres")

// RepositoryDB implements the Campaign read port with goqu and pgx.
type RepositoryDB struct{ db *pgxpool.Pool }

// NewRepositoryDB constructs a Campaign repository over the supplied pool.
func NewRepositoryDB(db *pgxpool.Pool) *RepositoryDB { return &RepositoryDB{db: db} }

// FindPublicDetail selects only an eligible Campaign and its ordered public
// media metadata. The explicit select list prevents schema growth from
// widening the public projection accidentally.
func (r *RepositoryDB) FindPublicDetail(ctx context.Context, campaignID uuid.UUID) (*DetailRecord, error) {
	sqlStr, args, err := pgDialect.From(goqu.T("campaigns").As("c")).
		LeftJoin(goqu.T("organizations").As("o"), goqu.On(goqu.I("c.organization_id").Eq(goqu.I("o.id")))).
		LeftJoin(goqu.T("campaign_media").As("m"), goqu.On(goqu.I("m.campaign_id").Eq(goqu.I("c.id")))).
		Select(
			goqu.I("c.id"), goqu.I("c.title"), goqu.I("c.purpose"), goqu.I("c.story"),
			goqu.I("o.id"), goqu.I("o.name"), goqu.I("c.published_at"), goqu.I("c.fundraising_ends_at"),
			goqu.I("c.target_amount"), goqu.I("c.collected_amount"), goqu.I("c.media_state"),
			goqu.I("m.id"), goqu.I("m.object_key"), goqu.I("m.content_type"), goqu.I("m.alt_text"), goqu.I("m.caption"), goqu.I("m.display_order"),
		).
		Where(goqu.Ex{"c.id": campaignID, "c.status": "published"}).
		Order(goqu.I("m.display_order").Asc()).
		Prepared(true).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("campaign: build public detail query: %w", err)
	}
	rows, err := r.db.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, fmt.Errorf("campaign: query public detail: %w", err)
	}
	defer rows.Close()

	var record *DetailRecord
	for rows.Next() {
		var row DetailRecord
		var mediaID *uuid.UUID
		var objectKey, contentType, altText, caption *string
		var displayOrder *int
		if err := rows.Scan(&row.ID, &row.Title, &row.Purpose, &row.Story, &row.StewardID, &row.StewardName,
			&row.PublishedAt, &row.FundraisingEndsAt, &row.TargetAmount, &row.CollectedAmount, &row.MediaState,
			&mediaID, &objectKey, &contentType, &altText, &caption, &displayOrder); err != nil {
			return nil, fmt.Errorf("campaign: scan public detail: %w", err)
		}
		if record == nil {
			record = &row
		}
		if mediaID != nil {
			record.Media = append(record.Media, MediaMetadata{ID: *mediaID, ObjectKey: derefString(objectKey), ContentType: derefString(contentType), AltText: derefString(altText), Caption: caption, DisplayOrder: derefInt(displayOrder)})
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("campaign: iterate public detail: %w", err)
	}
	return record, nil
}

// FindPublicMedia performs the parent eligibility and member association
// check in one filtered query before storage is contacted.
func (r *RepositoryDB) FindPublicMedia(ctx context.Context, campaignID, mediaID uuid.UUID) (*MediaMetadata, error) {
	sqlStr, args, err := pgDialect.From(goqu.T("campaign_media").As("m")).
		InnerJoin(goqu.T("campaigns").As("c"), goqu.On(goqu.I("m.campaign_id").Eq(goqu.I("c.id")))).
		Select(goqu.I("m.id"), goqu.I("m.object_key"), goqu.I("m.content_type"), goqu.I("m.alt_text"), goqu.I("m.caption"), goqu.I("m.display_order")).
		Where(goqu.Ex{"c.id": campaignID, "m.id": mediaID, "c.status": "published", "c.media_state": "available"}).
		Prepared(true).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("campaign: build public media query: %w", err)
	}
	var media MediaMetadata
	err = r.db.QueryRow(ctx, sqlStr, args...).Scan(&media.ID, &media.ObjectKey, &media.ContentType, &media.AltText, &media.Caption, &media.DisplayOrder)
	if err == nil {
		return &media, nil
	}
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return nil, fmt.Errorf("campaign: query public media: %w", err)
}

func derefString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
func derefInt(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}

// SeedCampaign persists a seed only through an explicit command path. A
// collision fails by default; replace is the only mutation path and returns
// the superseded object key for best-effort post-commit cleanup.
func (r *RepositoryDB) SeedCampaign(ctx context.Context, record SeedRecord, replace bool) (string, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("campaign: begin seed transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var oldObjectKey string
	if replace {
		findOldSQL, findOldArgs, err := pgDialect.From("campaign_media").Select("object_key").Where(goqu.Ex{"campaign_id": record.CampaignID}).Order(goqu.I("display_order").Asc()).Limit(1).Prepared(true).ToSQL()
		if err != nil {
			return "", fmt.Errorf("campaign: build replaced object query: %w", err)
		}
		if err := tx.QueryRow(ctx, findOldSQL, findOldArgs...).Scan(&oldObjectKey); err != nil && err != pgx.ErrNoRows {
			return "", fmt.Errorf("campaign: find replaced object: %w", err)
		}
		deleteSQL, deleteArgs, err := pgDialect.Delete("campaigns").Where(goqu.Ex{"id": record.CampaignID}).Prepared(true).ToSQL()
		if err != nil {
			return "", fmt.Errorf("campaign: build replace seed campaign: %w", err)
		}
		if _, err := tx.Exec(ctx, deleteSQL, deleteArgs...); err != nil {
			return "", fmt.Errorf("campaign: replace seed campaign: %w", err)
		}
	} else {
		findSQL, findArgs, err := pgDialect.From("campaigns").Select("id").Where(goqu.Ex{"id": record.CampaignID}).Prepared(true).ToSQL()
		if err != nil {
			return "", fmt.Errorf("campaign: build seed collision query: %w", err)
		}
		var existingID uuid.UUID
		if err := tx.QueryRow(ctx, findSQL, findArgs...).Scan(&existingID); err == nil {
			return "", fmt.Errorf("campaign: seed collision")
		} else if err != pgx.ErrNoRows {
			return "", fmt.Errorf("campaign: check seed collision: %w", err)
		}
	}

	orgSQL, orgArgs, err := pgDialect.Insert("organizations").Rows(goqu.Record{"id": record.OrganizationID, "name": record.StewardName}).OnConflict(goqu.DoNothing()).Prepared(true).ToSQL()
	if err != nil {
		return "", fmt.Errorf("campaign: build seed organization: %w", err)
	}
	if _, err := tx.Exec(ctx, orgSQL, orgArgs...); err != nil {
		return "", fmt.Errorf("campaign: seed organization: %w", err)
	}
	campaignSQL, campaignArgs, err := pgDialect.Insert("campaigns").Rows(goqu.Record{
		"id": record.CampaignID, "organization_id": record.OrganizationID, "title": record.Title, "purpose": record.Purpose, "story": record.Story,
		"status": "published", "published_at": record.PublishedAt, "fundraising_ends_at": record.FundraisingEndsAt,
		"target_amount": record.TargetAmount, "collected_amount": record.CollectedAmount, "media_state": record.MediaState,
	}).Prepared(true).ToSQL()
	if err != nil {
		return "", fmt.Errorf("campaign: build seed campaign: %w", err)
	}
	if _, err := tx.Exec(ctx, campaignSQL, campaignArgs...); err != nil {
		return "", fmt.Errorf("campaign: seed campaign: %w", err)
	}
	if record.Media != nil {
		mediaSQL, mediaArgs, err := pgDialect.Insert("campaign_media").Rows(goqu.Record{"id": record.Media.ID, "campaign_id": record.CampaignID, "object_key": record.Media.ObjectKey, "display_order": 0, "content_type": record.Media.ContentType, "alt_text": record.Media.AltText, "caption": record.Media.Caption}).Prepared(true).ToSQL()
		if err != nil {
			return "", fmt.Errorf("campaign: build seed media: %w", err)
		}
		if _, err := tx.Exec(ctx, mediaSQL, mediaArgs...); err != nil {
			return "", fmt.Errorf("campaign: seed media: %w", err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("campaign: commit seed: %w", err)
	}
	return oldObjectKey, nil
}

var _ Repository = (*RepositoryDB)(nil)
var _ SeedRepository = (*RepositoryDB)(nil)
