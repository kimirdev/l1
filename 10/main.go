package main

import "fmt"

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -60.3, 5.5, -5.5, 0, -10, -9.3, -11.3, -99.9, 99.9}

	// решил не использовать map, использую массив бакетов из 20 элементов
	// 0ой элемент - -90 градусов, 19ый элемент - +90 градусов
	var groups [20][]float32

	for i := 0; i < 20; i++ {
		groups[i] = make([]float32, 0)
	}

	for _, temp := range temps {
		if temp >= 100.0 || temp <= -100.0 {
			fmt.Println("Temperature range -99.9 - +99.9")
			return
		}
		// деление на 10 и отсечение нецелых частей (и +10), дает индекс массива
		t := int(temp / 10)
		groups[t+10] = append(groups[t+10], temp)
	}

	for i, gr := range groups {
		if len(gr) > 0 {
			fmt.Printf("%d:%v\n", (i-10)*10, gr)
		}
	}
}
