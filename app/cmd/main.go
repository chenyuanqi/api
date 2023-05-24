package main

import (
	"fmt"
	"time"
)

//生产数据
func send(ch chan int) {
	i := 0
	for i < 10 {
		ch <- i
		i++
	}
}

//接收并处理数据
func receive(ch chan int, rch chan string) {
	for {
		num := <-ch
		fmt.Println("接收到值：", num)
		rch <- fmt.Sprintf("处理后的编号：%d", num+1)
	}
}

func main() {
	ch := make(chan int)
	rch := make(chan string)
	go receive(ch, rch) //此时ch没有数据，子协程rec阻塞，直到其他协程向管道放入数据，阻塞解除
	go send(ch)
	go func(chan string) {
		for {
			fmt.Println("收到结果：", <-rch)
		}
	}(rch)

	time.Sleep(time.Second) //主协程等待1s
}
