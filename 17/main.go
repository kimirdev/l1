package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 6, 7, 8, 9, 13}

	fmt.Println(binarySearch(5, nums))
}

func binarySearch(num int, nums []int) bool {
	if len(nums) <= 1 {
		if nums[0] == num {
			return true
		}
		return false
	}

	mid := nums[len(nums)/2-1]
	if mid == num {
		return true
	} else if mid > num {
		return binarySearch(num, nums[:len(nums)/2])
	} else {
		return binarySearch(num, nums[len(nums)/2:])
	}
}
