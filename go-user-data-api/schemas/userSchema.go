package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}
