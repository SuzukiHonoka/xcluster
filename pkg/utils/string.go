package utils

func Empty(s string) bool {
	return s == ""
}

func EmptyAny(ss []string) bool {
	for _, s := range ss {
		if Empty(s) {
			return true
		}
	}
	return false
}
