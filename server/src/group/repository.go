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

func (repository *Repository) GetById(ctx *gin.Context, groupId string) (*Group, error) {
	userId, _ := ctx.Get("userId")

	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	group := &Group{}

	err = repository.collection.FindOne(ctx, bson.M{
		"id":      groupIdBson,
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
		UserIds:   make([]bson.ObjectID, 0),
		ImageIds:  make([]bson.ObjectID, 0),
		CreatedAt: bson.NewDateTimeFromTime(time.Now()),
		UpdatedAt: bson.NewDateTimeFromTime(time.Now()),
	}

	_, err := repository.collection.InsertOne(ctx, group)
	return err
}

func (repository *Repository) Update() (Group, error) {
	return Group{}, nil
}

func (repository *Repository) Delete(ctx *gin.Context, groupId string) error {
	userId, _ := ctx.Get("userId")
	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return err
	}

	_, err = repository.collection.DeleteOne(ctx, bson.M{
		"_id":       groupIdBson,
		"createdBy": userId,
	})
	return err
}

func (repository *Repository) AddImage(ctx *gin.Context, groupId string, imageId bson.ObjectID) error {
	userId, _ := ctx.Get("userId")
	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return err
	}

	_, err = repository.collection.UpdateOne(ctx,
		bson.M{
			"_id":       groupIdBson,
			"createdBy": userId,
		},
		bson.M{
			"$push": bson.M{
				"imageIds": imageId,
			},
		},
	)
	return err
}

func (repository *Repository) AddUser(ctx *gin.Context, groupId string, userIdToAdd bson.ObjectID) error {
	userId, _ := ctx.Get("userId")
	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return err
	}

	_, err = repository.collection.UpdateOne(ctx,
		bson.M{
			"_id":       groupIdBson,
			"createdBy": userId,
		},
		bson.M{
			"$push": bson.M{
				"userIds": userIdToAdd,
			},
		},
	)
	return err
}


func (repository *Repository) RemoveImage(ctx *gin.Context, groupId string, imageId string) error {
	userId, _ := ctx.Get("userId")
	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return err
	}


	imageIdBson, err := bson.ObjectIDFromHex(imageId)
	if err != nil {
		return err
	}



	_, err = repository.collection.UpdateOne(ctx,
		bson.M{
			"_id":       groupIdBson,
			"createdBy": userId,
		},
		bson.M{
			"$pull": bson.M{
				"imageIds": imageIdBson,
			},
		},
	)
	return err
}

func (repository *Repository) RemoveUser(ctx *gin.Context, groupId string, userIdToRemove string) error {
	userId, _ := ctx.Get("userId")
	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return err
	}

	userIdToRemoveBson, err := bson.ObjectIDFromHex(userIdToRemove)
	if err != nil {
		return err
	}

	_, err = repository.collection.UpdateOne(ctx,
		bson.M{
			"_id":       groupIdBson,
			"createdBy": userId,
		},
		bson.M{
			"$pull": bson.M{
				"userIds": userIdToRemoveBson,
			},
		},
	)
	return err
}

func (repository *Repository) GetImageIds(ctx *gin.Context, groupId string, pageInt, limitInt int) ([]bson.ObjectID, error) {
	userId, _ := ctx.Get("userId")
	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	var groupObj struct {
		ImageIds []bson.ObjectID
	}
	err = repository.collection.FindOne(
		ctx,
		bson.M{
			"_id":     groupIdBson,
			"userIds": userId,
		},
		options.FindOne().SetProjection(bson.M{
			"imageIds": bson.M{
				"$slice": []int{
					(pageInt - 1) * limitInt,
					limitInt,
				},
			},
		}),
	).Decode(&groupObj)

	return groupObj.ImageIds, nil
}

func (repository *Repository) GetUserIds(ctx *gin.Context, groupId string, pageInt, limitInt int) ([]bson.ObjectID, error) {
	userId, _ := ctx.Get("userId")
	groupIdBson, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	var groupObj struct {
		UserIds []bson.ObjectID
	}
	err = repository.collection.FindOne(
		ctx,
		bson.M{
			"_id":     groupIdBson,
			"userIds": userId,
		},
		options.FindOne().SetProjection(bson.M{
			"userIds": bson.M{
				"$slice": []int{
					(pageInt - 1) * limitInt,
					limitInt,
				},
			},
		}),
	).Decode(&groupObj)

	return groupObj.UserIds, nil
}
