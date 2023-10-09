package errors

type ErrorInterface struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (instance *ErrorInterface) Error() string {
	return instance.Message
}
func Exception(status int, message string) *ErrorInterface {
	return &ErrorInterface{
		status,
		message,
	}
}

func BadRequestException(message string) *ErrorInterface {
	return Exception(400, message)
}

func NotFoundException(message string) *ErrorInterface {
	return Exception(404, message)
}

func ConflictException(message string) *ErrorInterface {
	return Exception(409, message)
}

func InternalException(message string) *ErrorInterface {
	return Exception(500, message)
}
