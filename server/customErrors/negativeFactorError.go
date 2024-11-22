package customerrors

import "fmt"

type Factor struct {
	Size int
}

func (e *Factor) Error() string {
	return fmt.Sprintf("new factor %d, need positive", e.Size)
}

func NegativeFactorError(newSize int) *Factor {
	return &Factor{Size: newSize}
}

type WrongFactor struct {
	TextError string `json:"error"`
}

func (e *WrongFactor) Error() string {
	return e.TextError
}

func WrongFactorError(err error) *WrongFactor {
	return &WrongFactor{TextError: err.Error()}
}
