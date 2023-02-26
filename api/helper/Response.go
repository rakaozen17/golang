package apiError

type APIMessage struct {
	ResponseCode    int    `json:"code"`
	ResponseMessage string `json:"message"`
	Data            interface{}
}

func InitRes() *APIMessage {
	return &APIMessage{}
}
