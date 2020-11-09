package bidserver

import (
	"demo2/dgrpc/bidserver/bclient"
	"demo2/dgrpc/bidserver/bserver"
	"fmt"
	"testing"
	"time"
)

const (
	port = ":50051"
	address  = "localhost:50051"
)

func TestBidserver(t *testing.T) {
	//启动服务端
	go func() {
		s1 := bserver.NewGrpcServer(port)
		s1.Start()
		time.Sleep(time.Second*2)
		//优雅退出
		s1.GracefulStop()
	}()
	time.Sleep(time.Second)
	//启动客户端
	go func() {
		c1 := bclient.NewGrpcClient(address)
		c1.Execute()
		c1.Close()
	}()
	time.Sleep(time.Second*3)
	fmt.Println("结束")
}