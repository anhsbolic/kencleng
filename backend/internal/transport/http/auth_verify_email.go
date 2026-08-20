package http

import (
	"encoding/json"
	"net/http"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

type verifyEmailRequest struct {
	Token string `json:"token"`
}

type resendRequest struct {
	Email string `json:"email"`
}

// VerifyEmailHandler handles POST /auth/verify-email.
// R8: valid → 200. R9: expired → 410. R10/R11: not-found/used/revoked → 404.
func VerifyEmailHandler(svc *account.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req verifyEmailRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		// Reject empty token at the boundary — saves a DB round-trip
		// (techplan §13 row 9) and avoids any timing distinction between
		// "empty" and "not found".
		if req.Token == "" {
			MapServiceError(w, account.ErrTokenNotFound)
			return
		}

		if err := svc.VerifyEmail(r.Context(), req.Token); err != nil {
			MapServiceError(w, err) // ErrTokenExpired → 410, ErrTokenNotFound → 404
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Email verified.",
		})
	}
}

// ResendVerificationHandler handles POST /auth/verify-email/resend.
// R13/R14: always 202 generic — identical whether or not a token was
// issued.
func ResendVerificationHandler(svc *account.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req resendRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		if !looksLikeEmail(req.Email) {
			WriteValidationError(w, []fieldError{
				{Field: "email", Message: "must be a valid email"},
			})
			return
		}

		_ = svc.ResendVerification(r.Context(), req.Email)

		// Always 202 generic — the service returns nil for both
		// match and no-match branches (R13/R14).
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "If your email is registered and unverified, you will receive a new verification link.",
		})
	}
}
