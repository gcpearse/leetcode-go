package hard

import (
	"regexp"
)

func IsNumber(s string) bool {
	pattern := regexp.MustCompile(`(?i)^[+-]?\d*(\d+\.?|\.?\d+)(e[+-]?\d+)?$`)

	return pattern.MatchString(s)
}
