package utils

type ErrorHandler struct {
	Message string
}

func (v *ErrorHandler) Error() string {
	return v.Message
}

