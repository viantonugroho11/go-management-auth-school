package time

import (
	"time"
)

const (
	JAKARTA_TIME_LOCATION = "Asia/Jakarta"
	DATE_FORMAT           = "2006-01-02"
	DATE_FORMAT_DDMMMMYYYY 	= "02 January 2006"
	DATE_FORMAT_HHMM = "15:04"
)

// In ...
func InString(t, format, locName string) (time.Time, error) {
	date, err := time.Parse(format, t)
	if err != nil {
		return date, err
	}

	loc, err := time.LoadLocation(locName)
	if err == nil {
		date = date.In(loc)
	}
	return date, err
}

// In ...
func In(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

// Convert ...
func Convert(t, fromFormat, toFormat string) string {
	timeConvert, err := time.Parse(fromFormat, t)
	if err != nil {
		return ""
	}

	return timeConvert.Format(toFormat)
}

// ConvertWithTimezone ...
func ConvertWithTimezone(t, fromFormat, toFormat, timeZone string) string {
	timeConvert, err := time.Parse(fromFormat, t)
	if err != nil {
		return ""
	}

	localTime, err := In(timeConvert, timeZone)
	if err != nil {
		return ""
	}

	return localTime.Format(toFormat)
}

// Diff ...
func Diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
