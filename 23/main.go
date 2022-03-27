package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 4

	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice)
}
