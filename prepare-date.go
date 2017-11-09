package main

import (
	"time"
	"strings"
	"strconv"
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

func nextDate(dayName string, hourString string) (time.Time) {
	hourMinute := strings.Split(hourString, ":")
	hour, _ := strconv.Atoi(hourMinute[0])
	minute, _ := strconv.Atoi(hourMinute[1])

	date := time.Now()
	date = date.AddDate(0, 0, 6)

	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, 1)
	}

	for date.Weekday().String() != dayOfWeek {
		date = date.AddDate(0, 0, 1)
	}

	location, _ := time.LoadLocation("Europe/Madrid")
	date = time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, location);

	return date;
}