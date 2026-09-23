// Package campaign contains the Slice-1 public Campaign read model.
package campaign

import (
	"errors"
	"io"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrNotFound deliberately collapses malformed, absent, ineligible, and
	// non-member resources at the public transport boundary.
	ErrNotFound = errors.New("campaign public resource not found")
	// ErrUnavailable means an otherwise eligible Campaign dependency cannot
	// currently serve a truthful response.
	ErrUnavailable = errors.New("campaign public resource unavailable")
	// ErrObjectUnavailable is returned by an ObjectReader for private-object
	// lookup failures. Services translate it to ErrUnavailable.
	ErrObjectUnavailable = errors.New("campaign private object unavailable")
)

// MediaMetadata is the narrow persisted media data allowed into the public
// Campaign mapper. ObjectKey never crosses the HTTP projection.
type MediaMetadata struct {
	ID           uuid.UUID
	ObjectKey    string
	ContentType  string
	AltText      string
	Caption      *string
	DisplayOrder int
}

// DetailRecord is the narrow repository record for one eligible Campaign.
// It intentionally has no internal status, contact, legal, audit, or actor
// fields, so those values cannot accidentally reach the public projection.
type DetailRecord struct {
	ID                uuid.UUID
	Title             string
	Purpose           string
	Story             string
	StewardID         uuid.UUID
	StewardName       string
	PublishedAt       time.Time
	FundraisingEndsAt time.Time
	TargetAmount      *string
	CollectedAmount   *string
	MediaState        string
	Media             []MediaMetadata
}

// PublicDetail is the closed domain projection used by the HTTP wire mapper.
type PublicDetail struct {
	ID             uuid.UUID
	Title          string
	Purpose        OrganizerText
	Story          OrganizerText
	Steward        Steward
	Lifecycle      Lifecycle
	Funding        Funding
	Media          Media
	DonationAction DonationAction
}

// OrganizerText is organizer-authored plain text, not verification evidence.
type OrganizerText struct {
	Content string
	Source  string
}

// Steward is the public-safe Organization projection.
type Steward struct {
	ID   uuid.UUID
	Name string
}

// Lifecycle contains only the Slice-1 fundraising state.
type Lifecycle struct {
	PublicState       string
	PublishedAt       time.Time
	FundraisingEndsAt time.Time
}

// Funding is the tagged public funding truth.
type Funding struct {
	Availability    string
	Reason          string
	Currency        string
	TargetAmount    string
	CollectedAmount string
	Progress        FundingProgress
}

// FundingProgress carries the exact, uncapped percentage string.
type FundingProgress struct {
	State        string
	Percentage   *string
	Relationship string
}

// Media is the tagged public media projection.
type Media struct {
	State  string
	Reason string
	Items  []PublicMediaItem
}

// PublicMediaItem does not include ObjectKey and only has public fields.
type PublicMediaItem struct {
	ID          uuid.UUID
	ContentURL  string
	ContentType string
	AltText     string
	Caption     *string
	Source      string
}

// DonationAction records that no donation flow is active in Slice 1.
type DonationAction struct {
	Availability string
	Reason       string
}

// MediaContent is a validated private-object stream. Callers must Close it.
type MediaContent struct {
	Reader      io.ReadCloser
	ContentType string
}

// SeedRecord is the deliberately narrow operator-only persisted setup input.
// It is not a public creation workflow or HTTP request model.
type SeedRecord struct {
	CampaignID        uuid.UUID
	OrganizationID    uuid.UUID
	StewardName       string
	Title             string
	Purpose           string
	Story             string
	PublishedAt       time.Time
	FundraisingEndsAt time.Time
	TargetAmount      *string
	CollectedAmount   *string
	MediaState        string
	Media             *SeedMedia
}

// SeedMedia is accepted only when a validated private image is supplied.
type SeedMedia struct {
	ID          uuid.UUID
	ObjectKey   string
	ContentType string
	AltText     string
	Caption     *string
}
