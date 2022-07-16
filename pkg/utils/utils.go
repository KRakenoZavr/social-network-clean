package utils

func Contains[T comparable](arr []T, search T) bool {
	for _, el := range arr {
		if el == search {
			return true
		}
	}
	return false
}
