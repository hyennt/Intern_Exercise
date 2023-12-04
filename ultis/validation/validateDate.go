package validation

import (
	"fmt"
	"time"
)

const (
	CheckInDate = "2019-12-25"
	TimeFormat  = "2006-01-02"
	FiveDays    = 120 * time.Hour
)

func CheckValidDate(checkInDate string, date string) bool {
	checkInDateTime, err := time.Parse(TimeFormat, checkInDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}
	dateTime, err := time.Parse(TimeFormat, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}

	validTo := checkInDateTime.Add(FiveDays)

	if dateTime.Equal(validTo) || dateTime.After(validTo) {
		return true
	}
	return false
}
