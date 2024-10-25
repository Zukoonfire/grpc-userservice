package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/Zukoonfire/grpc-userservice/proto/github.com/Zukoonfire/grpc-userservice/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	mu    sync.Mutex
	users map[int32]*pb.User
}

// CreateUser implements pb.UserServiceServer
func (s *server) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.Id] = user
	log.Printf("User created:%v", user)
	return user, nil
}

// GetUser implements pb.UserServiceServer
func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//fetch the user by ID
	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user with ID %d not found", req.Id)
	}
	log.Printf("Fetched user:%v", user)
	return user, nil
}

func main() {
	s := &server{
		users: make(map[int32]*pb.User),
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s)
	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve :%v", err)
	}
}
