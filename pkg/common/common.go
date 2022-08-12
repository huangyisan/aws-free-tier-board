package common

import (
	"time"
)

const (
	layoutISO = "2006-01-02"
)

func FormatDateToRFC3339(date string) string {
	t, _ := time.Parse(layoutISO, date)
	nt := t.Format(time.RFC3339)
	return nt
}

func GetMonthFirstDay() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	// 2019-08-28T22:00:00Z
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	return firstOfMonth.Format(time.RFC3339)

}

func GetCurrentDay() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}
