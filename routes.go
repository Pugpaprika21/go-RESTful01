package main

import (
	"github.com/gin-gonic/gin"
	"go-RESTful01/controller"
)

func serveRoutes(r *gin.Engine) {
	//
	productController := controller.Product{}
	productGroup := r.Group("/products")
	productGroup.GET("/", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.POST("/:id", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)
}
