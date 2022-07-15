package area

import (
	"github.com/darabuchi/utils/language"
)

type (
	CountryInfo struct {
		Code     string                           `json:"code,omitempty"`
		Name     string                           `json:"name,omitempty"`
		I18nName map[language.LanguageCode]string `json:"i18n_name,omitempty"`
		TelCode  string                           `json:"tel_code,omitempty"`
		Banner   string                           `json:"banner,omitempty"`
	}
)
