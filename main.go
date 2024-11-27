package main

import (
	"log"
	"pr8_1/configs"
	"pr8_1/db"
	"pr8_1/models"
)

func main() {
	db, err := db.ConnectToDB("postgres", "qwe123", "5432", "tspo_pr7_2")
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

	route := configs.SetupRoutes(db)
	route.Run(":8080")
}
