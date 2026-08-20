// Package breachcheck checks passwords against the HaveIBeenPwned
// k-anonymity service. It fails open: if the API is unreachable,
// IsBreached returns (false, nil) and logs the fact.
package breachcheck

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
		// Fail-open: log the fact + category, not the password.
		log.Printf("breachcheck: API unreachable, proceeding without check: %v", err)
		return false, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
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
