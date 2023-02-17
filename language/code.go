package language

type LanguageCode string

//goland:noinspection ALL
const (
	Afar                      LanguageCode = "aa"         // 阿法尔语
	Abkhazian                 LanguageCode = "ab"         // 阿布哈兹语
	Avestan                   LanguageCode = "ae"         // 阿维斯陀语
	Afrikaans                 LanguageCode = "af"         // 南非荷兰语
	Akan                      LanguageCode = "ak"         // 阿肯语
	Amharic                   LanguageCode = "am"         // 阿姆哈拉语
	Arabic                    LanguageCode = "ar"         // 阿拉伯语
	Arabic_Bahrain            LanguageCode = "ar-bh"      // 阿拉伯语（巴林）
	Arabic_Algeria            LanguageCode = "ar-dz"      // 阿拉伯语（阿尔及利亚）
	Arabic_Egypt              LanguageCode = "ar-eg"      // 阿拉伯语（埃及）
	Arabic_Iraq               LanguageCode = "ar-iq"      // 阿拉伯语（伊拉克）
	Arabic_Jordan             LanguageCode = "ar-jo"      // 阿拉伯语（约旦）
	Arabic_Kuwait             LanguageCode = "ar-kw"      // 阿拉伯语（科威特）
	Arabic_Lebanon            LanguageCode = "ar-lb"      // 阿拉伯语（黎巴嫩）
	Arabic_Libya              LanguageCode = "ar-ly"      // 阿拉伯语（利比亚）
	Arabic_Morocco            LanguageCode = "ar-ma"      // 阿拉伯语（摩洛哥）
	Arabic_Oman               LanguageCode = "ar-om"      // 阿拉伯语（阿曼）
	Arabic_Qatar              LanguageCode = "ar-qa"      // 阿拉伯语（卡塔尔）
	Arabic_SaudiArabia        LanguageCode = "ar-sa"      // 阿拉伯语（沙特阿拉伯）
	Arabic_Sudan              LanguageCode = "ar-sd"      // 阿拉伯语（苏丹）
	Arabic_Syria              LanguageCode = "ar-sy"      // 阿拉伯语（叙利亚）
	Arabic_Tunisia            LanguageCode = "ar-tn"      // 阿拉伯语（突尼斯）
	Arabic_UAE                LanguageCode = "ar-ae"      // 阿拉伯语（阿拉伯联合酋长国）
	Arabic_Yemen              LanguageCode = "ar-ye"      // 阿拉伯语（也门）
	Assamese                  LanguageCode = "as"         // 阿萨姆语
	Avaric                    LanguageCode = "av"         // 阿瓦尔语
	Aymara                    LanguageCode = "ay"         // 阿亚马拉语
	Azerbaijani               LanguageCode = "az"         // 阿塞拜疆语
	Azerbaijani_Latin         LanguageCode = "az-l"       // 阿塞拜疆语（拉丁文）
	Bashkir                   LanguageCode = "ba"         // 巴什基尔语
	Belarusian                LanguageCode = "be"         // 白俄罗斯语
	Bulgarian                 LanguageCode = "bg"         // 保加利亚语
	Bihari                    LanguageCode = "bh"         // 比哈尔语
	Bislama                   LanguageCode = "bi"         // 比斯拉马语
	Bambara                   LanguageCode = "bm"         // 班巴拉语
	Bengali                   LanguageCode = "bn"         // 孟加拉语
	Bengali_Bangladesh        LanguageCode = "bn-bd"      // 孟加拉语（孟加拉国）
	Bengali_India             LanguageCode = "bn-in"      // 孟加拉语（印度）
	Tibetan                   LanguageCode = "bo"         // 藏语
	Tibetan_China             LanguageCode = "bo-cn"      // 藏语（中国）
	Breton                    LanguageCode = "br"         // 布列塔尼语
	Bosnian                   LanguageCode = "bs"         // 波斯尼亚语
	Catalan                   LanguageCode = "ca"         // 加泰罗尼亚语
	Chechen                   LanguageCode = "ce"         // 车臣语
	Chamorro                  LanguageCode = "ch"         // 查莫罗语
	Corsican                  LanguageCode = "co"         // 科西嘉语
	Cree                      LanguageCode = "cr"         // 克里语
	Czech                     LanguageCode = "cs"         // 捷克语
	ChurchSlavic              LanguageCode = "cu"         // 教会斯拉夫语
	Chuvash                   LanguageCode = "cv"         // 楚瓦什语
	Welsh                     LanguageCode = "cy"         // 威尔士语
	Danish                    LanguageCode = "da"         // 丹麦语
	German                    LanguageCode = "de"         // 德语
	German_Austria            LanguageCode = "de-at"      // 德语（奥地利）
	German_Liechtenstein      LanguageCode = "de-li"      // 德语（列支敦士登）
	German_Luxembourg         LanguageCode = "de-lu"      // 德语（卢森堡）
	German_Switzerland        LanguageCode = "de-ch"      // 德语（瑞士）
	Divehi                    LanguageCode = "dv"         // 迪维希语
	Dzongkha                  LanguageCode = "dz"         // 宗卡语
	Ewe                       LanguageCode = "ee"         // 埃维语
	Greek                     LanguageCode = "el"         // 希腊语
	English                   LanguageCode = "en"         // 英语
	English_Australia         LanguageCode = "en-au"      // 英语（澳大利亚）
	English_Belize            LanguageCode = "en-bz"      // 英语（伯利兹）
	English_Canada            LanguageCode = "en-ca"      // 英语（加拿大）
	English_Caribbean         LanguageCode = "en-cb"      // 英语（加勒比海）
	English_India             LanguageCode = "en-in"      // 英语（印度）
	English_Ireland           LanguageCode = "en-ie"      // 英语（爱尔兰）
	English_Jamaica           LanguageCode = "en-jm"      // 英语（牙买加）
	English_NewZealand        LanguageCode = "en-nz"      // 英语（新西兰）
	English_Philippines       LanguageCode = "en-ph"      // 英语（菲律宾）
	English_SouthAfrica       LanguageCode = "en-za"      // 英语（南非）
	English_Trinidad          LanguageCode = "en-tt"      // 英语（特立尼达和多巴哥）
	English_UnitedKingdom     LanguageCode = "en-gb"      // 英语（英国）
	English_UnitedStates      LanguageCode = "en-us"      // 英语（美国）
	English_Zimbabwe          LanguageCode = "en-zw"      // 英语（津巴布韦）
	Esperanto                 LanguageCode = "eo"         // 世界语
	Spanish                   LanguageCode = "es"         // 西班牙语
	Spanish_Argentina         LanguageCode = "es-ar"      // 西班牙语（阿根廷）
	Spanish_Bolivia           LanguageCode = "es-bo"      // 西班牙语（玻利维亚）
	Spanish_Chile             LanguageCode = "es-cl"      // 西班牙语（智利）
	Spanish_Colombia          LanguageCode = "es-co"      // 西班牙语（哥伦比亚）
	Spanish_CostaRica         LanguageCode = "es-cr"      // 西班牙语（哥斯达黎加）
	Spanish_DominicanRepublic LanguageCode = "es-do"      // 西班牙语（多米尼加共和国）
	Spanish_Ecuador           LanguageCode = "es-ec"      // 西班牙语（厄瓜多尔）
	Spanish_Spain             LanguageCode = "es-es"      // 西班牙语（西班牙）
	Spanish_Guatemala         LanguageCode = "es-gt"      // 西班牙语（危地马拉）
	Spanish_Honduras          LanguageCode = "es-hn"      // 西班牙语（洪都拉斯）
	Spanish_Mexico            LanguageCode = "es-mx"      // 西班牙语（墨西哥）
	Spanish_Nicaragua         LanguageCode = "es-ni"      // 西班牙语（尼加拉瓜）
	Spanish_Panama            LanguageCode = "es-pa"      // 西班牙语（巴拿马）
	Spanish_Paraguay          LanguageCode = "es-py"      // 西班牙语（巴拉圭）
	Spanish_Peru              LanguageCode = "es-pe"      // 西班牙语（秘鲁）
	Spanish_PuertoRico        LanguageCode = "es-pr"      // 西班牙语（波多黎各）
	Spanish_Uruguay           LanguageCode = "es-uy"      // 西班牙语（乌拉圭）
	Spanish_Venezuela         LanguageCode = "es-ve"      // 西班牙语（委内瑞拉）
	Estonian                  LanguageCode = "et"         // 爱沙尼亚语
	Basque                    LanguageCode = "eu"         // 巴斯克语
	Persian                   LanguageCode = "fa"         // 波斯语
	Fulah                     LanguageCode = "ff"         // 富拉语
	Finnish                   LanguageCode = "fi"         // 芬兰语
	Fijian                    LanguageCode = "fj"         // 斐济语
	Faroese                   LanguageCode = "fo"         // 法罗语
	French                    LanguageCode = "fr"         // 法语
	French_Belgium            LanguageCode = "fr-be"      // 法语（比利时）
	French_Canada             LanguageCode = "fr-ca"      // 法语（加拿大）
	French_France             LanguageCode = "fr-fr"      // 法语（法国）
	French_Luxembourg         LanguageCode = "fr-lu"      // 法语（卢森堡）
	French_Monaco             LanguageCode = "fr-mc"      // 法语（摩纳哥）
	French_Switzerland        LanguageCode = "fr-ch"      // 法语（瑞士）
	Frisian                   LanguageCode = "fy"         // 弗里西语
	Irish                     LanguageCode = "ga"         // 爱尔兰语
	ScotsGaelic               LanguageCode = "gd"         // 苏格兰盖尔语
	Galician                  LanguageCode = "gl"         // 加利西亚语
	Guarani                   LanguageCode = "gn"         // 瓜拉尼语
	Gujarati                  LanguageCode = "gu"         // 古吉拉特语
	Manx                      LanguageCode = "gv"         // 曼岛语
	Hausa                     LanguageCode = "ha"         // 豪萨语
	Hebrew                    LanguageCode = "he"         // 希伯来语
	Hindi                     LanguageCode = "hi"         // 印地语
	HiriMotu                  LanguageCode = "ho"         // 希里莫图语
	Croatian                  LanguageCode = "hr"         // 克罗地亚语
	Haitian                   LanguageCode = "ht"         // 海地克里奥尔语
	Hungarian                 LanguageCode = "hu"         // 匈牙利语
	Armenian                  LanguageCode = "hy"         // 亚美尼亚语
	Armenian_Eastern          LanguageCode = "hy-am"      // 亚美尼亚语（东方）
	Armenian_Western          LanguageCode = "hy-ws"      // 亚美尼亚语（西方）
	Armenian_Middle           LanguageCode = "hy-me"      // 亚美尼亚语（中部）
	Herero                    LanguageCode = "hz"         // 赫雷罗语
	Indonesian                LanguageCode = "id"         // 印尼语
	Interlingua               LanguageCode = "ia"         // 国际语
	Interlingue               LanguageCode = "ie"         // 国际语（E）
	Igbo                      LanguageCode = "ig"         // 伊博语
	SichuanYi                 LanguageCode = "ii"         // 四川彝语
	Inupiaq                   LanguageCode = "ik"         // 伊努皮克语
	Ido                       LanguageCode = "io"         // 伊多语
	Icelandic                 LanguageCode = "is"         // 冰岛语
	Italian                   LanguageCode = "it"         // 意大利语
	Italian_Italy             LanguageCode = "it-it"      // 意大利语（意大利）
	Italian_Switzerland       LanguageCode = "it-ch"      // 意大利语（瑞士）
	Inuktitut                 LanguageCode = "iu"         // 因纽特语
	Japanese                  LanguageCode = "ja"         // 日语
	Javanese                  LanguageCode = "jv"         // 爪哇语
	Georgian                  LanguageCode = "ka"         // 格鲁吉亚语
	Kongo                     LanguageCode = "kg"         // 刚果语
	Kikuyu                    LanguageCode = "ki"         // 基库尤语
	Kuanyama                  LanguageCode = "kj"         // 宽亚玛语
	Kazakh                    LanguageCode = "kk"         // 哈萨克语
	Kalaallisut               LanguageCode = "kl"         // 格陵兰语
	CentralKhmer              LanguageCode = "km"         // 中央高棉语
	Kannada                   LanguageCode = "kn"         // 卡纳达语
	Korean                    LanguageCode = "ko"         // 韩语
	Kanuri                    LanguageCode = "kr"         // 卡努里语
	Kashmiri                  LanguageCode = "ks"         // 克什米尔语
	Kashmiri_Arabic           LanguageCode = "ks-arab"    // 克什米尔语（阿拉伯文）
	Kashmiri_Arabic_IN        LanguageCode = "ks-arab-in" // 克什米尔语（阿拉伯文，印度）
	Kashmiri_Deva             LanguageCode = "ks-deva"    // 克什米尔语（梵文）
	Kashmiri_Deva_IN          LanguageCode = "ks-deva-in" // 克什米尔语（梵文，印度）
	Kurdish                   LanguageCode = "ku"         // 库尔德语
	Komi                      LanguageCode = "kv"         // 科米语
	Cornish                   LanguageCode = "kw"         // 康沃尔语
	Kirghiz                   LanguageCode = "ky"         // 吉尔吉斯语
	Latin                     LanguageCode = "la"         // 拉丁语
	Luxembourgish             LanguageCode = "lb"         // 卢森堡语
	Ganda                     LanguageCode = "lg"         // 干达语
	Limburgan                 LanguageCode = "li"         // 林堡语
	Lingala                   LanguageCode = "ln"         // 林加拉语
	Lao                       LanguageCode = "lo"         // 老挝语
	Lithuanian                LanguageCode = "lt"         // 立陶宛语
	LubaKatanga               LanguageCode = "lu"         // 卢巴-卡丹加语
	Latvian                   LanguageCode = "lv"         // 拉脱维亚语
	Malagasy                  LanguageCode = "mg"         // 马尔加什语
	Marshallese               LanguageCode = "mh"         // 马绍尔语
	Maori                     LanguageCode = "mi"         // 毛利语
	Macedonian                LanguageCode = "mk"         // 马其顿语
	Malayalam                 LanguageCode = "ml"         // 马拉雅拉姆语
	Mongolian                 LanguageCode = "mn"         // 蒙古语
	Moldavian                 LanguageCode = "mo"         // 摩尔多瓦语
	Marathi                   LanguageCode = "mr"         // 马拉地语
	Malay                     LanguageCode = "ms"         // 马来语
	Maltese                   LanguageCode = "mt"         // 马耳他语
	Burmese                   LanguageCode = "my"         // 缅甸语
	Nauru                     LanguageCode = "na"         // 瑙鲁语
	NorwegianBokmal           LanguageCode = "nb"         // 挪威语（博克马尔）
	NorthNdebele              LanguageCode = "nd"         // 北恩德贝勒语
	Nepali                    LanguageCode = "ne"         // 尼泊尔语
	Ndonga                    LanguageCode = "ng"         // 恩东加语
	Dutch                     LanguageCode = "nl"         // 荷兰语
	Dutch_Netherlands         LanguageCode = "nl-nl"      // 荷兰语（荷兰）
	Dutch_Belgium             LanguageCode = "nl-be"      // 荷兰语（比利时）
	NorwegianNynorsk          LanguageCode = "nn"         // 挪威语（尼诺斯克）
	Norwegian                 LanguageCode = "no"         // 挪威语
	SouthNdebele              LanguageCode = "nr"         // 南恩德贝勒语
	Navajo                    LanguageCode = "nv"         // 纳瓦霍语
	Chichewa                  LanguageCode = "ny"         // 齐切瓦语
	Occitan                   LanguageCode = "oc"         // 奥克语
	Ojibwa                    LanguageCode = "oj"         // 奥吉布瓦语
	Oromo                     LanguageCode = "om"         // 奥罗莫语
	Oriya                     LanguageCode = "or"         // 奥里亚语
	Ossetian                  LanguageCode = "os"         // 奥塞梯语
	Panjabi                   LanguageCode = "pa"         // 旁遮普语
	Pali                      LanguageCode = "pi"         // 帕利语
	Polish                    LanguageCode = "pl"         // 波兰语
	Pushto                    LanguageCode = "ps"         // 普什图语
	Portuguese                LanguageCode = "pt"         // 葡萄牙语
	Portuguese_Brazil         LanguageCode = "pt-br"      // 葡萄牙语（巴西）
	Portuguese_Portugal       LanguageCode = "pt-pt"      // 葡萄牙语（葡萄牙）
	Quechua                   LanguageCode = "qu"         // 克丘亚语
	Dari                      LanguageCode = "prs"        // 达里语
	Romansh                   LanguageCode = "rm"         // 罗曼什语
	Rundi                     LanguageCode = "rn"         // 隆迪语
	Romanian                  LanguageCode = "ro"         // 罗马尼亚语
	Russian                   LanguageCode = "ru"         // 俄语
	Kinyarwanda               LanguageCode = "rw"         // 基尼亚尔万达语
	Sanskrit                  LanguageCode = "sa"         // 梵语
	Sardinian                 LanguageCode = "sc"         // 萨丁尼亚语
	Sindhi                    LanguageCode = "sd"         // 信德语
	NorthernSami              LanguageCode = "se"         // 北萨米语
	Sango                     LanguageCode = "sg"         // 桑戈语
	Serbian                   LanguageCode = "sr"         // 塞尔维亚语
	Sinhalese                 LanguageCode = "si"         // 僧伽罗语
	Slovak                    LanguageCode = "sk"         // 斯洛伐克语
	Slovenian                 LanguageCode = "sl"         // 斯洛文尼亚语
	Samoan                    LanguageCode = "sm"         // 萨摩亚语
	Shona                     LanguageCode = "sn"         // 修纳语
	Somali                    LanguageCode = "so"         // 索马里语
	Albanian                  LanguageCode = "sq"         // 阿尔巴尼亚语
	SerbianLatin              LanguageCode = "sr-latn"    // 塞尔维亚语（拉丁文）
	Sesotho                   LanguageCode = "st"         // 塞索托语
	Sundanese                 LanguageCode = "su"         // 巽他语
	Swedish                   LanguageCode = "sv"         // 瑞典语
	Swahili                   LanguageCode = "sw"         // 斯瓦希里语
	Tamil                     LanguageCode = "ta"         // 泰米尔语
	Telugu                    LanguageCode = "te"         // 泰卢固语
	Tajik                     LanguageCode = "tg"         // 塔吉克语
	Thai                      LanguageCode = "th"         // 泰语
	Tigrinya                  LanguageCode = "ti"         // 提格利尼亚语
	Turkmen                   LanguageCode = "tk"         // 土库曼语
	Tagalog                   LanguageCode = "tl"         // 塔加路语
	Tswana                    LanguageCode = "tn"         // 茨瓦纳语
	Tonga                     LanguageCode = "to"         // 汤加语
	Turkish                   LanguageCode = "tr"         // 土耳其语
	Tsonga                    LanguageCode = "ts"         // 宗加语
	Tatar                     LanguageCode = "tt"         // 鞑靼语
	Twi                       LanguageCode = "tw"         // 塔威语
	Tahitian                  LanguageCode = "ty"         // 大溪地语
	Uyghur                    LanguageCode = "ug"         // 维吾尔语
	Ukrainian                 LanguageCode = "uk"         // 乌克兰语
	Urdu                      LanguageCode = "ur"         // 乌尔都语
	Uzbek                     LanguageCode = "uz"         // 乌兹别克语
	Venda                     LanguageCode = "ve"         // 文达语
	Vietnamese                LanguageCode = "vi"         // 越南语
	Volapuk                   LanguageCode = "vo"         // 沃拉普克语
	Walloon                   LanguageCode = "wa"         // 瓦隆语
	Wolof                     LanguageCode = "wo"         // 沃洛夫语
	Xhosa                     LanguageCode = "xh"         // 科萨语
	Yiddish                   LanguageCode = "yi"         // 意第绪语
	Yoruba                    LanguageCode = "yo"         // 约鲁巴语
	Zhuang                    LanguageCode = "za"         // 壮语
	Chinese                   LanguageCode = "zh"         // 中文
	Chinese_HongKong          LanguageCode = "zh-hk"      // 中文（香港）
	Chinese_Simplified        LanguageCode = "zh-cn"      // 中文（简体）
	Chinese_Taiwan            LanguageCode = "zh-tw"      // 中文（繁体）
	Zulu                      LanguageCode = "zu"         // 祖鲁语

	Chittagonian      LanguageCode = "ctg"    // 柴达木语
	Chittagonian_BD   LanguageCode = "ctg-bd" // 柴达木语（孟加拉国）
	Chittagonian_IN   LanguageCode = "ctg-in" // 柴达木语（印度）
	Hakka             LanguageCode = "hak"    // 客家语
	Hakka_CN          LanguageCode = "hak-cn" // 客家语（中国）
	KhumiChakma       LanguageCode = "kfr"    // 科弗拉语
	KhumiChakma_BD    LanguageCode = "kfr-bd" // 科弗拉语（孟加拉国）
	KhumiChakma_IN    LanguageCode = "kfr-in" // 科弗拉语（印度）
	Khasi             LanguageCode = "kha"    // 喀西语
	Khasi_IN          LanguageCode = "kha-in" // 喀西语（印度）
	Khasi_BD          LanguageCode = "kha-bd" // 喀西语（孟加拉国）
	Khasi_MV          LanguageCode = "kha-mv" // 喀西语（马尔代夫）
	KishoreganjChakma LanguageCode = "cja"    // 基什莫尔甘吉语
	Magahi            LanguageCode = "mag"    // 马加伊语
	Magahi_IN         LanguageCode = "mag-in" // 马加伊语（印度）
	Magahi_BD         LanguageCode = "mag-bd" // 马加伊语（孟加拉国）
	Magahi_MV         LanguageCode = "mag-mv" // 马加伊语（马尔代夫）
	Mundari           LanguageCode = "unr"    // 蒙达里语
	Mundari_IN        LanguageCode = "unr-in" // 蒙达里语（印度）
	Mundari_BD        LanguageCode = "unr-bd" // 蒙达里语（孟加拉国）
	Mundari_MV        LanguageCode = "unr-mv" // 蒙达里语（马尔代夫）
	Mundari_PK        LanguageCode = "unr-pk" // 蒙达里语（巴基斯坦）
	Mundari_NP        LanguageCode = "unr-np" // 蒙达里语（尼泊尔）
	Mundari_LK        LanguageCode = "unr-lk" // 蒙达里语（斯里兰卡）
	Mundari_BH        LanguageCode = "unr-bh" // 蒙达里语（巴林）
	Mundari_AE        LanguageCode = "unr-ae" // 蒙达里语（阿联酋）
	Mundari_OM        LanguageCode = "unr-om" // 蒙达里语（阿曼）
	Mundari_QA        LanguageCode = "unr-qa" // 蒙达里语（卡塔尔）
	Mundari_KW        LanguageCode = "unr-kw" // 蒙达里语（科威特）
	Mundari_SA        LanguageCode = "unr-sa" // 蒙达里语（沙特阿拉伯）
	Marwari           LanguageCode = "mwr"    // 马尔瓦里语
	Marwari_IN        LanguageCode = "mwr-in" // 马尔瓦里语（印度）
	Marwari_BD        LanguageCode = "mwr-bd" // 马尔瓦里语（孟加拉国）
	Marwari_MV        LanguageCode = "mwr-mv" // 马尔瓦里语（马尔代夫）
	Marwari_PK        LanguageCode = "mwr-pk" // 马尔瓦里语（巴基斯坦）
	Marwari_NP        LanguageCode = "mwr-np" // 马尔瓦里语（尼泊尔）
	Marwari_LK        LanguageCode = "mwr-lk" // 马尔瓦里语（斯里兰卡）
	Marwari_BH        LanguageCode = "mwr-bh" // 马尔瓦里语（巴林）
	Marwari_AE        LanguageCode = "mwr-ae" // 马尔瓦里语（阿联酋）
	Marwari_OM        LanguageCode = "mwr-om" // 马尔瓦里语（阿曼）
	Marwari_QA        LanguageCode = "mwr-qa" // 马尔瓦里语（卡塔尔）
	Marwari_KW        LanguageCode = "mwr-kw" // 马尔瓦里语（科威特）
	Marwari_SA        LanguageCode = "mwr-sa" // 马尔瓦里语（沙特阿拉伯）
	Rangpuri          LanguageCode = "rbb"    // 兰格普里语
	Rangpuri_IN       LanguageCode = "rbb-in" // 兰格普里语（印度）
	Rangpuri_BD       LanguageCode = "rbb-bd" // 兰格普里语（孟加拉国）
	Rangpuri_MV       LanguageCode = "rbb-mv" // 兰格普里语（马尔代夫）
	Rangpuri_PK       LanguageCode = "rbb-pk" // 兰格普里语（巴基斯坦）
	Rangpuri_NP       LanguageCode = "rbb-np" // 兰格普里语（尼泊尔）
	Rangpuri_LK       LanguageCode = "rbb-lk" // 兰格普里语（斯里兰卡）
	Rangpuri_BH       LanguageCode = "rbb-bh" // 兰格普里语（巴林）
	Rangpuri_AE       LanguageCode = "rbb-ae" // 兰格普里语（阿联酋）
	Rangpuri_OM       LanguageCode = "rbb-om" // 兰格普里语（阿曼）
	Rangpuri_QA       LanguageCode = "rbb-qa" // 兰格普里语（卡塔尔）
	Rangpuri_KW       LanguageCode = "rbb-kw" // 兰格普里语（科威特）
	Rangpuri_SA       LanguageCode = "rbb-sa" // 兰格普里语（沙特阿拉伯）
	Chhattisgarhi     LanguageCode = "hne"    // 查蒂斯加里语
	Chhattisgarhi_IN  LanguageCode = "hne-in" // 查蒂斯加里语（印度）
	Chhattisgarhi_BD  LanguageCode = "hne-bd" // 查蒂斯加里语（孟加拉国）
	Chhattisgarhi_MV  LanguageCode = "hne-mv" // 查蒂斯加里语（马尔代夫）
	Chhattisgarhi_PK  LanguageCode = "hne-pk" // 查蒂斯加里语（巴基斯坦）
	Chhattisgarhi_NP  LanguageCode = "hne-np" // 查蒂斯加里语（尼泊尔）
	Chhattisgarhi_LK  LanguageCode = "hne-lk" // 查蒂斯加里语（斯里兰卡）
	Chhattisgarhi_BH  LanguageCode = "hne-bh" // 查蒂斯加里语（巴林）
	Chhattisgarhi_AE  LanguageCode = "hne-ae" // 查蒂斯加里语（阿联酋）
	Chhattisgarhi_OM  LanguageCode = "hne-om" // 查蒂斯加里语（阿曼）
	Chhattisgarhi_QA  LanguageCode = "hne-qa" // 查蒂斯加里语（卡塔尔）
	Chhattisgarhi_KW  LanguageCode = "hne-kw" // 查蒂斯加里语（科威特）
	Chhattisgarhi_SA  LanguageCode = "hne-sa" // 查蒂斯加里语（沙特阿拉伯）
	Sylheti           LanguageCode = "syl"    // 锡尔赫蒂语
	Sylheti_IN        LanguageCode = "syl-in" // 锡尔赫蒂语（印度）
	Sylheti_BD        LanguageCode = "syl-bd" // 锡尔赫蒂语（孟加拉国）
	Sylheti_MV        LanguageCode = "syl-mv" // 锡尔赫蒂语（马尔代夫）
	Sylheti_PK        LanguageCode = "syl-pk" // 锡尔赫蒂语（巴基斯坦）
	Sylheti_NP        LanguageCode = "syl-np" // 锡尔赫蒂语（尼泊尔）
	Sylheti_LK        LanguageCode = "syl-lk" // 锡尔赫蒂语（斯里兰卡）
	Sylheti_BH        LanguageCode = "syl-bh" // 锡尔赫蒂语（巴林）
	Sylheti_AE        LanguageCode = "syl-ae" // 锡尔赫蒂语（阿联酋）
	Sylheti_OM        LanguageCode = "syl-om" // 锡尔赫蒂语（阿曼）
	Sylheti_QA        LanguageCode = "syl-qa" // 锡尔赫蒂语（卡塔尔）
	Sylheti_KW        LanguageCode = "syl-kw" // 锡尔赫蒂语（科威特）
	Sylheti_SA        LanguageCode = "syl-sa" // 锡尔赫蒂语（沙特阿拉伯）
	Chakma            LanguageCode = "ccp"    // 查克马语
	Chakma_IN         LanguageCode = "ccp-in" // 查克马语（印度）
	Chakma_BD         LanguageCode = "ccp-bd" // 查克马语（孟加拉国）
	Chakma_MV         LanguageCode = "ccp-mv" // 查克马语（马尔代夫）
	Chakma_PK         LanguageCode = "ccp-pk" // 查克马语（巴基斯坦）
	Chakma_NP         LanguageCode = "ccp-np" // 查克马语（尼泊尔）
	Chakma_LK         LanguageCode = "ccp-lk" // 查克马语（斯里兰卡）
	Chakma_BH         LanguageCode = "ccp-bh" // 查克马语（巴林）
	Chakma_AE         LanguageCode = "ccp-ae" // 查克马语（阿联酋）
	Chakma_OM         LanguageCode = "ccp-om" // 查克马语（阿曼）
	Chakma_QA         LanguageCode = "ccp-qa" // 查克马语（卡塔尔）
	Chakma_KW         LanguageCode = "ccp-kw" // 查克马语（科威特）
	Chakma_SA         LanguageCode = "ccp-sa" // 查克马语（沙特阿拉伯）
	Chiga             LanguageCode = "cgg"    // 奇加语
	Chiga_UG          LanguageCode = "cgg-ug" // 奇加语（乌干达）
	Cherokee          LanguageCode = "chr"    // 切罗基语
	Cherokee_US       LanguageCode = "chr-us" // 切罗基语（美国）
	Cherokee_IN       LanguageCode = "chr-in" // 切罗基语（印度）
	Cherokee_BD       LanguageCode = "chr-bd" // 切罗基语（孟加拉国）
	Cherokee_MV       LanguageCode = "chr-mv" // 切罗基语（马尔代夫）
	Cherokee_PK       LanguageCode = "chr-pk" // 切罗基语（巴基斯坦）
	Cherokee_NP       LanguageCode = "chr-np" // 切罗基语（尼泊尔）
	Cherokee_LK       LanguageCode = "chr-lk" // 切罗基语（斯里兰卡）
	Cherokee_BH       LanguageCode = "chr-bh" // 切罗基语（巴林）
	Cherokee_AE       LanguageCode = "chr-ae" // 切罗基语（阿联酋）
	Cherokee_OM       LanguageCode = "chr-om" // 切罗基语（阿曼）
	Cherokee_QA       LanguageCode = "chr-qa" // 切罗基语（卡塔尔）
	Cherokee_KW       LanguageCode = "chr-kw" // 切罗基语（科威特）
	Cherokee_SA       LanguageCode = "chr-sa" // 切罗基语（沙特阿拉伯）
	Church            LanguageCode = "cu"     // 教会斯拉夫语
	Church_RU         LanguageCode = "cu-ru"  // 教会斯拉夫语（俄罗斯）
	Chechen_RU        LanguageCode = "ce-ru"  // 车臣语（俄罗斯）
	Chechen_IN        LanguageCode = "ce-in"  // 车臣语（印度）
	Chechen_BD        LanguageCode = "ce-bd"  // 车臣语（孟加拉国）
	Chechen_MV        LanguageCode = "ce-mv"  // 车臣语（马尔代夫）
	Chechen_PK        LanguageCode = "ce-pk"  // 车臣语（巴基斯坦）
	Chechen_NP        LanguageCode = "ce-np"  // 车臣语（尼泊尔）

	Tripuri    LanguageCode = "trip"    // 三�
	Tripuri_IN LanguageCode = "trip-in" // 三�
	Tripuri_BD LanguageCode = "trip-bd" // 三�
	Tripuri_MV LanguageCode = "trip-mv" // 三�
	Tripuri_PK LanguageCode = "trip-pk" // 三�
	Tripuri_NP LanguageCode = "trip-np" // 三�
	Tripuri_LK LanguageCode = "trip-lk" // 三�
	Tripuri_BH LanguageCode = "trip-bh" // 三�
	Tripuri_AE LanguageCode = "trip-ae" // 三�
	Tripuri_OM LanguageCode = "trip-om" // 三�
	Tripuri_QA LanguageCode = "trip-qa" // 三�
	Tripuri_KW LanguageCode = "trip-kw" // 三�
	Tripuri_SA LanguageCode = "trip-sa" // 三�

	Pashto            LanguageCode = "ps"  // 普什图语
	Kikongo           LanguageCode = "kg"  // 基金�
	Kwanyama          LanguageCode = "kj"  // 基金�
	Lubwa             LanguageCode = "lu"  // 基金�
	Lunda             LanguageCode = "lun" // 基金�
	NyanekaNkhumbi    LanguageCode = "nnh" // 基金�
	NyanekaLuvale     LanguageCode = "ny"  // 基金�
	NyanekaLuchazi    LanguageCode = "nyl" // 基金�
	NyanekaChewa      LanguageCode = "nyn" // 基金�
	NyanekaLuchongwe  LanguageCode = "nyo" // 基金�
	NyanekaLukonde    LanguageCode = "nzi" // 基金�
	NyanekaLunyankole LanguageCode = "nyn" // 基金�
	NyanekaLutumbuka  LanguageCode = "nyt" // 基金�
	NyanekaLuchewa    LanguageCode = "nyn" // 基金�
	NyanekaLuchiluba  LanguageCode = "nyn" // 基金�

	Mapudungun LanguageCode = "arn" // 马普�
	Tehuelche  LanguageCode = "arn" // 马普�
	Wichi      LanguageCode = "arn" // 马普�

	Kirundi    LanguageCode = "rn"    // 基隆迪语
	Kirundi_BI LanguageCode = "rn-bi" // 基隆迪语（布隆迪）
	Kirundi_CD LanguageCode = "rn-cd" // 基隆迪语（刚果民主共和国）
	Kirundi_RW LanguageCode = "rn-rw" // 基隆迪语（卢旺达）
	Kirundi_TZ LanguageCode = "rn-tz" // 基隆迪语（坦桑尼亚）

	German_DE  LanguageCode = "de-de" // 德语（德国）
	German_CH  LanguageCode = "de-ch" // 德语（瑞士）
	German_AT  LanguageCode = "de-at" // 德语（奥地利）
	German_LU  LanguageCode = "de-lu" // 德语（卢森堡）
	German_LI  LanguageCode = "de-li" // 德语（列支敦士登）
	German_BE  LanguageCode = "de-be" // 德语（比利时）
	Papiamento LanguageCode = "pap"   // 帕皮亚门托语

	English_US     LanguageCode = "en-us"       // 英语（美国）
	English_GB     LanguageCode = "en-gb"       // 英语（英国）
	English_AU     LanguageCode = "en-au"       // 英语（澳大利亚）
	English_CA     LanguageCode = "en-ca"       // 英语（加拿大）
	English_NZ     LanguageCode = "en-nz"       // 英语（新西兰）
	English_IE     LanguageCode = "en-ie"       // 英语（爱尔兰）
	English_ZA     LanguageCode = "en-za"       // 英语（南非）
	English_JM     LanguageCode = "en-jm"       // 英语（牙买加）
	English_BZ     LanguageCode = "en-bz"       // 英语（伯利兹）
	English_TT     LanguageCode = "en-tt"       // 英语（特立尼达和多巴哥）
	English_ZW     LanguageCode = "en-zw"       // 英语（津巴布韦）
	English_PH     LanguageCode = "en-ph"       // 英语（菲律宾）
	English_IN     LanguageCode = "en-in"       // 英语（印度）
	English_MY     LanguageCode = "en-my"       // 英语（马来西亚）
	English_SG     LanguageCode = "en-sg"       // 英语（新加坡）
	English_HK     LanguageCode = "en-hk"       // 英语（香港）
	English_MO     LanguageCode = "en-mo"       // 英语（澳门）
	English_ID     LanguageCode = "en-id"       // 英语（印度尼西亚）
	English_TH     LanguageCode = "en-th"       // 英语（泰国）
	English_TR     LanguageCode = "en-tr"       // 英语（土耳其）
	English_VN     LanguageCode = "en-vn"       // 英语（越南）
	English_KR     LanguageCode = "en-kr"       // 英语（韩国）
	English_IL     LanguageCode = "en-il"       // 英语（以色列）
	English_MX     LanguageCode = "en-mx"       // 英语（墨西哥）
	English_CO     LanguageCode = "en-co"       // 英语（哥伦比亚）
	English_PE     LanguageCode = "en-pe"       // 英语（秘鲁）
	English_AR     LanguageCode = "en-ar"       // 英语（阿根廷）
	English_CL     LanguageCode = "en-cl"       // 英语（智利）
	English_VE     LanguageCode = "en-ve"       // 英语（委内瑞拉）
	English_EC     LanguageCode = "en-ec"       // 英语（厄瓜多尔）
	English_UY     LanguageCode = "en-uy"       // 英语（乌拉圭）
	English_PY     LanguageCode = "en-py"       // 英语（巴拉圭）
	English_PA     LanguageCode = "en-pa"       // 英语（巴拿马）
	English_CR     LanguageCode = "en-cr"       // 英语（哥斯达黎加）
	English_DO     LanguageCode = "en-do"       // 英语（多米尼加共和国）
	English_PR     LanguageCode = "en-pr"       // 英语（波多黎各）
	English_US_POS LanguageCode = "en-us-posix" // 英语（美国，POSIX）
	English_GB_POS LanguageCode = "en-gb-posix" // 英语（英国，POSIX）
	English_CA_POS LanguageCode = "en-ca-posix" // 英语（加拿大，POSIX）
	English_AU_POS LanguageCode = "en-au-posix" // 英语（澳大利亚，POSIX）
	English_NZ_POS LanguageCode = "en-nz-posix" // 英语（新西兰，POSIX）
	English_IE_POS LanguageCode = "en-ie-posix" // 英语（爱尔兰，POSIX）
	English_ZA_POS LanguageCode = "en-za-posix" // 英语（南非，POSIX）
	English_JM_POS LanguageCode = "en-jm-posix" // 英语（牙买加，POSIX）
	English_BZ_POS LanguageCode = "en-bz-posix" // 英语（伯利兹，POSIX）
	English_TT_POS LanguageCode = "en-tt-posix" // 英语（特立尼达和多巴哥，POSIX）
	English_ZW_POS LanguageCode = "en-zw-posix" // 英语（津巴布韦，POSIX）
	English_PH_POS LanguageCode = "en-ph-posix" // 英语（菲律宾，POSIX）
	English_IN_POS LanguageCode = "en-in-posix" // 英语（印度，POSIX）
	English_MY_POS LanguageCode = "en-my-posix" // 英语（马来西亚，POSIX）
	English_SG_POS LanguageCode = "en-sg-posix" // 英语（新加坡，POSIX）
	English_HK_POS LanguageCode = "en-hk-posix" // 英语（香港，POSIX）
	English_MO_POS LanguageCode = "en-mo-posix" // 英语（澳门，POSIX）
	English_TH_POS LanguageCode = "en-th-posix" // 英语（泰国，POSIX）
	English_ID_POS LanguageCode = "en-id-posix" // 英语（印度尼西亚，POSIX）
	English_TW_POS LanguageCode = "en-tw-posix" // 英语（台湾，POSIX）
	English_CN_POS LanguageCode = "en-cn-posix" // 英语（中国，POSIX）
	English_KR_POS LanguageCode = "en-kr-posix" // 英语（韩国，POSIX）
	English_JP_POS LanguageCode = "en-jp-posix" // 英语（日本，POSIX）
	English_CZ_POS LanguageCode = "en-cz-posix" // 英语（捷克，POSIX）
	English_NL_POS LanguageCode = "en-nl-posix" // 英语（荷兰，POSIX）
	English_BE_POS LanguageCode = "en-be-posix" // 英语（比利时，POSIX）
	English_DK_POS LanguageCode = "en-dk-posix" // 英语（丹麦，POSIX）
	English_FI_POS LanguageCode = "en-fi-posix" // 英语（芬兰，POSIX）
	English_GR_POS LanguageCode = "en-gr-posix" // 英语（希腊，POSIX）
	English_HU_POS LanguageCode = "en-hu-posix" // 英语（匈牙利，POSIX）
	English_NO_POS LanguageCode = "en-no-posix" // 英语（挪威，POSIX）
	English_PL_POS LanguageCode = "en-pl-posix" // 英语（波兰，POSIX）
	English_PT_POS LanguageCode = "en-pt-posix" // 英语（葡萄牙，POSIX）
	English_RO_POS LanguageCode = "en-ro-posix" // 英语（罗马尼亚，POSIX）
	English_RU_POS LanguageCode = "en-ru-posix" // 英语（俄罗斯，POSIX）
	English_SK_POS LanguageCode = "en-sk-posix" // 英语（斯洛伐克，POSIX）
	English_SE_POS LanguageCode = "en-se-posix" // 英语（瑞典，POSIX）
	English_CH_POS LanguageCode = "en-ch-posix" // 英语（瑞士，POSIX）
	English_AT_POS LanguageCode = "en-at-posix" // 英语（奥地利，POSIX）
	English_IL_POS LanguageCode = "en-il-posix" // 英语（以色列，POSIX）
	English_ES_POS LanguageCode = "en-es-posix" // 英语（西班牙，POSIX）
	English_MX_POS LanguageCode = "en-mx-posix" // 英语（墨西哥，POSIX）
	English_GT_POS LanguageCode = "en-gt-posix" // 英语（危地马拉，POSIX）
	English_CR_POS LanguageCode = "en-cr-posix" // 英语（哥斯达黎加，POSIX）
	English_PA_POS LanguageCode = "en-pa-posix" // 英语（巴拿马，POSIX）
	English_DO_Pos LanguageCode = "en-do-posix" // 英语（多米尼加共和国，POSIX）
	English_VE_POS LanguageCode = "en-ve-posix" // 英语（委内瑞拉，POSIX）
	English_CO_POS LanguageCode = "en-co-posix" // 英语（哥伦比亚，POSIX）
	English_PE_POS LanguageCode = "en-pe-posix" // 英语（秘鲁，POSIX）
	English_AR_POS LanguageCode = "en-ar-posix" // 英语（阿根廷，POSIX）
	English_CL_POS LanguageCode = "en-cl-posix" // 英语（智利，POSIX）
	English_UY_POS LanguageCode = "en-uy-posix" // 英语（乌拉圭，POSIX）
	English_PY_POS LanguageCode = "en-py-posix" // 英语（巴拉圭，POSIX）
	English_BO_POS LanguageCode = "en-bo-posix" // 英语（玻利维亚，POSIX）
	English_EC_POS LanguageCode = "en-ec-posix" // 英语（厄瓜多尔，POSIX）
	English_SV_POS LanguageCode = "en-sv-posix" // 英语（萨尔瓦多，POSIX）
	English_HN_POS LanguageCode = "en-hn-posix" // 英语（洪都拉斯，POSIX）
	English_NI_POS LanguageCode = "en-ni-posix" // 英语（尼加拉瓜，POSIX）
	English_PR_POS LanguageCode = "en-pr-posix" // 英语（波多黎各，POSIX）

	Spanish_ES     LanguageCode = "es-es"       // 西班牙语（西班牙）
	Spanish_MX     LanguageCode = "es-mx"       // 西班牙语（墨西哥）
	Spanish_GT     LanguageCode = "es-gt"       // 西班牙语（危地马拉）
	Spanish_CR     LanguageCode = "es-cr"       // 西班牙语（哥斯达黎加）
	Spanish_PA     LanguageCode = "es-pa"       // 西班牙语（巴拿马）
	Spanish_DO     LanguageCode = "es-do"       // 西班牙语（多米尼加共和国）
	Spanish_VE     LanguageCode = "es-ve"       // 西班牙语（委内瑞拉）
	Spanish_CO     LanguageCode = "es-co"       // 西班牙语（哥伦比亚）
	Spanish_PE     LanguageCode = "es-pe"       // 西班牙语（秘鲁）
	Spanish_AR     LanguageCode = "es-ar"       // 西班牙语（阿根廷）
	Spanish_EC     LanguageCode = "es-ec"       // 西班牙语（厄瓜多尔）
	Spanish_CL     LanguageCode = "es-cl"       // 西班牙语（智利）
	Spanish_UY     LanguageCode = "es-uy"       // 西班牙语（乌拉圭）
	Spanish_PY     LanguageCode = "es-py"       // 西班牙语（巴拉圭）
	Spanish_BO     LanguageCode = "es-bo"       // 西班牙语（玻利维亚）
	Spanish_SV     LanguageCode = "es-sv"       // 西班牙语（萨尔瓦多）
	Spanish_HN     LanguageCode = "es-hn"       // 西班牙语（洪都拉斯）
	Spanish_NI     LanguageCode = "es-ni"       // 西班牙语（尼加拉瓜）
	Spanish_PR     LanguageCode = "es-pr"       // 西班牙语（波多黎各）
	Spanish_US     LanguageCode = "es-us"       // 西班牙语（美国）
	Spanish_UY_POS LanguageCode = "es-uy-posix" // 西班牙语（乌拉圭，POSIX）
	Spanish_PY_POS LanguageCode = "es-py-posix" // 西班牙语（巴拉圭，POSIX）
	Spanish_AR_POS LanguageCode = "es-ar-posix" // 西班牙语（阿根廷，POSIX）
	Spanish_CL_POS LanguageCode = "es-cl-posix" // 西班牙语（智利，POSIX）
	Spanish_EC_POS LanguageCode = "es-ec-posix" // 西班牙语（厄瓜多尔，POSIX）
	Spanish_PE_POS LanguageCode = "es-pe-posix" // 西班牙语（秘鲁，POSIX）
	Spanish_CO_POS LanguageCode = "es-co-posix" // 西班牙语（哥伦比亚，POSIX）
	Spanish_VE_POS LanguageCode = "es-ve-posix" // 西班牙语（委内瑞拉，POSIX）
	Spanish_MX_POS LanguageCode = "es-mx-posix" // 西班牙语（墨西哥，POSIX）
	Spanish_GT_POS LanguageCode = "es-gt-posix" // 西班牙语（危地马拉，POSIX）
	Spanish_CR_POS LanguageCode = "es-cr-posix" // 西班牙语（哥斯达黎加，POSIX）
	Spanish_PA_POS LanguageCode = "es-pa-posix" // 西班牙语（巴拿马，POSIX）
	Spanish_DO_POS LanguageCode = "es-do-posix" // 西班牙语（多米尼加共和国，POSIX）
	Spanish_SV_POS LanguageCode = "es-sv-posix" // 西班牙语（萨尔瓦多，POSIX）
	Spanish_HN_POS LanguageCode = "es-hn-posix" // 西班牙语（洪都拉斯，POSIX）
	Spanish_NI_POS LanguageCode = "es-ni-posix" // 西班牙语（尼加拉瓜，POSIX）
	Spanish_PR_POS LanguageCode = "es-pr-posix" // 西班牙语（波多黎各，POSIX）
	Spanish_US_POS LanguageCode = "es-us-posix" // 西班牙语（美国，POSIX）

	Estonian_EE  LanguageCode = "et-ee"    // 爱沙尼亚语（爱沙尼亚）
	Estonian_POS LanguageCode = "et-posix" // 爱沙尼亚语（POSIX）

	French_FR     LanguageCode = "fr-fr"       // 法语（法国）
	French_CA     LanguageCode = "fr-ca"       // 法语（加拿大）
	French_BE     LanguageCode = "fr-be"       // 法语（比利时）
	French_CH     LanguageCode = "fr-ch"       // 法语（瑞士）
	French_LU     LanguageCode = "fr-lu"       // 法语（卢森堡）
	French_MC     LanguageCode = "fr-mc"       // 法语（摩纳哥）
	French_POS    LanguageCode = "fr-posix"    // 法语（POSIX）
	French_FR_POS LanguageCode = "fr-fr-posix" // 法语（法国，POSIX）
	French_CA_POS LanguageCode = "fr-ca-posix" // 法语（加拿大，POSIX）
	French_BE_POS LanguageCode = "fr-be-posix" // 法语（比利时，POSIX）
	French_CH_POS LanguageCode = "fr-ch-posix" // 法语（瑞士，POSIX）
	French_LU_POS LanguageCode = "fr-lu-posix" // 法语（卢森堡，POSIX）
	French_MC_POS LanguageCode = "fr-mc-posix" // 法语（摩纳哥，POSIX）

	Galician_ES       LanguageCode = "gl-es"       // 加利西亚语（西班牙）
	Galician_POS      LanguageCode = "gl-posix"    // 加利西亚语（POSIX）
	Galician_ES_POS   LanguageCode = "gl-es-posix" // 加利西亚语（西班牙，POSIX）
	Galician_ES_POSIX LanguageCode = "gl-es-posix" // 加利西亚语（西班牙，POSIX）

	Italian_IT     LanguageCode = "it-it"       // 意大利语（意大利）
	Italian_CH     LanguageCode = "it-ch"       // 意大利语（瑞士）
	Italian_POS    LanguageCode = "it-posix"    // 意大利语（POSIX）
	Italian_IT__OS LanguageCode = "it-it-posix" // 意大利语（意大利，POSIX）
	Italian_CH__OS LanguageCode = "it-ch-posix" // 意大利语（瑞士，POSIX）

	Japanese_JP  LanguageCode = "ja-jp"    // 日语（日本）
	Japanese_POS LanguageCode = "ja-posix" // 日语（POSIX）

	Georgian_GE  LanguageCode = "ka-ge"    // 格鲁吉亚语（格鲁吉亚）
	Georgian_POS LanguageCode = "ka-posix" // 格鲁吉亚语（POSIX）

	Korean_KR  LanguageCode = "ko-kr"    // 韩语（韩国）
	Korean_POS LanguageCode = "ko-posix" // 韩语（POSIX）

	Latvian_LV  LanguageCode = "lv-lv"    // 拉脱维亚语（拉脱维亚）
	Latvian_POS LanguageCode = "lv-posix" // 拉脱维亚语（POSIX）

	Lithuanian_LT  LanguageCode = "lt-lt"    // 立陶宛语（立陶宛）
	Lithuanian_POS LanguageCode = "lt-posix" // 立陶宛语（POSIX）

	Macedonian_MK  LanguageCode = "mk-mk"    // 马其顿语（马其顿）
	Macedonian_POS LanguageCode = "mk-posix" // 马其顿语（POSIX）

	Malay_MY     LanguageCode = "ms-my"       // 马来语（马来西亚）
	Malay_SG     LanguageCode = "ms-sg"       // 马来语（新加坡）
	Malay_POS    LanguageCode = "ms-posix"    // 马来语（POSIX）
	Malay_MY_POS LanguageCode = "ms-my-posix" // 马来语（马来西亚，POSIX）
	Malay_SG_POS LanguageCode = "ms-sg-posix" // 马来语（新加坡，POSIX）

	Norwegian_NO  LanguageCode = "no-no"    // 挪威语（挪威）
	Norwegian_POS LanguageCode = "no-posix" // 挪威语（POSIX）

	Polish_PL  LanguageCode = "pl-pl"    // 波兰语（波兰）
	Polish_POS LanguageCode = "pl-posix" // 波兰语（POSIX）

	Portuguese_PT     LanguageCode = "pt-pt"       // 葡萄牙语（葡萄牙）
	Portuguese_BR     LanguageCode = "pt-br"       // 葡萄牙语（巴西）
	Portuguese_POS    LanguageCode = "pt-posix"    // 葡萄牙语（POSIX）
	Portuguese_PT_POS LanguageCode = "pt-pt-posix" // 葡萄牙语（葡萄牙，POSIX）
	Portuguese_BR_POS LanguageCode = "pt-br-posix" // 葡萄牙语（巴西，POSIX）

	Romanian_RO  LanguageCode = "ro-ro"    // 罗马尼亚语（罗马尼亚）
	Romanian_POS LanguageCode = "ro-posix" // 罗马尼亚语（POSIX）

	Russian_RU  LanguageCode = "ru-ru"    // 俄语（俄罗斯）
	Russian_POS LanguageCode = "ru-posix" // 俄语（POSIX）

	Serbian_RS  LanguageCode = "sr-rs"    // 塞尔维亚语（塞尔维亚）
	Serbian_POS LanguageCode = "sr-posix" // 塞尔维亚语（POSIX）

	Slovak_SK  LanguageCode = "sk-sk"    // 斯洛伐克语（斯洛伐克）
	Slovak_POS LanguageCode = "sk-posix" // 斯洛伐克语（POSIX）

	Slovenian_SI  LanguageCode = "sl-si"    // 斯洛文尼亚语（斯洛文尼亚）
	Slovenian_POS LanguageCode = "sl-posix" // 斯洛文尼亚语（POSIX）

	ChineseHans    LanguageCode = "zh-hans"    // 中文（简体）
	ChineseHant    LanguageCode = "zh-hant"    // 中文（繁体）
	ChineseHans_HK LanguageCode = "zh-hans-hk" // 中文（简体，香港）
	ChineseHans_CN LanguageCode = "zh-hans-cn" // 中文（简体，中国）
	ChineseHant_TW LanguageCode = "zh-hant-tw" // 中文（繁体，台湾）
	ChineseHant_HK LanguageCode = "zh-hant-hk" // 中文（繁体，香港）
	ChineseHant_MO LanguageCode = "zh-hant-mo" // 中文（繁体，澳门）
	ChineseHant_CN LanguageCode = "zh-hant-cn" // 中文（繁体，中国）
	Chineseans_SG  LanguageCode = "zh-hans-sg" // 中文（简体，新加坡）
	Chineseans_MO  LanguageCode = "zh-hans-mo" // 中文（简体，澳门）
	Chineseans_TW  LanguageCode = "zh-hans-tw" // 中文（简体，台湾）
	ChineseHant_SG LanguageCode = "zh-hant-sg" // 中文（繁体，新加坡）
)
