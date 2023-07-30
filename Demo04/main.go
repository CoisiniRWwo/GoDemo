package main

import "strings"

func main() {
	str := "123-456-789"
	split := strings.Split(str, "-")
	println(split[0])
}
