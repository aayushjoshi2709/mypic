package group

import "github.com/gin-gonic/gin"

type GetGroupResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (getGroupResponse *GetGroupResponse) Set(ctx *gin.Context, group *Group) {
	getGroupResponse.Id = group.Id.String()
	getGroupResponse.Name = group.Name
}

type CreateGroupRequest struct {
	Name string `json:"name"`
}
