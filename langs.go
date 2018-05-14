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

func IsValid(lang string) bool {
	for _, l := range All {
		if l == lang {
			return true
		}
	}
	return false
}
