package presign

type PresignedObjectRequest struct{
	OriginalName string `json:"originalName"`
	Type string `json:"type"`
}

type PresignedObjectResponse struct {
	URL string `json:"url"`
	BucketName string `json:"bucketName"`
	Key string `json:"key"`
}