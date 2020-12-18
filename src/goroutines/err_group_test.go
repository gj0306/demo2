package goroutines

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

/*
ErrGroup 分组执行一批相同的或类似的任务则是任务编排中一类情形
*/

//简单模式，返回第一个错误
func TestErrGroupOen(t *testing.T) {
	var g errgroup.Group
	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})
	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})

	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return nil
	})
	// 等待三个任务都完成
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}
}

//返回所有子任务的错误
func TestErrGroupTwo(t *testing.T) {
	var (
		g      errgroup.Group
		result []error
	)

	result = make([]error, 3)

	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		result[0] = nil // 保存成功或者失败的结果
		return nil
	})

	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")

		result[1] = errors.New("failed to exec #2") // 保存成功或者失败的结果
		return result[1]
	})

	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		result[2] = nil // 保存成功或者失败的结果
		return nil
	})

	if err := g.Wait(); err == nil {
		fmt.Printf("Successfully exec all. result: %v\n", result)
	} else {
		fmt.Printf("failed: %v\n", result)
	}

}

//任务流水线pipeline
func TestErrGroupPipeline(t *testing.T) {
	m, err := MD5All(context.Background(), ".")
	if err != nil {
		log.Fatal(err)
	}

	for k, sum := range m {
		fmt.Printf("%s:\t%x\n", k, sum)
	}
}

type result struct {
	path string
	sum  [md5.Size]byte
}

// 遍历根目录下所有的文件和子文件夹,计算它们的md5的值.
func MD5All(ctx context.Context, root string) (map[string][md5.Size]byte, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string) // 文件路径channel

	g.Go(func() error {
		defer close(paths) // 遍历完关闭paths chan
		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			//将文件路径放入到paths
			paths <- root
			return nil
		})
	})

	// 启动20个goroutine执行计算md5的任务，计算的文件由上一阶段的文件遍历子任务生成.
	c := make(chan result)
	const numDigesters = 20
	for i := 0; i < numDigesters; i++ {
		g.Go(func() error {
			for path := range paths { // 遍历直到paths chan被关闭
				//计算path的md5值，放入到c中
				c <- result{
					path: path,
				} //
			}
			return nil
		})
	}
	go func() {
		g.Wait() // 20个goroutine以及遍历文件的goroutine都执行完
		close(c) // 关闭收集结果的chan
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c { // 将md5结果从chan中读取到map中,直到c被关闭才退出
		m[r.path] = r.sum
	}

	// 再次调用Wait，依然可以得到group的error信息
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
}
