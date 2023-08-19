package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义一个远程调用的方法
type Hello struct {
}

// 方法只有两个可序列化的参数，其中第二个参数是指针类型
func (receiver Hello) SayHello(request string, response *string) error {
	*response = "你好" + request
	return nil
}

func main() {
	//注册RPC服务
	err1 := rpc.RegisterName("hello", Hello{})
	if err1 != nil {
		fmt.Println(err1)
	}
	//监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println(err2)
	}

	//应用推出的时候关闭监听端口
	defer listener.Close()

	for {
		fmt.Println("开始建立链接")
		//建立连接
		accept, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}

		//绑定服务
		rpc.ServeConn(accept)
	}
}
