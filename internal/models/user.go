package models

import (
	"encoding/json"
	// "fmt"
	"time"
)

const (
	// TODO: уточнить

)

type User struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Email    string `json:"email" binding:"required,email"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
}

type UserViewDTO struct {
	ID    uint
	Email string `json:"email"`
	Login string `json:"login"`
}

func (u *User) Convert() (UserViewDTO, error) {

	jsonModel, err := json.Marshal(u)
	if err != nil {
		return UserViewDTO{}, err
	}

	var userDTO UserViewDTO
	dErr := json.Unmarshal(jsonModel, &userDTO)
	if dErr != nil {
		return UserViewDTO{}, dErr
	}

	return userDTO, nil
}
