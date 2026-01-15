package uploadurlgetter

import (
	"context"
	"time"
)

type UploadURLSigner interface {
	SignUpload(
		ctx context.Context,
		objectKey string,
		contentType string,
		metadata map[string]string,
		ttl time.Duration,
	) (string, error)
}
