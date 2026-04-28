package group

import "github.com/gin-gonic/gin"

type GetGroupResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func (getGroupResponse *GetGroupResponse) Set(ctx *gin.Context, cloudFrontUrl string, group *Group) {
	getGroupResponse.Id = group.Id.Hex()
	getGroupResponse.Name = group.Name
	getGroupResponse.ImageUrl = cloudFrontUrl + group.ImageKey
}

type CreateGroupRequest struct {
	Name string `json:"name"`
	ImageKey string `json:"imageKey"`
}

type AddImageRequest struct {
	ImageId string `json:"imageId"`
}

type AddUserRequest struct {
	UserId string `json:"userId"`
}
