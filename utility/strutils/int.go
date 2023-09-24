package strutils

import "strconv"

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsUint(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
