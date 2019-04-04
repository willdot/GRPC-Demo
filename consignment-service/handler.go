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

// Create is a method that creates a consignment
func (s *service) Create(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	defer s.GetRepo().Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	err = s.GetRepo().Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// Get is a method that gets all consignments
func (s *service) Get(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	defer s.GetRepo().Close()

	consignments, err := s.GetRepo().GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
