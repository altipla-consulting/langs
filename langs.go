package langs

// Lang represents a language code.
type Lang string

// String returns the code of the language as string.
func (lang Lang) String() string {
	return string(lang)
}

const (
	CA = Lang("ca")
	DE = Lang("de")
	EN = Lang("en")
	ES = Lang("es")
	EU = Lang("eu")
	FR = Lang("fr")
	IT = Lang("it")
	JA = Lang("ja")
	PT = Lang("pt")
	RU = Lang("ru")
)

// All contains all the known languages of this library.
var All = []Lang{
	CA,
	DE,
	EN,
	ES,
	EU,
	FR,
	IT,
	JA,
	PT,
	RU,
}

var native = map[Lang]string{
	"CA": "Català",
	"DE": "Deutsch",
	"EN": "English",
	"ES": "Español",
	"EU": "Euskera",
	"FR": "Français",
	"IT": "Italiano",
	"JA": "日本語",
	"PT": "Portugues",
	"RU": "русский",
}

// IsValid checks if the lang code is a known one.
func IsValid(lang string) bool {
	for _, l := range All {
		if string(l) == lang {
			return true
		}
	}
	return false
}

// NativeName returns the native name of the language.
func NativeName(lang Lang) string {
	return native[lang]
}
