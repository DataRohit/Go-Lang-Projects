package db

import (
	"errors"

	"github.com/datarohit/go-jwt-csrf-project/db/models"
	"github.com/datarohit/go-jwt-csrf-project/randomStrings"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]models.User{}
var refreshTokens map[string]string

func InitDB() {
	refreshTokens = make(map[string]string)
}

func StoreUser(username, password, role string) (uuid string, err error) {
	uuid, err = randomStrings.GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	for _, exists := users[uuid]; exists; _, exists = users[uuid] {
		uuid, err = randomStrings.GenerateRandomString(32)
		if err != nil {
			return "", err
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	users[uuid] = models.User{Username: username, PasswordHash: string(passwordHash), Role: role}

	return uuid, nil
}

func DeleteUser(uuid string) {
	delete(users, uuid)
}

func FetchUserById(uuid string) (models.User, error) {
	u, exists := users[uuid]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return u, nil
}

func FetchUserByUsername(username string) (models.User, string, error) {
	for k, v := range users {
		if v.Username == username {
			return v, k, nil
		}
	}
	return models.User{}, "", errors.New("user not found")
}

func StoreRefreshToken() (jti string, err error) {
	jti, err = randomStrings.GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	for exists := refreshTokens[jti] != ""; exists; exists = refreshTokens[jti] != "" {
		jti, err = randomStrings.GenerateRandomString(32)
		if err != nil {
			return "", err
		}
	}

	refreshTokens[jti] = "valid"
	return jti, nil
}

func DeleteRefreshToken(jti string) {
	delete(refreshTokens, jti)
}

func CheckRefreshToken(jti string) bool {
	return refreshTokens[jti] != ""
}

func LogUserIn(username, password string) (models.User, string, error) {
	user, uuid, err := FetchUserByUsername(username)
	if err != nil {
		return models.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return user, uuid, err
}
