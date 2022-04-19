package utils

import (
	"strconv"
	"strings"
	"time"
)

type Dates struct {
}

/**
input format : "DD.MM.YYYYThh:mm:ss"
*/
func (d Dates) ParseWithDefault(input string, defaultValue time.Time) time.Time {

	datetimeList := strings.Split(input, "T")
	if len(datetimeList) != 2 {
		return defaultValue
	}

	Date := strings.Split(datetimeList[0], ".")
	Time := strings.Split(datetimeList[1], ":")

	if len(Date) != 3 || len(Time) != 3 {
		return defaultValue
	}

	convertedDateList := make([]int, 0)
	convertedTimeList := make([]int, 0)

	for _, strNumber := range Date {
		converted, err := strconv.ParseInt(strNumber, 10, 64)
		if err != nil {
			return defaultValue
		}
		convertedDateList = append(convertedDateList, int(converted))
	}

	for _, strNumber := range Time {
		converted, err := strconv.ParseInt(strNumber, 10, 64)
		if err != nil {
			return defaultValue
		}
		convertedTimeList = append(convertedTimeList, int(converted))
	}

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return defaultValue
	}

	return time.Date(convertedDateList[2], time.Month(convertedDateList[1]), convertedDateList[0],
		convertedTimeList[0], convertedTimeList[1], convertedTimeList[2], 0, location)
}
