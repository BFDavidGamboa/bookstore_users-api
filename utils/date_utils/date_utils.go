package date_utils

import "time"

const (
	apiDateLayout = "01-02-02T15:04:05Z-07:00"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
