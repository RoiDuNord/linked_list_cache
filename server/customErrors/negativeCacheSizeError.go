package customerrors

import "fmt"

type CacheSize struct {
	Size int
}

func (e *CacheSize) Error() string {
	return fmt.Sprintf("new cache size %d, need positive", e.Size)
}

func NegativeCacheSizeError(newSize int) *CacheSize {
	return &CacheSize{Size: newSize}
}
