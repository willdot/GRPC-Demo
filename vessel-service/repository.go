package main

import (
	"errors"

	pb "github.com/willdot/GRPC-Demo/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository ...
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

// VesselRepository is a fake datastore
type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable checks a provided specification against all vessels and returns one that is under the capacity and max weight
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {

	var vessels []*pb.Vessel

	err := repo.collection().Find(nil).All(&vessels)
	if err != nil {
		return nil, err
	}
	for _, vessel := range vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

// Create will create a new vessel
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// Close closes the session after each query has run
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
