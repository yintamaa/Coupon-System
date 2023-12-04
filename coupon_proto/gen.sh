protoc --go_out=. --go-grpc_out=. *.proto

# protoc -I . *.proto --go_out=plugins=grpc:pb
