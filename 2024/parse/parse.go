package parse

import "strconv"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckError(err)
	return i
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
