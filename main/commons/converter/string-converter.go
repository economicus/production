package converter

import (
	"strconv"
	"time"
)

func StrToTime(str, format string) (time.Time, error) {
	return time.Parse(format, str)
}

func StrToUint(str string) (uint, error) {
	res, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(res), nil
}
