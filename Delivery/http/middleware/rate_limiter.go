package middleware

import (
	"log"
	"net/http"

	"a2sv.org/hub/infrastructure/ip_services"
	"a2sv.org/hub/infrastructure/rate_limit_services"
	"github.com/gin-gonic/gin"
)

func UpstashRateLimiter(limit int, windowSeconds int, upstashURL, upstashToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := ip_services.GetClientIP(c)

		limited, err := rate_limit_services.IsRateLimited(ip, limit, windowSeconds, upstashURL, upstashToken)
		if err != nil {
			log.Println("Rate limit check failed:", err)
			c.Next() // Proceed without rate limiting
			return
		}

		if limited {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			return
		}
		c.Next()
	}
}
