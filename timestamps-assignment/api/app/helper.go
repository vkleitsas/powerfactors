package app

import (
	"time"
)

func addInterval(interval string, timestamp time.Time) time.Time {

	switch interval {
	case "1h":
		timestamp = timestamp.Add(time.Hour)
	case "1d":
		timestamp = timestamp.AddDate(0, 0, 1)
	case "1mo":
		timestamp = timestamp.AddDate(0, 1, 0)
	case "1y":
		timestamp = timestamp.AddDate(1, 0, 0)
	}
	return timestamp
}

func truncateToHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

func lastDateOfMonth(t time.Time) time.Time {
	tsYear, tsMonth, _ := t.Date()
	firstOfMonth := time.Date(tsYear, tsMonth, 1, t.Hour(), 0, 0, 0, t.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return lastOfMonth
}

func lastDateOfYear(t time.Time) time.Time {
	tsYear, _, _ := t.Date()
	lastOfYear := time.Date(tsYear, 12, 31, t.Hour(), 0, 0, 0, t.Location())
	return lastOfYear
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
