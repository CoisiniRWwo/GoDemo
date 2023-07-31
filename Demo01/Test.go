package main

import "fmt"

type SayInterface interface {
	say()
}
type MoveInterface interface {
	move()
}

// 接口嵌套
type Animal interface {
	SayInterface
	MoveInterface
}
type Cat struct {
	name string
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}
func (c Cat) move() {
	fmt.Println("猫会动")
}
func main() {
	var x Animal
	x = Cat{name: "花花"}
	x.move()
	x.say()
}
