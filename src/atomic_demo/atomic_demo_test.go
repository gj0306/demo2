package atomicdemo

import (
	"fmt"
	"go.uber.org/atomic"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestAtomic(t *testing.T)  {
	//创建一个 原子操作的 int64 变量
	num := atomic.NewInt64(0)
	//加减操作
	v := num.Add(1)
	fmt.Println(v)
	/*
	CAS方法
	比较当前 addr 地址里的值是不是 old，如果不等于 old，就返回 false；
	如果等于 old，就把此地址的值替换成 new 值，返回 true
	*/
	b := num.CAS(1,1)
	fmt.Println(b)
	/*
	Swap方法 替换原值
	*/
	num.Swap(3)
	/*
	Load 取出 addr 地址中的值 在多处理器、多核、有 CPU cache 的情况下，这个操作也能保证 Load 是一个原子操
	*/
	v = num.Load()
	fmt.Println(v)
	/*
		Store 把一个值存入到指定的 addr 地址中，即使在多处理器、多核、有 CPU cache 的情况下，这个操作也能保证 Store 是一个原子操作
		别的 goroutine 通过 Load 读取出来，不会看到存取了一半的值
	*/
	num.Store(5)

	/*
	Value
	原子地存取对象类型，但也只能存取，不能 CAS 和 Swap，常常用在配置变更等场景中
	*/
	var config atomic.Value
	config.Store("loadConf")
	var cond = sync.NewCond(&sync.Mutex{})
	go func() {
		// 设置新的config
		go func() {
			for {
				time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
				config.Store("loadConf")
				cond.Broadcast() // 通知等待着配置已变更
			}
		}()

		go func() {
			for {
				cond.L.Lock()
				cond.Wait()                 // 等待变更信号
				c := config.Load().(string) // 读取新的配置
				fmt.Printf("new config: %+v\n", c)
				cond.L.Unlock()
			}
		}()

	}()


}

