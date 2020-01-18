package utils

import "time"

func DateDelta(delta int) time.Time {
	currenTime := time.Now()
	deltaTime := time.Date(currenTime.Year(), currenTime.Month(), currenTime.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, delta)
	return deltaTime
}


func DateFormat(t time.Time) string {
	return t.Format("2006-01-02")
}
