package customerrors

import "fmt"

type NegativeCacheSize struct {
	Size int
}

func (e *NegativeCacheSize) Error() string {
	return fmt.Sprintf("new cache size (%d) can't be negative", e.Size)
}

func NewNegativeCacheSize(newSize int) *NegativeCacheSize {
	return &NegativeCacheSize{Size: newSize}
}
