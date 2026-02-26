package utils

func NotEmpty(values ...string) bool {
	for _, v := range values {
		if v == "" {
			return false
		}
	}
	return true
}
