// Command campaign-seed creates one explicitly supplied persisted public
// Campaign. It is an operator tool, not a startup seed or product workflow.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/shopspring/decimal"

	"github.com/anhsbolic/kencleng/backend/internal/domain/campaign"
	"github.com/anhsbolic/kencleng/backend/internal/platform/db"
	"github.com/anhsbolic/kencleng/backend/internal/platform/storage"
)

type manifest struct {
	CampaignID        string         `json:"campaign_id"`
	OrganizationID    string         `json:"organization_id"`
	StewardName       string         `json:"steward_name"`
	Title             string         `json:"title"`
	Purpose           string         `json:"purpose"`
	Story             string         `json:"story"`
	PublishedAt       string         `json:"published_at"`
	FundraisingEndsAt string         `json:"fundraising_ends_at"`
	TargetAmount      *string        `json:"target_amount"`
	CollectedAmount   *string        `json:"collected_amount"`
	MediaState        string         `json:"media_state"`
	Media             *manifestMedia `json:"media"`
}

type manifestMedia struct {
	ID      string  `json:"id"`
	AltText string  `json:"alt_text"`
	Caption *string `json:"caption"`
}

func main() {
	manifestPath := flag.String("manifest", "", "path to a Campaign seed JSON manifest")
	mediaPath := flag.String("media-file", "", "local JPEG or PNG for media_state=available")
	replace := flag.Bool("replace", false, "replace a Campaign with the supplied campaign_id")
	flag.Parse()
	if *manifestPath == "" {
		fail(errors.New("--manifest is required"))
	}
	if err := run(context.Background(), *manifestPath, *mediaPath, *replace); err != nil {
		fail(err)
	}
}

func fail(err error) { fmt.Fprintln(os.Stderr, "campaign-seed:", err); os.Exit(1) }

func run(ctx context.Context, manifestPath, mediaPath string, replace bool) error {
	_ = godotenv.Load()
	seed, media, err := loadSeedManifest(manifestPath, mediaPath)
	if err != nil {
		return err
	}
	pool, err := db.Open(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer pool.Close()
	client, err := minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{Creds: credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""), Secure: os.Getenv("MINIO_USE_SSL") == "true"})
	if err != nil {
		return fmt.Errorf("create minio client: %w", err)
	}
	reader := storage.NewPrivateReader(client, os.Getenv("MINIO_BUCKET_PRIVATE"))
	if media != nil {
		file, err := os.Open(mediaPath) // #nosec G304 -- explicit operator-only --media-file input is validated as a local JPEG/PNG.
		if err != nil {
			return fmt.Errorf("open media file: %w", err)
		}
		defer file.Close()
		stat, err := file.Stat()
		if err != nil {
			return fmt.Errorf("stat media file: %w", err)
		}
		if err := reader.Put(ctx, media.ObjectKey, file, stat.Size(), media.ContentType); err != nil {
			return err
		}
	}
	oldKey, err := campaign.NewRepositoryDB(pool).SeedCampaign(ctx, seed, replace)
	if err != nil {
		// The new object is private and never referenced when persistence fails;
		// best-effort cleanup avoids leaving an avoidable orphan.
		if media != nil {
			if cleanupErr := reader.Remove(ctx, media.ObjectKey); cleanupErr != nil {
				fmt.Fprintln(os.Stderr, "campaign-seed: new private object cleanup required")
			}
		}
		return err
	}
	if oldKey != "" && oldKey != mediaObjectKey(media) {
		if err := reader.Remove(ctx, oldKey); err != nil {
			fmt.Fprintln(os.Stderr, "campaign-seed: old private object cleanup required")
		}
	}
	fmt.Printf("persisted campaign %s\n", seed.CampaignID)
	return nil
}

