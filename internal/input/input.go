package input

import "strings"

func CleanInput(text string) []string {
	if text == "" {
		return []string{}
	}

	wsCleanedString := strings.Trim(text, " ")

	stringContents := []string{}

	for subString := range strings.SplitSeq(wsCleanedString, " ") {
		if len(subString) > 0 {
			stringContents = append(stringContents, strings.ToLower(subString))
		}
	}

	return stringContents
}
