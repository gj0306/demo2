package sidclien

import (
	csclient "demo2/dgrpc/sidclien/csclient"
	"demo2/dgrpc/sidclien/csserver"
	"fmt"
	"testing"
	"time"
)

const (
	port = ":50051"
	address  = "localhost:50051"
)

func TestSidclien(t *testing.T) {
	//启动服务端
	go func() {
		s1 := csserver.NewGrpcServer(port)
		s1.Start()
		time.Sleep(time.Second*2)
		//优雅退出
		s1.GracefulStop()
	}()
	time.Sleep(time.Second)
	//启动客户端
	go func() {
		c1 := csclient.NewGrpcClient(address)
		c1.Execute()
		c1.Close()
	}()
	time.Sleep(time.Second*3)
	fmt.Println("结束")
}