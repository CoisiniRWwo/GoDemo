package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(x int) {
	for num := (x-1)*300000 + 1; num < x*300000; num++ {
		if num > 1 {
			flag := true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				//fmt.Println(num, "是素数")
			}
		}
	}
	defer wg.Done()
}

func main() {
	start := time.Now().Unix()
	for i := 1; i < 15; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
	fmt.Println("执行完毕")
}
