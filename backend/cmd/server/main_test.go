package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEnvironmentWithoutDotenvUsesExplicitContainerSettings(t *testing.T) {
	t.Chdir(t.TempDir())
	t.Setenv("APP_PORT", "8090")
	t.Setenv("DATABASE_URL", "postgres://container-example")
	if err := loadEnvironment(); err != nil {
		t.Fatalf("explicit runtime settings: %v", err)
	}
}

func TestEnvironmentStillRequiresDatabaseAndPort(t *testing.T) {
	for _, missing := range []string{"APP_PORT", "DATABASE_URL"} {
		t.Run(missing, func(t *testing.T) {
			t.Chdir(t.TempDir())
			t.Setenv("APP_PORT", "8090")
			t.Setenv("DATABASE_URL", "postgres://container-example")
			t.Setenv(missing, "")
			if err := loadEnvironment(); err == nil || !strings.Contains(err.Error(), missing) {
				t.Fatalf("want missing %s, got %v", missing, err)
			}
		})
	}
}

func TestDotenvDoesNotOverrideExplicitRuntimeDatabase(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	t.Setenv("APP_PORT", "8090")
	t.Setenv("DATABASE_URL", "postgres://container-example")
	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte("DATABASE_URL=postgres://host-example\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := loadEnvironment(); err != nil {
		t.Fatal(err)
	}
	if actual := os.Getenv("DATABASE_URL"); actual != "postgres://container-example" {
		t.Fatal("dotenv replaced the explicit runtime database")
	}
}

func TestMalformedDotenvFailsEvenWithRuntimeSettings(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	t.Setenv("APP_PORT", "8090")
	t.Setenv("DATABASE_URL", "postgres://container-example")
	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte("BROKEN='unterminated\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := loadEnvironment(); err == nil || !strings.Contains(err.Error(), "load .env") {
		t.Fatalf("want malformed dotenv failure, got %v", err)
	}
}
