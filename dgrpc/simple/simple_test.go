package simple

import (
	"demo2/dgrpc/simple/client"
	"demo2/dgrpc/simple/server"
	"fmt"
	"testing"
	"time"

)

const (
	port = ":50051"
	address  = "localhost:50051"
)

func TestSimple(t *testing.T) {
	//启动服务端
	go func() {
		s1 := server.NewGrpcServer(port)
		s1.Start()
		time.Sleep(time.Second*3)
		//优雅退出
		s1.GracefulStop()
	}()
	time.Sleep(time.Second)
	//启动客户端
	go func() {
		c1 := client.NewGrpcAuthClient(address)
		for i:=0;i<3;i++{
			c1.Execute(i)
		}
		c1.Close()
	}()
	time.Sleep(time.Second*5)
	fmt.Println("结束")
}