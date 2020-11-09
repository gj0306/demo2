package server

//服务端
import (
	pb "demo2/dgrpc/protocol/advanced/protocol"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"runtime"
	"time"
	"google.golang.org/grpc/credentials"
)


type Server struct{}

type GrpcServer struct {
	port        string
	Server      *grpc.Server
	Credentials *credentials.TransportCredentials
}


//获取server
func GetGrpcServer(prot string)*grpc.Server{
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	return s
}

//普通grpc
func NewGrpcServer(port string)*GrpcServer{
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	return &GrpcServer{
		port:   port,
		Server: s,
	}
}
//TLC grpc
func NewTlcGrpcServer(c credentials.TransportCredentials,port string)*GrpcServer {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	return &GrpcServer{
		port:        port,
		Server:      s,
		Credentials: &c,
	}
}


func ListenPort(port string,server *grpc.Server)  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	server.Serve(lis)
}

//开启
func (this *GrpcServer)Start(){
	if this.Credentials != nil{
		this.Server = grpc.NewServer(grpc.Creds(*this.Credentials))
	}else {
		lis, err := net.Listen("tcp", this.port)
		if err != nil {
			log.Fatal("failed to listen: %v", err)
			return
		}
		go this.Server.Serve(lis)
	}
}

//关闭
func (this *GrpcServer)Stop(){
	this.Server.Stop()
}

func (this *GrpcServer)GracefulStop(){
	this.Server.GracefulStop()
}

//注册服务
func (this *GrpcServer)Register(){
	pb.RegisterGreeterServer(this.Server, &Server{})
}


/*server方法*/
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if ok{
		fmt.Printf("信息:%+v\n",md)
	}
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
//超时方法
func (s *Server)HandleDuration(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//获取认证信息
	data := make(chan *pb.HelloReply, 1)
	go func() {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			runtime.Goexit() //超时后退出该Go协程
		case <-time.After(4 * time.Second): // 模拟耗时操作
			res := pb.HelloReply{
				Message: "结果",
			}
			//修改数据库前进行超时判断
			if ctx.Err() == context.Canceled{
				//如果已经超时，则退出
				runtime.Goexit()
			}
			data <- &res
		}

	}()


	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
}

//自定义认证
func (s *Server)HandleAuth(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if ok{
		fmt.Printf("token信息:%+v\n",md)
	}
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

