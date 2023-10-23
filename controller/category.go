package controller

import (
	"errors"
	"go-RESTful01/db"
	"go-RESTful01/dto"
	"go-RESTful01/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct{}

func (c *Category) FindAll(ctx *gin.Context) {
	var categorys []model.Category
	db.Conn.Find(&categorys)

	var result []dto.CategoryResponse
	for _, category := range categorys {
		result = append(result, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *Category) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})

}

func (c *Category) Create(ctx *gin.Context) {
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		Name: form.Name,
	}

	if err := db.Conn.Create(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func (c *Category) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	category.Name = form.Name
	db.Conn.Save(&category)

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func (c *Category) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	//db.Conn.Unscoped().Delete(&model.Category{}, id) // delete
	db.Conn.Delete(&model.Category{}, id) // soft delete

	ctx.JSON(http.StatusOK, gin.H{"deleteAt": time.Now()})
}
