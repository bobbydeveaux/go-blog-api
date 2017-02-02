package main

import (
	"github.com/bobbydeveaux/go-blog-api/api"

	"log"
	"net"

	pb "github.com/bobbydeveaux/go-blog-proto/post"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetPosts(ctx context.Context, in *pb.PostRequest) (*pb.PostReply, error) {
	api := new(api.API)
	posts := api.GetPosts()
	return &pb.PostReply{Message: posts}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
