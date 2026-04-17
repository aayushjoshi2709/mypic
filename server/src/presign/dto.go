package presign

type PresignedObjectRequest struct{
	OriginalName string `json:"originalName" validate:"required"`
	Type string `json:"type" validate:"required" enums:"video,image"`
}

type PresignedObjectResponse struct {
	URL string `json:"url"`
	Key string `json:"key"`
}