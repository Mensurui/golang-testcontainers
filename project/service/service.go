package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/Mensurui/golang-testcontainers/project/data"
	protos "github.com/Mensurui/golang-testcontainers/protobuf/gen"
)

type Service struct {
	protos.UnimplementedUserServiceServer
	db     *sql.DB
	models *data.Models
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db:     db,
		models: data.NewModels(db),
	}
}

func (s *Service) CheckHealth(ctx context.Context, req *protos.CheckHealthRequest) (*protos.CheckHealthResponse, error) {
	return &protos.CheckHealthResponse{
		Message: "Working",
	}, nil
}

func (s *Service) AddUser(ctx context.Context, req *protos.AddUserRequest) (*protos.AddUserResponse, error) {
	username := req.UserName
	email := req.Email
	age := req.Age

	err := s.models.User.AddUser(username, email, age)
	if err != nil {
		log.Printf("Error registering the user")
		return &protos.AddUserResponse{
			Message: "Unable to register",
		}, err
	}
	return &protos.AddUserResponse{
		Message: "Registered Successfully",
	}, err
}

func (s *Service) CheckUser(ctx context.Context, req *protos.CheckUserRequest) (*protos.CheckUserResponse, error) {
	userID := req.UserID

	user, err := s.models.User.CheckUser(userID)
	if err != nil {
		log.Printf("Error fetching data")
		return &protos.CheckUserResponse{}, nil
	}

	return &protos.CheckUserResponse{
		User: user.UserName,
		Age:  user.Email,
	}, nil
}
