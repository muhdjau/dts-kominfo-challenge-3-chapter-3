package controllers

import (
	"challenge-chapter-3-sesi-3/config"
	"challenge-chapter-3-sesi-3/helpers"
	"challenge-chapter-3-sesi-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	appJson = "application/json"
)

func RegisterUser(ctx *gin.Context) {
	db := config.GetDB()
	var user models.Users
	var role models.Roles

	contentType := helpers.GetContentType(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	err := db.Debug().First(&role, user.RoleID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Role ID not found",
		})
		return
	}

	err = db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"full_name": user.FullName,
			"email":     user.Email,
			"role":      role.RoleName,
		},
		"message": "User success created",
	})

}

func LoginUser(ctx *gin.Context) {
	db := config.GetDB()
	var user models.Users

	contentType := helpers.GetContentType(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	password := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email",
		})
		return
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.RoleID, user.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
