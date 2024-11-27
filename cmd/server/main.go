package main

import (
	"log"
	_ "tspo_final/docs"
	"tspo_final/internal/db"
	"tspo_final/internal/models"
	"tspo_final/internal/routes"

	"github.com/ermos/dotenv"
)

func main() {
	db, err := db.ConnectToDB(dotenv.GetString("DB_ALIAS"), dotenv.GetString("DB_USER"), dotenv.GetString("DB_PASSWORD"), dotenv.GetString("DB_PORT"), dotenv.GetString("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Favorite{})
	db.AutoMigrate(&models.Basket{})
	db.AutoMigrate(&models.UserOrder{})
	db.AutoMigrate(&models.UserRole{})
	db.AutoMigrate(&models.GoodFeature{})
	db.AutoMigrate(&models.GoodVendor{})
	db.AutoMigrate(&models.OrderGood{})
	db.AutoMigrate(&models.Employee{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Feature{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Vendor{})
	db.AutoMigrate(&models.Good{})
	db.AutoMigrate(&models.Order{})

	route := routes.SetupRoutes(db)

	route.Run(":8080")
}
