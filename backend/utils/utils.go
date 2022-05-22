package utils

import (
	"time"
)

// FormatUnixToUTCString returns a UTC, ISO-8601 formatted string of the given time.Time.
func FormatUnixToUTCString(i int64) string {
	return time.Unix(i, 0).UTC().Format(time.RFC3339)
}
