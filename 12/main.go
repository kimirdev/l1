package main

import "fmt"

func stringSet(strings []string) []string {
	set := make(map[string]struct{})
	ret := make([]string, 0)

	// закидываем все значения из массива в map
	for _, el := range strings {
		set[el] = struct{}{}
	}

	// достаем все ключи(уникальные) из map
	for key := range set {
		ret = append(ret, key)
	}
	return ret
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := stringSet(strings)

	fmt.Println(set)
}
