package deduplication

import supportfuncs "server/server/convertingResponseToIntSlice/convertingFuncs/parameterValueToIntSlice/parameterValueToStringSlice/deduplication/deduplicationSupportFunctions"

func DeduplicateStringSlice(stringSlice []string) []string {
	set := supportfuncs.CreateSet(stringSlice)
	return set
}
