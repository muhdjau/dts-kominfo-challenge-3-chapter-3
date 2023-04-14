package controllers

import (
	"challenge-chapter-3-sesi-3/config"
	"challenge-chapter-3-sesi-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRole(ctx *gin.Context) {
	db := config.GetDB()
	var role models.Roles

	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Debug().Create(&role).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid create data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

func GetRole(ctx *gin.Context) {
	db := config.GetDB()
	var roleDatas []models.Roles

	err := db.Debug().Find(&roleDatas).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": roleDatas,
	})
}
