package services

import (
	"errors"
	"log"
	"sticker/src/config/database"
	users_entity "sticker/src/entities/users"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(user *users_entity.User) users_entity.User {
	password, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal("Error to encrypt password")
	}

	userToInsert := users_entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}

	database.Instance.Create(&userToInsert)

	return userToInsert
}

func UpdateUser(userId int, user *users_entity.User) (users_entity.User, error) {
	var userOnDB users_entity.User
	database.Instance.First(&userOnDB, userId)

	if userOnDB.ID == 0 {
		return userOnDB, errors.New("user not found")
	}

	userToUpdate := user
	userToUpdate.ID = uint(userId)

	database.Instance.Save(&userToUpdate)

	return users_entity.User{
		ID:    userToUpdate.ID,
		Name:  userToUpdate.Name,
		Email: userToUpdate.Email,
	}, nil
}

func ValidateUserByEmailAndPassword(email string, password string) error {
	var userOnDB users_entity.User
	// database.Instance.Find(&userOnDB, email)
	database.Instance.Where("email = ?", email).First(&userOnDB)

	if userOnDB.ID == 0 {
		return errors.New("user not found")
	}

	hasMatch := CheckPasswordHash(password, userOnDB.Password)
	if !hasMatch {
		return errors.New("credentials does not match")

	}

	return nil
}
