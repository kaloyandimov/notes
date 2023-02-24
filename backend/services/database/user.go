package database

import (
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"notes/backend/models"
	"notes/backend/utilities/token"
)

func GetUserByID(uid uint) (models.User, error) {
	var user models.User

	if err := database.Omit("password").First(&user, uid).Error; err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func AddUser(email, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)

	if err != nil {
		return err
	}

	var userRoleID uint
	userRoleID, err = GetRoleID("user")

	if err != nil {
		return err
	}

	user := models.User{
		Email:    html.EscapeString(strings.TrimSpace(email)),
		Username: html.EscapeString(strings.TrimSpace(username)),
		Password: string(hashedPassword),
		RoleID:   userRoleID,
	}

	return database.Create(&user).Error
}

func LoginUser(username, password string) (string, error) {
	var user models.User
	var err error

	if err = database.Where("username = ?", username).Take(&user).Error; err != nil {
		return "", err
	}

	if err = verifyPassword(user.Password, password); err != nil {
		return "", err
	}

	var tkn string
	tkn, err = token.Generate(user.ID, user.RoleID)

	if err != nil {
		return "", err
	}

	return tkn, nil
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
