package apiError

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func InitErr() *APIError {
	return &APIError{}
}
