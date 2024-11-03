package middleware

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter defines the structure for our rate limiter.
type RateLimiter struct {
	mu        sync.Mutex
	clients   map[string]time.Time // Track client requests
	rateLimit time.Duration        // Time duration limit per request
}

// NewRateLimiter initializes a rate limiter with a given limit duration.
func NewRateLimiter(rateLimit time.Duration) *RateLimiter {
	return &RateLimiter{
		clients:   make(map[string]time.Time),
		rateLimit: rateLimit,
	}
}

// LimitMiddleware is the middleware function to enforce rate limiting.
func (rl *RateLimiter) LimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr // You could also use a user ID or other identifier

		rl.mu.Lock()
		defer rl.mu.Unlock()

		lastRequest, exists := rl.clients[clientIP]
		now := time.Now()

		// Check if this request is within the rate limit duration
		if exists && now.Sub(lastRequest) < rl.rateLimit {
			http.Error(w, "Too many requests. Please wait before retrying.", http.StatusTooManyRequests)
			return
		}

		// Update last request time for the client
		rl.clients[clientIP] = now

		next.ServeHTTP(w, r)
	})
}
