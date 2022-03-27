package main

import (
	"fmt"
	"strings"
)

var justString string

var test string

func createHugeString(i int) string {
	test = "вфы"

	sb := strings.Builder{}

	sb.Grow(3 * i)

	for j := 0; j < i; j++ {
		sb.WriteString(test)
	}

	return sb.String()
}

func someFunc() {
	v := createHugeString(1 << 10)

	// Возможна потеря данных, если символы в строке не ASCII символ
	// justString = v[:101]

	justString = string([]rune(v)[:101])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
