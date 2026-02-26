package utils

import "strconv"

func MustAtoi(v string) int {
	n, _ := strconv.Atoi(v)
	return n
}
