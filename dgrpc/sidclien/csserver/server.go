package csserver

//服务端
import (
	pb "demo2/dgrpc/protocol/csgrpc/protocol"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

// SimpleService 定义我们的服务
type SimpleService struct{}

type GrpcServer struct {
	port string
	server *grpc.Server
}


//获取server
func GetGrpcServer(prot string)*grpc.Server{
	s := grpc.NewServer()
	pb.RegisterStreamClientServer(s, &SimpleService{})
	return s
}

//普通grpc
func NewGrpcServer(port string)*GrpcServer{
	s := grpc.NewServer()
	pb.RegisterStreamClientServer(s, &SimpleService{})
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
	_ = server.Serve(lis)
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

// RouteList 实现RouteList方法
func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.StreamData)
	}
}