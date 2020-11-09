package midd


/*
把gRPC中的panic转成error，从而恢复程序
*/

//服务端
/*
1.直接把grpc_recovery拦截器添加到服务端

grpcServer := grpc.NewServer(cred.TLSInterceptor(),
	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	        grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_recovery.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		    grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
            grpc_recovery.UnaryServerInterceptor(),
		)),
	)


2.自定义错误返回
当panic时候，自定义错误码并返回。

// RecoveryInterceptor panic时返回Unknown错误吗
func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return grpc.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}

//添加grpc_recovery拦截器到服务端
grpcServer := grpc.NewServer(cred.TLSInterceptor(),
	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	        grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		    grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
            grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)

*/

