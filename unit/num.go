package unit

import (
	"strconv"
)

func FormatInt64(val int64, prec int) string {
	// 小于 1万 的直接展示
	if val < 10000 {
		return strconv.FormatInt(val, 10)
	}

	if val < 10000000 {
		if val%10000 == 0 {
			return strconv.FormatInt(val/10000, 10) + "万"
		}
		return strconv.FormatFloat(float64(val)/10000, 'f', prec, 64) + "万"
	}

	if val < 100000000000 {
		if val%100000000 == 0 {
			return strconv.FormatInt(val/100000000, 10) + "亿"
		}
		return strconv.FormatFloat(float64(val)/100000000, 'f', prec, 64) + "亿"
	}

	return FormatInt64(val, 10)
}
