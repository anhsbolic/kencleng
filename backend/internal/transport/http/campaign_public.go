package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/campaign"
)

const publicCampaignCacheControl = "private, no-store"

type publicCampaignService interface {
	GetPublicDetail(ctx context.Context, campaignID uuid.UUID) (*campaign.PublicDetail, error)
	GetPublicMediaContent(ctx context.Context, campaignID, mediaID uuid.UUID) (*campaign.MediaContent, error)
}

type publicCampaignTextResponse struct {
	Content string `json:"content"`
	Source  string `json:"source"`
}
type publicCampaignStewardResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type publicCampaignLifecycleResponse struct {
	PublicState       string    `json:"public_state"`
	PublishedAt       time.Time `json:"published_at"`
	FundraisingEndsAt time.Time `json:"fundraising_ends_at"`
}
type publicCampaignFundingProgressResponse struct {
	State        string  `json:"state"`
	Percentage   *string `json:"percentage"`
	Relationship string  `json:"relationship"`
}
type publicCampaignFundingResponse struct {
	Availability    string                                 `json:"availability"`
	Reason          string                                 `json:"reason,omitempty"`
	Currency        string                                 `json:"currency,omitempty"`
	TargetAmount    string                                 `json:"target_amount,omitempty"`
	CollectedAmount string                                 `json:"collected_amount,omitempty"`
	Progress        *publicCampaignFundingProgressResponse `json:"progress,omitempty"`
}
type publicCampaignMediaItemResponse struct {
	ID          uuid.UUID `json:"id"`
	ContentURL  string    `json:"content_url"`
	ContentType string    `json:"content_type"`
	AltText     string    `json:"alt_text"`
	Caption     *string   `json:"caption"`
	Source      string    `json:"source"`
}
type publicCampaignMediaResponse struct {
	State  string                            `json:"state"`
	Reason string                            `json:"reason,omitempty"`
	Items  []publicCampaignMediaItemResponse `json:"items"`
}
type publicCampaignDonationActionResponse struct {
	Availability string `json:"availability"`
	Reason       string `json:"reason"`
}
type publicCampaignDetailResponse struct {
	ID             uuid.UUID                            `json:"id"`
	Title          string                               `json:"title"`
	Purpose        publicCampaignTextResponse           `json:"purpose"`
	Story          publicCampaignTextResponse           `json:"story"`
	Steward        publicCampaignStewardResponse        `json:"steward"`
	Lifecycle      publicCampaignLifecycleResponse      `json:"lifecycle"`
	Funding        publicCampaignFundingResponse        `json:"funding"`
	Media          publicCampaignMediaResponse          `json:"media"`
	DonationAction publicCampaignDonationActionResponse `json:"donation_action"`
}

// PublicCampaignDetailHandler handles the unauthenticated detail operation.
// Authorization is intentionally neither parsed nor consulted.
func PublicCampaignDetailHandler(svc publicCampaignService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		campaignID, err := uuid.Parse(r.PathValue("campaignId"))
		if err != nil {
			writePublicCampaignNotFound(w)
			return
		}
		detail, err := svc.GetPublicDetail(r.Context(), campaignID)
		if err != nil {
			writePublicCampaignError(w, err)
			return
		}
		if detail == nil {
			writePublicCampaignUnavailable(w)
			return
		}
		writePublicCampaignJSON(w, http.StatusOK, toPublicCampaignDetailResponse(detail))
	}
}

