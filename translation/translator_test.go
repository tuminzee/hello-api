package translation_test

import (
	"testing"

	"github.com/tuminzee/hello-api/translation"
)

func TestTranslate(t *testing.T) {

	tt := []struct {
		word        string
		language    string
		translation string
	}{
		{"hello", "English", "hello"},
		{"Hello", "English", "hello"},
		{"hello", "French", "bonjour"},
		{"hello", "Spanish", "hola"},
		{"hello ", "German", "hallo"},
		{"hello", "Unknown", ""},
		{"hello", "Hindi", ""},
		{"bye", "English", ""},
	}

	for _, tc := range tt {
		res := translation.Translate(tc.word, tc.language)
		if res != tc.translation {
			t.Errorf("Translate(%q, %q) = %q; want %q but got %q", tc.word, tc.language, res, tc.translation, res)
		}
	}
}
