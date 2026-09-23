package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/anhsbolic/kencleng/backend/internal/domain/campaign"
	"github.com/google/uuid"
)

type fakePublicCampaignService struct {
	detail     *campaign.PublicDetail
	detailErr  error
	content    *campaign.MediaContent
	contentErr error
}

// deterministicLatencyService models a fixed dependency budget without using
// wall-clock assertions. It makes the representative 404 paths comparable in
// a repeatable handler/service-seam test.
type deterministicLatencyService struct {
	detailErr, contentErr error
	dependencyLatency     time.Duration
	elapsed               time.Duration
}

func (s *deterministicLatencyService) GetPublicDetail(context.Context, uuid.UUID) (*campaign.PublicDetail, error) {
	s.elapsed += s.dependencyLatency
	return nil, s.detailErr
}

func (s *deterministicLatencyService) GetPublicMediaContent(context.Context, uuid.UUID, uuid.UUID) (*campaign.MediaContent, error) {
	s.elapsed += s.dependencyLatency
	return nil, s.contentErr
}

type trackedReadCloser struct {
	*strings.Reader
	closed bool
}

func (r *trackedReadCloser) Close() error { r.closed = true; return nil }

func (f fakePublicCampaignService) GetPublicDetail(context.Context, uuid.UUID) (*campaign.PublicDetail, error) {
	return f.detail, f.detailErr
}
func (f fakePublicCampaignService) GetPublicMediaContent(context.Context, uuid.UUID, uuid.UUID) (*campaign.MediaContent, error) {
	return f.content, f.contentErr
}

func wirePublicCampaignHandler(svc publicCampaignService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /campaigns/{campaignId}", PublicCampaignDetailHandler(svc))
	mux.HandleFunc("GET /campaigns/{campaignId}/media/{mediaId}/content", PublicCampaignMediaContentHandler(svc))
	return mux
}

func publicDetail() *campaign.PublicDetail {
	id := uuid.New()
	mediaID := uuid.New()
	percentage := "125.50"
	return &campaign.PublicDetail{ID: id, Title: "Title", Purpose: campaign.OrganizerText{Content: "Purpose", Source: "organizer"}, Story: campaign.OrganizerText{Content: "Story", Source: "organizer"}, Steward: campaign.Steward{ID: uuid.New(), Name: "Steward"}, Lifecycle: campaign.Lifecycle{PublicState: "fundraising", PublishedAt: time.Now().UTC(), FundraisingEndsAt: time.Now().UTC().Add(time.Hour)}, Funding: campaign.Funding{Availability: "available", Currency: "IDR", TargetAmount: "100.00", CollectedAmount: "125.50", Progress: campaign.FundingProgress{State: "computed", Percentage: &percentage, Relationship: "above_target"}}, Media: campaign.Media{State: "available", Items: []campaign.PublicMediaItem{{ID: mediaID, ContentURL: "/api/campaigns/" + id.String() + "/media/" + mediaID.String() + "/content", ContentType: "image/jpeg", AltText: "Alt", Source: "organizer"}}}, DonationAction: campaign.DonationAction{Availability: "unavailable", Reason: "donation_flow_not_available"}}
}

func TestPublicCampaignDetailHandler_ExactClosedWireShape(t *testing.T) {
	detail := publicDetail()
	req := httptest.NewRequest(http.MethodGet, "/campaigns/"+detail.ID.String(), nil)
	req.Header.Set("Authorization", "not consulted")
	recorder := httptest.NewRecorder()
	wirePublicCampaignHandler(fakePublicCampaignService{detail: detail}).ServeHTTP(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d", recorder.Code)
	}
	if recorder.Header().Get("Cache-Control") != publicCampaignCacheControl {
		t.Errorf("cache = %q", recorder.Header().Get("Cache-Control"))
	}
	var body map[string]json.RawMessage
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if len(body) != 9 {
		t.Errorf("top-level keys = %d, want 9: %s", len(body), recorder.Body.String())
	}
	for _, forbidden := range []string{"status", "object_key", "donor_count", "organization_id"} {
		if _, exists := body[forbidden]; exists {
			t.Errorf("forbidden key %q serialized", forbidden)
		}
	}
}

func TestPublicCampaignHandlers_PublicNotFoundParity(t *testing.T) {
	svc := fakePublicCampaignService{detailErr: campaign.ErrNotFound, contentErr: campaign.ErrNotFound}
	handler := wirePublicCampaignHandler(svc)
	missingDetail := httptest.NewRecorder()
	handler.ServeHTTP(missingDetail, httptest.NewRequest(http.MethodGet, "/campaigns/not-a-uuid", nil))
	missingMedia := httptest.NewRecorder()
	handler.ServeHTTP(missingMedia, httptest.NewRequest(http.MethodGet, "/campaigns/not-a-uuid/media/not-a-uuid/content", nil))
	if missingDetail.Code != http.StatusNotFound || missingMedia.Code != http.StatusNotFound {
		t.Fatalf("statuses = %d, %d", missingDetail.Code, missingMedia.Code)
	}
	if missingDetail.Body.String() != missingMedia.Body.String() {
		t.Errorf("404 bodies differ: %q / %q", missingDetail.Body, missingMedia.Body)
	}
	if missingDetail.Header().Get("Cache-Control") != publicCampaignCacheControl || missingMedia.Header().Get("Cache-Control") != publicCampaignCacheControl {
		t.Error("404 missing no-store")
	}
}

