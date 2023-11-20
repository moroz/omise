package helpers_test

import (
	"strings"
	"testing"

	"github.com/moroz/omise/helpers"
)

func TestSlugify(t *testing.T) {
	src := "Zażółć gęślą jaźń"
	actual := helpers.Slugify(src)
	expected := "zazolc-gesla-jazn"
	if actual != expected {
		t.Errorf("Expected %s to be slugified to %s, got %s", src, expected, actual)
	}
}

func TestGenerateSlug(t *testing.T) {
	src := "Zażółć gęślą jaźń"
	actual := helpers.GenerateSlug(src)
	prefix := "zazolc-gesla-jazn"
	if !strings.HasPrefix(actual, prefix) {
		t.Errorf("Expected %s to begin with %s", actual, prefix)
	}
}
