package database

import (
	"errors"
	"log"
	"passIt/internal/models"

	"github.com/google/uuid"
)

func (s *service) CreateUser(user *models.User) error {
	log.Printf("Creating user: %+v\n", user)
	result := s.GetGormDB().Create(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected, user not created")
	}
	return nil
}

func (s *service) FindUserById(id uuid.UUID) (models.User, error) {
	var user models.User
	result := s.GetGormDB().First(&user, "id = ?", id)
	if result.Error != nil {
		log.Println("Error finding user by ID:", result.Error)
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *service) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	result := s.GetGormDB().Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Println("Error finding user by ID:", result.Error)
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *service) UpdateUserById(user *models.User) error {
	// result := s.GetGormDB().First(&user, "id = ?", user.ID)
	// if result.Error != nil {
	// 	log.Println("Error finding user by ID:", result.Error)
	// 	return result.Error
	// }
	result := s.GetGormDB().Save(&user)
	if result.Error != nil {
		log.Println("Error Updating user by ID:", result.Error)
		return result.Error
	}
	return nil
}

func (s *service) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.GetGormDB().Find(&users)
	if result.Error != nil {
		log.Println("Error scanning user:", result.Error)
		return nil, result.Error
	}

	return users, nil
}
