package user

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	Name     string        `bson:"name,omitempty"`
	Username string        `bson:"username,omitempty,unique"`
	Password string        `bson:"password,omitempty"`
}
