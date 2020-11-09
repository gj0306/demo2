package decorator

import (
	"fmt"
	"time"
)

type funcDemo func(int64) int64

//被装饰函数
func f1(i int64)int64{
	fmt.Println("函数执行中")
	return i*i
}

//装饰器
func Decorator(f funcDemo)func(int64) int64{
	fn := func(n int64)int64{
		fmt.Println("被装饰 函数运行前","参数:",n,"时间:",time.Now().Unix())
		var value int64
		defer func() {
			fmt.Println("被装饰 函数运行后","参数:",n,"时间:",time.Now().Unix(),"结果:",value)
		}()
		//结果
		value = f(n)
		return value
	}
	return fn
}