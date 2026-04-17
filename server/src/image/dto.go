package image

import (
	"context"
)

type GetImageResponse struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	UserId    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateImageRequest struct {
	Key string `json:"key"`
}

type UpdateImageRequest struct {
	Key string `json:"key"`
}


func (getImageResponse *GetImageResponse) Set(ctx context.Context, image *Image) {
	getImageResponse.ID = image.Id.Hex()
	getImageResponse.Key = image.Key
	getImageResponse.UserId = image.UserId.Hex()
	getImageResponse.CreatedAt = image.CreatedAt.Time().String()
	getImageResponse.UpdatedAt = image.UpdatedAt.Time().String()
}
