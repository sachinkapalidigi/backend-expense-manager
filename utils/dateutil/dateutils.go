package dateutil

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNowUTC : gets current time in terms of utc
func GetNowUTC() time.Time {
	return time.Now().UTC()
}

// GetNowStringUTC : gets current time in normal format with t in between to indicate time utc
func GetNowStringUTC() string {

	return GetNowUTC().Format(apiDateLayout)
}

// GetNowDBFormatUTC : gives current time format supported for db utc
func GetNowDBFormatUTC() string {
	return GetNowUTC().Format(apiDbLayout)
}

// non utc

// GetNow : get now without utc
func GetNow() time.Time {
	return time.Now()
}

// GetNowString : get now without utc in normal format
func GetNowString() string {

	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat : get now without utc in db format
func GetNowDBFormat() string {

	return GetNow().Format(apiDbLayout)
}
