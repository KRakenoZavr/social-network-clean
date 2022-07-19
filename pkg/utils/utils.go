package utils

import "fmt"

func Contains[T comparable](arr []T, search T) bool {
	for _, el := range arr {
		if el == search {
			return true
		}
	}
	return false
}

func ArrayToSingle(arr []error) error {
	var err string
	for _, el := range arr {
		err += el.Error()
	}
	return fmt.Errorf(err)
}

func ErrArrayToStringArray(err []error) []string {
	var arrStr []string
	for _, el := range err {
		arrStr = append(arrStr, el.Error())
	}
	return arrStr
}
