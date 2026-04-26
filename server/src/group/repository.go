package group

import (
	"github.com/aayushjoshi2709/mypic/src/utils/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func (repository *Repository) GetById(ctx *gin.Context, id string) (Group, error) {
	return Group{}, nil
}

func (repository *Repository) getAll(ctx *gin.Context, userId bson.ObjectID, page, size int) ([]Group, error) {
	return []Group{}, nil
}

func (repository *Repository) add(ctx *gin.Context) (Group, error) {
	return Group{}, nil
}

func (repository *Repository) update() (Group, error) {
	return Group{}, nil
}

func (repository *Repository) delete(ctx *gin.Context) {

}
