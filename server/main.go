package main // import "github.com/mtodd/grpc-play/server"

import (
	"fmt"
	"log"
	"net"

	pb "github.com/mtodd/grpc-play/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = 8000
)

type helloServer struct {
}

func (s *helloServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println(*req)
	return &pb.GreetResponse{Message: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening...")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	log.Println(grpcServer)

	s := new(helloServer)
	pb.RegisterHelloServer(grpcServer, s)
	log.Println(*s)
	grpcServer.Serve(lis)
}
