package service

import (
	"context"
	"database/sql"

	protos "github.com/Mensurui/golang-testcontainers/protobuf/gen"
)

type Service struct {
	protos.UnimplementedUserServiceServer
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CheckHealth(ctx context.Context, req *protos.CheckHealthRequest) (*protos.CheckHealthResponse, error) {
	return &protos.CheckHealthResponse{
		Message: "Working",
	}, nil
}
