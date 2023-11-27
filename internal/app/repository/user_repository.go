// Package repository internal/repository/user_repository.go
package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"workspace/internal/app/model"
	"workspace/pkg/database"
)

// UserRepository handles database operations related to users.
type UserRepository struct {
	db *mongo.Database
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.Database,
	}
}

// GetUserByID retrieves a user by ID from MongoDB.
func (r *UserRepository) GetUserByID(userUUID string) (*model.User, error) {
	collection := r.db.Collection("users")
	filter := bson.M{"uuid": userUUID}

	var user model.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUsers retrieves all users from MongoDB.
func (r *UserRepository) GetUsers() ([]model.User, error) {
	collection := r.db.Collection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []model.User
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser creates a new user in the database.
func (r *UserRepository) CreateUser(user model.User) (*model.User, error) {
	// Check email uniqueness
	if err := r.isEmailUnique(user.Email); err != nil {
		return nil, err
	}

	// Generate UUID if not provided
	if user.UUID == "" {
		user.UUID = uuid.New().String()
	}

	collection := r.db.Collection("users")

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// isEmailUnique checks if the given email is unique in the database.
func (r *UserRepository) isEmailUnique(email string) error {
	collection := r.db.Collection("users")

	// Check if a user with the given email already exists
	filter := bson.M{"email": email}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}

	if count > 0 {
		// Email is not unique, return an error
		return ErrEmailNotUnique
	}

	return nil
}

// ErrEmailNotUnique is a custom error indicating that the email is not unique.
var ErrEmailNotUnique = errors.New("email is not unique")

// UpdateUser updates a user in MongoDB.
func (r *UserRepository) UpdateUser(userUUID string, updatedUser model.User) (*model.User, error) {
	updatedUser.UUID = userUUID
	collection := r.db.Collection("users")
	filter := bson.M{"uuid": userUUID}
	update := bson.M{"$set": updatedUser}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

// DeleteUser deletes a user from MongoDB.
func (r *UserRepository) DeleteUser(userUUID string) error {
	collection := r.db.Collection("users")
	filter := bson.M{"uuid": userUUID}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return nil
}
