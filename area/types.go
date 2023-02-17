package area

import (
	"github.com/darabuchi/utils/language"
)

type (
	LanguageInfo struct {
		Code     string `json:"code,omitempty"`
		Name     string `json:"name,omitempty"`
		ISO6391  string `json:"iso6391,omitempty"`
		ISO6392T string `json:"iso6392t,omitempty"`
		ISO6392B string `json:"iso6392b,omitempty"`
		ISO6393  string `json:"iso6393,omitempty"`

		I18nName map[language.LanguageCode]string `json:"i18n_name,omitempty"`
	}

	CurrencyInfo struct {
		Code   string `json:"code,omitempty"`
		Name   string `json:"name,omitempty"`
		Symbol string `json:"symbol,omitempty"`

		I18nName map[language.LanguageCode]string `json:"i18n_name,omitempty"`
	}

	CountryInfo struct {
		Code string `json:"code,omitempty"`
		Name string `json:"name,omitempty"`

		TelCode string `json:"tel_code,omitempty"`
		Domain  string `json:"domain,omitempty"`
		Banner  string `json:"banner,omitempty"`

		ISO2    string `json:"iso2,omitempty"`
		ISO3    string `json:"iso3,omitempty"`
		ISO3166 string `json:"iso3166,omitempty"`
		STANAG  string `json:"stanag,omitempty"`
		GEC     string `json:"gec,omitempty"` //

		Capital string `json:"capital,omitempty"`

		NationalLanguage language.LanguageCode   `json:"national_language,omitempty"` // 民族语言,国家语言,国语,官方语言
		Languages        []language.LanguageCode `json:"languages,omitempty"`
		Currency         string                  `json:"currency,omitempty"` // 货币
		Timezone         Timezone                `json:"timezone,omitempty"`

		I18nName map[language.LanguageCode]string `json:"i18n_name,omitempty"`
	}

	AirPortInfo struct {
		Code            string                           `json:"code,omitempty"`
		Icao            string                           `json:"icao,omitempty"`
		Name            string                           `json:"name,omitempty"`
		CountryCode     string                           `json:"country_code,omitempty"`
		Timezone        Timezone                         `json:"timezone,omitempty"`
		I18nName        map[language.LanguageCode]string `json:"i18n_name,omitempty"`
		CountryI18nName map[language.LanguageCode]string `json:"country_i18n_name,omitempty"`
	}
)
