package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	pb "github.com/willdot/GRPC-Demo/user-service/proto/user"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (s *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with: ", req.Email, req.Password)

	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	res.Token = user.Name
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := s.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
