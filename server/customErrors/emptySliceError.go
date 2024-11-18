package customerrors

func EmptySliceError(slice []int, parameterValue string) error {
	if len(slice) == 0 {
		err := ZeroLenError(parameterValue)
		return err
	}
	return nil
}
