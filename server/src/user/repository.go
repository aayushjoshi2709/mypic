package user

import (
	"context"
	"time"

	"github.com/aayushjoshi2709/mypic/src/utils/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func (repo *Repository) getUserCollection() {
	if repo.collection == nil {
		repo.collection = db.GetConn().Collection("users")
	}
}

func (repo *Repository) createIndexes() {
	collection := repo.collection
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("username_unique_index"),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}
}

func (repo *Repository) Init() {
	repo.getUserCollection()
	repo.createIndexes()
}

func (repository *Repository) GetById(ctx context.Context, id string) (*User, error) {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = repository.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (repository *Repository) Add(ctx context.Context, name, username, password string) (*User, error) {
	user := User{}
	user.Id = bson.NewObjectID()
	user.Name = name
	user.Username = username
	user.Password = password
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err := repository.collection.InsertOne(ctx, &user)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (repository *Repository) Update(ctx context.Context, id string, name, username string) (*User, error) {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	updatedFields := bson.M{}

	if name != "" {
		updatedFields["name"] = name
	}

	if username != "" {
		updatedFields["username"] = username
	}

	if len(updatedFields) == 0 {
		return nil, nil
	}

	update := bson.M{
		"$set": updatedFields,
	}

	update["updatedAt"] = primitive.NewDateTimeFromTime(time.Now())

	user := &User{}

	err = repository.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectId},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (repository *Repository) Delete(ctx context.Context, id string) error {
	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = repository.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}
