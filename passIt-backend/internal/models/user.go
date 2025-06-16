package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// User represents a user in the system
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Username    string         `gorm:"unique;not null" json:"username"`
	Email       string         `gorm:"unique;not null" json:"email"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	DOB         time.Time      `json:"dob"`
	PhoneNumber string         `gorm:"not null" json:"phone_number"`
	Address     string         `gorm:"not null" json:"address"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	IsAdmin     bool           `gorm:"default:false" json:"is_admin"`
}

// CustomTime handles custom date formats
// type CustomTime struct {
// 	time.Time
// }

// // Custom UnmarshalJSON function to parse "YYYY-MM-DD"
// func (ct *CustomTime) UnmarshalJSON(b []byte) error {
// 	str := string(b)
// 	str = str[1 : len(str)-1] // Remove quotes from JSON string
// 	t, err := time.Parse("2006-01-02", str)
// 	if err != nil {
// 		return err
// 	}
// 	ct.Time = t
// 	return nil
// }

// // Scan handles SQL input (time.Time)
// func (ct *CustomTime) Scan(value interface{}) error {
// 	t, ok := value.(time.Time)
// 	if !ok {
// 		return fmt.Errorf("cannot scan type %T into CustomTime", value)
// 	}
// 	ct.Time = t
// 	return nil
// }

// // Value converts CustomTime to a format SQL understands
// func (ct CustomTime) Value() (driver.Value, error) {
// 	return ct.Time, nil
// }
