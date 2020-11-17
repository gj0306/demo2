package goroutines

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"testing"
	"time"
)

func TestChannelBaseDemo(t *testing.T)  {
	/*
	// 可以发送接收int
	var c1 chan int
	// 只能发送int
	var c2 chan<- int
	// 只能从chan接收int
	var c3 <-chan int
	*/
	ch := make(chan int,10)
	send := func(ch chan<- int) {
		for i:=0;i<10;i++{
			ch <- i
		}
	}
	receive := func(ch <- chan int) {
		for {
			i := <-ch
			fmt.Println(i)
		}
	}
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second)
	fmt.Println("结束")
}

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


//不定长chan
func TestChannelSelectCase(t *testing.T)  {
	//创建SelectCase 函数
	createCases := func(chs ...chan int) []reflect.SelectCase {
		var cases []reflect.SelectCase

		// 创建recv case
		for _, ch := range chs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}

		// 创建send case
		for i, ch := range chs {
			v := reflect.ValueOf(i)
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(ch),
				Send: v,
			})
		}

		return cases
	}

	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建SelectCase
	var cases = createCases(ch1, ch2)

	// 执行10次select
	for i := 0; i < 10; i++ {
		//从 cases 中选择一个 case 执行
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

//数据传递
func TestMessaging(t *testing.T)  {
	type Token struct{}
	newWorker := func(id int, ch chan Token, nextCh chan Token) {
		for {
			token := <-ch         // 取得令牌
			fmt.Println((id + 1)) // id从1开始
			time.Sleep(time.Second)
			nextCh <- token
		}
	}

	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
	// 创建4个worker
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	//首先把令牌交给第一个worker
	chs[0] <- struct{}{}

	select {}
}

//信号通知
func TestSignal(t *testing.T)  {
	var closing = make(chan struct{})
	var closed = make(chan struct{})
	//清理动作
	doCleanup := func(closed chan struct{}) {
		time.Sleep((time.Minute))
		close(closed)
	}

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default:
				// ....... 业务计算
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("优雅退出")
}