package supportfuncs

func CreateSet(stringSlice []string) []string {
	stringSet := make(map[string]struct{})
	for _, el := range stringSlice {
		stringSet[el] = struct{}{}
	}

	deduplicatedSlice := make([]string, 0, len(stringSet))
	for keys := range stringSet {
		deduplicatedSlice = append(deduplicatedSlice, keys)
	}
	return deduplicatedSlice
}
