package services

import (
	"errors"
	"server/api/template"
	"server/api/utils"
	"server/models"
	"server/setup"

	"gorm.io/gorm"
)

var (
	RecordsNotFound    error = errors.New("records not found")
	InvalidCredentials error = errors.New("invalid credentials")
	ServerError        error = errors.New("server error")
	DuplicateData      error = errors.New("duplicate data")
)

func Login(r template.LoginRequest) (*models.User, string, error) {
	db := setup.DB
	username := r.Username
	password := r.Password

	var user models.User
	res := db.Select("id", "name", "username", "password").First(&user, "username = ?", username)
	if err := res.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", RecordsNotFound
	}

	if !utils.VerifyPassword(password, user.Password) {
		return nil, "", InvalidCredentials
	}

	token, err := utils.MakeJWT(user.ID)
	if err != nil {
		return nil, "", ServerError
	}

	return &user, token, nil
}

func RegisterNewUser(r template.RegisterRequest) (*models.User, error) {
	db := setup.DB
	name := r.Name
	username := r.Username
	password := r.Password

	var cnt int64
	db.First(&models.User{}, "username = ?", username).Count(&cnt)
	if cnt != 0 {
		return nil, DuplicateData
	}

	var newUser models.User
	newUser.Name = name
	newUser.Username = username
	newUser.Password = utils.HashPassword(password)
	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}
