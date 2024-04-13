package date_time

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetCurrentTimeString() string {
	return GetNow().Format(apiDbLayout)
}

func GetNowDbFormat() string {
	return GetNow().Format(apiDbLayout)
}

func SubTimes(t1 time.Time) int {
	t2, _ := time.Parse("2006-01-02 15:04:05", GetNowDbFormat())
	return int(t2.Sub(t1))
}
