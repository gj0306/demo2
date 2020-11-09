package client

//客户端

import (
	pb "demo2/dgrpc/protocol/grpc/protocol"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

// 实现grpc.PerRPCCredentials接⼝
type Authentication struct {
	Login    string
	Password string
}

//获取当前请求认证所需的元数据
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"login": a.Login, "password": a.Password}, nil
}
//是否需要基于 TLS 认证进行安全传输
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}


type GrpcClient struct {
	address string
	client *grpc.ClientConn
	greeterClient pb.GreeterClient
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
		greeterClient:pb.NewGreeterClient(conn),
	}
}

//带认证的Client
func NewGrpcAuthClient(address string)*GrpcClient{
	auth := Authentication{
		Login:    "gopher",
		Password: "password",
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil{
		fmt.Printf("NewGrpcAuthServer err:%s \n",err.Error())
		return nil
	}
	return &GrpcClient{
		address:address,
		client:conn,
		greeterClient:pb.NewGreeterClient(conn),
	}
}

//

func (this *GrpcClient)Execute(id int){
	switch id {
	case 0:
	default:
		r, err := this.greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name: "姓名:"+strconv.Itoa(id)})
		if err != nil {
			log.Fatal("could not greet: %v", err)
			return
		}
		fmt.Println("结果:",r.Message,"时间:",time.Now().Unix())
	}
}

func (this *GrpcClient)Close() {
	_ = this.client.Close()
}