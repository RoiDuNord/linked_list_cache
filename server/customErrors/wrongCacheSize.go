package customerrors

type WrongSize struct {
	TextError string `json:"error"`
}

func (e *WrongSize) Error() string {
	return e.TextError
}

func WrongSizeError(err error) *WrongSize {
	return &WrongSize{TextError: err.Error()}
}
