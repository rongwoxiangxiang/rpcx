package util

import "strconv"

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return i
}

func StringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		i = 0
	}
	return i
}

func BoolenToInt8(bool2 bool) int8 {
	if bool2 {
		return 1
	} else {
		return 2
	}
}
