package midd

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"log"
	"gopkg.in/natefinch/lumberjack.v2"
)
/*
字段含义
{
	  "level": "info",						// string  zap log levels
	  "msg": "finished unary call",			// string  log message

	  "grpc.code": "OK",						// string  grpc status code
	  "grpc.method": "Ping",					/ string  method name
	  "grpc.service": "mwitkow.testproto.TestService",              // string  full name of the called service
	  "grpc.start_time": "2006-01-02T15:04:05Z07:00",               // string  RFC3339 representation of the start time
	  "grpc.request.deadline": "2006-01-02T15:04:05Z07:00",         // string  RFC3339 deadline of the current request if supplied
	  "grpc.request.value": "something",		// string  value on the request
	  "grpc.time_ms": 1.345,					// float32 run time of the call in ms

	  "peer.address": {
	    "IP": "127.0.0.1",					// string  IP address of calling party
	    "Port": 60216,						// int     port call is coming in on
	    "Zone": ""							// string  peer zone for caller
	  },
	  "span.kind": "server",				// string  client | server
	  "system": "grpc",						// string

	  "custom_field": "custom_value",		// string  user defined field
	  "custom_tags.int": 1337,				// int     user defined tag on the ctx
	  "custom_tags.string": "something"		// string  user defined tag on the ctx
}
*/


//1.创建zap.Logger实例
func ZapInterceptor() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}

//日志写到硬盘
func ZapRomInterceptor() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "log/debug.log",
		MaxSize:   1024, //MB
		LocalTime: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.NewAtomicLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}


func LogGrpc()  {
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(ZapInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(ZapInterceptor()),
		)),
	)
	//优雅关闭grpc
	grpcServer.GracefulStop()
}