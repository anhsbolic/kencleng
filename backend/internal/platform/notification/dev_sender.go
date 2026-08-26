package notification

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DevSender is a development-only Sender that appends each "email" to an
// outbox file — a local stand-in for a user's inbox. It is the dev
// equivalent of FakeSender: FakeSender drops the token entirely (safe
// default), DevSender surfaces it in a dev-only file so a developer can
// complete manual verification flows (there is no SMTP in v1).
//
// DevSender never writes to the structured log stream (log.Printf) — the
// outbox file is a simulated inbox, not a log, so the "no tokens in
// logs" golden rule is not violated. Use NewDevSender only when
// APP_ENV=development; prefer NewFakeSender in every other environment.
type DevSender struct {
	outboxPath string
	mu         sync.Mutex
}

// NewDevSender returns a Sender that appends simulated emails to the
// given outbox file path. The file is opened in append mode with mode
// 0600 (owner-only) for hygiene.
func NewDevSender(outboxPath string) *DevSender {
	return &DevSender{outboxPath: filepath.Clean(outboxPath)}
}

// SendVerificationEmail appends a verification "email" line carrying the
// plaintext token to the outbox file. The token must be retrievable for
// manual testing of POST /auth/verify-email — there is no SMTP inbox in
// v1, so this file is the inbox.
func (s *DevSender) SendVerificationEmail(ctx context.Context, to, token string) error {
	return s.append(fmt.Sprintf("[verification] recipient=%s token=%s\n", to, token))
}

// SendNudgeEmail appends a nudge "email" line to the outbox file. The
// nudge carries no secret, but the recipient is recorded (as it would
// be in a real inbox).
func (s *DevSender) SendNudgeEmail(ctx context.Context, to, nudgeType string) error {
	return s.append(fmt.Sprintf("[nudge:%s] recipient=%s\n", nudgeType, to))
}

// SendPasswordResetEmail appends a password-reset "email" line carrying
// the plaintext token to the outbox file. Same rationale as the
// verification line: with no SMTP in v1, this file IS the inbox a
// developer uses to complete POST /auth/reset-password manually.
func (s *DevSender) SendPasswordResetEmail(ctx context.Context, to, token string) error {
	return s.append(fmt.Sprintf("[password_reset] recipient=%s token=%s\n", to, token))
}

// append writes one line to the outbox file under a lock (the server is
// concurrent). It creates the file (and parent dir) on first use.
func (s *DevSender) append(line string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.MkdirAll(filepath.Dir(s.outboxPath), 0o700); err != nil {
		return fmt.Errorf("dev outbox mkdir: %w", err)
	}

	// O_APPEND is atomic for small writes on local filesystems; the lock
	// above serializes goroutines within this process regardless.
	f, err := os.OpenFile(s.outboxPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("dev outbox open: %w", err)
	}
	defer f.Close()

	stamp := time.Now().Format("2006/01/02 15:04:05")
	_, err = fmt.Fprintf(f, "%s %s", stamp, line)
	return err
}
