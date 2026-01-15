package r2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	AccountID       string
}

func NewClient(params Config) (*s3.PresignClient, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				params.AccessKeyID,
				params.SecretAccessKey,
				"",
			),
		),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(
		"https://%s.r2.cloudflarestorage.com",
		params.AccountID,
	)

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})

	presignClient := s3.NewPresignClient(s3Client)

	return presignClient, nil

}
