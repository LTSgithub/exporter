package time

import "time"

func GetDateString() string {
	return time.Now().Format("2006-01-02")
}

func GetDateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}