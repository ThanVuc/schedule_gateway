package utils

import (
	"strconv"
	"time"
)

func Int32PtrToString(num *int32) string {
	if num == nil {
		return "nil"
	}
	return strconv.FormatInt(int64(*num), 10)
}

func ParseStringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StartAndEndOfDayTimestamp(t time.Time) (int64, int64) {
	startOfDay := time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0,
		t.Location(),
	)
	endOfDay := time.Date(
		t.Year(), t.Month(), t.Day(),
		23, 59, 59, 999999999,
		t.Location(),
	)
	return startOfDay.UnixMilli(), endOfDay.UnixMilli()
}

func SafeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func SafeInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func SafeInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}
