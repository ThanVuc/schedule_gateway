package utils

func Int32PtrToString(num *int32) string {
	if num == nil {
		return "nil"
	}
	return string(*num)
}
