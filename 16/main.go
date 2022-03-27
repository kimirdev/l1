package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 2, 6, 3, 1, 8, 5, 9, 0}

	quickSort(nums)

	fmt.Println(nums)
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// в качестве pivot'a беру правый элемент
	left, right := 0, len(nums)-1

	// все числа меньше pivot'a отправляю в левую часть
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	// отправляю pivot на свое место
	nums[left], nums[right] = nums[right], nums[left]

	// саб-слайсы указывают на один и тот же массив, поэтому нет необходимости передавать целый слайс с индексами
	// рекурсивно иду в левую часть слайса
	quickSort(nums[:left])
	// рекурсивно иду в правую часть слайса
	quickSort(nums[left+1:])
}
