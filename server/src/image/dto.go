package image

type GetImageResponse struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	User      string `json:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateImageRequest struct {
	URL string `json:"url"`
}

type UpdateImageRequest struct {
	URL string `json:"url"`
}

func (getImageResponse *GetImageResponse) Set(image *Image) {
	getImageResponse.ID = image.ID.Hex()
	getImageResponse.URL = image.URL
	getImageResponse.User = image.User.Username
	getImageResponse.CreatedAt = image.CreatedAt.Time().String()
	getImageResponse.UpdatedAt = image.UpdatedAt.Time().String()
}
