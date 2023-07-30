package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title    string
	Students []Student
}

func main() {
	str := `{"Title":"001","Students":[{"ID":0,"Gender":" 男 ","Name":"stu00"},{"ID":1,"Gender":" 男 ","Name":"stu01"},{"ID":2,"Gender":" 男 ","Name":"stu02"},{"ID":3,"Gender":" 男","Name":"stu03"},{"ID":4,"Gender":" 男 ","Name":"stu04"},{"ID":5,"Gender":" 男","Name":"stu05"},{"ID":6,"Gender":" 男 ","Name":"stu06"},{"ID":7,"Gender":" 男","Name":"stu07"},{"ID":8,"Gender":" 男 ","Name":"stu08"},{"ID":9,"Gender":" 男","Name":"stu09"}]}`
	c1 := &Class{}
	err := json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
