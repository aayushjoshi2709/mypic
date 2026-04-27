package group

import (
	"time"

	"github.com/aayushjoshi2709/mypic/src/utils/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func (repository *Repository) getGroupCollection() {
	if repository.collection == nil {
		repository.collection = db.GetConn().Collection("group")
	}
}

func (repository *Repository) createIndexes() {

}

func (repository *Repository) Init() {
	repository.getGroupCollection()
	repository.createIndexes()
}

func (repository *Repository) GetById(ctx *gin.Context, id string) (*Group, error) {
	userId, _ := ctx.Get("userId")

	ObjectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	group := &Group{}

	err = repository.collection.FindOne(ctx, bson.M{
		"id":      ObjectId,
		"userIds": userId,
	},
		options.
			FindOne().
			SetProjection(
				bson.M{
					"_id":  1,
					"name": 1,
				},
			),
	).Decode(group)

	if err != nil {
		return nil, err
	}
	return group, nil
}

func (repository *Repository) GetAll(ctx *gin.Context, page, limit int) ([]Group, error) {
	userId, _ := ctx.Get("userId")
	cursor, err := repository.collection.Find(
		ctx,
		bson.M{
			"userIds": userId,
		},
		options.Find().
			SetProjection(bson.M{
				"_id":  1,
				"name": 1,
			}).
			SetSort(bson.M{"created_at": 1}).
			SetSkip(int64(limit*(page-1))).
			SetLimit(int64(limit)),
	)

	if err != nil {
		return nil, err
	}

	images := []Group{}
	err = cursor.All(ctx, &images)
	return images, err
}

func (repository *Repository) Add(ctx *gin.Context, name string) error {
	userId, _ := ctx.Get("userId")
	
	group := Group{
		Id:        bson.NewObjectID(),
		Name:      name,
		createdBy: userId.(bson.ObjectID),
		CreatedAt: bson.NewDateTimeFromTime(time.Now()),
		UpdatedAt: bson.NewDateTimeFromTime(time.Now()),
	}

	_, err := repository.collection.InsertOne(ctx, group)
	return err
}

func (repository *Repository) Update() (Group, error) {
	return Group{}, nil
}

func (repository *Repository) Delete(ctx *gin.Context, id string) error {
	userId, _ := ctx.Get("userId")
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = repository.collection.DeleteOne(ctx, bson.M{
		"_id":       objectId,
		"createdBy": userId,
	})
	return err
}
