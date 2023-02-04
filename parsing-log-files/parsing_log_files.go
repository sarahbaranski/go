package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)

}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`\<[~*=-]*\>`).Split(text, -1)

}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)

	var result int
	for _, line := range lines {
		result += len(re.FindAllStringIndex(line, -1))
	}
	return result

}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			line = "[USR] " + match[1] + " " + line
		}
		lines[i] = line
	}
	return lines

}
