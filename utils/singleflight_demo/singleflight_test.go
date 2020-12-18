package singleflight_demo


import (
	"fmt"
"sync"
"sync/atomic"
	"testing"
	"time"

"golang.org/x/sync/singleflight"
)

var (
	count = int64(0)
)

// 模拟接口方法
func f() (interface{}, error) {
	time.Sleep(time.Millisecond * 500)
	return atomic.AddInt64(&count, 1), nil
}

func TestSingleflight(t *testing.T)  {
	//创建一个SingleFlight
	g := singleflight.Group{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			//Do 传入key，以及回调函数，如果key相同，fn方法只会执行一次，同步等待
			val, err, shared := g.Do("f1", f)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("index: %d, val: %d, shared: %v\n", j, val, shared)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			select {
			//DoChan 类似Do方法，只是返回的是一个chan
			case v := <-g.DoChan("f2", f):
				fmt.Println(i,v)
			case <-time.After(time.Millisecond*1000):
				fmt.Println("超时")
			}
		}(i)
	}
	//调用这个方法的 goroutine 会一直阻塞，直到 WaitGroup 的计数值变为 0
	wg.Wait()

	//Forget 控制key关联的值是否失效
	g.Forget("f1")

}

