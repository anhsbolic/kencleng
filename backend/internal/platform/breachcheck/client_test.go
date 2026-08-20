package breachcheck

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestIsBreached_NotInList_ReturnsFalse verifies a password whose suffix
// is absent from the API response is reported as not breached.
func TestIsBreached_NotInList_ReturnsFalse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// A suffix that will not match the test password.
		_, _ = w.Write([]byte("00000000000000000000000000000000000:1\n"))
	}))
	defer ts.Close()

	c := NewClient(5 * time.Second)
	c.baseURL = ts.URL

	got, err := c.IsBreached(context.Background(), "a-very-unique-non-breached-password-xyz")
	if err != nil {
		t.Fatalf("IsBreached: unexpected error: %v", err)
	}
	if got {
		t.Fatal("IsBreached: expected false for non-breached password")
	}
}

// TestIsBreached_InList_ReturnsTrue verifies a password whose suffix is
// present in the API response is reported as breached.
func TestIsBreached_InList_ReturnsTrue(t *testing.T) {
	password := "breached-password-test"
	sum := sha1.Sum([]byte(password))
	full := strings.ToUpper(hex.EncodeToString(sum[:]))
	suffix := full[5:]

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(suffix + ":5\n"))
	}))
	defer ts.Close()

	c := NewClient(5 * time.Second)
	c.baseURL = ts.URL

	got, err := c.IsBreached(context.Background(), password)
	if err != nil {
		t.Fatalf("IsBreached: unexpected error: %v", err)
	}
	if !got {
		t.Fatal("IsBreached: expected true for breached password")
	}
}

// TestIsBreached_API5xx_FailOpen verifies a non-OK status triggers the
// fail-open path (false, nil) and the log line contains neither the
// password nor its full SHA-1 hash.
func TestIsBreached_API5xx_FailOpen(t *testing.T) {
	password := "some-password-for-failopen"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	c := NewClient(5 * time.Second)
	c.baseURL = ts.URL

	got, err := c.IsBreached(context.Background(), password)
	if err != nil {
		t.Fatalf("IsBreached: unexpected error on fail-open: %v", err)
	}
	if got {
		t.Fatal("IsBreached: expected false on fail-open")
	}

	logged := logBuf.String()
	if strings.Contains(logged, password) {
		t.Fatalf("fail-open log must not contain the password; got: %q", logged)
	}
	sum := sha1.Sum([]byte(password))
	fullHash := strings.ToUpper(hex.EncodeToString(sum[:]))
	if strings.Contains(logged, fullHash) {
		t.Fatalf("fail-open log must not contain the password hash; got: %q", logged)
	}
}

// TestIsBreached_ConnectionError_FailOpen verifies a closed server
// (connection refused) triggers the fail-open path (false, nil).
func TestIsBreached_ConnectionError_FailOpen(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	ts.Close() // closed -> connection error

	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	c := NewClient(5 * time.Second)
	c.baseURL = ts.URL

	got, err := c.IsBreached(context.Background(), "irrelevant-password")
	if err != nil {
		t.Fatalf("IsBreached: unexpected error on connection failure: %v", err)
	}
	if got {
		t.Fatal("IsBreached: expected false on connection failure (fail-open)")
	}
}
