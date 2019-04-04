package main

import (
	"log"

	"github.com/jinzhu/gorm"
	pb "github.com/willdot/GRPC-Demo/user-service/proto/auth"
)

// Repository ..
type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
}

// UserRepository is a datastore
type UserRepository struct {
	db *gorm.DB
}

// GetAll will get all users from database
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User

	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Get will get a single user
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id

	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmail will get a user by email
func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create will create a new user
func (repo *UserRepository) Create(user *pb.User) error {
	log.Println("User")
	log.Println(user)
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
