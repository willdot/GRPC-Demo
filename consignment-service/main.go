package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro/client"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	pb "github.com/willdot/GRPC-Demo/consignment-service/proto/consignment"
	userService "github.com/willdot/GRPC-Demo/user-service/proto/user"
	vesselProto "github.com/willdot/GRPC-Demo/vessel-service/proto/vessel"
	"golang.org/x/net/context"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// AuthWrapper is a wrapper for authorising using a JWT
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)

		if !ok {
			return errors.New("no auth meta data found in request")
		}

		token := meta["token"]
		log.Println("Authenticating token: ", token)

		authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)

		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})

		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
