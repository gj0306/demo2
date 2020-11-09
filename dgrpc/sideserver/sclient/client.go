package sclient

//客户端

import (
	pb "demo2/dgrpc/protocol/ssgrpc/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
)


type GrpcClient struct {
	address string
	client *grpc.ClientConn
	greeterClient pb.StreamServerClient
}

func NewGrpcClient(address string)*GrpcClient{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
		return nil
	}
	return &GrpcClient{
		address:address,
		client:conn,
		greeterClient:pb.NewStreamServerClient(conn),
	}
}


func (this *GrpcClient)Close() {
	_ = this.client.Close()
}

func (this *GrpcClient)Execute(){
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc ",
	}
	// 调用我们的服务(ListValue方法)
	stream, err := this.greeterClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)
	}
}

