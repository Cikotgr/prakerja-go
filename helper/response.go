package helper

type ResponseMeta struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
type Response struct {
	Meta ResponseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}

func ResponseData(message string, status int, data interface{}) Response {
	return Response{
		Meta: ResponseMeta{
			Message: message,
			Status:  status,
		},
		Data: data,
	}
}
