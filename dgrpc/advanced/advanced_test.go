package advanced

import (
	pb "demo2/dgrpc/protocol/advanced/protocol"
	"demo2/dgrpc/advanced/client"
	"demo2/dgrpc/advanced/server"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"testing"
	"time"
)

const (
	port = ":50051"
	address  = "localhost:50051"
)

func TestDuration(t *testing.T) {
	//启动服务端
	go func() {
		s1 := server.NewGrpcServer(port)
		s1.Start()
		time.Sleep(time.Second*3)
		//优雅退出
		s1.GracefulStop()
	}()
	time.Sleep(time.Second)
	//启动客户端
	go func() {
		c1 := client.NewGrpcAuthClient(address)
		c1.DurationExecute(1)
		c1.Close()
	}()
	time.Sleep(time.Second*10)
	fmt.Println("结束")
}

func TestTls(t *testing.T)  {
	path:= "D:\\work\\go\\demo2\\dgrpc\\advanced\\"

	//Tls服务器启动
	go func() {
		// 监听本地端口
		listener, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("net.Listen err: %v", err)
		}
		// 从输入证书文件和密钥文件为服务端构造TLS凭证
		creds, err := credentials.NewServerTLSFromFile(path + "server.pem", path + "server.key")
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		// 新建gRPC服务器实例,并开启TLS认证
		s1 := server.NewTlcGrpcServer(creds,port)
		s1.Start()
		//注册服务
		s1.Register()
		//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
		err = s1.Server.Serve(listener)
		if err != nil {
			log.Fatalf("grpcServer.Serve err: %v", err)
		}
	}()
	time.Sleep(time.Second)
	//tlc客户端
	go func() {
		//从输入的证书文件中为客户端构造TLS凭证
		creds, err := credentials.NewClientTLSFromFile(path + "server.pem", "go-grpc-example")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		c1 := client.NewGrpcTlcClient(creds,address)
		//执行
		c1.AuthExecute()
	}()
	time.Sleep(time.Second*3)
}

func TestToken(t *testing.T)  {
	//启动服务端
	go func() {
		s1 := server.NewGrpcServer(port)
		s1.Start()
		time.Sleep(time.Second*3)
		//优雅退出
		s1.GracefulStop()
	}()
	time.Sleep(time.Second)
	//启动客户端
	go func() {
		c1 := client.NewGrpcAuthClient(address)
		c1.Execute()
		c1.Close()
	}()
	time.Sleep(time.Second*5)
	fmt.Println("结束")
}

//一元拦截器 只会拦截简单RPC方法  流式拦截器 StreamInterceptor
func TestInterceptor(t *testing.T)  {
	// 监听本地端口
	listener, err := net.Listen("tpc", address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证Token
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			/* ...拦截方法... */
			if len(md["name"]) ==0{
				//失败
				return
			}
			fmt.Printf("token信息:%+v\n",md)
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	// 新建gRPC服务器实例,并开启TLS认证和Token认证
	//creds, err := credentials.NewServerTLSFromFile(path + "server.pem", path + "server.key")
	//if err != nil {
	//	log.Fatalf("Failed to generate credentials %v", err)
	//}
	//grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))

	// 新建gRPC服务器实例,并开启Token认证
	grpcServer := grpc.NewServer( grpc.UnaryInterceptor(interceptor))
	// 在gRPC服务器注册我们的服务
	pb.RegisterGreeterServer(grpcServer, &server.Server{})
	log.Println(address + " net.Listing whth TLS and token...")
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
