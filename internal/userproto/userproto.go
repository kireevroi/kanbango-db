package userproto

import (
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/kireevroi/kanbango/auth/internal/db"
	"github.com/kireevroi/kanbango/auth/pkg/hash"
)

type Server struct {
	UnimplementedUserServiceServer
	DB *db.DB
}

func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	h, err := hash.HashPassword(in.Password)
	if err != nil {
		return &CreateUserResponse{Status: false}, status.Error(codes.InvalidArgument, "password not hashable")
	}
	err = s.DB.CreateUser(db.User{Username: in.Username, PasswordHash: h})
	if err != nil {
		return &CreateUserResponse{Status: false}, status.Error(codes.AlreadyExists, "user already exists / failed to write to db")
	}
	return &CreateUserResponse{Status: true}, nil
}

func (s *Server) LoginUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	u, err := s.DB.GetUser(in.Username)
	if err != nil {
		return &CreateUserResponse{Status: false}, status.Error(codes.Internal, "Username doesn't exist")
	}
	if !hash.CheckPassword(u.PasswordHash, in.Password) {
		return &CreateUserResponse{Status: false}, nil
	}
	return &CreateUserResponse{Status: true}, nil
}