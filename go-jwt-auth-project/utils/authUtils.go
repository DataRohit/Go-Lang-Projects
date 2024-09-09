package utils

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "email or password is incorrect"
		check = false
	}

	return check, msg
}

func CheckUserType(ginCtx *gin.Context, role string) (err error) {
	userType := ginCtx.GetString("userType")
	err = nil

	if userType != role {
		err = errors.New("access to this resource is restricted")
		return err
	}

	return err
}

func MatchUserTypeToUserId(ginCtx *gin.Context, userId string) (err error) {
	userType := ginCtx.GetString("userType")
	uid := ginCtx.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("access to this resource is restricted")
		return err
	}

	err = CheckUserType(ginCtx, userType)
	return err
}
