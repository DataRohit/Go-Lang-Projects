package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       string             `json:"userId"`
	FirstName    *string            `json:"firstName" validate:"required,min=2,max=100"`
	LastName     *string            `json:"lastName" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"email,required"`
	Phone        *string            `json:"phone" validate:"required"`
	UserType     *string            `json:"userType" validate:"required,eq=ADMIN|eq=USER"`
	AccessToken  *string            `json:"accessToken"`
	RefreshToken *string            `json:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
}
