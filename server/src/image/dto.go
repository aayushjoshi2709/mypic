package image

import "github.com/gin-gonic/gin"

type GetImageResponse struct {
	ID           string `json:"id"`
	Url          string `json:"url"`
	Key          string `json:"key"`
	OriginalName string `json:"originalName"`
	UserId       string `json:"userId"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type CreateImageRequest struct {
	Key          string `json:"key"`
	OriginalName string `json:"originalName"`
}

type UpdateImageRequest struct {
	Key          string `json:"key"`
	OriginalName string `json:"originalName"`
}

func (getImageResponse *GetImageResponse) Set(ctx *gin.Context, cloudFrontUrl string, image *Image) {
	getImageResponse.ID = image.Id.Hex()
	getImageResponse.Url = cloudFrontUrl + image.Key
	getImageResponse.UserId = image.UserId.Hex()
	getImageResponse.OriginalName = image.OriginalName
	getImageResponse.Key = image.Key
	getImageResponse.CreatedAt = image.CreatedAt.Time().String()
	getImageResponse.UpdatedAt = image.UpdatedAt.Time().String()
}
