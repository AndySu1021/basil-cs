package report

import (
	"time"
)

func GetTimeRange(dateInput string) (startTime, endTime time.Time, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	if dateInput == "" {
		dateInput = time.Now().In(loc).Format("2006-01-02")
	}

	startTime, err = time.ParseInLocation("2006-01-02", dateInput, loc)
	if err != nil {
		return
	}
	
	startTime = startTime.UTC()
	endTime = startTime.Add(24 * time.Hour)

	return
}
