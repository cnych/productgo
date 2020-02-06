package utils

import "time"

func DateDelta(current time.Time, delta int) time.Time {
	deltaTime := time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, delta)
	return deltaTime
}

func DateFormat(t time.Time) string {
	return t.Format("2006-01-02")
}

func Str2Date(str string) (time.Time, error) {
	layout := "2006-01-02"
	return time.ParseInLocation(layout, str, time.Local)
}
