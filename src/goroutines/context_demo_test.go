package goroutines

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)
/*
Context 的调用应该是链式的，通过WithCancel，WithDeadline，WithTimeout或WithValue派生出新的 Context。
当父 Context 被取消时，其派生的所有 Context 都将取消。

通过context.WithXXX都将返回新的 Context 和 CancelFunc。调用 CancelFunc 将取消子代，
移除父代对子代的引用，并且停止所有定时器。未能调用 CancelFunc 将泄漏子代，直到父代被取消或定时器触发。
*/


func TestContextDemo(t *testing.T){
	/* 生成顶层 Context */

	//1 生成一个非nil 的、空的 Context
	ctx := context.Background()
	//2 生成一个非nil的、空的 Context，没有任何值
	ctx = context.TODO()

	//创建继承Background的子节点Context
	ctx1,cancle1 := context.WithCancel(ctx)
	ctx2,cancle2 := context.WithCancel(ctx)
	ctx3,_ := context.WithCancel(ctx)
	/* 1 监听context 是否关闭 */
	go t1(ctx1)
	time.Sleep(time.Second)
	cancle1()

	/* 2 context 增加上下文信息 */
	ch := make(chan string,1)
	ch <- "key"
	//添加上下文信息
	ctx2 = context.WithValue(ctx2, "key", "value")
	go t2(ctx2,ch)
	time.Sleep(time.Second)
	cancle2()


	/* 3 超时 Timeout WithTimeout   */
	ctx3,cf3 := context.WithTimeout(ctx3,time.Second*3)
	go t3(ctx3)
	go func() {
		//告诉操作放弃其工作
		time.Sleep(time.Second*5)
		cf3()
	}()


	/* 4 定义截至时间 Timeout WithDeadline   */
	ctx4,cf4 := context.WithDeadline(ctx3,time.Now().Add(time.Second))
	go t3(ctx4)
	go func() {
		//告诉操作放弃其工作
		time.Sleep(time.Second*5)
		cf4()
	}()


	time.Sleep(time.Second*3)

}

//监听context 是否关闭
func t1(ctx context.Context)  {
	for {
		select {
		case <- time.After(time.Second*2):
			fmt.Println("模拟操作，所用时间")
		case _ = <- ctx.Done():
			fmt.Println("context 已经被取消")
			//结束当前goroutine
			runtime.Goexit()
		}
	}
}
//context 增加上下文信息
func t2(ctx context.Context,ch chan string)  {
	for {
		if len(ch)== 0{
			break
		}
		k := <- ch
		val := ctx.Value(k)
		fmt.Println("context 传参 ",val)
	}
}

//context 设定超时时间，超时后会自己执行cancle()方法
func t3(ctx context.Context)  {
	for {
		select {
		case _ = <- ctx.Done():
			fmt.Println("context 已经超时被取消")
			//结束当前goroutine
			runtime.Goexit()
		default:

		}
	}
}