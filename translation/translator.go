package translation

import (
	"log/slog"
	"strings"
)

func Translate(word string, language string) string {
	slog.Info("translating word", "word", word, "language", language)
	word = sanitize(word)
	language = sanitize(language)
	if word != "hello" {
		return ""
	}
	switch language {
	case "english":
		return "hello"
	case "french":
		return "bonjour"
	case "spanish":
		return "hola"
	case "german":
		return "hallo"
	default:
		return ""
	}
}

func sanitize(word string) string {
	return strings.TrimSpace(strings.ToLower(word))
}
