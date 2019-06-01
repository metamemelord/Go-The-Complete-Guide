// Package programming enumerates various programming languages
package programming

// Language is the low level implementation of a programming language
type Language struct {
	Name string
	Code int
}

// GetLanguageAndCode returns Name and Code of a language
func (l Language) GetLanguageAndCode() (string, int) {
	return l.Name, l.Code
}

// Sum takes two integers and returns their sum
func Sum(a int, b int) int {
	return a + b
}
