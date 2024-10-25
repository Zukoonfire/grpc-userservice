package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Zukoonfire/grpc-userservice/proto/github.com/Zukoonfire/grpc-userservice/proto"
	"google.golang.org/grpc"
)

func main() {
	//Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect:%v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	//create a new user
	user := &pb.User{

		Id:    3,
		Name:  "Jemin",
		Email: "jemin.smith@example.com",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	createdUser, err := c.CreateUser(ctx, user)
	if err != nil {
		log.Fatalf("Could not create user:%v", err)
	}
	log.Printf("Created USer:%v", createdUser)
	//Call GetUser
	userRequest := pb.UserRequest{Id: 3}
	fetchedUser, err := c.GetUser(ctx, &userRequest)

	if err != nil {
		log.Fatalf("could not fetch user:%v", err)
	}
	
	log.Printf("Fetched User:%v", fetchedUser)
	// userRequest1 := pb.UserRequest{Id: 2}
	// fetchedUser2, err := c.GetUser(ctx, &userRequest1)
	// if err != nil {
	// 	log.Fatalf("could not fetch user:%v", err)
	// }
	// log.Printf("Fetched User:%v", fetchedUser2)
}
