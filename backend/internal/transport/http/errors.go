package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// problem is the RFC 9457 Problem Details base shape (openapi.yaml
// components.schemas.Problem). ValidationProblem extends it with the
// errors array — both use the same Content-Type.
type problem struct {
	Type   string       `json:"type"`
	Title  string       `json:"title"`
	Status int          `json:"status"`
	Detail string       `json:"detail,omitempty"`
	Errors []fieldError `json:"errors,omitempty"`
}

type fieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// WriteProblem writes a generic RFC 9457 Problem Details response.
// The Content-Type is always application/problem+json. The detail string
// must never leak internals (stack traces, raw SQL, file paths) — it is
// a stable, human-safe sentence per the AGENTS.md golden rule.
func WriteProblem(w http.ResponseWriter, status int, problemType, title, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(problem{
		Type:   problemType,
		Title:  title,
		Status: status,
		Detail: detail,
	})
}

// WriteValidationError writes a 422 ValidationProblem with per-field
// errors. The field names are not sensitive; the password value is never
// echoed.
func WriteValidationError(w http.ResponseWriter, errs []fieldError) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	_ = json.NewEncoder(w).Encode(problem{
		Type:   "https://kencleng.dev/problems/validation-failed",
		Title:  "Validation Failed",
		Status: http.StatusUnprocessableEntity,
		Errors: errs,
	})
}

// Login/session error vocabulary (task #03). The anti-enumeration rule is
// load-bearing here: wrong-email, wrong-password, and lockout share the
// SAME title+detail — only the status code differs between 401 and 429
// (openapi LockedOutGenericCredentials, amended 2026-08-26). The type URI
// is the one machine-readable distinction.
const (
	problemTypeInvalidCredentials = "https://kencleng.dev/errors/invalid-credentials" // #nosec G101 — problem-type URI, not a credential
	problemTypeTooManyRequests    = "https://kencleng.dev/errors/too-many-requests"   // #nosec G101 — problem-type URI, not a credential

	// #nosec G101 — user-facing Indonesian error TEXT (the word "password"
	// inside the sentence trips the heuristic); these are response strings,
	// never credentials.
	problemTitleInvalidCredentials = "Invalid Credentials"
	problemDetailGenericCredential = "Email atau password salah." // #nosec G101
)

// MapServiceError maps a service sentinel error to the appropriate HTTP
// status + Problem Details. Unknown errors map to 500 with a generic
// detail — never the raw error string.
func MapServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, account.ErrValidation):
		WriteProblem(w, http.StatusUnprocessableEntity,
			"https://kencleng.dev/problems/validation-failed",
			"Validation Failed", "The request was invalid.")
	case errors.Is(err, account.ErrTokenExpired):
		WriteProblem(w, http.StatusGone,
			"https://kencleng.dev/problems/token-expired",
			"Token Expired", "The verification token has expired.")
	case errors.Is(err, account.ErrTokenNotFound):
		WriteProblem(w, http.StatusNotFound,
			"https://kencleng.dev/problems/token-not-found",
			"Token Not Found", "The verification token was not found.")
	case errors.Is(err, account.ErrLockedOut):
		WriteProblem(w, http.StatusTooManyRequests,
			problemTypeTooManyRequests,
			problemTitleInvalidCredentials, problemDetailGenericCredential)
	case errors.Is(err, account.ErrMfaPendingInvalid):
		WriteProblem(w, http.StatusUnauthorized,
			problemTypeInvalidCredentials,
			problemTitleInvalidCredentials, problemDetailGenericCredential)
	case errors.Is(err, account.ErrInvalidCredentials):
		WriteProblem(w, http.StatusUnauthorized,
			problemTypeInvalidCredentials,
			problemTitleInvalidCredentials, problemDetailGenericCredential)
	default:
		// Do NOT leak err.Error() — log it server-side, return generic.
		log.Printf("transport: unhandled service error: %v", err)
		WriteProblem(w, http.StatusInternalServerError,
			"https://kencleng.dev/problems/internal",
			"Internal Error", "An unexpected error occurred.")
	}
}

// isErrValidation checks if err wraps account.ErrValidation.
func isErrValidation(err error) bool {
	return errors.Is(err, account.ErrValidation)
}

// write400InvalidJSON writes a 400 Problem for malformed JSON bodies.
func write400InvalidJSON(w http.ResponseWriter) {
	WriteProblem(w, http.StatusBadRequest,
		"https://kencleng.dev/problems/invalid-json",
		"Invalid Request", "The request body was not valid JSON.")
}
