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

func ValidateDateFormat(date string) (time.Time, error) {
	dateTime, err := time.Parse(TimeFormat, date)
	if err != nil {
		fmt.Println("Error parsing date. Wrong Format:", err)
		return time.Time{}, err
	}
	return dateTime, nil
}

//func ValidateDateFormat(date string) bool {
//	_, err := time.Parse(TimeFormat, date)
//	if err != nil {
//		fmt.Println("Error parsing date. Need to follow format YYYY-MM-DD", err)
//		return false
//	}
//	return true
//}

func CheckValidDate(checkInDate string, date string) bool {
	checkInDateTime, err := time.Parse(TimeFormat, checkInDate)
	if err != nil {
		fmt.Println("Error parsing date. Wrong Format:", err)
		return false
	}

	//checkInDateTime, err := ValidateDateFormat3(checkInDate)
	//if err != nil {
	//	fmt.Println("Error parsing date. Wrong Format:", err)
	//	return false
	//}

	dateTime, err := time.Parse(TimeFormat, date)
	if err != nil {
		fmt.Println("Error parsing date. Wrong Format:", err)
		return false
	}

	//dateTime, err := ValidateDateFormat3(date)
	//if err != nil {
	//	fmt.Println("Error parsing date. Wrong Format:", err)
	//	return false
	//}

	validTo := checkInDateTime.Add(FiveDays)

	if dateTime.Equal(validTo) || dateTime.After(validTo) {
		return true
	}
	return false
}
