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
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
	pb "github.com/willdot/GRPC-Demo/consignment-service/proto/consignment"
	authService "github.com/willdot/GRPC-Demo/user-service/proto/auth"
	vesselProto "github.com/willdot/GRPC-Demo/vessel-service/proto/vessel"
	"golang.org/x/net/context"
)

const (
	defaultHost = "localhost:27017"
)

func init() {
	plugin.Register(cors.NewPlugin())
}

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
		micro.Name("shippy.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	vesselClient := vesselProto.NewVesselServiceClient("shippy.vessel", srv.Client())

	srv.Init()

	pb.RegisterConsignmentServiceHandler(srv.Server(), &service{session, vesselClient})

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

		token := meta["Token"]
		log.Println("Authenticating token: ", token)

		authClient := authService.NewAuthClient("shippy.auth", client.DefaultClient)

		authResp, err := authClient.ValidateToken(ctx, &authService.Token{
			Token: token,
		})

		log.Println("Auth response: ", authResp)
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
