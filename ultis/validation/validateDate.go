package validation

import (
	"time"
)

const (
	TimeFormat = "2006-01-02"
)

func ValidDate(date string) (time.Time, bool) {
	dateTime, err := time.Parse(TimeFormat, date)
	if err != nil {
		return time.Time{}, false
	}
	return dateTime, true
}


func CheckValidDate(checkInDate time.Time, date string) bool {
	dateTime, ok := ValidDate(date)
	if !ok {
		return false
	}

	validTo := checkInDate.AddDate(0, 0, 5)

	if dateTime.Equal(validTo) || dateTime.After(validTo) {
		return true
	}
	return false
}
