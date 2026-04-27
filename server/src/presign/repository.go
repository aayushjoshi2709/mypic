package presign

import (
	"errors"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type Repository struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
}

func (repo *Repository) Init(ctx *gin.Context) *Repository {
	client, err := repo.getS3Client(ctx)
	if err != nil {
		log.Printf("Error creating S3 client: %v", err)
		panic("Error connecting to S3")
	}
	repo.client = client
	repo.presignClient = s3.NewPresignClient(client)

	repo.bucketName = os.Getenv("AWS_S3_BUCKET_NAME")

	if repo.bucketName == "" {
		slog.Error("Error loading aws bucket name")
		panic("Error loading aws bucket name")
	}

	return repo
}

func (repo *Repository) getS3Client(ctx *gin.Context) (*s3.Client, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")

	if accessKey == "" {
		slog.Error("Error loading aws access key")
		panic("Error loading aws access key")
	}

	secretKey := os.Getenv("AWS_SECRET_KEY")

	if secretKey == "" {
		slog.Error("Error loading aws secret key")
		panic("Error loading aws secret key")
	}

	region := os.Getenv("AWS_REGION")

	if region == "" {
		slog.Error("Error loading aws region")
		panic("Error loading aws region")
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)
	if err != nil {
		log.Printf("Error loading AWS config: %v", err)
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func (repo *Repository) PutObject(
	ctx *gin.Context,
	objectKey string,
	expirationInMin int,
) (*PresignedObjectResponse, error) {

	bucketName := os.Getenv("AWS_S3_BUCKET_NAME")

	if bucketName == "" {
		slog.Error("Error loading aws bucket name")
		panic("Error loading aws bucket name")
	}

	presignParams := &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}

	presignOpts := func(o *s3.PresignOptions) {
		o.Expires = time.Duration(expirationInMin) * time.Minute
	}
	presignResult, err := repo.presignClient.PresignPutObject(ctx, presignParams, presignOpts)
	if err != nil {
		log.Printf("Error generating presigned URL: %v", err)
		return nil, errors.New("Error generating presigned URL")
	}

	return &PresignedObjectResponse{
		URL: presignResult.URL,
		Key: objectKey,
	}, nil
}

func (repository *Repository) GetObjectStream(ctx *gin.Context, key string) (*s3.GetObjectOutput, error) {
	out, err := repository.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &repository.bucketName,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}
