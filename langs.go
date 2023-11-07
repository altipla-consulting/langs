package langs

import "fmt"

// Lang represents a language
type Lang struct {
	Code   string
	Native string
	Group  string
}

// String returns the code of the language as string.
func (lang Lang) String() string {
	return lang.Code
}

func (lang Lang) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, lang.Code)), nil
}

func (lang *Lang) UnmarshalJSON(b []byte) error {
	code := string(b[1 : len(b)-1])
	for _, l := range All {
		if l.Code == code {
			*lang = l
		}
	}
	return nil
}

func (lang Lang) MarshalText() ([]byte, error) {
	return []byte(lang.Code), nil
}

func (lang *Lang) UnmarshalText(text []byte) error {
	for _, l := range All {
		if l.Code == string(text) {
			*lang = l
		}
	}
	return nil
}

var (
	CA = Lang{Code: "ca", Native: "Català", Group: "ca"}
	DE = Lang{Code: "de", Native: "Deutsch", Group: "de"}
	EN = Lang{Code: "en", Native: "English", Group: "en"}
	ES = Lang{Code: "es", Native: "Español", Group: "es"}
	EU = Lang{Code: "eu", Native: "Euskera", Group: "eu"}
	FR = Lang{Code: "fr", Native: "Français", Group: "fr"}
	IT = Lang{Code: "it", Native: "Italiano", Group: "it"}
	JA = Lang{Code: "ja", Native: "日本語", Group: "ja"}
	PT = Lang{Code: "pt", Native: "Portugues", Group: "pt"}
	RU = Lang{Code: "ru", Native: "Русский", Group: "ru"}

	EsES = Lang{Code: "es-ES", Native: "Español", Group: "es"}
	EnGB = Lang{Code: "en-GB", Native: "English", Group: "en"}
	EnUS = Lang{Code: "en-US", Native: "English", Group: "en"}
	FrFR = Lang{Code: "fr-FR", Native: "Français", Group: "fr"}
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

	EsES,
	EnGB,
	EnUS,
	FrFR,
}

// IsValid checks if the lang code is a known one.
func IsValid(lang string) bool {
	for _, l := range All {
		if l.Code == lang {
			return true
		}
	}
	return false
}

// Parse returns the Lang for a given language.
func Parse(lang string) (Lang, error) {
	for _, l := range All {
		if l.Code == lang {
			return l, nil
		}
	}
	return Lang{}, fmt.Errorf("unknown lang: %s", lang)
}

// Deprecated: use Lang.Native instead.
func NativeName(lang Lang) string {
	return lang.Native
}
