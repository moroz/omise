package helpers_test

import (
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
