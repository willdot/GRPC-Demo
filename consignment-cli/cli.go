package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"

	pb "github.com/willdot/GRPC-Demo/consignment-service/proto/consignment"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	cmd.Init()
	client := pb.NewConsignmentServiceClient("shippy.consignment", microclient.DefaultClient)

	file := defaultFilename

	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecting file and token"))
	}

	file = os.Args[1]
	token = os.Args[2]

	fmt.Println("Token: ", token)

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.Create(ctx, consignment)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("created: %t", r.Created)

	getAll, err := client.Get(ctx, &pb.GetRequest{})

	if err != nil {
		log.Fatalf("could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
