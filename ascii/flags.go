package web

import "strings"

func SpecialCharacters(str string) (string, bool) {
	// Handle newline character
	if str == "\\n" {
		return "", false
	}

	// Handle special characters
	if strings.Contains(str, "\\a") {
		return "Error: Bell Character", true
	}
	if strings.Contains(str, "\\v") {
		return "Error: Vertical tab character", true
	}
	if strings.Contains(str, "\\f") {
		return "Error: Form feed character", true
	}
	if strings.Contains(str, "\\r") {
		return "Error: Carriage return character", true
	}

	// Handle tab characters
	str = strings.ReplaceAll(str, "\\t", "    ")

	// Handle backspace tabs
	str = strings.ReplaceAll(str, "\\b", "\b")

	for {
		index := strings.Index(str, "\b")

		if index == -1 {
			break
		}
		if index > 0 {
			str = str[:index-1] + str[index+1:]
		} else {
			str = str[index+1:]
		}
	}

	return str, false
}
