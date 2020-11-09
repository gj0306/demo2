package bclient

//客户端

import (
	pb "demo2/dgrpc/protocol/bsgrpc/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

type GrpcClient struct {
	address string
	client *grpc.ClientConn
	greeterClient pb.StreamClient
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
		greeterClient:pb.NewStreamClient(conn),
	}
}


func (this *GrpcClient)Close() {
	_ = this.client.Close()
}

func (this *GrpcClient)Execute(){
	//调用服务端RouteList方法，获流
	stream, err := this.greeterClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}

