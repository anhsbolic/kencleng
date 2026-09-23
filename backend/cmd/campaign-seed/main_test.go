package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSeedManifest_RejectsInvalidOperatorInputBeforeAnyWrite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "manifest.json")
	content := `{"campaign_id":"d1312b3f-8a0e-4f05-9202-aeacf14053a7","organization_id":"11906b77-6087-4df9-9688-6e61e65917d5","steward_name":"Steward","title":"Title","purpose":"Purpose","story":"Story","published_at":"2026-01-01T00:00:00Z","fundraising_ends_at":"2026-02-01T00:00:00Z","target_amount":"1.00","collected_amount":null,"media_state":"absent"}`
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, _, err := loadSeedManifest(path, ""); err == nil {
		t.Fatal("expected partial funding pair rejection")
	}
}

func TestLoadSeedManifest_AcceptsUnavailableMediaWithoutFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "manifest.json")
	content := `{"campaign_id":"d1312b3f-8a0e-4f05-9202-aeacf14053a7","organization_id":"11906b77-6087-4df9-9688-6e61e65917d5","steward_name":"Steward","title":"Title","purpose":"Purpose","story":"Story","published_at":"2026-01-01T00:00:00Z","fundraising_ends_at":"2026-02-01T00:00:00Z","target_amount":null,"collected_amount":null,"media_state":"unavailable"}`
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	seed, media, err := loadSeedManifest(path, "")
	if err != nil {
		t.Fatalf("loadSeedManifest: %v", err)
	}
	if media != nil || seed.Media != nil || seed.MediaState != "unavailable" {
		t.Errorf("seed = %#v media = %#v", seed, media)
	}
}
