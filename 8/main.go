package main

import (
	"fmt"
	"strconv"
)

func SetBit(num int64, i int, one bool) int64 {
	i--
	if one {
		num = num | 1<<i
	} else {
		act := int64(1 << i)
		num = num & ^(act)
	}
	return num
}

func main() {
	var num int64
	var bit, oneZero int

	fmt.Println("write number: ")

	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("wrong input")
		return
	}

	//fmt.Println(strconv.FormatInt(-1, 2))
	fmt.Println("write bit position:")
	_, err = fmt.Scan(&bit)

	if err != nil || bit < 1 || bit > 64 {
		fmt.Println("wrong input")
		return
	}

	fmt.Println("1 or 0:")
	_, err = fmt.Scan(&oneZero)

	if err != nil || (oneZero != 1 && oneZero != 0) {
		fmt.Println("wrong input")
		return
	}

	if oneZero == 1 {
		num = SetBit(num, bit, true)

		fmt.Printf("Binary: %s\nDec: %d\n", strconv.FormatInt(num, 2), num)
	} else if oneZero == 0 {
		num = SetBit(num, bit, false)

		fmt.Printf("Binary: %s\nDec: %d\n", strconv.FormatInt(num, 2), num)
	}
}
