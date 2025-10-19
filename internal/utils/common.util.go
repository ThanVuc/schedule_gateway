package utils

import "strconv"

func Int32PtrToString(num *int32) string {
	if num == nil {
		return "nil"
	}
	return strconv.FormatInt(int64(*num), 10)
}
