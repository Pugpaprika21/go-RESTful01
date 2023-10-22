package controller

import (
	"go-RESTful01/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category struct{}

func (c *Category) FindAll(ctx *gin.Context) {

}

func (c *Category) FindOne(ctx *gin.Context) {

}

func (c *Category) Create(ctx *gin.Context) {
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Name": form.Name})

}

func (c *Category) Update(ctx *gin.Context) {

}

func (c *Category) Delete(ctx *gin.Context) {

}
