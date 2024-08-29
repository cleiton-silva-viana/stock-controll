package validation

import (
	"regexp"
	"strings"
	"time"
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

func DateFormatIsValid(date string) (bool, *time.Time) {
	layout := "2006-01-02"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		return false, nil
	}
	return true, &dateTime
}

func GetWithoutSpecialChars(words string, specialChars []string) string {
	var wordsReplaced = words
	for _, char := range specialChars {
		wordsReplaced = strings.ReplaceAll(wordsReplaced, char, "")
	}
	return wordsReplaced
}
