package image

import (
	"github.com/aayushjoshi2709/mypic/src/user"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Image struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	URL       string             `bson:"url,omitempty"`
	CreatedAt bson.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt bson.DateTime `bson:"updatedAt,omitempty"`
	User      *user.User         `bson:"user,omitempty"`
}
