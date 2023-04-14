package controllers

import (
	"challenge-chapter-3-sesi-3/config"
	"challenge-chapter-3-sesi-3/helpers"
	"challenge-chapter-3-sesi-3/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateProduct(ctx *gin.Context) {
	db := config.GetDB()
	var newProduct models.Products

	contentType := helpers.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		ctx.ShouldBindJSON(&newProduct)
	} else {
		ctx.ShouldBind(&newProduct)
	}

	newProduct.UserID = userID

	err := db.Debug().Create(&newProduct).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newProduct)

}

func GetProductById(ctx *gin.Context) {
	db := config.GetDB()
	var product models.Products

	convProductID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	err = db.Debug().First(&product, "id = ?", convProductID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func GetAllProducts(ctx *gin.Context) {
	db := config.GetDB()
	var allProducts []models.Products

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	roleID := uint(userData["role"].(float64))

	if roleID == 1 {
		err := db.Debug().Find(&allProducts).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	if roleID == 2 {
		err := db.Debug().Where("user_id", userID).Find(&allProducts).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data_product": allProducts,
	})

}

func UpdateProduct(ctx *gin.Context) {
	db := config.GetDB()
	var updatedProduct, findProduct models.Products

	contentType := helpers.GetContentType(ctx)

	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&updatedProduct)
	} else {
		ctx.ShouldBind(&updatedProduct)
	}

	updatedProduct = models.Products{
		Title:       updatedProduct.Title,
		Description: updatedProduct.Description,
	}

	err = db.Where("id = ?", productID).First(&findProduct).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	updatedProduct.ID = uint(productID)
	updatedProduct.UserID = findProduct.UserID
	updatedProduct.CreatedAt = findProduct.CreatedAt

	err = db.Model(&updatedProduct).Where("id = ?", productID).Updates(updatedProduct).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, updatedProduct)
}

func DeleteProduct(ctx *gin.Context) {
	db := config.GetDB()
	var product models.Products

	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	err = db.Debug().Where("id = ?", productID).First(&product).Delete(&product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Product %s success deleted", product.Title),
	})
}
