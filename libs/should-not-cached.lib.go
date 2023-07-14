package libs

import "strings"

func ShouldNotCached(URL string) bool {
	switch {
	case strings.Contains(URL, "auth"):
		return true
	case strings.Contains(URL, "user/profile"):
		return true
	default:
		return false
	}
}
