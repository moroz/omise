package helpers

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Slugify(text string) string {
	transformer := runes.Remove(runes.In(unicode.Mn))
	chain := transform.Chain(norm.NFD, transformer, norm.NFC)
	deaccented, _, _ := transform.String(chain, text)
	lower := strings.ToLower(deaccented)
	lower = strings.ReplaceAll(lower, "ł", "l")
	return strings.ReplaceAll(lower, " ", "-")
}
