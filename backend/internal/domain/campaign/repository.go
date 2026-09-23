package campaign

import (
	"context"

	"github.com/google/uuid"
)

// SeedRepository persists an explicitly operator-supplied Campaign seed.
type SeedRepository interface {
	SeedCampaign(ctx context.Context, record SeedRecord, replace bool) (replacedObjectKey string, err error)
}

// Repository is the persistence port for the public Campaign read model.
// Implementations must issue parameterized queries and apply the public
// eligibility predicate before returning a record.
type Repository interface {
	FindPublicDetail(ctx context.Context, campaignID uuid.UUID) (*DetailRecord, error)
	FindPublicMedia(ctx context.Context, campaignID, mediaID uuid.UUID) (*MediaMetadata, error)
}

// ObjectReader reads a private object by opaque key. It never offers public
// URLs, signed URLs, or redirects.
type ObjectReader interface {
	Read(ctx context.Context, objectKey string) (*MediaContent, error)
}
