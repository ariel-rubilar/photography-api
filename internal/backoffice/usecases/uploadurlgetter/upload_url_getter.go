package uploadurlgetter

import (
	"context"
	"fmt"
	"time"

	"github.com/ariel-rubilar/photography-api/internal/shared/aplication/clock"
	"github.com/google/uuid"
)

type Getter struct {
	signer        UploadURLSigner
	clock         clock.Clock
	publicBaseURL string
}

func New(publicBaseURL string, signer UploadURLSigner, clock clock.Clock) *Getter {
	return &Getter{
		signer:        signer,
		clock:         clock,
		publicBaseURL: publicBaseURL,
	}
}

func (uc *Getter) Execute(
	ctx context.Context,
	cmd GetUploadURLCommand,
) (Response, error) {

	now := uc.clock.Now().UTC()

	objectKey := fmt.Sprintf(
		"photos/%04d/%02d/%s%s",
		now.Year(),
		now.Month(),
		uuid.NewString(),
		cmd.Extension,
	)

	metadata := map[string]string{}
	if cmd.TakenAt != nil {
		metadata["taken-at"] = cmd.TakenAt.Format(time.RFC3339)
	}

	url, err := uc.signer.SignUpload(
		ctx,
		objectKey,
		cmd.ContentType,
		metadata,
		10*time.Minute,
	)
	if err != nil {
		return Response{}, err
	}

	return Response{
		UploadURL: url,
		ObjectKey: objectKey,
		PublicURL: uc.publicBaseURL + "/" + objectKey,
	}, nil
}
