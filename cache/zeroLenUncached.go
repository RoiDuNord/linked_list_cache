package cache

import "fmt"

type ZeroUncachedSize struct {
	Request []int `json:"cache"`
}

func (e *ZeroUncachedSize) Error() string {
	return fmt.Sprintf("requested values %d fully cached", e.Request)
}

func ZeroLenUncached(inputNumbers []int) *ZeroUncachedSize {
	return &ZeroUncachedSize{Request: inputNumbers}
}
