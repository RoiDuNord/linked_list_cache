package customerrors

import "fmt"

type URLWithoutNumbers struct {
	Parameter string `json:"error"`
}

func (e *URLWithoutNumbers) Error() string {
	return fmt.Sprintf("this URL doesn't contain requested parameter '%s'", e.Parameter)
}

func URLWithoutParameterNumbersError(parameter string) *URLWithoutNumbers {
	return &URLWithoutNumbers{Parameter: parameter}
}
