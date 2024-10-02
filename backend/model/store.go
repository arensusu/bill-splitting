package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func InitGorm() *gorm.DB {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return db
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Migrate() {
	s.db.AutoMigrate(&User{}, &Group{}, &Member{}, &Expense{}, &Settlement{})
}
