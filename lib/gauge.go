package lib

import "fmt"

func PrintProgress(curr int, char rune) {
	const Max int = 100;
	var a = make([]rune, Max)
	a[0] = '['
	a[Max-1] = ']'
	for i := 1; i <= curr; i++ {
		if i > Max-5 {
			break
		}
		a[i] = char
	}
	str := string(a)
	switch {
	case len(str) < 80:

		str = "\x1B[0m\x1B[33m" + str + "\x1B[0m"
	case len(str) > 80:
		str = "\x1B[0m\x1B[31m" + str + "\x1B[0m"
	}
	fmt.Print("\033[2K\r", curr, str)
}
