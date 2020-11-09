package bserver

//服务端
import (
	pb "demo2/dgrpc/protocol/bsgrpc/protocol"
	"google.golang.org/grpc"
	"io"
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
	pb.RegisterStreamServer(s, &StreamService{})
	return s
}

//普通grpc
func NewGrpcServer(port string)*GrpcServer{
	s := grpc.NewServer()
	pb.RegisterStreamServer(s, &StreamService{})
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

// Conversations 实现Conversations方法
func (s *StreamService) Conversations(srv pb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}