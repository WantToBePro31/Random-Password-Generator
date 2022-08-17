package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RandPass struct {
	Id             primitive.ObjectID `bson:"_id"`
	PasswordLength int                `bson:"password_length"`
	SpecialChar    int                `bson:"special_char"`
	Number         int                `bson:"number"`
	Uppercase      int                `bson:"uppercase"`
	Password       string             `bson:"password"`
}
