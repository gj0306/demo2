package sserver

//服务端
import (
	pb "demo2/dgrpc/protocol/ssgrpc/protocol"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

// SimpleService 定义我们的服务
type StreamService struct{}

type GrpcServer struct {
	port string
	server *grpc.Server
}


//获取server
func GetGrpcServer(prot string)*grpc.Server{
	s := grpc.NewServer()
	pb.RegisterStreamServerServer(s, &StreamService{})
	return s
}

//普通grpc
func NewGrpcServer(port string)*GrpcServer{
	s := grpc.NewServer()
	pb.RegisterStreamServerServer(s, &StreamService{})
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

// ListValue 实现ListValue方法
func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}