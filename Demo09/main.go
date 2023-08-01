package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 1200000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for val := range intChan {
		flag := true
		for i := 2; i < val; i++ {
			if val%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//fmt.Println(num, "是素数")
			primeChan <- val
		}
	}
	exitChan <- true
	wg.Done()
}

func printPrime(primeChan chan int) {
	for val := range primeChan {
		fmt.Println(val)
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16)

	wg.Add(1)
	go putNum(intChan)

	for i := 1; i < 15; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 1; i < 15; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
	fmt.Println("执行完毕")

}
