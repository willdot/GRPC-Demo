package main

import (
	"fmt"

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
	CassandraSession := Session
	defer CassandraSession.Close()

	repo := &UserRepository{CassandraSession}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.auth"),
	)

	srv.Init()

	publisher := micro.NewPublisher("user.created", srv.Client())

	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService, publisher})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
