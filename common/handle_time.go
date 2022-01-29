package common

import (
	"fmt"
	"time"
)

// check the job deadline with the current time, see if it's expired, expired is true
func CheckTimeBefore(jobDeadline string) (bool, error) {
	timeDeadline, err := ParseTime(jobDeadline)
	if err != nil {
		fmt.Println(err)
	}

	timeToday, err := ParseTime(time.Now().Format("02/01/2006"))
	if err != nil {
		fmt.Println(err)
	}

	res := timeDeadline.Before(timeToday)

	return res, nil
}

// parses a format time string to time time
func ParseTime(timeString string) (time.Time, error) {
	timeTime, err := time.Parse("02/01/2006", timeString)
	if err != nil {
		fmt.Println(err)
	}
	return timeTime, nil
}
