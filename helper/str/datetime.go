package str

import (
	"log"
	"time"
)

func DateStringToDate(dateStr, format string) time.Time {
	date, err := time.Parse(format, dateStr)
	log.Println(err)
	return date
}
