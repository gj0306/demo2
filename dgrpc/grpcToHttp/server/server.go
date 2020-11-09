package server

//服务端
import (
	pb "demo2/dgrpc/protocol/grpc/protocol"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type server struct{}

type GrpcServer struct {
	port string
	server *grpc.Server
}


//获取server
func GetGrpcServer(prot string)*grpc.Server{
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	return s
}

//普通grpc
func NewGrpcServer(port string)*GrpcServer{
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	return &GrpcServer{
		port:port,
		server:s,
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
	lis, err := net.Listen("tcp", this.port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
		return
	}
	go this.server.Serve(lis)
}
//关闭
func (this *GrpcServer)Stop(){
	this.server.Stop()
}

func (this *GrpcServer)GracefulStop(){
	this.server.GracefulStop()
}

/*server方法*/
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if ok{
		fmt.Printf("信息:%+v\n",md)
	}
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}