package value_object

import (
	"regexp"
)

func ContainsSpecialChars(word string) bool {
	re := regexp.MustCompile(`[\W]`)
	return re.MatchString(word)
}
func ContainsNumbers(word string) bool {
	re := regexp.MustCompile(`[0-9]`)
	return re.MatchString(word)
}

func ContainsLetters(word string) bool {
	hasLowerCase := ContainsLowerCaseLetters(word)
	hasUpperCase := ContainsUpperCaseLetters(word)
	return hasLowerCase || hasUpperCase
}

func ContainsLowerCaseLetters(word string) bool {
	re := regexp.MustCompile(`[a-z]`)
	return re.MatchString(word)
}

func ContainsUpperCaseLetters(word string) bool {
	re := regexp.MustCompile(`[A-Z]`)
	return re.MatchString(word)
}
