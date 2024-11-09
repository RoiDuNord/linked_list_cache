package customerrors

import "fmt"

type ExceedTenValues struct {
	stringSlice []string
}

func (e *ExceedTenValues) Error() string {
	return fmt.Sprintf("you request contains more values(%d) than maximum(10)", len(e.stringSlice))
}

func NewExceedTenValues(slice []string) *ExceedTenValues {
	return &ExceedTenValues{stringSlice: slice}
}
