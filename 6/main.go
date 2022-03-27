package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Основные способы остановки горутин:
	контекст:
		context.WithCancel - горутина будет остановлена после вызова ф-ции cancel;
		context.WithTimeout - надстройка над WithDeadline, оба способа останавливают горутину через N времени
	канал:
		передаем quit-канал в горутину, select'ом ожидаем закрытия канала
*/

func main() {

	ctxTimeout, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(2))

	ctxCancel, cancel := context.WithCancel(context.Background())

	quit := make(chan struct{})

	// context.WithCancel
	go func(ctx context.Context) {
		fmt.Println("timeout goroutine started")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timeout goroutine done")
				return
			}
		}
	}(ctxTimeout)

	// context.WithTimeout
	go func(ctx context.Context) {
		fmt.Println("cancel goroutine started")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel goroutine done")
				return
			}
		}
	}(ctxCancel)

	// quit-канал
	go func(quit chan struct{}) {
		fmt.Println("quit channel goroutine started")
		for {
			select {
			case <-quit:
				fmt.Println("quit channel goroutine done")
				return
			}
		}
	}(quit)

	time.Sleep(time.Second * time.Duration(3))
	cancel()
	close(quit)
	time.Sleep(time.Second * time.Duration(1))
}
