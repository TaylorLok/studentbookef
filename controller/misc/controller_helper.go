package misc

import "time"

const (
	YYYYMMDD_FORMAT    = "2006-01-02"
	YYYMMDDTIME_FORMAT = "2006-01-02 15:04:05"
)

/**
Format date in yyyy-MM-dd HH:mm:ss
*/

func FormatDateTime(date time.Time) string {
	return date.Format(YYYMMDDTIME_FORMAT)
}

/**
format date in yyyy-MM-dd
*/
func FormatDate(date time.Time) string {
	return date.Format(YYYYMMDD_FORMAT)
}
