package main

import (
	"net/http"
	"testing"
)

func TestNewHTTPServer_ReadHeaderTimeout(t *testing.T) {
	srv := newHTTPServer(":8080", http.NewServeMux())
	if srv.ReadHeaderTimeout != serverReadHeaderTimeout {
		t.Fatalf("ReadHeaderTimeout = %s, want %s", srv.ReadHeaderTimeout, serverReadHeaderTimeout)
	}
}
