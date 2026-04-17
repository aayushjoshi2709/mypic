package image

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Image struct {
	Id        bson.ObjectID      `bson:"_id,omitempty"`
	Key       string             `bson:"key,omitempty"`
	CreatedAt bson.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt bson.DateTime `bson:"updatedAt,omitempty"`
	UserId    bson.ObjectID  `bson:"userId,omitempty"`
}
