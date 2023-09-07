package utils

import (
	"time"
)

const (
	dateDMY = "02.01.2006"
	dateYMD = "2006.01.02"
	timeHMS = "15:04:05"
)

func ParseDate(in string) (time.Time, error) {
	data, err := time.Parse(dateDMY, in)
	if err == nil {
		return data, nil
	}
	return time.Parse(dateYMD, in)
}

func ParseTime(in string) (time.Time, error) {
	return time.Parse(timeHMS, in)
}
