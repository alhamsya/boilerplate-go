package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
)

func (l *thing) CurrentTimeF(format string) (string, error) {
	location, err := time.LoadLocation(constCommon.TimeLocalJakarta)
	if err != nil {
		return "", err
	}
	return time.Now().In(location).Format(format), nil
}

func (l *thing) GenerateRangeDate(fromYear, fromMonth int) (result []map[string]time.Time, err error) {
	var mut sync.Mutex
	mut.Lock()
	defer mut.Unlock()

	if fromYear < 2000 {
		return nil, fmt.Errorf("request year less than 2000")
	}

	if fromMonth < 1 || fromMonth > 12 {
		return nil, fmt.Errorf("request month less than 1 and more than 12")
	}

	location, err := time.LoadLocation(constCommon.TimeLocalJakarta)
	if err != nil {
		return nil, err
	}

	currentYear := time.Now().In(location).Year()
	currentMonth := int(time.Now().In(location).Month())

	limit := false
	for y := fromYear; y <= currentYear; y++ {
		for m := 1; m <= 12; m++ {
			if m < fromMonth && limit == false {
				continue
			}

			if m == fromMonth {
				limit = true
			}

			result = append(result, map[string]time.Time{
				"start_date": time.Date(y, time.Month(m), 1, 0, 0, 0, 0, location),
				"end_date":   time.Date(y, time.Month(m+1), 1, 0, 0, 0, 0, location),
			})

			if y == currentYear && m == currentMonth {
				break
			}
		}
	}

	return result, nil
}
