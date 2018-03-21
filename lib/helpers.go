package lib

import "strings"

func Intptr(i int) *int {
	tmp := i
	return &tmp
}

func Capitalize(s *string) *string {
	if s == nil {
		return s
	}
	var newStr = strings.Title(*s)
	return &newStr
}
