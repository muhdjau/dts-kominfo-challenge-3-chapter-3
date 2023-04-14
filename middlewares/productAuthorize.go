package middlewares

import (
	"challenge-chapter-3-sesi-3/config"
	"challenge-chapter-3-sesi-3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := config.GetDB()

		productID, _ := strconv.Atoi(ctx.Param("productID"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		roleID := uint(userData["role"].(float64))

		if roleID == 1 {
			ctx.Next()
		} else {
			var product models.Products
			err := db.Select("user_id").First(&product, uint(productID)).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Data not found",
				})
				return
			}
			if product.UserID != userID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "You are not allowed to access this data",
				})
				return
			} else {
				ctx.Next()
			}
		}
	}
}
