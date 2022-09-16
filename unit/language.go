package unit

import (
	"github.com/darabuchi/utils/language"
)

type UnitCode string

const (
	Day         UnitCode = "day"
	Hour        UnitCode = "hour"
	Minute      UnitCode = "minute"
	Second      UnitCode = "second"
	Millisecond UnitCode = "millisecond"
)

var i18N = map[UnitCode]map[language.LanguageCode]string{
	Day: {
		language.Chinese:  "天",
		language.English:  "day",
		language.Japanese: "日",
	},
	Hour: {
		language.Chinese:  "小时",
		language.English:  "hour",
		language.Japanese: "時間",
	},
	Minute: {
		language.Chinese:  "分钟",
		language.English:  "minute",
		language.Japanese: "分",
	},
	Second: {
		language.Chinese:  "秒",
		language.English:  "second",
		language.Japanese: "秒",
	},
	Millisecond: {
		language.Chinese:  "毫秒",
		language.English:  "millisecond",
		language.Japanese: "ミリ秒",
	},
}

func GetUnitCode2Show(unitCode UnitCode, languageCode language.LanguageCode) string {
	if u, ok := i18N[unitCode]; ok {
		if s, ok := u[languageCode]; ok {
			return s
		}
	}

	return string(unitCode)
}
