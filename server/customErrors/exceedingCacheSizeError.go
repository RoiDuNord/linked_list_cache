package customerrors

import "fmt"

type NewSizeError struct {
	CurSize, NewSize int
}

func (e *NewSizeError) Error() string {
	return fmt.Sprintf("new cache size %d, max size %d", e.NewSize, e.CurSize)
}

func ExceedingCacheSizeError(curSize, newSize int) *NewSizeError {
	return &NewSizeError{
		CurSize: curSize,
		NewSize: newSize,
	}
}
