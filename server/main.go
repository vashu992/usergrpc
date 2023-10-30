package main

import (
	"context"
	"fmt"
	"net"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"practice/usergrpc/userservice"
)

var users = map[int64]*userservice.UserResponse{
	1: {
		Id:      1,
		Fname:   "Steve",
		City:    "LA",
		Phone:   "1234567890",
		Height:  5.8,
		Married: true,
	},
	2: {
		Id:      2,
		Fname:   "Alice",
		City:    "NY",
		Phone:   "9876543210",
		Height:  5.6,
		Married: false,
	},
	// Add more user data here
}

type server struct{
	userservice.UnimplementedUserServiceServer
}

func (s *server) GetUserById(ctx context.Context, req *userservice.GetUserRequest) (*userservice.UserResponse, error) {
	user, ok := users[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "User with ID %d not found", req.Id)
	}
	return user, nil
}

func (s *server) GetUsersByIds(ctx context.Context, req *userservice.GetUsersRequest) (*userservice.UsersResponse, error) {
	var response userservice.UsersResponse
	for _, id := range req.Ids {
		user, ok := users[id]
		if !ok {
			return nil, status.Errorf(codes.NotFound, "User with ID %d not found", id)
		}
		response.Users = append(response.Users, user)
	}
	return &response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userservice.RegisterUserServiceServer(s, &server{})
	fmt.Println("gRPC server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}