package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// golang中空接口和类型断言使用细节
func main() {
	userinfo := make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 22
	userinfo["hobby"] = []string{"敲代码", "看番"}

	fmt.Println(userinfo["username"])
	fmt.Println(userinfo["age"])
	assertHobby, _ := userinfo["hobby"].([]string)
	fmt.Println(assertHobby[0])

	address := new(Address)
	address.Name = "李四"
	address.Phone = 110

	fmt.Println(address)

	userinfo["address"] = *address
	assertAddress, _ := userinfo["address"].(Address)
	fmt.Println(assertAddress.Name, assertAddress.Phone)
}
