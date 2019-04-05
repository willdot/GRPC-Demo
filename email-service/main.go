package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"

	pb "github.com/willdot/GRPC-Demo/user-service/proto/auth"
)

const topic = "user.created"
const topic2 = "user.created2"

// Subscriber ..
type Subscriber struct{}

// Process ..
func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("New messaged received")
	log.Println("Sending email to: ", user.Name)
	return nil
}

// Subscriber ..
type Subscriber2 struct{}

// Process ..
func (sub *Subscriber2) Process(ctx context.Context, user *pb.User) error {
	log.Println("New messaged received from second")
	log.Println("Sending email to second: ", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))
	micro.RegisterSubscriber(topic2, srv.Server(), new(Subscriber2))

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("Sending email to: ", user.Name)
	return nil
}
