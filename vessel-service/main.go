package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/willdot/GRPC-Demo/vessel-service/proto/vessel"
)

// Repository ...
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository is a fake datastore
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable checks a provided specification against all vessels and returns one that is under the capacity and max weight
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

// service should implement all the methods in the protobuf definition
type service struct {
	repo Repository
}

// FindAvailable checks a provided specification against all vessels and returns ones that are under the capacity and max weight
func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatFace", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
