package account

import (
	"context"

	"github.com/google/uuid"
)

// GetProfile returns the resource owner's own profile view — the same
// LoginUserView login assembles, including the decrypted primary email
// (self-view only; the MaskedField concern applies to other users' PII,
// never this one). Returns (nil, nil) when the session's user no longer
// exists; the transport maps that to 401 (session no longer identifies a
// valid account).
func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*LoginUserView, error) {
	return s.repo.GetLoginUserView(ctx, userID)
}
