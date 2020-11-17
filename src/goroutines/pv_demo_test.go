package goroutines

import (
	"context"
	"github.com/marusama/semaphore/v2"
	"testing"
	"time"
)

func TestPvDome(t *testing.T)  {
	//新建一个5的信号量
	sem := semaphore.New(5)
	ctx := context.Background()

	/* 获取信号量 */

	//用上下文获取n
	_ = sem.Acquire(ctx, 1)
	//无阻塞信号量
	_ = sem.TryAcquire(1)
	//超时信号量
	ctx1,_ := context.WithTimeout(context.Background(), time.Second)
	_ = sem.Acquire(ctx1, 1)

	/* 释放信号量 */
	sem.Release(1)


	/* 修改信号量数量 */
	sem.SetLimit(6)

}