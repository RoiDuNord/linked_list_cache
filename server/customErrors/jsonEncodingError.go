package customerrors

type JSONHandlerError struct {
	ErrorText string `json:"error"`
}

func (e *JSONHandlerError) Error() string {
	return e.ErrorText
}

func JSONEncodingError(err error) *JSONHandlerError {
	return &JSONHandlerError{ErrorText: err.Error()}
}
