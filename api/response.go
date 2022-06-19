package api

func NewGenericResponse(stsCd, isError int, messages []string, data interface{}) *GenericResponse {

	return &GenericResponse{
		Status:   stsCd,
		Success:  isError == 0,
		Messages: messages,
		Data:     data,
	}
}

type GenericResponse struct {
	Status   int         `json:"status"`
	Success  bool        `json:"success"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}
