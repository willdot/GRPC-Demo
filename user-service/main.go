package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
	pb "github.com/willdot/GRPC-Demo/user-service/proto/auth"
)

func init() {
	plugin.Register(cors.NewPlugin())
}

func main() {
	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.auth"),
	)

	srv.Init()

	publisher := micro.NewPublisher("user.created", srv.Client())
	publisher2 := micro.NewPublisher("user.created2", srv.Client())

	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService, publisher, publisher2})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
