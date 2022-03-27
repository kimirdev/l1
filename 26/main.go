package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {
	str = strings.ToLower(str)

	set := make(map[rune]struct{})
	for _, el := range str {
		_, exist := set[el]
		if exist {
			return false
		}
		set[el] = struct{}{}
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(checkUnique(str1))
	fmt.Println(checkUnique(str2))
	fmt.Println(checkUnique(str3))
}
