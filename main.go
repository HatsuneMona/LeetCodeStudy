package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		for true {
			if _, ok := <-ch; ok != false {
				fmt.Println("终止！！！！！！！！！！！")
				runtime.Goexit()
			}
			fmt.Println("For循环")
		}
	}()
	time.Sleep(10000)
	//给一个信号
	ch <- struct{}{}
	for true{
		a:=1+1
		a++
	}
}
