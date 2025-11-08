package error

func ErrMapping(err error) bool {
	allErrors := make([]error, 8)
	allErrors = append(append(GeneralErrors[:]))
	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true

		}
	}
	return false
}
