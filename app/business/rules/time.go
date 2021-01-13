package rules

import (
	"strconv"
	"time"
)

func GetTimeFromUnixNano(t string) (time.Time, error) {
	u, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, u), nil
}

func IsIn24HoursFromNow(t time.Time) bool {
	return t.After(time.Now().Add(-time.Hour * 24))
}
