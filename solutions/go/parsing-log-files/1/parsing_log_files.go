package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, l := range lines {
		if re.MatchString(l) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	newLines := []string{}

	for _, l := range lines {
		if m := re.FindStringSubmatch(l); m != nil {
			newLines = append(newLines, fmt.Sprintf("[USR] %s %s", m[1], l))
		} else {
			newLines = append(newLines, l)
		}
	}

	return newLines
}
