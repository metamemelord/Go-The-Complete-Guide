package tests

import (
	"testing"

	programming "../src"
)

func TestSum(t *testing.T) {
	var a, b int = 10, 20
	result := programming.Sum(a, b)
	if result != 30 {
		t.Error("Result expected 30, got", result)
	}
}

func TestLangauge(t *testing.T) {
	var language programming.Language
	language = programming.Language{Name: "C++", Code: 1}
	if language.Name != "C++" {
		t.Error("Expected name to be C++, but got", language.Name)
	}
	if language.Code != 1 {
		t.Error("Expected code to be 1, but got", language.Code)
	}
}

func TestGetLanguageAndCode(t *testing.T) {
	var language programming.Language
	language = programming.Language{Name: "C++", Code: 1}
	lang, code := language.GetLanguageAndCode()
	if lang != "C++" || code != 1 {
		t.Errorf("Expected Language to be C++: 1 but got %s: %d", lang, code)
	}
}
