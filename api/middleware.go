package api

import (
	"bill-splitting/token"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func authMiddleware(tokenMaker *token.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is not provided"})
			c.Abort()
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unsupported authorization type"})
			c.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		fmt.Println(payload)
		c.Set("payload", payload)
		c.Next()
	}
}
