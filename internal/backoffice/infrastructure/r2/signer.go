package r2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Signer struct {
	bucket string
	client *s3.PresignClient
}

func NewSigner(bucket string, client *s3.PresignClient) *Signer {
	return &Signer{
		bucket: bucket,
		client: client,
	}
}

func (s *Signer) SignUpload(
	ctx context.Context,
	objectKey string,
	contentType string,
	metadata map[string]string,
	ttl time.Duration,
) (string, error) {

	input := &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         &objectKey,
		ContentType: &contentType,
		Metadata:    metadata,
	}

	req, err := s.client.PresignPutObject(
		ctx,
		input,
		s3.WithPresignExpires(ttl),
	)
	if err != nil {
		return "", err
	}

	return req.URL, nil
}
