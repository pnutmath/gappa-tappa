package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository ...
type Repository struct {
	DB *mongo.Database
}

// NewUserRepository ...
func NewUserRepository(db *mongo.Database) *Repository {
	return &Repository{DB: db}
}

// Get ...
func (repo *Repository) Get(username string) (*User, error) {
	return &User{ID: 1, Username: username}, nil
}
