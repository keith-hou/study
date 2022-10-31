package time

import (
	"time"
)

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currTime time.Time, d string) (time.Time, error) {
	duration, err :=  time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currTime.Add(duration), nil
}