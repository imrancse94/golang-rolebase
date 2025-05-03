package services

type ServiceResponse struct {
	StatusCode int
	HTTPCode   int
	Message    string
	Data       any
}

func NewServiceResponse(statusCode int, httpCode int, message string, data any) ServiceResponse {
	if data == nil {
		data = []any{}
	}
	return ServiceResponse{
		StatusCode: statusCode,
		HTTPCode:   httpCode,
		Message:    message,
		Data:       data,
	}
}
