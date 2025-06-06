package langs

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	CA   = Lang{Code: "ca", Native: "Català", Group: "ca"}
	DE   = Lang{Code: "de", Native: "Deutsch", Group: "de"}
	DeDE = Lang{Code: "de-DE", Native: "Deutsch", Group: "de"}
	EN   = Lang{Code: "en", Native: "English", Group: "en"}
	EnGB = Lang{Code: "en-GB", Native: "English", Group: "en"}
	EnUS = Lang{Code: "en-US", Native: "English", Group: "en"}
	ES   = Lang{Code: "es", Native: "Español", Group: "es"}
	EsES = Lang{Code: "es-ES", Native: "Español", Group: "es"}
	EU   = Lang{Code: "eu", Native: "Euskera", Group: "eu"}
	FR   = Lang{Code: "fr", Native: "Français", Group: "fr"}
	FrFR = Lang{Code: "fr-FR", Native: "Français", Group: "fr"}
	IT   = Lang{Code: "it", Native: "Italiano", Group: "it"}
	ItIT = Lang{Code: "it-IT", Native: "Italiano", Group: "it"}
	JA   = Lang{Code: "ja", Native: "日本語", Group: "ja"}
	PT   = Lang{Code: "pt", Native: "Portugues", Group: "pt"}
	PtPT = Lang{Code: "pt-PT", Native: "Portugues", Group: "pt"}
	RU   = Lang{Code: "ru", Native: "Русский", Group: "ru"}

	Empty = Lang{}
)

// All contains all the known languages of this library.
var All = []Lang{
	CA,
	DE,
	DeDE,
	EN,
	EnGB,
	EnUS,
	ES,
	EsES,
	EU,
	FR,
	FrFR,
	IT,
	ItIT,
	JA,
	PT,
	PtPT,
	RU,
}

var _ json.Marshaler = Lang{}
var _ json.Unmarshaler = (*Lang)(nil)
var _ encoding.TextMarshaler = Lang{}
var _ encoding.TextUnmarshaler = (*Lang)(nil)
var _ sql.Scanner = (*Lang)(nil)
var _ driver.Valuer = Lang{}

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
		if strings.EqualFold(l.Code, code) {
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
		if strings.EqualFold(l.Code, string(text)) {
			*lang = l
		}
	}
	return nil
}

// Value implements driver.Valuer.
func (lang Lang) Value() (driver.Value, error) {
	return lang.Code, nil
}

// Scan implements sql.Scanner.
func (lang *Lang) Scan(src any) error {
	if src == nil {
		return fmt.Errorf("langs: cannot scan nil into %T", lang)
	}
	switch src := src.(type) {
	case []byte:
		return lang.UnmarshalText(src)
	case string:
		return lang.UnmarshalText([]byte(src))
	}
	return fmt.Errorf("langs: cannot scan %T into %T", src, lang)
}

// IsValid checks if the lang code is a known one.
func IsValid(lang string) bool {
	for _, l := range All {
		if strings.EqualFold(l.Code, lang) {
			return true
		}
	}
	return false
}

// Empty returns if the lang is empty.
func (l Lang) Empty() bool {
	return l == Empty
}

// Parse returns the Lang for a given language.
func Parse(lang string) (Lang, error) {
	for _, l := range All {
		if strings.EqualFold(l.Code, lang) {
			return l, nil
		}
	}
	return Lang{}, fmt.Errorf("langs: unknown code %q", lang)
}

// Deprecated: use Lang.Native instead.
func NativeName(lang Lang) string {
	return lang.Native
}
