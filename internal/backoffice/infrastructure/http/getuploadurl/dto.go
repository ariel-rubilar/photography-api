package getuploadurl

type Response struct {
	UploadURL string `json:"uploadUrl"`
	ObjectKey string `json:"objectKey"`
}
