package http

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// profileService is the subset of *account.Service the account-profile
// handler depends on. *account.Service satisfies it; tests inject a stub
// so the transport contract (status codes, response shape) is exercisable
// without the full domain — same seam philosophy as securityService.
type profileService interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (*account.LoginUserView, error)
}

// userResponse mirrors openapi components.schemas.User exactly — snake_case
// keys are deliberate (the User schema's contract), unlike the untagged
// LoginUserView whose Go field names would leak camelCase onto the wire.
// uuid.UUID marshals via MarshalText (uuid string) and time.Time via RFC
// 3339 — both schema formats satisfied by encoding/json defaults.
type userResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Roles         []string  `json:"roles"`
	AuthProviders []string  `json:"auth_providers"`
	MFAEnabled    bool      `json:"mfa_enabled"`
	CreatedAt     time.Time `json:"created_at"`
}

// toUserResponse maps a domain LoginUserView onto the wire contract. It
// returns nil for a nil view (so a `*userResponse` field with omitempty
// stays omitted, mirroring the previous pointer semantics) and normalizes
// nil Roles/AuthProviders to empty slices so the arrays serialize as `[]`,
// never `null` (nil-and-zero-values §2 — the boundary owns this guarantee
// deliberately, rather than inheriting it from a repository implementation
// detail).
func toUserResponse(v *account.LoginUserView) *userResponse {
	if v == nil {
		return nil
	}
	roles := v.Roles
	if roles == nil {
		roles = []string{}
	}
	providers := v.AuthProviders
	if providers == nil {
		providers = []string{}
	}
	return &userResponse{
		ID:            v.ID,
		Name:          v.Name,
		Email:         v.Email,
		EmailVerified: v.EmailVerified,
		Roles:         roles,
		AuthProviders: providers,
		MFAEnabled:    v.MFAEnabled,
		CreatedAt:     v.CreatedAt,
	}
}

// AccountMeHandler handles GET /account/me — returns the authenticated
// user's own profile. No ID parameter; the resource is keyed entirely by
// the session (no IDOR surface, threat-model component 5). A session whose
// user no longer exists maps to 401 with the same generic body as a missing
// token — the spec documents only 401 for this endpoint.
func AccountMeHandler(svc profileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := UserIDFromContext(r.Context())
		if !ok {
			WriteProblem(w, http.StatusUnauthorized,
				"https://kencleng.dev/errors/unauthorized",
				"Unauthorized", "Authentication required.")
			return
		}

		view, err := svc.GetProfile(r.Context(), userID)
		if err != nil {
			MapServiceError(w, err)
			return
		}
		if view == nil {
			WriteProblem(w, http.StatusUnauthorized,
				"https://kencleng.dev/errors/unauthorized",
				"Unauthorized", "Authentication required.")
			return
		}

		writeJSON(w, http.StatusOK, toUserResponse(view))
	}
}
