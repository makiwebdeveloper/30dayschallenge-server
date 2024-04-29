package repository

import (
	"log"
	"os"
)

type Repository struct {
	User
}

func NewRepository() *Repository {
	db, err := NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to postgres database. \n", err)
		os.Exit(2)
	}

	return &Repository{
		User: NewUserRepository(db),
	}
}
