package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var seconds int

	fmt.Println("write seconds")
	_, err := fmt.Scan(&seconds)

	if err != nil {
		fmt.Println("wrong input")
	}

	data := make(chan string)

	// контекст закроет канал, после N секунд
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(seconds))

	var wg sync.WaitGroup
	wg.Add(2)

	// Горутина отправляет данные в канал data
	go func(data chan string, ctx context.Context) {
		for {
			select {
			// канал ожидает закрытия канала через N секунд
			case <-ctx.Done():
				fmt.Println("writer: done")
				wg.Done()
				return
			default:
				data <- time.Now().GoString()
			}
			time.Sleep(time.Millisecond * time.Duration(100))
		}
	}(data, ctx)

	// Горутина читает данные из канала data
	go func(data chan string, ctx context.Context) {
		for {
			select {
			// канал ожидает закрытия канала через N секунд
			case <-ctx.Done():
				fmt.Println("reader: done")
				wg.Done()
				return
			case msg := <-data:
				fmt.Println(msg)
			}
		}
	}(data, ctx)
	wg.Wait()
}
