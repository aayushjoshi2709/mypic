package image

import (
	"context"

	"github.com/aayushjoshi2709/mypic/src/presign"
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
	getImageResponse.Key, _ = presign.GeneratePublicUrl(ctx, image.Key)
	getImageResponse.UserId = image.UserId.Hex()
	getImageResponse.CreatedAt = image.CreatedAt.Time().String()
	getImageResponse.UpdatedAt = image.UpdatedAt.Time().String()
}
