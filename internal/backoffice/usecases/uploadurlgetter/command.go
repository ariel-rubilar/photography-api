package uploadurlgetter

import "time"

type GetUploadURLCommand struct {
	ContentType string
	Extension   string
	TakenAt     *time.Time
}
