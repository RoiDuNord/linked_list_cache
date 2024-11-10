package intslicetostringslice

import "strconv"

func IntSliceToStringSlice(ints []int) []string {
	stringSlice := make([]string, len(ints))
	for i, v := range ints {
		stringSlice[i] = strconv.Itoa(v)
	}
	return stringSlice
}
