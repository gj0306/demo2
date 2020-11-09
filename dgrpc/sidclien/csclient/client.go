package csclient

//客户端

import (
	pb "demo2/dgrpc/protocol/csgrpc/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

type GrpcClient struct {
	address string
	client *grpc.ClientConn
	greeterClient pb.StreamClientClient
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
		greeterClient:pb.NewStreamClientClient(conn),
	}
}


func (this *GrpcClient)Close() {
	_ = this.client.Close()
}

func (this *GrpcClient)Execute(){
	//调用服务端RouteList方法，获流
	stream, err := this.greeterClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}
	for n := 0; n < 5; n++ {
		//向流中发送消息
		err := stream.Send(&pb.StreamRequest{StreamData: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}

