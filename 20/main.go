package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun abc cde eee"

	strs := strings.Split(str, " ")

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	fmt.Println(strings.Join(strs, " "))
}
