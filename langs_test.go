package langs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidSuccess(t *testing.T) {
	require.True(t, IsValid("en"))
}

func TestIsValidFailure(t *testing.T) {
	require.False(t, IsValid("foo"))
}

func TestNativeName(t *testing.T) {
	require.Equal(t, "English", NativeName(EN))
}

func TestLangString(t *testing.T) {
	require.Equal(t, EN.String(), "en")
}

func TestLangJSONMarshal(t *testing.T) {
	content, err := json.Marshal(ES)
	require.NoError(t, err)
	require.JSONEq(t, string(content), `"es"`)

	content, err = json.Marshal(map[Lang]string{
		ES: "es-content",
		EN: "en-content",
	})
	require.NoError(t, err)
	require.JSONEq(t, string(content), `{"en": "en-content", "es": "es-content"}`)
}

func TestLangJSONUnmarshal(t *testing.T) {
	var lang Lang
	require.NoError(t, json.Unmarshal([]byte(`"es"`), &lang))
	require.Equal(t, ES, lang)

	var content map[Lang]string
	require.NoError(t, json.Unmarshal([]byte(`{"en": "en-content", "es": "es-content"}`), &content))
	require.Equal(t, content, map[Lang]string{
		ES: "es-content",
		EN: "en-content",
	})
}

func TestLangParse(t *testing.T) {
	lang, _ := Parse("es")
	require.Equal(t, ES, lang)

	_, err := Parse("foo")
	require.EqualError(t, err, "unknown lang: foo")
}
