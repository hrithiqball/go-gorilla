package middlewares

import (
	"local_my_api/pkg/utils"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiterMap map[string]*rate.Limiter
	mu         sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiterMap: make(map[string]*rate.Limiter),
	}
}

func (rl *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if limiter, exists := rl.limiterMap[ip]; exists {
		return limiter
	}

	limiter := rate.NewLimiter(1, 5)
	rl.limiterMap[ip] = limiter
	return limiter
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := rl.GetLimiter(ip)

		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			utils.ResponseWithError(w, http.StatusTooManyRequests, "rate limit exceeded")
		}
	})
}
