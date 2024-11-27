package utils

import "time"

func CurrentDate() time.Time {
	currentTime := time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	return currentDate
}
