package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	numChan := make(chan int)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(3))

	var wg sync.WaitGroup
	wg.Add(2)

	// горутина читает числа из канала numChan
	go func(numChan chan int, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case num := <-numChan:
				fmt.Printf("x:%d, x^2:%d\n", num, num*num)
			}
		}
	}(numChan, ctx)

	// горутина отправляет числа в канал numChan
	go func(numChan chan int, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				numChan <- rand.Intn(100)
			}
			time.Sleep(time.Millisecond * time.Duration(200))
		}
	}(numChan, ctx)

	wg.Wait()
}
