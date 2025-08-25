package utils

import "time"

var (
	MapMonths = map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}
)

func GetParsedDate(str string) (*time.Time, error) {
	timezone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	parsedDate, err := time.ParseInLocation("2006-01-02", str, timezone)
	if err != nil {
		return nil, err
	}

	return &parsedDate, nil
}

func GetParsedDateTime(str string) (*time.Time, error) {
	timezone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	parsedDateTime, err := time.ParseInLocation("2006-01-02 15:04:05", str, timezone)
	if err != nil {
		return nil, err
	}

	return &parsedDateTime, nil
}
