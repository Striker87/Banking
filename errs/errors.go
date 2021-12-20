package errs

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    404,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    500,
	}
}
