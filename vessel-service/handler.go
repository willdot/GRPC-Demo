package main

import (
	"context"

	pb "github.com/willdot/GRPC-Demo/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

// service should implement all the methods in the protobuf definition
type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

// FindAvailable checks a provided specification against all vessels and returns ones that are under the capacity and max weight
func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

// Create will create a vessel
func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true

	return nil
}
