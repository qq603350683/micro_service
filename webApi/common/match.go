package common

import (
	"errors"
	"regexp"
	"time"
)

// 判断是否日期格式
func IsDate(str string) bool {
	b, err := regexp.MatchString(`(\d{4})([//-]{1})(\d{1,2})([//-]{1})(\d{1,2})`, str)
	if err != nil {
		return false
	}

	return b
}

// 判断是否时间格式
func IsTime() bool {
	return true
}

// 判断时间大小
func CompareTime(date1 string, symbol string, date2 string) (bool, error) {
	loc, _ := time.LoadLocation("Local")
	d1, err := time.ParseInLocation("2006-01-02 15:04:05", date1, loc)
	if err != nil {
		return false, err
	}

	d2, err := time.ParseInLocation("2006-01-02 15:04:05", date2, loc)
	if err != nil {
		return false, err
	}

	u1 := d1.Unix()
	u2 := d2.Unix()

	switch symbol {
	case ">":
		if u1 > u2 {
			return true, nil
		} else {
			return false, nil
		}
	case "=":
		if u1 == u2 {
			return true, nil
		} else {
			return false, nil
		}
	case "<":
		if u1 < u2 {
			return true, nil
		} else {
			return false, nil
		}
	default:
		return false, errors.New("symbol: " + symbol + " 不在指定的值[<>=]内")
	}
}


