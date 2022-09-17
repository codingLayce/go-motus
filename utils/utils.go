package utils

func Contains[T comparable](list []T, element T) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}

	return false
}

func IndexOf(str []rune, ch rune, skip int) int {
	for i := 0; i < len(str); i++ {
		if str[i] == ch {
			if skip == 0 {
				return i
			}
			skip -= 1
		}
	}

	return -1
}
