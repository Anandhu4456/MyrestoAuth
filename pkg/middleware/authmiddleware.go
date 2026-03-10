package middleware

import (
	"myresto/pkg/cfg"
	"myresto/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserID = "user_id"
	ContextEmail  = "user_email"
)

func AuthMiddleware(cfg *cfg.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header missing",
			})
			ctx.Abort()
			return
		}

		tokenStr := strings.Split(authHeader, " ")
		if len(tokenStr) != 2 || tokenStr[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorrization header format",
			})
			ctx.Abort()
			return
		}

		token := tokenStr[1] // removng 'Bearer'

		claims, err := jwt.ValidateAccessToken(token, cfg)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			ctx.Abort()
			return
		}

		// setting user details into context
		ctx.Set(ContextUserID, claims.UserID)
		ctx.Set(ContextEmail, claims.Email)

		ctx.Next()
	}
}
