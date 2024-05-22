package http

import (
	"net/url"
	"regexp"
)

func isURLPatternMatch(u *url.URL, patterns []regexp.Regexp) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(u.String()) {
			return true
		}
	}
	return false
}
