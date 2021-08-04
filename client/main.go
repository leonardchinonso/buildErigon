package main

import (
	"context"
	model "es/model"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var (
	server     = "localhost:8000" // TODO: Use environment variables to hide this detail later on
	branchName string
)

func main() {
	// Receive inputs from command line
	flag.StringVar(&server, "s", server, "gRPC server address host:port")
	flag.StringVar(&branchName, "b", branchName, "branch name to check out erigon from")
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	// Make a connection to the server using grpc dial options
	conn, err := grpc.Dial(server, opts...)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to connect to gRPC service: %v", err))
		return
	}

	// defer a function to close the connection when done
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(fmt.Errorf("cannot close connection: %v", err))
		}
	}(conn)

	client := model.NewErigonServiceClient(conn)

	// Initialise context and branch request message
	ctx := context.Background()
	in := &model.BranchRequest{BranchName: branchName}

	// Call BuildFrom method to check out erigon on server and build it
	result, err := client.BuildFrom(ctx, in)

	if err != nil {
		log.Fatal(fmt.Errorf("BuildFrom rpc failed: %v", err))
	}

	fmt.Printf("BuildFrom(%v) => %v\n", in.BranchName, result)
}
