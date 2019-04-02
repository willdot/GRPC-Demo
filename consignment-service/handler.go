package main

import (
	"log"

	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"

	pb "github.com/willdot/GRPC-Demo/consignment-service/proto/consignment"
	vesselProto "github.com/willdot/GRPC-Demo/vessel-service/proto/vessel"
)

// Service should implement all of the methods in the protobuf definition
type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

// CreateConsignment is a method that creates a consignment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	err = repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments is a method that gets all consignments
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
