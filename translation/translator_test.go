package translation_test

import (
	"testing"

	"github.com/tuminzee/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	word := "Hello"
	language := "French"
	res := translation.Translate(word, language)
	if res != word {
		t.Errorf("Translate(%q, %q) = %q; want %q", word, language, res, word)
	}
}
