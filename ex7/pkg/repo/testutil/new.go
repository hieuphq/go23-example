package testutil

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	connectionStr := "postgres://postgres:postgres@localhost:5432/ex7_local?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
