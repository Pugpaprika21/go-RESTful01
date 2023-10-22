package main

import (
	"go-RESTful01/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	db.ConnectDB()

	os.MkdirAll("upload/products", 0755)
	r := gin.Default()
	serveRoutes(r)
	r.Run()
}
