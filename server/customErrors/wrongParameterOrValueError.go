package customerrors

type WrongPair struct {
	ErrorText string `json:"error"`
}

func (e *WrongPair) Error() string {
	return e.ErrorText
}

func ParameterIsNotNumbersOrEmptyValueError(err error) *WrongPair {
	return &WrongPair{ErrorText: err.Error()}
}
