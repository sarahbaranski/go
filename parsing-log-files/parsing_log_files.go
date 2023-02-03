package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	logArray := [6]string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	match, _ := regexp.MatchString(`^[][A-Z]{3}`, text)
	if !match {
		return match
	}
	input := strings.Split(text, " ")
	for _, log := range logArray {
		if input[0] == log {
			return true
		}
	}
	return false

}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`\<[~*=-]*\>`).Split(text, -1)

}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)\".*password.*\"`)

	var result int
	for _, line := range lines {
		result += len(re.FindAllStringIndex(line, -1))
	}
	return result

}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-ofline\d+`)
	return re.ReplaceAllLiteralString(text, "")

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
