gen:
	protoc --go_out=. model/messages.proto
	protoc --go-grpc_out=. -I model messages.proto

install:
	go get google.golang.org/grpc
	go get google.golang.org/grpc/codes
	go get google.golang.org/grpc/status
	go get google.golang.org/protobuf/reflect/protoreflect
	go get google.golang.org/protobuf/runtime/protoimpl

build:
	cd client
	go build
	cd ..
	cd server
	go build

run-c:
	cd client
	go run main.go

run-s:
	cd server
	go run main.go

tidy:
	go mod tidy
