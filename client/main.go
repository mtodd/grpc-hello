package main

import (
	"log"
	"os"

	pb "../proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	serverAddr  = "127.0.0.1:8000"
	defaultName = "mtodd"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	req := pb.GreetRequest{Name: name}
	log.Printf("GreetRequest%v", req)

	res, err := client.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("%v.Greet(_) = _, %v", client, err)
	}
	log.Println(res)
}
