package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"

	email "github.com/willdot/GRPC-Demo/email-service/proto/email"
	pb "github.com/willdot/GRPC-Demo/user-service/proto/auth"
)

const topic = "user.created"
const topic2 = "user.created2"

// Subscriber ..
type Subscriber struct{}

// Process ..
func (sub *Subscriber) Process(ctx context.Context, email *email.Message) error {
	log.Println("New messaged received")
	log.Println(email.Subject)
	log.Println(email.Content)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.email"),
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
