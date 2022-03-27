package main

import (
	"fmt"
	"time"
)

func CustomSleep(dur time.Duration) {
	finishAt := time.Now().Add(dur)

	for time.Now().Before(finishAt) {
	}
}

func main() {

	dur := time.Second * time.Duration(1)

	fmt.Printf("time.Sleep start: %d\n", time.Now().UnixMilli())
	time.Sleep(dur)
	fmt.Printf("time.Sleep over:  %d\n", time.Now().UnixMilli())

	fmt.Printf("CustomSleep start: %d\n", time.Now().UnixMilli())
	CustomSleep(dur)
	fmt.Printf("CustomSleep over:  %d\n", time.Now().UnixMilli())

}
