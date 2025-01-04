package slice

import "slices"

func Remove[T any](arr []T, idx int) []T {
	var newArr []T
	for i, n := range arr {
		if i != idx {
			newArr = append(newArr, n)
		}
	}
	return newArr
}

func RemoveDuplicates[T comparable](arr []T) []T {
	var newArr []T
	for _, obj := range arr {
		if slices.Contains(newArr, obj) {
			continue
		} else {
			newArr = append(newArr, obj)
		}
	}
	return newArr
}
