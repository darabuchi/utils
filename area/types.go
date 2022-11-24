package area

import (
	"github.com/darabuchi/utils/language"
)

type (
	CountryInfo struct {
		Code     string                           `json:"code,omitempty"`
		Name     string                           `json:"name,omitempty"`
		TelCode  string                           `json:"tel_code,omitempty"`
		Domain   string                           `json:"domain,omitempty"`
		Banner   string                           `json:"banner,omitempty"`
		ISO3     string                           `json:"iso3,omitempty"`
		Capital  string                           `json:"capital,omitempty"`
		Timezone string                           `json:"timezone,omitempty"`
		I18nName map[language.LanguageCode]string `json:"i18n_name,omitempty"`
	}

	AirPortInfo struct {
		Code            string                           `json:"code,omitempty"`
		Icao            string                           `json:"icao,omitempty"`
		Name            string                           `json:"name,omitempty"`
		CountryCode     string                           `json:"country_code,omitempty"`
		Timezone        string                           `json:"timezone,omitempty"`
		I18nName        map[language.LanguageCode]string `json:"i18n_name,omitempty"`
		CountryI18nName map[language.LanguageCode]string `json:"country_i18n_name,omitempty"`
	}
)