// PublicCampaignMediaContentHandler handles the controlled byte operation.
// It has no session middleware and never exposes an object-store location.
func PublicCampaignMediaContentHandler(svc publicCampaignService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		campaignID, err := uuid.Parse(r.PathValue("campaignId"))
		if err != nil {
			writePublicCampaignNotFound(w)
			return
		}
		mediaID, err := uuid.Parse(r.PathValue("mediaId"))
		if err != nil {
			writePublicCampaignNotFound(w)
			return
		}
		content, err := svc.GetPublicMediaContent(r.Context(), campaignID, mediaID)
		if err != nil {
			writePublicCampaignError(w, err)
			return
		}
		if content == nil || content.Reader == nil {
			writePublicCampaignUnavailable(w)
			return
		}
		defer func() {
			if closeErr := content.Reader.Close(); closeErr != nil {
				log.Printf("campaign media stream close failed")
			}
		}()
		w.Header().Set("Cache-Control", publicCampaignCacheControl)
		w.Header().Set("Content-Type", content.ContentType)
		w.WriteHeader(http.StatusOK)
		if _, err := io.Copy(w, content.Reader); err != nil {
			// Headers/body may already be committed; writing a Problem now would
			// lie about the response. Do not include object keys or provider text.
			log.Printf("campaign media stream interrupted")
		}
	}
}

func toPublicCampaignDetailResponse(detail *campaign.PublicDetail) publicCampaignDetailResponse {
	items := make([]publicCampaignMediaItemResponse, 0, len(detail.Media.Items))
	for _, item := range detail.Media.Items {
		items = append(items, publicCampaignMediaItemResponse{ID: item.ID, ContentURL: item.ContentURL, ContentType: item.ContentType, AltText: item.AltText, Caption: item.Caption, Source: item.Source})
	}
	funding := publicCampaignFundingResponse{Availability: detail.Funding.Availability, Reason: detail.Funding.Reason, Currency: detail.Funding.Currency, TargetAmount: detail.Funding.TargetAmount, CollectedAmount: detail.Funding.CollectedAmount}
	if detail.Funding.Availability == "available" {
		funding.Progress = &publicCampaignFundingProgressResponse{State: detail.Funding.Progress.State, Percentage: detail.Funding.Progress.Percentage, Relationship: detail.Funding.Progress.Relationship}
	}
	return publicCampaignDetailResponse{
		ID: detail.ID, Title: detail.Title,
		Purpose: publicCampaignTextResponse{Content: detail.Purpose.Content, Source: detail.Purpose.Source}, Story: publicCampaignTextResponse{Content: detail.Story.Content, Source: detail.Story.Source},
		Steward:   publicCampaignStewardResponse{ID: detail.Steward.ID, Name: detail.Steward.Name},
		Lifecycle: publicCampaignLifecycleResponse{PublicState: detail.Lifecycle.PublicState, PublishedAt: detail.Lifecycle.PublishedAt, FundraisingEndsAt: detail.Lifecycle.FundraisingEndsAt},
		Funding:   funding, Media: publicCampaignMediaResponse{State: detail.Media.State, Reason: detail.Media.Reason, Items: items},
		DonationAction: publicCampaignDonationActionResponse{Availability: detail.DonationAction.Availability, Reason: detail.DonationAction.Reason},
	}
}

func writePublicCampaignJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Cache-Control", publicCampaignCacheControl)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writePublicCampaignNotFound(w http.ResponseWriter) {
	writePublicCampaignProblem(w, http.StatusNotFound, "https://kencleng.dev/errors/campaign-not-found", "Campaign Not Found", "The requested public campaign resource was not found.")
}

func writePublicCampaignUnavailable(w http.ResponseWriter) {
	writePublicCampaignProblem(w, http.StatusServiceUnavailable, "https://kencleng.dev/errors/campaign-unavailable", "Campaign Unavailable", "The requested public campaign resource is temporarily unavailable.")
}

func writePublicCampaignProblem(w http.ResponseWriter, status int, problemType, title, detail string) {
	w.Header().Set("Cache-Control", publicCampaignCacheControl)
	WriteProblem(w, status, problemType, title, detail)
}

func writePublicCampaignError(w http.ResponseWriter, err error) {
	if errors.Is(err, campaign.ErrNotFound) {
		writePublicCampaignNotFound(w)
		return
	}
	writePublicCampaignUnavailable(w)
}
