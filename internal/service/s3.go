package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct {
	client *s3.Client
	bucket string
}

func NewS3Service() (*S3Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	bucket := os.Getenv("S3_BUCKET_NAME")

	return &S3Service{client: s3Client, bucket: bucket}, nil
}

func (s *S3Service) UploadReceipt(transactionID string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fmt.Sprintf("receipts/%s.json", transactionID)),
		Body:   bytes.NewReader(jsonData),
	})

	if err != nil {
		return fmt.Errorf("failed to upload receipt: %w", err)
	}

	log.Printf("Receipt uploaded to S3: receipts/%s.json", transactionID)
	return nil
}
