package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var sum int
	nums := []int{2, 4, 6, 8, 10}

	wg.Add(len(nums))

	for _, num := range nums {
		go func(num int) {
			// используем мютекс, чтобы не получить гонку данных в sum
			mutex.Lock()
			defer mutex.Unlock()

			sum += num * num
			wg.Done()
		}(num)
	}

	wg.Wait()

	fmt.Println(sum)
}
