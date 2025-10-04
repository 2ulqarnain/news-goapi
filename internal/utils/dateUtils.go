package utils

import "time"

func ConvertDateToISO(date string) string {
	parsedDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return ""
	}
	return parsedDate.Format("2006-01-02")
}
