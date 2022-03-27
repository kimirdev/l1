package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var workerCount int

	_, err := fmt.Scan(&workerCount)

	if err != nil {
		fmt.Println("wrong input")
		return
	}

	data := make(chan string)

	// контекст закроет канал Done, после вызова ф-ции cancel
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// запускаем воркеры в цикле
	for i := 0; i < workerCount; i++ {
		go func(ctx context.Context, data chan string, num int) {
			for {
				select {
				// канал ожидает вызова ф-ции cancel
				case <-ctx.Done():
					fmt.Printf("worker #%d: done\n", num)
					wg.Done()
					return
				case msg := <-data:
					fmt.Printf("worker #%d: %s\n", num, msg)
				default:
				}
			}
		}(ctx, data, i+1)
	}

	// ретранслируем входящий сигнал в канал signalCh
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		for {
			select {
			// канал ожидает системного вызова, для вызова ф-ции cancel()
			case <-signalCh:
				cancel()
				return
			default:
				data <- time.Now().GoString()
			}
		}
	}()

	wg.Wait()
}
