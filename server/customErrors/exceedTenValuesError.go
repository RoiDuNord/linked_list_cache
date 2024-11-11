package customerrors

import "fmt"

type ExceedTenValues struct {
	stringSlice []string
}

// здесь молодец
func (e *ExceedTenValues) Error() string {
	return fmt.Sprintf("got %d elems, expected max 10", len(e.stringSlice))
}

func NewExceedTenValues(slice []string) *ExceedTenValues {
	return &ExceedTenValues{stringSlice: slice}
}
