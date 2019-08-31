protoc -I ./session/ --go_out=plugins=grpc:./session/ ./session/session.proto
protoc -I ./redirect/ --go_out=plugins=grpc:./redirect/ ./redirect/redirect.proto