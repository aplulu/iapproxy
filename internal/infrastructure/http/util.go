package http

import (
	"regexp"
)

func isURLPatternMatch(url string, patterns []regexp.Regexp) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(url) {
			return true
		}
	}
	return false
}
