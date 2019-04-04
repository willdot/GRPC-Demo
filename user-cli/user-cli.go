package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	pb "github.com/willdot/GRPC-Demo/user-service/proto/auth"
)

func main() {

	srv := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)

	srv.Init()

	client := pb.NewAuthClient("shippy.auth", microclient.DefaultClient)

	name := "william andrews"
	email := "will@hotmail.com"
	password := "password"
	company := "civca"

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})

	if err != nil {
		log.Fatalf("could not create: %v", err)
	}

	log.Printf("Created: %v", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})

	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: "password",
	})

	if err != nil {
		log.Fatalf("could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	os.Exit(0)
}
