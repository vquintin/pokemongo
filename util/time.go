package util

import "time"

func TimeInMilliseconds() int64 {
	return ConvertToMilliseconds(time.Now())
}

func ConvertToMilliseconds(date time.Time) int64 {
	return date.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
