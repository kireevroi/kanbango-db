package userproto

import (
	"context"

	"github.com/kireevroi/kanbango/auth/internal/db"
	"github.com/kireevroi/kanbango/auth/internal/cache"
	"github.com/kireevroi/kanbango/auth/pkg/hash"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedUserServiceServer
	DB *db.DB
	Cache *cache.Cache
}

func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	h, err := hash.HashPassword(in.Password)
	if err != nil {
		return &CreateUserResponse{Status: Status_STATUS_BADPASSWD}, status.Error(codes.InvalidArgument, "password not hashable")
	}
	err = s.DB.CreateUser(db.User{Username: in.Username, PasswordHash: h})
	if err != nil {
		return &CreateUserResponse{Status: Status_STATUS_USEREX}, status.Error(codes.AlreadyExists, "user already exists / failed to write to db")
	}
	return &CreateUserResponse{Status: Status_STATUS_OK}, nil
}

func (s *Server) LoginUser(ctx context.Context, in *LoginUserRequest) (*LoginUserResponse, error) {
	var uid string
	var err error
	if in.Uuid == "" {
		u, err := s.DB.GetUser(in.Username)
		if err != nil {
			return &LoginUserResponse{Status: Status_STATUS_NOUSER, Uuid: ""}, status.Error(codes.Internal, "Username doesn't exist")
		}
		if !hash.CheckPassword(u.PasswordHash, in.Password) {
			return &LoginUserResponse{Status: Status_STATUS_WRGPASSWD, Uuid: ""}, nil
		}
		uid, err = hash.GenerateUUID()
		if err != nil {
			return &LoginUserResponse{Status: Status_STATUS_UNSPECIFIED, Uuid: ""}, status.Error(codes.Internal, "Unknown")
		}
		s.Cache.NewSession(uid, in.Username)
		err = s.DB.SetSession(db.Session{Session: uid, User_ID: u.ID})
		if err != nil {
			return &LoginUserResponse{Status: Status_STATUS_UNSPECIFIED, Uuid: ""}, status.Error(codes.Internal, "Cache probs")
		}
	} else {
		_, err = s.Cache.GetSession(in.Uuid)
		if err != nil {
			return &LoginUserResponse{Status: Status_STATUS_NOUSER, Uuid: ""}, nil
		}
		uid = in.Uuid
	}
	return &LoginUserResponse{Status: Status_STATUS_OK, Uuid: uid}, nil
}
