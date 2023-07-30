package main

type Animal struct {
	Name string
}

func (receiver Animal) run() {
	println("动物", receiver.Name, "在运动")
}

type Dog struct {
	Sound string
	Animal
}

func (dog Dog) Bark() {
	println(dog.Name, "狗在叫", dog.Sound)
}

func main() {
	doga := Dog{
		Sound: "汪汪汪",
		Animal: Animal{
			Name: "小明",
		},
	}

	doga.run()
	doga.Bark()
}
