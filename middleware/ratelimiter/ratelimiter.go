package ratelimiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func RateLimiter(rate int, interval time.Duration, tokenExpiration time.Duration) gin.HandlerFunc {
	rateLimiters := cache.New(tokenExpiration, 2*tokenExpiration)

	return func(ctx *gin.Context) {
		userToken := ctx.GetHeader("X-User-Token")
		if userToken == "" {
			ctx.AbortWithStatus(401)
			return
		}

		limiter, ok := rateLimiters.Get(userToken)
		if !ok {
			limiter = newRateLimiter(rate, interval)
			rateLimiters.Set(userToken, limiter, tokenExpiration)
		}

		if !limiter.(*rateLimiter).acquire() {
			ctx.AbortWithStatus(429)
			return
		}

		ctx.Next()

		limiter.(*rateLimiter).lastAccess = time.Now()
	}
}

type rateLimiter struct {
	rate       int
	interval   time.Duration
	tokens     chan time.Time
	lastAccess time.Time
}

func newRateLimiter(rate int, interval time.Duration) *rateLimiter {
	return &rateLimiter{
		rate:       rate,
		interval:   interval,
		tokens:     make(chan time.Time, rate),
		lastAccess: time.Now(),
	}
}

func (lr *rateLimiter) acquire() bool {
	select {
	case lr.tokens <- time.Now():
		return true
	default:
		return false
	}
}
