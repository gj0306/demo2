package goroutines

import (
	"fmt"
	"testing"
	"time"
)


func TestChannelDemo(t *testing.T)  {
	send := func(ch chan int, begin int) {
		// 循环向 channel 发送消息
		for i :=begin ; i< begin + 10 ;i++{
			ch <- i
		}
	}
	receive := func(ch <-chan int) {
		val := <- ch
		fmt.Println("receive:", val)
	}


	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 0)
	go receive(ch2)
	// 主 goroutine 休眠 1s，保证调度成功
	time.Sleep(time.Second)
	for {
		select {
		case val := <- ch1: // 从 ch1 读取数据
			fmt.Printf("get value %d from ch1\n", val)
		case ch2 <- 2 : // 使用 ch2 发送消息
			fmt.Println("send value by ch2")
		case <-time.After(2 * time.Second): // 超时设置
			fmt.Println("Time out")
			return
		}
	}

}
