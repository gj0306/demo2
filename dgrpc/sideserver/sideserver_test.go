package sideserver

import (
	sclient "demo2/dgrpc/sideserver/sclient"
	"demo2/dgrpc/sideserver/sserver"
	"fmt"
	"testing"
	"time"
)

const (
	port = ":50051"
	address  = "localhost:50051"
)

func TestSideserver(t *testing.T) {
	//启动服务端
	go func() {
		s1 := sserver.NewGrpcServer(port)
		s1.Start()
		time.Sleep(time.Second*2)
		//优雅退出
		s1.GracefulStop()
	}()
	time.Sleep(time.Second)
	//启动客户端
	go func() {
		c1 := sclient.NewGrpcClient(address)
		c1.Execute()
		c1.Close()
	}()
	time.Sleep(time.Second*3)
	fmt.Println("结束")
}