package models

import (
	"context"
	"errors"
	"task/config"
	"task/helpers"
)

type User struct {
	ID          string `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}

func ValidateUser(phoneNumber, userID string, isRegisterFunction bool) error {

	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if isRegisterFunction {
		var phone_number string
		db.QueryRow(context.Background(), "SELECT phone_number FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber).Scan(&phone_number)
		if phone_number != "" {
			return errors.New("this user already exists")
		}
	} else {
		if userID == "" {
			return errors.New("user_id is required")
		}

		// database - de request body - den gelen id bilen gabat gelyan user barmy ya-da yokmy sol barlanyar
		// eger yok bolsa onda error return edilyar
		var id string
		if err := db.QueryRow(context.Background(), "SELECT id FROM users WHERE id = $1 AND deleted_at IS NULL", userID).Scan(&id); err != nil {
			return errors.New("user not found")
		}

		var user_id string
		db.QueryRow(context.Background(), "SELECT id FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber).Scan(&user_id)
		if user_id != userID && user_id != "" {
			return errors.New("this user already exists")
		}

	}

	if !helpers.ValidatePhoneNumber(phoneNumber) {
		return errors.New("invalid phone number")
	}

	return nil

}
