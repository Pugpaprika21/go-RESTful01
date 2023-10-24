package main

import (
	"go-RESTful01/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	productController := controller.Product{}
	productGroup := r.Group("/products")
	productGroup.GET("/", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.POST("/", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)

	categoryController := controller.Category{}
	categoryGroup := r.Group("/categorys")
	categoryGroup.GET("/", categoryController.FindAll)
	categoryGroup.GET("/:id", categoryController.FindOne)
	categoryGroup.POST("/", categoryController.Create)
	categoryGroup.PATCH("/:id", categoryController.Update)
	categoryGroup.DELETE("/:id", categoryController.Delete)

	orderController := controller.Order{}
	orderGroup := r.Group("/orders")
	orderGroup.GET("/", orderController.FindAll)
	orderGroup.GET("/:id", orderController.FindOne)
	orderGroup.POST("/", orderController.Create)
}
