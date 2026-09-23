package campaign

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Service owns public eligibility classification and semantic projection.
type Service struct {
	repository Repository
	objects    ObjectReader
}

// NewService constructs the public Campaign service with its narrow seams.
func NewService(repository Repository, objects ObjectReader) *Service {
	return &Service{repository: repository, objects: objects}
}

// GetPublicDetail returns only a closed public projection for an eligible
// Campaign. Repository errors are intentionally classified as unavailable.
func (s *Service) GetPublicDetail(ctx context.Context, campaignID uuid.UUID) (*PublicDetail, error) {
	record, err := s.repository.FindPublicDetail(ctx, campaignID)
	if err != nil {
		return nil, unavailable("get public campaign detail", err)
	}
	if record == nil {
		return nil, ErrNotFound
	}
	detail, err := toPublicDetail(record)
	if err != nil {
		return nil, unavailable("map public campaign detail", err)
	}
	return detail, nil
}

// GetPublicMediaContent rechecks eligible parent/member metadata before any
// private-object read. Object failures stay distinct from absent resources.
func (s *Service) GetPublicMediaContent(ctx context.Context, campaignID, mediaID uuid.UUID) (*MediaContent, error) {
	metadata, err := s.repository.FindPublicMedia(ctx, campaignID, mediaID)
	if err != nil {
		return nil, unavailable("get public campaign media", err)
	}
	if metadata == nil {
		return nil, ErrNotFound
	}
	if !allowedMediaType(metadata.ContentType) {
		return nil, ErrUnavailable
	}
	content, err := s.objects.Read(ctx, metadata.ObjectKey)
	if err != nil {
		if errors.Is(err, ErrObjectUnavailable) {
			return nil, unavailable("read public campaign media", err)
		}
		return nil, unavailable("read public campaign media", err)
	}
	if content == nil || content.Reader == nil || content.ContentType != metadata.ContentType || !allowedMediaType(content.ContentType) {
		if content != nil && content.Reader != nil {
			_ = content.Reader.Close()
		}
		return nil, ErrUnavailable
	}
	return content, nil
}

func unavailable(operation string, cause error) error {
	return fmt.Errorf("%s: %w", operation, errors.Join(ErrUnavailable, cause))
}

func toPublicDetail(record *DetailRecord) (*PublicDetail, error) {
	if record == nil || record.ID == uuid.Nil || record.StewardID == uuid.Nil ||
		strings.TrimSpace(record.Title) == "" || strings.TrimSpace(record.Purpose) == "" ||
		strings.TrimSpace(record.Story) == "" || strings.TrimSpace(record.StewardName) == "" ||
		record.PublishedAt.IsZero() || record.FundraisingEndsAt.IsZero() {
		return nil, errors.New("invalid public campaign record")
	}
	funding, err := mapFunding(record.TargetAmount, record.CollectedAmount)
	if err != nil {
		return nil, err
	}
	media, err := mapMedia(record.ID, record.MediaState, record.Media)
	if err != nil {
		return nil, err
	}
	return &PublicDetail{
		ID: record.ID, Title: record.Title,
		Purpose:   OrganizerText{Content: record.Purpose, Source: "organizer"},
		Story:     OrganizerText{Content: record.Story, Source: "organizer"},
		Steward:   Steward{ID: record.StewardID, Name: record.StewardName},
		Lifecycle: Lifecycle{PublicState: "fundraising", PublishedAt: record.PublishedAt, FundraisingEndsAt: record.FundraisingEndsAt},
		Funding:   funding, Media: media,
		DonationAction: DonationAction{Availability: "unavailable", Reason: "donation_flow_not_available"},
	}, nil
}

func mapFunding(targetRaw, collectedRaw *string) (Funding, error) {
	if targetRaw == nil && collectedRaw == nil {
		return Funding{Availability: "unavailable", Reason: "not_available"}, nil
	}
	if targetRaw == nil || collectedRaw == nil {
		return Funding{}, errors.New("partial funding pair")
	}
	target, err := decimal.NewFromString(*targetRaw)
	if err != nil || target.IsNegative() {
		return Funding{}, errors.New("invalid target amount")
	}
	collected, err := decimal.NewFromString(*collectedRaw)
	if err != nil || collected.IsNegative() {
		return Funding{}, errors.New("invalid collected amount")
	}
	funding := Funding{Availability: "available", Currency: "IDR", TargetAmount: target.StringFixed(2), CollectedAmount: collected.StringFixed(2)}
	if target.IsZero() {
		funding.Progress = FundingProgress{State: "not_computable", Relationship: "not_computable"}
		return funding, nil
	}
	percentage := collected.Div(target).Mul(decimal.NewFromInt(100)).Round(2).StringFixed(2)
	relationship := "below_target"
	switch collected.Cmp(target) {
	case 0:
		relationship = "target_reached"
	case 1:
		relationship = "above_target"
	}
	funding.Progress = FundingProgress{State: "computed", Percentage: &percentage, Relationship: relationship}
	return funding, nil
}

func mapMedia(campaignID uuid.UUID, state string, metadata []MediaMetadata) (Media, error) {
	switch state {
	case "absent":
		if len(metadata) != 0 {
			return Media{}, errors.New("absent media has metadata")
		}
		return Media{State: "absent", Items: []PublicMediaItem{}}, nil
	case "unavailable":
		if len(metadata) != 0 {
			return Media{}, errors.New("unavailable media has metadata")
		}
		return Media{State: "unavailable", Reason: "temporarily_unavailable", Items: []PublicMediaItem{}}, nil
	case "available":
		if len(metadata) == 0 {
			return Media{}, errors.New("available media lacks metadata")
		}
		items := make([]PublicMediaItem, 0, len(metadata))
		for _, item := range metadata {
			if item.ID == uuid.Nil || !allowedMediaType(item.ContentType) || strings.TrimSpace(item.AltText) == "" {
				return Media{}, errors.New("invalid available media metadata")
			}
			items = append(items, PublicMediaItem{ID: item.ID, ContentURL: fmt.Sprintf("/api/campaigns/%s/media/%s/content", campaignID, item.ID), ContentType: item.ContentType, AltText: item.AltText, Caption: item.Caption, Source: "organizer"})
		}
		return Media{State: "available", Items: items}, nil
	default:
		return Media{}, errors.New("invalid media state")
	}
}

func allowedMediaType(contentType string) bool {
	return contentType == "image/jpeg" || contentType == "image/png"
}
