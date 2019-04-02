package main

import (
	"fmt"
	"log"

	pb "github.com/willdot/GRPC-Demo/consignment-service/proto/consignment"
	vesselProto "github.com/willdot/GRPC-Demo/vessel-service/proto/vessel"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"
)

// Repository ..
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository is a fake datastore
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create will create a consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll will get all consignments
func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods in the protobuf definition
type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment is a method that creates a consignment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments is a method that gets all consignments
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()

	res.Consignments = consignments
	return nil
}

func main() {
	repo := &ConsignmentRepository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
