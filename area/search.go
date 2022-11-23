package area

import (
	"strings"

	"github.com/darabuchi/utils/language"
)

func GetCountryName(code string, languageCode language.LanguageCode) string {
	code = strings.ToUpper(code)
	if c, ok := CountryCodeMap[code]; ok {
		if n, ok := c.I18nName[languageCode]; ok {
			return n
		}

		return c.Name
	}

	return code
}

func GetCountryBanner(code string) string {
	code = strings.ToUpper(code)
	if c, ok := CountryCodeMap[code]; ok {
		return c.Banner
	}

	return code
}

func GetCountry(code string) *CountryInfo {
	if c, ok := CountryCodeMap[code]; ok {
		return c
	}

	return CountryCodeMap[""]
}
