package solutions

import "strconv"

// must panics if err is not nil. Use for parsing operations where
// invalid input should be a fatal error.
func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// mustAtoi is a convenience wrapper for strconv.Atoi that panics on error.
func mustAtoi(s string) int {
	return must(strconv.Atoi(s))
}

// mustParseInt64 is a convenience wrapper for strconv.ParseInt that panics on error.
func mustParseInt64(s string) int64 {
	return must(strconv.ParseInt(s, 10, 64))
}
