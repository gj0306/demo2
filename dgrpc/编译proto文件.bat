rem 标准protobuf
.\protoc.exe --go_out=.\protocol\base\ .\protocol\gs.proto
rem grpc版协议
.\protoc.exe --go_out=plugins=grpc:.\protocol\grpc\ .\protocol\gs.proto
rem server_stream
.\protoc.exe --go_out=plugins=grpc:.\protocol\ssgrpc\ .\protocol\server_stream.proto
rem client_stream
.\protoc.exe --go_out=plugins=grpc:.\protocol\csgrpc\ .\protocol\client_stream.proto
rem both_stream
.\protoc.exe --go_out=plugins=grpc:.\protocol\bsgrpc\ .\protocol\both_stream.proto
rem advanced
.\protoc.exe --go_out=plugins=grpc:.\protocol\advanced\ .\protocol\advanced.proto
rem GrpcToHttp
.\protoc.exe --go_out=plugins=grpc:.\protocol\ht\ .\protocol\ht.proto
pause