func TestPublicCampaignHandlers_NotFoundAuthorizationMatrix(t *testing.T) {
	campaignID := uuid.New()
	mediaID := uuid.New()
	type requestClass struct {
		name string
		path string
		svc  fakePublicCampaignService
	}
	cases := []requestClass{
		{name: "absent campaign", path: "/campaigns/" + campaignID.String(), svc: fakePublicCampaignService{detailErr: campaign.ErrNotFound}},
		{name: "non-published campaign", path: "/campaigns/" + campaignID.String(), svc: fakePublicCampaignService{detailErr: campaign.ErrNotFound}},
		{name: "absent media", path: "/campaigns/" + campaignID.String() + "/media/" + mediaID.String() + "/content", svc: fakePublicCampaignService{contentErr: campaign.ErrNotFound}},
		{name: "non-member media", path: "/campaigns/" + campaignID.String() + "/media/" + mediaID.String() + "/content", svc: fakePublicCampaignService{contentErr: campaign.ErrNotFound}},
	}
	authorizations := []struct {
		name  string
		value string
	}{{name: "none"}, {name: "garbage", value: "Bearer not-a-token"}, {name: "valid-shaped", value: "Bearer eyJhbGciOiJFUzI1NiJ9.eyJzdWIiOiJub3QtY29uc3VsdGVkIn0.signature"}}

	var baseline *httptest.ResponseRecorder
	for _, requestClass := range cases {
		for _, authorization := range authorizations {
			t.Run(requestClass.name+"/"+authorization.name, func(t *testing.T) {
				recorder := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, requestClass.path, nil)
				if authorization.value != "" {
					req.Header.Set("Authorization", authorization.value)
				}
				wirePublicCampaignHandler(requestClass.svc).ServeHTTP(recorder, req)
				if recorder.Code != http.StatusNotFound {
					t.Fatalf("status = %d, want 404", recorder.Code)
				}
				if recorder.Header().Get("Content-Type") != "application/problem+json" || recorder.Header().Get("Cache-Control") != publicCampaignCacheControl {
					t.Fatalf("headers = Content-Type %q, Cache-Control %q", recorder.Header().Get("Content-Type"), recorder.Header().Get("Cache-Control"))
				}
				if baseline == nil {
					baseline = recorder
					return
				}
				if recorder.Code != baseline.Code || recorder.Body.String() != baseline.Body.String() || recorder.Header().Get("Content-Type") != baseline.Header().Get("Content-Type") || recorder.Header().Get("Cache-Control") != baseline.Header().Get("Cache-Control") {
					t.Fatalf("public 404 differs from baseline: status=%d body=%q headers=%q/%q", recorder.Code, recorder.Body.String(), recorder.Header().Get("Content-Type"), recorder.Header().Get("Cache-Control"))
				}
			})
		}
	}
}

func TestPublicCampaignHandlers_RepresentativeNotFoundTimingParity(t *testing.T) {
	const samples = 128
	const dependencyLatency = 2 * time.Millisecond
	campaignID, mediaID := uuid.New(), uuid.New()
	paths := []struct {
		name    string
		path    string
		service *deterministicLatencyService
	}{
		{name: "absent campaign", path: "/campaigns/" + campaignID.String(), service: &deterministicLatencyService{detailErr: campaign.ErrNotFound, dependencyLatency: dependencyLatency}},
		{name: "non-published campaign", path: "/campaigns/" + campaignID.String(), service: &deterministicLatencyService{detailErr: campaign.ErrNotFound, dependencyLatency: dependencyLatency}},
		{name: "non-member media", path: "/campaigns/" + campaignID.String() + "/media/" + mediaID.String() + "/content", service: &deterministicLatencyService{contentErr: campaign.ErrNotFound, dependencyLatency: dependencyLatency}},
	}

	for _, path := range paths {
		t.Run(path.name, func(t *testing.T) {
			for range samples {
				recorder := httptest.NewRecorder()
				wirePublicCampaignHandler(path.service).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, path.path, nil))
				if recorder.Code != http.StatusNotFound {
					t.Fatalf("status = %d, want 404", recorder.Code)
				}
			}
			if want := time.Duration(samples) * dependencyLatency; path.service.elapsed != want {
				t.Fatalf("logical elapsed = %s, want %s", path.service.elapsed, want)
			}
		})
	}
}

func TestPublicCampaignMediaHandler_StreamsValidatedBytesAndUnavailable(t *testing.T) {
	id := uuid.New()
	mediaID := uuid.New()
	reader := &trackedReadCloser{Reader: strings.NewReader("jpeg-bytes")}
	handler := wirePublicCampaignHandler(fakePublicCampaignService{content: &campaign.MediaContent{Reader: reader, ContentType: "image/jpeg"}})
	ok := httptest.NewRecorder()
	handler.ServeHTTP(ok, httptest.NewRequest(http.MethodGet, "/campaigns/"+id.String()+"/media/"+mediaID.String()+"/content", nil))
	if ok.Code != http.StatusOK || ok.Header().Get("Content-Type") != "image/jpeg" || ok.Body.String() != "jpeg-bytes" {
		t.Errorf("media response = status %d type %q body %q", ok.Code, ok.Header().Get("Content-Type"), ok.Body.String())
	}
	if !reader.closed {
		t.Error("media reader was not closed")
	}
	bad := httptest.NewRecorder()
	wirePublicCampaignHandler(fakePublicCampaignService{contentErr: errors.New("private object key must not leak")}).ServeHTTP(bad, httptest.NewRequest(http.MethodGet, "/campaigns/"+id.String()+"/media/"+mediaID.String()+"/content", nil))
	if bad.Code != http.StatusServiceUnavailable || strings.Contains(bad.Body.String(), "private object key") || bad.Header().Get("Cache-Control") != publicCampaignCacheControl {
		t.Errorf("unavailable response = %d %q", bad.Code, bad.Body.String())
	}
}
