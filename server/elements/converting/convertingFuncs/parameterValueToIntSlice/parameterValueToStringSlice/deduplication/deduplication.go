package deduplication

// самому не стращно от длины это залупы?)
// ещё и дедупликация какая-то, упаси господи
import supportfuncs "workWithCache/server/elements/converting/convertingFuncs/parameterValueToIntSlice/parameterValueToStringSlice/deduplication/deduplicationSupportFunctions"

func DeduplicateStringSlice(stringSlice []string) []string {
	set := supportfuncs.CreateSet(stringSlice)
	return set
}
