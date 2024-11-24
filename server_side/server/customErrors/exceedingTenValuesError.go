package customerrors

import "fmt"

type ExceedTenValues struct {
	StringSlice []string
}

func (e *ExceedTenValues) Error() string {
	return fmt.Sprintf("got %d elems, expected max 10", len(e.StringSlice))
}

func ExceedingTenValuesError(slice []string) *ExceedTenValues {
	return &ExceedTenValues{StringSlice: slice}
}
