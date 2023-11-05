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

// IsValid checks if the lang code is a known one.
func IsValid(lang string) bool {
	for _, l := range All {
		if string(l.Code) == lang {
			return true
		}
	}
	return false
}

// NativeName returns the native name of the language.
func NativeName(lang string) (string, error) {
	for _, l := range All {
		if string(l.Code) == lang {
			return l.Native, nil
		}
	}
	return "", fmt.Errorf("unknown lang %q", lang)
}

// LangGroup returns the group of the language.
func LangGroup(lang string) (string, error) {
	for _, l := range All {
		if string(l.Code) == lang {
			return l.Group, nil
		}
	}
	return "", fmt.Errorf("unknown lang %q", lang)
}
