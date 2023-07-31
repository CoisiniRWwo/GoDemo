package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	var num1 float64 = 8.2
	var num2 float64 = 3.8

	float1 := decimal.NewFromFloat(8.2)
	float2 := decimal.NewFromFloat(3.8)

	//decimal.New
	fmt.Printf("%v\n", num1-num2)
	fmt.Println(float1.Sub(float2))
}