func loadSeedManifest(path, mediaPath string) (campaign.SeedRecord, *campaign.SeedMedia, error) {
	file, err := os.Open(path) // #nosec G304 -- explicit operator-only --manifest input is decoded with a size limit and strict schema.
	if err != nil {
		return campaign.SeedRecord{}, nil, fmt.Errorf("open manifest: %w", err)
	}
	defer file.Close()
	var input manifest
	decoder := json.NewDecoder(io.LimitReader(file, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		return campaign.SeedRecord{}, nil, fmt.Errorf("decode manifest: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return campaign.SeedRecord{}, nil, errors.New("manifest contains multiple JSON values")
	}
	campaignID, err := uuid.Parse(input.CampaignID)
	if err != nil {
		return campaign.SeedRecord{}, nil, errors.New("campaign_id must be a UUID")
	}
	organizationID, err := uuid.Parse(input.OrganizationID)
	if err != nil {
		return campaign.SeedRecord{}, nil, errors.New("organization_id must be a UUID")
	}
	publishedAt, err := parseTime(input.PublishedAt, "published_at")
	if err != nil {
		return campaign.SeedRecord{}, nil, err
	}
	endsAt, err := parseTime(input.FundraisingEndsAt, "fundraising_ends_at")
	if err != nil {
		return campaign.SeedRecord{}, nil, err
	}
	if !endsAt.After(publishedAt) {
		return campaign.SeedRecord{}, nil, errors.New("fundraising_ends_at must be after published_at")
	}
	if !nonEmpty(input.StewardName, input.Title, input.Purpose, input.Story) {
		return campaign.SeedRecord{}, nil, errors.New("steward_name, title, purpose, and story must be non-empty")
	}
	if err := validateAmounts(input.TargetAmount, input.CollectedAmount); err != nil {
		return campaign.SeedRecord{}, nil, err
	}
	seed := campaign.SeedRecord{CampaignID: campaignID, OrganizationID: organizationID, StewardName: input.StewardName, Title: input.Title, Purpose: input.Purpose, Story: input.Story, PublishedAt: publishedAt, FundraisingEndsAt: endsAt, TargetAmount: input.TargetAmount, CollectedAmount: input.CollectedAmount, MediaState: input.MediaState}
	switch input.MediaState {
	case "absent", "unavailable":
		if input.Media != nil || mediaPath != "" {
			return campaign.SeedRecord{}, nil, errors.New("media is only accepted for media_state=available")
		}
		return seed, nil, nil
	case "available":
		if input.Media == nil || mediaPath == "" || strings.TrimSpace(input.Media.AltText) == "" {
			return campaign.SeedRecord{}, nil, errors.New("available media requires media metadata, alt_text, and --media-file")
		}
		mediaID, err := uuid.Parse(input.Media.ID)
		if err != nil {
			return campaign.SeedRecord{}, nil, errors.New("media.id must be a UUID")
		}
		contentType, err := localImageContentType(mediaPath)
		if err != nil {
			return campaign.SeedRecord{}, nil, err
		}
		media := &campaign.SeedMedia{ID: mediaID, ObjectKey: fmt.Sprintf("campaign/%s/%s%s", campaignID, uuid.New(), extensionFor(contentType)), ContentType: contentType, AltText: input.Media.AltText, Caption: input.Media.Caption}
		seed.Media = media
		return seed, media, nil
	default:
		return campaign.SeedRecord{}, nil, errors.New("media_state must be absent, unavailable, or available")
	}
}

func parseTime(value, field string) (time.Time, error) {
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("%s must be RFC3339: %w", field, err)
	}
	return parsed, nil
}
func nonEmpty(values ...string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}
func validateAmounts(target, collected *string) error {
	if (target == nil) != (collected == nil) {
		return errors.New("target_amount and collected_amount must both be present or null")
	}
	for _, raw := range []*string{target, collected} {
		if raw == nil {
			continue
		}
		amount, err := decimal.NewFromString(*raw)
		if err != nil || amount.IsNegative() || !amount.Equal(amount.Truncate(2)) {
			return errors.New("amounts must be non-negative decimal values with at most two places")
		}
	}
	return nil
}
func localImageContentType(path string) (string, error) {
	file, err := os.Open(path) // #nosec G304 -- this is the already validated operator media path, re-opened only to inspect its image format.
	if err != nil {
		return "", fmt.Errorf("open media file: %w", err)
	}
	defer file.Close()
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return "", errors.New("media file must be a valid JPEG or PNG")
	}
	switch format {
	case "jpeg":
		return "image/jpeg", nil
	case "png":
		return "image/png", nil
	default:
		return "", errors.New("media file must be a valid JPEG or PNG")
	}
}
func extensionFor(contentType string) string {
	if contentType == "image/png" {
		return ".png"
	}
	return ".jpg"
}
func mediaObjectKey(media *campaign.SeedMedia) string {
	if media == nil {
		return ""
	}
	return media.ObjectKey
}
