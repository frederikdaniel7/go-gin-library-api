package exception

type ErrorType struct {
	StatusCode int
	Message    string
}

func NewErrorType(statusCode int, message string) *ErrorType {
	return &ErrorType{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e *ErrorType) Error() string {
	return e.Message
}
