package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"practice/usergrpc/userservice"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client using the connection.
	client := userservice.NewUserServiceClient(conn)

	// Call the GetUserById RPC.
	user, err := GetUserByID(client, 1)
	if err != nil {
		log.Fatalf("Error calling GetUserById: %v", err)
	}
	fmt.Printf("User by ID: %v\n", user)

	// Call the GetUsersByIds RPC.
	users, err := GetUsersByIds(client, []int64{1, 2})
	if err != nil {
		log.Fatalf("Error calling GetUsersByIds: %v", err)
	}
	fmt.Println("Users by IDs:")
	for _, u := range users {
		fmt.Println(u)
	}
}

// GetUserByID calls the GetUserById RPC and returns the user details.
func GetUserByID(client userservice.UserServiceClient, userID int64) (*userservice.UserResponse, error) {
	req := &userservice.GetUserRequest{
		Id: userID,
	}
	user, err := client.GetUserById(context.Background(), req)
	return user, err
}

// GetUsersByIds calls the GetUsersByIds RPC and returns a list of user details.
func GetUsersByIds(client userservice.UserServiceClient, userIDs []int64) ([]*userservice.UserResponse, error) {
	req := &userservice.GetUsersRequest{
		Ids: userIDs,
	}
	users, err := client.GetUsersByIds(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return users.Users, nil
}
