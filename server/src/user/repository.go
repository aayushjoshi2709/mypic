package user

import (
	"context"

	"github.com/aayushjoshi2709/mypic/src/utils/db"
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

func (repository *Repository) GetById(id bson.ObjectID) (*User, error) {
	user := &User{}
	err := repository.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (repository *Repository) Add(name, username, password string) (*User, error) {
	user := User{}
	user.Id = bson.NewObjectID()
	user.Name = name
	user.Username = username
	user.Password = password

	_, err := repository.collection.InsertOne(context.TODO(), &user)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (repository *Repository) Update(id string, name, username string) (*User, error) {
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

	update := bson.M{
		"$set": updatedFields,
	}

	_, err = repository.collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = repository.collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (repository *Repository) Delete(id string) error {
	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = repository.collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	return err
}
