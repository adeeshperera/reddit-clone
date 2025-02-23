package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient(host, user, password, dbname, port string) (*gorm.DB, error) {
	log.Println("Initializing PostgreSQL connection...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v\n", err)
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}
