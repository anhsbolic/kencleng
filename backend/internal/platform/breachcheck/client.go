// Package breachcheck checks passwords against the HaveIBeenPwned
// k-anonymity service. It fails open: if the API is unreachable,
// IsBreached returns (false, nil) and logs the fact.
package breachcheck

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"time"
)

// Client is a HaveIBeenPwned k-anonymity breach-check client. It is safe
// for concurrent use and must be constructed once and reused — do not
// construct one per call (an explicit timeout is set at construction).
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient constructs a breach-check client with the given per-call
// timeout. The returned client is safe for concurrent use and should
// be reused — do not construct one per call.
func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		baseURL:    "https://api.pwnedpasswords.com/range",
	}
}

// IsBreached reports whether the password has appeared in a known
// breach. On API unreachable it returns (false, nil) — fail-open,
// logged via the standard logger. Never returns the password or its
// hash in any error or log line.
func (c *Client) IsBreached(ctx context.Context, password string) (bool, error) {
	sum := sha1.Sum([]byte(password))
	full := strings.ToUpper(hex.EncodeToString(sum[:]))
	prefix, suffix := full[:5], full[5:]

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+prefix, nil)
	if err != nil {
		return false, fmt.Errorf("breachcheck: build request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Fail-open: log a sanitized category, never the raw error. A
		// *url.Error from http.Client.Do embeds the request URL, which
		// contains the 5-char SHA-1 prefix of the password — partial
		// credential-derived data. The k-anonymity prefix is safe to
		// send to the server, not to log. Per
		// go/secrets-and-sensitive-logging.md §1.
		log.Printf("breachcheck: API unreachable (%s), proceeding without check",
			breachErrorCategory(err))
		return false, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// Drain the body so the connection returns to the pool for reuse
		// (go/http-client-and-transport.md §2).
		_, _ = io.Copy(io.Discard, resp.Body)
		log.Printf("breachcheck: API returned status %d, proceeding without check", resp.StatusCode)
		return false, nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("breachcheck: read response: %w", err)
	}
	for _, line := range strings.Split(string(body), "\n") {
		if strings.HasPrefix(line, suffix+":") {
			return true, nil
		}
	}
	return false, nil
}

// breachErrorCategory reduces a breachcheck HTTP error to a safe,
// PII-free category string for logging. It never returns the request
// URL (which carries the 5-char SHA-1 password-hash prefix) or any
// password-derived string. Per go/secrets-and-sensitive-logging.md §1.
func breachErrorCategory(err error) string {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		// urlErr.Op ("Get"/"Post") is safe; urlErr.URL is NOT (contains
		// the SHA-1 prefix). Extract only the op + a coarse net category.
		if urlErr.Timeout() {
			return "timeout"
		}
		return fmt.Sprintf("%s: %s", urlErr.Op, classifyNetErr(urlErr.Err))
	}
	// Fallback: a coarse category, not the verbatim message.
	return "transport error"
}

// classifyNetErr maps a low-level net error to a short, safe category.
func classifyNetErr(err error) string {
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return "timeout"
	}
	if errors.Is(err, context.Canceled) {
		return "canceled"
	}
	if errors.Is(err, io.EOF) || errors.Is(err, syscall.ECONNRESET) {
		return "connection reset"
	}
	return "network error"
}
