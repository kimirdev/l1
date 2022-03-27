package main

import (
	"fmt"
	"math/big"
)

// func main() {
// 	var a, b, res *big.Float = new(big.Float), new(big.Float), new(big.Float)

// 	a.SetString("3214343242356567576")

// 	b.SetString("1231244545635645646")

// 	fmt.Println(res.Add(a, b))
// 	fmt.Println(res.Sub(a, b))
// 	fmt.Println(res.Quo(a, b))
// 	fmt.Println(res.Mul(a, b))
// }

func main() {
	var a, b, res *big.Int = new(big.Int), new(big.Int), new(big.Int)

	a.SetString("3214343242356567576", 10)

	b.SetString("1231244545635645646", 10)

	fmt.Println(res.Add(a, b))
	fmt.Println(res.Sub(a, b))
	fmt.Println(res.Quo(a, b))
	fmt.Println(res.Mul(a, b))
}
