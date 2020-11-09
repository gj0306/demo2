package client

//客户端
import (
	pb "demo2/dgrpc/protocol/advanced/protocol"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
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

//带tlc的client
func NewGrpcTlcClient(c credentials.TransportCredentials,address string) *GrpcClient{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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

func (this *GrpcClient)Close() {
	_ = this.client.Close()
}

//
func (this *GrpcClient)Execute(){
	r, err := this.greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name: "姓名"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	fmt.Println("结果:",r.Message,"时间:",time.Now().Unix())

}

//设置超时时间
func (this *GrpcClient)DurationExecute(tm int64){
	ctx := context.Background()
	clientDeadline := time.Now().Add(time.Duration(time.Duration(tm) * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	// 传入超时时间为tm秒的ctx
	r, err := this.greeterClient.HandleDuration(ctx, &pb.HelloRequest{Name: "姓名"})
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout!")
			}
		}
		log.Fatalf("Call Route err: %v", err)
	}
	fmt.Println("结果:",r.Message,"时间:",time.Now().Unix())
}

//
func (this *GrpcClient)AuthExecute(){
	r, err := this.greeterClient.HandleAuth(context.Background(), &pb.HelloRequest{Name: "姓名"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	fmt.Println("结果:",r.Message,"时间:",time.Now().Unix())
}

