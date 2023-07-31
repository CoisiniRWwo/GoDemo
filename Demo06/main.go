package main

import (
	"Demo06/calc"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

func main() {
	result := calc.Add(1, 2)
	println(result)
	println(calc.Public)

	quantity := decimal.NewFromInt(3)

	fmt.Println(quantity)

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}
