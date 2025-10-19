package helpers

import "time"

func GetWeekRange(now time.Time) (time.Time, time.Time) {
	now = now.Truncate(24 * time.Hour)

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -7
	}

	return now.AddDate(0, 0, offset), now.AddDate(0, 0, offset).AddDate(0, 0, 6)
}
