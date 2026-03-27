package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id        bson.ObjectID      `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Username  string             `bson:"username,omitempty,unique"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
}
