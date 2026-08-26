package notification

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDevSender_SendPasswordResetEmail_AppendsOutboxLine verifies the
// password-reset line lands in the outbox file with recipient + token,
// alongside the pre-existing line types, and that the file keeps its
// 0600 mode.
func TestDevSender_SendPasswordResetEmail_AppendsOutboxLine(t *testing.T) {
	outbox := filepath.Join(t.TempDir(), "outbox", "kencleng-dev-outbox.log")
	s := NewDevSender(outbox)
	ctx := context.Background()

	if err := s.SendVerificationEmail(ctx, "u@example.com", "verify-tok"); err != nil {
		t.Fatalf("SendVerificationEmail: %v", err)
	}
	if err := s.SendNudgeEmail(ctx, "u@example.com", NudgeGoogleOnly); err != nil {
		t.Fatalf("SendNudgeEmail: %v", err)
	}
	if err := s.SendPasswordResetEmail(ctx, "u@example.com", "reset-tok"); err != nil {
		t.Fatalf("SendPasswordResetEmail: %v", err)
	}

	data, err := os.ReadFile(outbox)
	if err != nil {
		t.Fatalf("read outbox: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 outbox lines, got %d: %q", len(lines), string(data))
	}
	if !strings.Contains(lines[0], "[verification]") || !strings.Contains(lines[0], "verify-tok") {
		t.Fatalf("line 0 should be a verification line carrying the token; got %q", lines[0])
	}
	if !strings.Contains(lines[1], "[nudge:google_only]") {
		t.Fatalf("line 1 should be a google_only nudge line; got %q", lines[1])
	}
	if !strings.Contains(lines[2], "[password_reset]") || !strings.Contains(lines[2], "reset-tok") {
		t.Fatalf("line 2 should be a password_reset line carrying the token; got %q", lines[2])
	}

	fi, err := os.Stat(outbox)
	if err != nil {
		t.Fatalf("stat outbox: %v", err)
	}
	if fi.Mode().Perm() != 0o600 {
		t.Fatalf("outbox mode = %v, want 0600", fi.Mode().Perm())
	}
}
