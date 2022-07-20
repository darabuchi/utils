package language

type LanguageCode string

const (
	Chinese                  LanguageCode = "zh"
	English                  LanguageCode = "en"
	Arabic                   LanguageCode = "ar"
	Farsi                    LanguageCode = "fa"
	Sindhi                   LanguageCode = "sd"
	Kashmiri                 LanguageCode = "ks"
	Assamese                 LanguageCode = "as"
	Oriya                    LanguageCode = "or"
	Burmese                  LanguageCode = "my"
	Afrikaans                LanguageCode = "af"
	Albanian                 LanguageCode = "sq"
	Amharic                  LanguageCode = "am"
	Armenian                 LanguageCode = "hy"
	Azerbaijani              LanguageCode = "az"
	Basque                   LanguageCode = "eu"
	Belarusian               LanguageCode = "be"
	Bengali                  LanguageCode = "bn"
	Bosnian                  LanguageCode = "bs"
	Bulgarian                LanguageCode = "bg"
	Catalan                  LanguageCode = "ca"
	Cebuano                  LanguageCode = "ceb"
	Chichewa                 LanguageCode = "ny"
	ChineseCN                LanguageCode = "zh-cn"
	ChineseHK                LanguageCode = "zh-hk"
	ChineseTW                LanguageCode = "zh-tw"
	ChineseSG                LanguageCode = "zh-sg"
	ChineseMO                LanguageCode = "zh-mo"
	ChineseHans              LanguageCode = "zh-hans"
	ChineseHant              LanguageCode = "zh-hant"
	ChineseSimplifiedLegacy  LanguageCode = "zh-chs"
	ChineseTraditionalLegacy LanguageCode = "zh-cht"
	ChineseSimplified        LanguageCode = "zh-cn"
	ChineseTraditional       LanguageCode = "zh-tw"
	Chuvash                  LanguageCode = "cv"
	Croatian                 LanguageCode = "hr"
	Czech                    LanguageCode = "cs"
	Danish                   LanguageCode = "da"
	Dutch                    LanguageCode = "nl"
	EnglishAU                LanguageCode = "en-au"
	EnglishCA                LanguageCode = "en-ca"
	EnglishGB                LanguageCode = "en-gb"
	EnglishIN                LanguageCode = "en-in"
	EnglishIE                LanguageCode = "en-ie"
	EnglishNZ                LanguageCode = "en-nz"
	EnglishUS                LanguageCode = "en-us"
	EnglishJM                LanguageCode = "en-jm"
	EnglishPH                LanguageCode = "en-ph"
	EnglishSG                LanguageCode = "en-sg"
	EnglishZA                LanguageCode = "en-za"
	Estonian                 LanguageCode = "et"
	Filipino                 LanguageCode = "tl"
	Finnish                  LanguageCode = "fi"
	French                   LanguageCode = "fr"
	FrenchCA                 LanguageCode = "fr-ca"
	FrenchBE                 LanguageCode = "fr-be"
	FrenchCH                 LanguageCode = "fr-ch"
	FrenchCD                 LanguageCode = "fr-cd"
	FrenchKM                 LanguageCode = "fr-km"
	FrenchCG                 LanguageCode = "fr-cg"
	FrenchML                 LanguageCode = "fr-ml"
	FrenchCI                 LanguageCode = "fr-ci"
	FrenchFR                 LanguageCode = "fr-fr"
	FrenchLU                 LanguageCode = "fr-lu"
	FrenchMC                 LanguageCode = "fr-mc"
	FrenchSN                 LanguageCode = "fr-sn"
	FrenchTN                 LanguageCode = "fr-tn"
	FrenchGP                 LanguageCode = "fr-gp"
	FrenchMG                 LanguageCode = "fr-mg"
	FrenchRE                 LanguageCode = "fr-re"
	FrenchRW                 LanguageCode = "fr-rw"
	FrenchBL                 LanguageCode = "fr-bl"
	FrenchMF                 LanguageCode = "fr-mf"
	FrenchYT                 LanguageCode = "fr-yt"
	FrenchPM                 LanguageCode = "fr-pm"
	FrenchWF                 LanguageCode = "fr-wf"
	FrenchMQ                 LanguageCode = "fr-mq"
	FrenchTF                 LanguageCode = "fr-tf"
	Frisian                  LanguageCode = "fy"
	Galician                 LanguageCode = "gl"
	Georgian                 LanguageCode = "ka"
	German                   LanguageCode = "de"
	GermanAT                 LanguageCode = "de-at"
	GermanCH                 LanguageCode = "de-ch"
	GermanLI                 LanguageCode = "de-li"
	GermanLU                 LanguageCode = "de-lu"
	GermanDE                 LanguageCode = "de-de"
	Greek                    LanguageCode = "el"
	Gujarati                 LanguageCode = "gu"
	Haitian                  LanguageCode = "ht"
	Hausa                    LanguageCode = "ha"
	Hebrew                   LanguageCode = "he"
	Hindi                    LanguageCode = "hi"
	Hungarian                LanguageCode = "hu"
	Icelandic                LanguageCode = "is"
	Igbo                     LanguageCode = "ig"
	Indonesian               LanguageCode = "id"
	Irish                    LanguageCode = "ga"
	Italian                  LanguageCode = "it"
	Japanese                 LanguageCode = "ja"
	Kannada                  LanguageCode = "kn"
	Kazakh                   LanguageCode = "kk"
	Khmer                    LanguageCode = "km"
	Korean                   LanguageCode = "ko"
	Kurdish                  LanguageCode = "ku"
	Kyrgyz                   LanguageCode = "ky"
	Lao                      LanguageCode = "lo"
	Latin                    LanguageCode = "la"
	Latvian                  LanguageCode = "lv"
	Lithuanian               LanguageCode = "lt"
	Luxembourgish            LanguageCode = "lb"
	Macedonian               LanguageCode = "mk"
	Malagasy                 LanguageCode = "mg"
	Malay                    LanguageCode = "ms"
	Malayalam                LanguageCode = "ml"
	Maltese                  LanguageCode = "mt"
	Maori                    LanguageCode = "mi"
	Marathi                  LanguageCode = "mr"
	Mongolian                LanguageCode = "mn"
	Nepali                   LanguageCode = "ne"
	Norwegian                LanguageCode = "no"
	Pashto                   LanguageCode = "ps"
	Persian                  LanguageCode = "fa"
	Polish                   LanguageCode = "pl"
	Portuguese               LanguageCode = "pt"
	PortugueseBR             LanguageCode = "pt-br"
	PortuguesePT             LanguageCode = "pt-pt"
	Punjabi                  LanguageCode = "pa"
	Romanian                 LanguageCode = "ro"
	Russian                  LanguageCode = "ru"
	Samoan                   LanguageCode = "sm"
	Serbian                  LanguageCode = "sr"
	Sinhalese                LanguageCode = "si"
	Slovak                   LanguageCode = "sk"
	Slovenian                LanguageCode = "sl"
	Somali                   LanguageCode = "so"
	Spanish                  LanguageCode = "es"
	SpanishAR                LanguageCode = "es-ar"
	SpanishBO                LanguageCode = "es-bo"
	SpanishCL                LanguageCode = "es-cl"
	SpanishCO                LanguageCode = "es-co"
	SpanishCR                LanguageCode = "es-cr"
	SpanishCU                LanguageCode = "es-cu"
	SpanishDO                LanguageCode = "es-do"
	SpanishEC                LanguageCode = "es-ec"
	SpanishSV                LanguageCode = "es-sv"
	SpanishGT                LanguageCode = "es-gt"
	SpanishHN                LanguageCode = "es-hn"
	SpanishMX                LanguageCode = "es-mx"
	SpanishNI                LanguageCode = "es-ni"
	SpanishPA                LanguageCode = "es-pa"
	SpanishES                LanguageCode = "es-es"
	SpanishPY                LanguageCode = "es-py"
	SpanishUS                LanguageCode = "es-us"
	SpanishPE                LanguageCode = "es-pe"
	SpanishPR                LanguageCode = "es-pr"
	SpanishSL                LanguageCode = "es-sl"
	SpanishUY                LanguageCode = "es-uy"
	SpanishVE                LanguageCode = "es-ve"
	Swahili                  LanguageCode = "sw"
	Swedish                  LanguageCode = "sv"
	Tamil                    LanguageCode = "ta"
	Telugu                   LanguageCode = "te"
	Thai                     LanguageCode = "th"
	Turkish                  LanguageCode = "tr"
	Ukrainian                LanguageCode = "uk"
	Urdu                     LanguageCode = "ur"
	Uzbek                    LanguageCode = "uz"
	Vietnamese               LanguageCode = "vi"
	Welsh                    LanguageCode = "cy"
	Xhosa                    LanguageCode = "xh"
	Yiddish                  LanguageCode = "yi"
	Yoruba                   LanguageCode = "yo"
	Zulu                     LanguageCode = "zu"
)
