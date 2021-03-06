package main

import "fmt"

// 使用通道传送通道

var counter = func(n int) chan<- chan<- int {
	//定义一个chan<-int 类型的chan
	request := make(chan chan<- int)
	go func() {
		for request := range request {
			if request == nil {
				n++ //递增计数
			} else {
				request <- n //返回当前
			}
		}
	}()
	return request //隐式转换到类型chan<-(chan<-int)
}(0)

func main() {
	increase100 := func(done chan<- struct{}) {
		for i := 0; i < 100; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}
	done := make(chan struct{})
	go increase100(done)
	go increase100(done)
	<-done
	<-done
	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request) // 200
}
