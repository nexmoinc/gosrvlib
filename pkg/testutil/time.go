package testutil

import (
	"regexp"
)

// ReplaceDateTime replaces a datetime value.
// This function is useful to compare JSON responses containing variable values.
func ReplaceDateTime(src, repl string) string {
	re := regexp.MustCompile("([0-9]{4}\\-[0-9]{2}\\-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}[^\"]*)")
	return re.ReplaceAllString(src, repl)
}

// ReplaceUnixTimestamp replaces a unix timestamp.
// This function is useful to compare JSON responses containing variable values.
func ReplaceUnixTimestamp(src, repl string) string {
	re := regexp.MustCompile("([0-9]{19})")
	return re.ReplaceAllString(src, repl)
}
