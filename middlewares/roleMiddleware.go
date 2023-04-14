package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RoleMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		roleID := uint(userData["role"].(float64))

		if roleID != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not admin",
			})
			return
		}

		ctx.Next()
	}
}
