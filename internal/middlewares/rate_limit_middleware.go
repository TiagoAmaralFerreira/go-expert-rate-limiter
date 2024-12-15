package middleware

import (
	"net/http"

	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/service"
)

type RateLimitMiddleware struct {
	service *service.RateLimiterService
}

func NewRateLimitMiddleware(service *service.RateLimiterService) *RateLimitMiddleware {
	return &RateLimitMiddleware{service: service}
}

func (m *RateLimitMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.RemoteAddr // Use IP address as the default key

		// Check for API_KEY in headers
		if token := r.Header.Get("API_KEY"); token != "" {
			key = "token:" + token
		}

		// Verify rate limiting
		rateLimited, err := m.service.IsRateLimited(key)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if rateLimited {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// If not rate-limited, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
