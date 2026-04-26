package group

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Group struct {
	Id bson.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	UserIds []bson.ObjectID `bson:"userIds,omitempty"`
	ImageIds []bson.ObjectID `bson:"imageIds,omitempty"`
	createdBy bson.ObjectID `bson:"createdBy,omitempty"`
	CreatedAt bson.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt bson.DateTime `bson:"updatedAt,omitempty"`
}