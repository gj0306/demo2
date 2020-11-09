package snowflake

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateIdBuilder(t *testing.T)  {
	num := 12
	ch := make(chan int64,num)
	ids := make([]int64,0,12*100)
	f := func(ch chan int64) {
		select {
		case sn := <- ch:
			idf := CreateIdBuilder(sn)
			for i:=0;i<100;i++{
				id := idf()
				ids = append(ids, id)
			}
		}
	}
	for i:=1;i<=num;i++{
		go f(ch)
	}
	for i:=1;i<=num;i++{
		ch <- int64(i)
	}
	time.Sleep(time.Second*3)
	fmt.Println(ids)
}
