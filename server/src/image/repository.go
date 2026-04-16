package image

import (
	"context"
	"time"

	"github.com/aayushjoshi2709/mypic/src/user"
	"github.com/aayushjoshi2709/mypic/src/utils/db"
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

func (repository *Repository) GetById(ctx context.Context, id string) (*Image, error) {
	ObjectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	image := &Image{}
	err = repository.collection.FindOne(ctx, bson.M{"_id": ObjectId}).Decode(image)

	if err != nil {
		return nil, err
	}

	return image, err
}

func (repository *Repository) GetAll(ctx context.Context, page, limit int) ([]Image, error) {
	cursor, err := repository.collection.Find(
		ctx,
		bson.M{},
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

func (repository *Repository) Add(ctx context.Context, url string, user *user.User) (*Image, error) {
	image := &Image{
		Id:        bson.NewObjectID(),
		URL:       url,
		User:      user,
		CreatedAt: bson.NewDateTimeFromTime(time.Now()),
		UpdatedAt: bson.NewDateTimeFromTime(time.Now()),
	}

	_, err := repository.collection.InsertOne(ctx, image)
	if err != nil {
		return nil, err
	}

	return image, err
}

func (repository *Repository) Update(ctx context.Context, id string, url string) (*Image, error) {
	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	updateFields := bson.M{}

	if url != "" {
		updateFields["url"] = url
	}

	if len(updateFields) == 0 {
		return nil, nil
	}

	updateFields["UpdatedAt"] = bson.NewDateTimeFromTime(time.Now())

	image := &Image{}

	err = repository.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectId},
		updateFields,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(image)

	if err != nil {
		return nil, err
	}

	return image, err
}

func (repository *Repository) Delete(ctx context.Context, id string) error {
	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = repository.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}
