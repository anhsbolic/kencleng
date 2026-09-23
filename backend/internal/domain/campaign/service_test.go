package campaign

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

type fakeRepository struct {
	detail    *DetailRecord
	detailErr error
	media     *MediaMetadata
	mediaErr  error
}

func (f fakeRepository) FindPublicDetail(context.Context, uuid.UUID) (*DetailRecord, error) {
	return f.detail, f.detailErr
}
func (f fakeRepository) FindPublicMedia(context.Context, uuid.UUID, uuid.UUID) (*MediaMetadata, error) {
	return f.media, f.mediaErr
}

type fakeObjects struct {
	content *MediaContent
	err     error
}

func (f fakeObjects) Read(context.Context, string) (*MediaContent, error) { return f.content, f.err }

func publicRecord() *DetailRecord {
	amount := "100.00"
	collected := "125.505"
	caption := "Photo"
	return &DetailRecord{ID: uuid.New(), Title: "Campaign", Purpose: "Purpose", Story: "Story", StewardID: uuid.New(), StewardName: "Steward", PublishedAt: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC), FundraisingEndsAt: time.Date(2026, 2, 2, 3, 4, 5, 0, time.UTC), TargetAmount: &amount, CollectedAmount: &collected, MediaState: "available", Media: []MediaMetadata{{ID: uuid.New(), ObjectKey: "private-key", ContentType: "image/jpeg", AltText: "Alt", Caption: &caption}}}
}

func TestGetPublicDetail_ClosedProjectionAndExactFunding(t *testing.T) {
	record := publicRecord()
	detail, err := NewService(fakeRepository{detail: record}, fakeObjects{}).GetPublicDetail(context.Background(), record.ID)
	if err != nil {
		t.Fatalf("GetPublicDetail: %v", err)
	}
	if detail.Funding.CollectedAmount != "125.51" || detail.Funding.Progress.Percentage == nil || *detail.Funding.Progress.Percentage != "125.51" {
		t.Fatalf("funding = %#v, want half-up factual values", detail.Funding)
	}
	if detail.Funding.Progress.Relationship != "above_target" {
		t.Errorf("relationship = %q", detail.Funding.Progress.Relationship)
	}
	if detail.Media.Items[0].ContentURL != "/api/campaigns/"+record.ID.String()+"/media/"+record.Media[0].ID.String()+"/content" {
		t.Errorf("content URL = %q", detail.Media.Items[0].ContentURL)
	}
	if strings.Contains(detail.Media.Items[0].ContentURL, record.Media[0].ObjectKey) {
		t.Error("object key leaked into content URL")
	}
}

func TestMapFunding_ZeroAndUnavailable(t *testing.T) {
	zero := "0.00"
	cases := []struct {
		name                string
		target, collected   *string
		state, relationship string
		percentage          *string
	}{
		{"unavailable", nil, nil, "", "", nil},
		{"zero target", &zero, &zero, "not_computable", "not_computable", nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			funding, err := mapFunding(tc.target, tc.collected)
			if err != nil {
				t.Fatal(err)
			}
			if tc.target == nil {
				if funding.Availability != "unavailable" {
					t.Errorf("availability = %q", funding.Availability)
				}
				return
			}
			if funding.Progress.State != tc.state || funding.Progress.Relationship != tc.relationship || funding.Progress.Percentage != tc.percentage {
				t.Errorf("progress = %#v", funding.Progress)
			}
		})
	}
}

func TestGetPublicMediaContent_ClassifiesNotFoundAndUnavailable(t *testing.T) {
	id := uuid.New()
	cases := []struct {
		name    string
		repo    fakeRepository
		objects fakeObjects
		want    error
	}{
		{"missing", fakeRepository{}, fakeObjects{}, ErrNotFound},
		{"repository failure", fakeRepository{mediaErr: errors.New("db")}, fakeObjects{}, ErrUnavailable},
		{"storage failure", fakeRepository{media: &MediaMetadata{ID: id, ObjectKey: "key", ContentType: "image/png"}}, fakeObjects{err: ErrObjectUnavailable}, ErrUnavailable},
		{"wrong returned type", fakeRepository{media: &MediaMetadata{ID: id, ObjectKey: "key", ContentType: "image/png"}}, fakeObjects{content: &MediaContent{Reader: io.NopCloser(strings.NewReader("x")), ContentType: "image/jpeg"}}, ErrUnavailable},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewService(tc.repo, tc.objects).GetPublicMediaContent(context.Background(), id, id)
			if !errors.Is(err, tc.want) {
				t.Errorf("error = %v, want %v", err, tc.want)
			}
		})
	}
}
