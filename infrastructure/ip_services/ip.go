package ip_services

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	ip := c.GetHeader("X-Forwarded-For")
	if ip != "" {
		return strings.Split(ip, ",")[0]
	}
	ip = c.GetHeader("X-Real-IP")
	if ip != "" {
		return ip
	}
	ip, _, _ = net.SplitHostPort(c.Request.RemoteAddr)
	return ip
}
