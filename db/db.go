package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dbUser string, dbPassword string, dbPort string, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbPort, dbName,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
