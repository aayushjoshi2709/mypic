package image

import (
	"github.com/aayushjoshi2709/mypic/src/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserName  string             `bson:"username,omitempty"`
	URL       string             `bson:"url,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
	User      *user.User         `bson:"user,omitempty"`
}
