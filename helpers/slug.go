package helpers

import (
	"crypto/rand"
	"encoding/hex"
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

// / GenerateSlug converts a title to a kebab-case slug with a random
// / hexadecimal string at the end
func GenerateSlug(title string) string {
	randomBytes := make([]byte, 2)
	rand.Read(randomBytes)
	hexStr := hex.EncodeToString(randomBytes)
	return Slugify(title) + "-" + hexStr
}
