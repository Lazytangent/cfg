package utils

import (
	"fmt"
	"strings"
)

const defaultLen = 80

func CreateDelimiter(msg string) string {
	if msg == "" || msg == "-" {
		return strings.Repeat("-", defaultLen)
	}

	n := len(msg)
	remaining := defaultLen - n
	halfLen := remaining / 2
	lenOfDash := halfLen - 1
	padding := halfLen * 2 < remaining

	left := strings.Repeat("-", lenOfDash)
	mid := fmt.Sprintf(" %s ", msg)
	right := strings.Repeat("-", lenOfDash)
	if padding {
		right = fmt.Sprintf("%s-", right)
	}

	return fmt.Sprintf("%s%s%s", left, mid, right)
}
