package presign

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Repository struct {
	presignClient	*s3.PresignClient 
}

func (repo *Repository)Init(ctx context.Context) (*Repository) {
	s3Client, err := repo.getS3Client(ctx)
	if err != nil {
		log.Printf("Error creating S3 client: %v", err)
		panic("Error connecting to S3")
	}
	repo.presignClient = s3.NewPresignClient(s3Client)
	return repo
}

func (repo *Repository) getS3Client(ctx context.Context) (*s3.Client, error){
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


	cfg, err := config.LoadDefaultConfig(ctx,
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
	ctx context.Context,
	objectKey string,
	expirationInMin int,
) (string, error) {

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
		return "", errors.New("Error generating presigned URL")
	}

	return presignResult.URL, nil
}
