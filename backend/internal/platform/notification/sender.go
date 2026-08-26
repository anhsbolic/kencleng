// Package notification defines the email-sending seam used by the
// account domain. v1 ships FakeSender (logged, no SMTP); a real SMTP
// implementation will be added in the notification domain phase.
package notification

import (
	"context"
	"log"
)

// Sender abstracts outbound email delivery.
// Implementations must not block on network I/O inside a caller's
// DB transaction — callers are responsible for sending after commit.
type Sender interface {
	SendVerificationEmail(ctx context.Context, to, token string) error
	SendNudgeEmail(ctx context.Context, to, nudgeType string) error
	SendPasswordResetEmail(ctx context.Context, to, token string) error
}

// FakeSender logs the recipient and message type instead of sending.
// Suitable for v1 and for tests. Does not perform any network I/O.
type FakeSender struct{}

// NewFakeSender returns a Sender that logs rather than sends.
func NewFakeSender() *FakeSender { return &FakeSender{} }

// Nudge type constants — keep in sync with service calls.
const (
	NudgePasswordReset = "password_reset"
	NudgeGoogleOnly    = "google_only"
)

// SendVerificationEmail logs the fact that a verification email was
// queued. It never logs the recipient address or the token.
func (FakeSender) SendVerificationEmail(ctx context.Context, to, token string) error {
	log.Printf("notification: verification email queued (recipient redacted)")
	return nil
}

// SendNudgeEmail logs the fact that a nudge email was queued along with
// its type. It never logs the recipient address.
func (FakeSender) SendNudgeEmail(ctx context.Context, to, nudgeType string) error {
	log.Printf("notification: nudge email queued type=%s (recipient redacted)", nudgeType)
	return nil
}

// SendPasswordResetEmail logs the fact that a password-reset email was
// queued. It never logs the recipient address or the token.
func (FakeSender) SendPasswordResetEmail(ctx context.Context, to, token string) error {
	log.Printf("notification: password reset email queued (recipient redacted)")
	return nil
}
