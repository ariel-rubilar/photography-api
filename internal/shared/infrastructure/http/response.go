package http

type SuccessResponse struct {
	Data any `json:"data"`
	Meta any `json:"meta,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func NewSuccessResponse(data any, meta any) *SuccessResponse {
	return &SuccessResponse{
		Data: data,
		Meta: meta,
	}
}

func NewErrorResponse(code string, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: Error{
			Code:    code,
			Message: message,
		},
	}
}
