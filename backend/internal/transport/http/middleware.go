package http

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// RateLimit returns middleware that applies a per-IP token bucket of rps
// requests/second with burst capacity burst. Idle limiters are evicted
// by a background sweeper to bound memory (techplan §7 risk row 3).
//
// Per-IP keying is appropriate for public anonymous endpoints
// (register/verify/resend) — there is no authenticated identity to key
// on. Behind a reverse proxy r.RemoteAddr is the proxy IP; a
// future iteration should handle X-Forwarded-For if a proxy is added.
func RateLimit(rps float64, burst int) func(http.Handler) http.Handler {
	var (
		mu       sync.Mutex
		limiters = make(map[string]*rate.Limiter)
		lastSeen = make(map[string]time.Time)
	)

	// Background eviction: drop entries not seen in the last TTL.
	const ttl = 10 * time.Minute
	go func() {
		t := time.NewTicker(time.Minute)
		defer t.Stop()
		for range t.C {
			mu.Lock()
			now := time.Now()
			for ip, seen := range lastSeen {
				if now.Sub(seen) > ttl {
					delete(limiters, ip)
					delete(lastSeen, ip)
				}
			}
			mu.Unlock()
		}
	}()

	getLimiter := func(ip string) *rate.Limiter {
		mu.Lock()
		defer mu.Unlock()
		l, ok := limiters[ip]
		if !ok {
			l = rate.NewLimiter(rate.Limit(rps), burst)
			limiters[ip] = l
		}
		lastSeen[ip] = time.Now()
		return l
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr // fallback for non-host:port forms
			}
			if !getLimiter(ip).Allow() {
				WriteProblem(w, http.StatusTooManyRequests,
					"https://kencleng.dev/problems/rate-limited",
					"Rate Limited", "Too many requests. Try again later.")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
