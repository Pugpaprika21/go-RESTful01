package main

import (
	"go-RESTful01/db"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	db.ConnectDB()
	db.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	os.MkdirAll("uploads/products", 0755)
	r := gin.Default()
	r.Use(cors.New(corsConfig))
	r.Static("/uploads", "./uploads")
	serveRoutes(r)
	r.Run(":" + os.Getenv("PORT"))
}
