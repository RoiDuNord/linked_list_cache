package deduplication

import supportfuncs "workWithCache/server/elements/converting/convertingFuncs/parameterValueToIntSlice/parameterValueToStringSlice/deduplication/deduplicationSupportFunctions"

func DeduplicateStringSlice(stringSlice []string) []string {
	set := supportfuncs.CreateSet(stringSlice)
	return set
}
