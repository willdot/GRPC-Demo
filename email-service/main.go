package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"

	pb "github.com/willdot/GRPC-Demo/user-service/proto/user"
)

const topic = "user.created"

// Subscriber ..
type Subscriber struct{}

// Process ..
func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("New messaged received")
	log.Println("Sending email to: ", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("Sending email to: ", user.Name)
	return nil
}
