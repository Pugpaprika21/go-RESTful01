package db

import (
	"go-RESTful01/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Connot connect to database")
	}

	Conn = db
}

func Migrate() {
	Conn.AutoMigrate(
		&model.Category{},
		&model.Order{},
		&model.OrderItem{},
		&model.Product{},
	)
}
