package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID  `json:"id"`
	KeycloackID string     `json:"keycloak_id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DOB         CustomTime `json:"date_of_birth"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	LastLogin   time.Time  `json:"last_login"`
	IsActive    bool       `json:"is_active"`
	IsAdmin     bool       `json:"is_admin"`
}

// CustomTime handles custom date formats
type CustomTime struct {
	time.Time
}

// Custom UnmarshalJSON function to parse "YYYY-MM-DD"
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // Remove quotes from JSON string
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// Scan handles SQL input (time.Time)
func (ct *CustomTime) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}
	ct.Time = t
	return nil
}

// Value converts CustomTime to a format SQL understands
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}
