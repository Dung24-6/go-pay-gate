package aws

import (
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/Dung24-6/go-pay-gate/internal/config"
)

type AWSClients struct {
	S3Client  *s3.Client
	SQSClient *sqs.Client
}

// NewAWSClients creates new AWS service clients
func NewAWSClients(cfg *config.AWSConfig) (*AWSClients, error) {
	// Use awsconfig instead of config to avoid naming conflict
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(cfg.Region),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKeyID,
			cfg.SecretAccessKey,
			cfg.SessionToken,
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &AWSClients{
		S3Client:  s3.NewFromConfig(awsCfg),
		SQSClient: sqs.NewFromConfig(awsCfg),
	}, nil
}
