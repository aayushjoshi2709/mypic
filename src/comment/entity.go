package comment

import (
	"github.com/aayushjoshi2709/mypic/src/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	User      *user.User         `bson:"User,omitempty"`
	Parent    *Comment           `bson:"parentId,omitempty"`
	Content   string             `bson:"content,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
}

