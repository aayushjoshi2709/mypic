package image

import (
	"log/slog"
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

func (repository *Repository) getImageCollection() {
	if repository.collection == nil {
		repository.collection = db.GetConn().Collection("images")
	}
}

func (repository *Repository) Init() {
	repository.getImageCollection()
}

func (repository *Repository) GetById(ctx *gin.Context, id string) (*Image, error) {
	userId, _ := ctx.Get("userId")
	ObjectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	image := &Image{}
	err = repository.collection.FindOne(ctx, bson.M{
		"_id":    ObjectId,
		"userId": userId,
	}).Decode(image)

	if err != nil {
		return nil, err
	}

	return image, err
}

func (repository *Repository) GetAll(ctx *gin.Context, page, limit int) ([]Image, error) {
	userId, _ := ctx.Get("userId")

	slog.Info("Getting all images for user", "userId", userId.(bson.ObjectID).String(), "page", page, "limit", limit)
	cursor, err := repository.collection.Find(
		ctx,
		bson.M{"userId": userId},
		options.
			Find().
			SetSort(bson.M{"created_at": -1}).
			SetSkip(int64(limit*(page-1))).
			SetLimit(int64(limit)),
	)

	if err != nil {
		return nil, err
	}

	images := []Image{}
	err = cursor.All(ctx, &images)
	return images, err
}

func (repository *Repository) Add(ctx *gin.Context, key string, originalName string) error {
	userId, _ := ctx.Get("userId")
	slog.Info("Adding image with key", "key", key, "userId", userId.(bson.ObjectID).String())
	image := &Image{
		Id:           bson.NewObjectID(),
		OriginalName: originalName,
		Key:          key,
		UserId:       userId.(bson.ObjectID),
		CreatedAt:    bson.NewDateTimeFromTime(time.Now()),
		UpdatedAt:    bson.NewDateTimeFromTime(time.Now()),
	}

	_, err := repository.collection.InsertOne(ctx, image)
	return err
}

func (repository *Repository) Update(ctx *gin.Context, id string, key string, originalName string) (*Image, error) {
	objectId, err := bson.ObjectIDFromHex(id)
	userId, _ := ctx.Get("userId")

	if err != nil {
		return nil, err
	}

	updateFields := bson.M{}

	if key != "" {
		updateFields["key"] = key
	}

	if originalName != "" {
		updateFields["originalName"] = originalName
	}

	if len(updateFields) == 0 {
		return nil, nil
	}

	updateFields["UpdatedAt"] = bson.NewDateTimeFromTime(time.Now())

	image := &Image{}

	err = repository.collection.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id":    objectId,
			"userId": userId,
		},
		updateFields,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(image)

	if err != nil {
		return nil, err
	}

	return image, err
}

func (repository *Repository) Delete(ctx *gin.Context, id string) error {
	userId, _ := ctx.Get("userId")
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = repository.collection.DeleteOne(ctx, bson.M{
		"_id":    objectId,
		"userId": userId,
	})
	return err
}
