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
