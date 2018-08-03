package langs

const (
	CA = "ca"
	DE = "de"
	EN = "en"
	ES = "es"
	FR = "fr"
	IT = "it"
	JA = "ja"
	PT = "pt"
	RU = "ru"
)

var All = []string{
	CA,
	DE,
	EN,
	ES,
	FR,
	IT,
	JA,
	PT,
	RU,
}

var native = map[string]string{
	"CA": "català",
	"DE": "deutsch",
	"EN": "english",
	"ES": "español",
	"FR": "français",
	"IT": "italiano",
	"JA": "日本" ,
	"PT": "portugues",
	"RU": "русский",
}

func IsValid(lang string) bool {
	for _, l := range All {
		if l == lang {
			return true
		}
	}
	return false
}

func NativeNames(lang string) string {
	return native[lang]
}
