package userproto

import (
	"context"

	"github.com/kireevroi/kanbango/auth/internal/db"
	"github.com/kireevroi/kanbango/auth/internal/cache"
	"github.com/kireevroi/kanbango/auth/pkg/hash"
)

type Server struct {
	UnimplementedUserServiceServer
	DB *db.DB
	Cache *cache.Cache
}




// CreateUser functions handles user creation
func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	h, err := hash.HashPassword(in.Password)

	if err != nil {
		return &CreateUserResponse{Status: Status_STATUS_BADPASSWD}, nil
	}

	err = s.DB.CreateUser(db.User{Username: in.Username, PasswordHash: h})
	if err != nil {
		return &CreateUserResponse{Status: Status_STATUS_USEREX}, nil
	}

	return &CreateUserResponse{Status: Status_STATUS_OK}, nil
}

// LoginUser logins user and returns uuid
func (s *Server) LoginUser(ctx context.Context, in *LoginUserRequest) (*AuthUserResponse, error) {
	var uid string
	var err error
	if in.Uuid == "" {

		u, err := s.DB.GetUser(in.Username)
		if err != nil {
			return &AuthUserResponse{Status: Status_STATUS_NOUSER, Uuid: ""}, nil
		}

		if !hash.CheckPassword(u.PasswordHash, in.Password) {
			return &AuthUserResponse{Status: Status_STATUS_WRGPASSWD, Uuid: ""}, nil
		}

		uid, err = hash.GenerateUUID()
		if err != nil {
			return &AuthUserResponse{Status: Status_STATUS_UNSPECIFIED, Uuid: ""}, nil
		}

		s.Cache.NewSession(uid, in.Username)
		err = s.DB.SetSession(db.Session{Session: uid, User_ID: u.ID})
		if err != nil {
			return &AuthUserResponse{Status: Status_STATUS_UNSPECIFIED, Uuid: ""}, nil
		}

	} else {

		_, err = s.Cache.GetSession(in.Uuid)
		if err != nil {
			return &AuthUserResponse{Status: Status_STATUS_NOUSER, Uuid: ""}, nil
		}

		return &AuthUserResponse{Status: Status_STATUS_ALRLOGGED, Uuid: in.Uuid}, nil
	}

	return &AuthUserResponse{Status: Status_STATUS_OK, Uuid: uid}, nil
}

// LogoutUser logouts the user and removes from cache
func (s *Server)LogoutUser(ctx context.Context, in *LogoutUserRequest) (*AuthUserResponse, error) {

	if (in.Uuid != "") {
		s.Cache.DeleteSession(in.Uuid)
		s.DB.DeleteSession(db.Session{Session: in.Uuid})
	}

	return &AuthUserResponse{Status: Status_STATUS_OK, Uuid: ""}, nil
}

func (s *Server)DeleteUser(ctx context.Context, in *LogoutUserRequest) (*AuthUserResponse, error) {
	if (in.Uuid == "") {
		return &AuthUserResponse{Status: Status_STATUS_UNSPECIFIED, Uuid: ""}, nil
	}

	user_id, _ := s.DB.GetUserIdByUuid(in.Uuid)
	sessions, _ := s.DB.GetAllSessions(user_id)
	for i, _ := range sessions {
		s.Cache.DeleteSession(sessions[i].Session)
	}
	s.DB.DeleteUser(user_id)
	s.DB.DeleteAllSessions(user_id)

	return &AuthUserResponse{Status: Status_STATUS_OK, Uuid: ""}, nil
}