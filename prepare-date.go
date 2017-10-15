package main

import (
	"time"
)

func prepareDate(dayOfWeek string, hour string) (time.Time) {
	date := time.Now()
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, 1)
	}

	for date.Weekday().String() != dayOfWeek {
		date = date.AddDate(0, 0, 1)
	}

	datestr := date.Format("2006-01-02") + "T" + hour + ":00"
	date, _ = time.Parse("2006-01-02T15:04:05", datestr)

	return date
}