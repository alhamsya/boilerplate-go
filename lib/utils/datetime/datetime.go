package datetime

import (
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
)

func CurrentTimeF(format string) (string, error) {
	location, err := time.LoadLocation(constCommon.TimeLocalJakarta)
	if err != nil {
		return "", err
	}
	return time.Now().In(location).Format(format), nil
}
