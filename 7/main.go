package main

import (
	"fmt"
	"sync"
)

type AtomicMap struct {
	// используем мьютекс, чтобы избежать гонки данных
	rwm    sync.RWMutex
	intMap map[int]int
}

func NewAtomicMap() *AtomicMap {
	return &AtomicMap{
		intMap: make(map[int]int),
	}
}

func (m *AtomicMap) Set(key, val int) {
	m.rwm.Lock()
	defer m.rwm.Unlock()

	m.intMap[key] = val
}

func (m *AtomicMap) Get(key int) (int, bool) {
	m.rwm.RLock()
	defer m.rwm.RUnlock()
	val, exist := m.intMap[key]
	return val, exist
}

func main() {
	amap := NewAtomicMap()

	amap.Set(0, 0)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			amap.Set(i, i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(amap.Get(i))
		}
		wg.Done()
	}()

	wg.Wait()
}
