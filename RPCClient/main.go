package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//通过rpc.Dial和rpc微服务端建立连接
	dial, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	//当客户端退出的时候关闭连接
	defer dial.Close()

	//调用远程函数
	var reply string
	err2 := dial.Call("hello.SayHello", "客户端", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}

	//获取微服务返回的数据
	fmt.Println(reply)
}
