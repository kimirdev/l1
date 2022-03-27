package main

import "fmt"

func intersection(arr1, arr2 []int) []int {
	set := make(map[int]struct{})
	ret := make([]int, 0)

	// закидываем все числа из arr1 в map
	for _, el := range arr1 {
		set[el] = struct{}{}
	}

	// проходим по arr2 и находим пересечения
	for _, el := range arr2 {
		_, exist := set[el]
		if exist {
			ret = append(ret, el)
		}
	}
	return ret
}

func main() {
	arr1 := []int{1, 3, 2, 5, 6, 9, 13}
	arr2 := []int{13, 8, 15, 7, 6, 2}

	arr3 := intersection(arr1, arr2)
	fmt.Println(arr3)
}
