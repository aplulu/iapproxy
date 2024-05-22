package http

import (
	"net/url"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsURLPatternMatch(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		patterns []string
		expected bool
	}{
		{
			name:     "URL matches pattern",
			input:    "http://example.com",
			patterns: []string{"http://.*"},
			expected: true,
		},
		{
			name:     "URL does not match pattern",
			input:    "http://notmatching.com",
			patterns: []string{"https://.*"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := url.Parse(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse URL: %v", err)
			}

			var patterns []regexp.Regexp
			for _, pattern := range tc.patterns {
				re, err := regexp.Compile(pattern)
				if err != nil {
					t.Fatalf("Failed to compile pattern: %v", err)
				}
				patterns = append(patterns, *re)
			}

			result := isURLPatternMatch(u, patterns)
			assert.Equal(t, tc.expected, result)
		})
	}
}
