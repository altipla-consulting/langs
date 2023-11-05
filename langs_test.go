package langs

import (
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
