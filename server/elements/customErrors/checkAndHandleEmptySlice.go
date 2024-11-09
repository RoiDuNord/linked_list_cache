package customerrors

func CheckAndHandleEmptySlice(slice []int, parameterValue string) error {
	if len(slice) == 0 {
		err := NewZeroLenError(parameterValue)
		return err
	}
	return nil
}
