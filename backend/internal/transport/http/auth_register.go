package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler handles POST /auth/register.
// On success it writes 202 with a generic accepted message — identical
// for all four internal branches (anti-enumeration, R7). The service
// returns nil on every branch; the handler does not know which branch
// ran.
func RegisterHandler(svc *account.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		// Boundary validation — reject malformed input before reaching
		// the service. Field names in errors are not sensitive; values
		// are never echoed.
		var fieldErrs []fieldError
		if len(req.Name) < 1 || len(req.Name) > 255 {
			fieldErrs = append(fieldErrs, fieldError{Field: "name", Message: "must be 1-255 characters"})
		}
		if !looksLikeEmail(req.Email) {
			fieldErrs = append(fieldErrs, fieldError{Field: "email", Message: "must be a valid email"})
		}
		if len(req.Password) < 8 { // defense-in-depth; service also checks (R5)
			fieldErrs = append(fieldErrs, fieldError{Field: "password", Message: "must be at least 8 characters"})
		}
		if len(fieldErrs) > 0 {
			WriteValidationError(w, fieldErrs)
			return
		}

		if err := svc.Register(r.Context(), req.Name, req.Email, req.Password); err != nil {
			// Only ErrValidation is expected here; everything else
			// would be a service bug (the four register branches
			// return nil). Map defensively.
			if isErrValidation(err) {
				// Service-level validation (e.g. breach-list hit) —
				// surface as field error on password.
				WriteValidationError(w, []fieldError{
					{Field: "password", Message: "password is not allowed"},
				})
				return
			}
			MapServiceError(w, err)
			return
		}

		// Anti-enumeration: identical 202 generic for all branches.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "If your email is not already registered, you will receive a verification link.",
		})
	}
}

// looksLikeEmail performs a minimal email shape check. Full RFC 5322
// validation is out of scope — the service's crypto.HMAC + DB lookup is
// the real authority on "is this email known". This check exists only to
// reject obviously malformed input at the handler boundary.
func looksLikeEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	at := strings.IndexByte(email, '@')
	if at < 1 || at >= len(email)-2 {
		return false
	}
	return strings.IndexByte(email[at+1:], '.') >= 0
}
