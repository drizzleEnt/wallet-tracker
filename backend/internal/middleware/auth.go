package middleware

import (
	"net/http"

	"github.com/drizzleent/wallet-tracker/backend/internal/service"
	"github.com/gin-gonic/gin"
)

const prefix = "Bearer "

func AuthMiddleware(srv service.AuthService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		headerValue := ctx.GetHeader("Authorization")
		if len(headerValue) < len(prefix) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := headerValue[len(prefix):]
		if len(tokenString) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		u, err := srv.Verify(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", u)
		ctx.Next()
	}
}
