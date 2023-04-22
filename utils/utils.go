package utils

import (
	"time"
)

var layout string = "15:04"

func ConvertStringToTime(t string) (time.Time, error) {
	formattedTime, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, err
	}
	return formattedTime, nil
}
