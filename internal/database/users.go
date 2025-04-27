package database

import (
	"log"
	"passIt/internal/models"

	uuid "github.com/satori/go.uuid"
)

func (s *service) CreateUser(user *models.User) error {
	// Use the correct query format
	err := s.db.QueryRow(`
		INSERT INTO users (username, email, first_name, last_name, date_of_birth, phone_number, address, is_active, is_admin, last_login, created_at, updated_at, keycloak_id) 
		VALUES ($1, $2, $3, $4, $5::DATE, $6, $7, $8, $9, NOW(), NOW(), NOW(), $10) 
		RETURNING id, created_at, updated_at, last_login`,
		user.Username, user.Email, user.FirstName, user.LastName, user.DOB, user.PhoneNumber, user.Address, user.IsActive, user.IsAdmin, user.KeycloackID).
		Scan(
			&user.ID,        // Scan UUID into user.ID
			&user.CreatedAt, // Scan timestamp into user.CreatedAt
			&user.UpdatedAt, // Scan timestamp into user.UpdatedAt
			&user.LastLogin, // Scan timestamp into user.LastLogin
		)

	if err != nil {
		log.Println("Error creating user in Database:", err)
		return err
	}
	return nil
}

func (s *service) FindUserById(id uuid.UUID) (models.User, error) {
	var user models.User
	err := s.db.QueryRow(`
		SELECT id, username, email, first_name, last_name, date_of_birth, phone_number, address, is_active, is_admin, last_login, created_at, updated_at 
		FROM users WHERE id = $1`, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DOB,
		&user.PhoneNumber,
		&user.Address,
		&user.IsActive,
		&user.IsAdmin,
		&user.LastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return models.User{}, err
	}
	return user, nil
}

func (s *service) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := s.db.QueryRow(`
		SELECT id, username, email, first_name, last_name, date_of_birth, phone_number, address, is_active, is_admin, last_login, created_at, updated_at 
		FROM users WHERE email = $1`, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DOB,
		&user.PhoneNumber,
		&user.Address,
		&user.IsActive,
		&user.IsAdmin,
		&user.LastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Println("Error finding user by email:", err)
		return models.User{}, err
	}
	return user, nil
}

func (s *service) UpdateUserById(user *models.User) error {
	err := s.db.QueryRow(`
		UPDATE users 
		SET username = $1, email = $2, first_name = $3, last_name = $4, date_of_birth = $5::DATE, phone_number = $6, address = $7, is_active = $8, is_admin = $9, updated_at = NOW() 
		WHERE id = $10
		RETURNING created_at, updated_at, last_login`,
		user.Username,
		user.Email,
		user.FirstName,
		user.LastName,
		user.DOB,
		user.PhoneNumber,
		user.Address,
		user.IsActive,
		user.IsAdmin,
		user.ID,
	).Scan(
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	return nil
}

func (s *service) GetAllUsers() ([]models.User, error) {
	rows, err := s.db.Query(`
		SELECT id, username, email, first_name, last_name, date_of_birth, phone_number, address, is_active, is_admin, last_login, created_at, updated_at 
		FROM users`)
	if err != nil {
		log.Println("Error getting all users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.DOB,
			&user.PhoneNumber,
			&user.Address,
			&user.IsActive,
			&user.IsAdmin,
			&user.LastLogin,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
