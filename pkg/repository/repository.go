package repository

import (
	"log"
	"os"
)

type Repository struct {
}

func NewRepository() *Repository {
	_, err := NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to postgres database. \n", err)
		os.Exit(2)
	}

	return &Repository{}
}
