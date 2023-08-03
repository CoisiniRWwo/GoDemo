package main

import (
	"fmt"
	"io"
	"os"
)

// 自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (err error) {
	source, _ := os.Open(srcFileName)
	destination, _ := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	buf := make([]byte, 128)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return
}

func main() {
	//调用 CopyFile 完成文件拷贝
	srcFile := "c:/000.avi"
	dstFile := "D:/000.avi"
	err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}
