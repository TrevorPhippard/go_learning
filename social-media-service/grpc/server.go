package grpc

import (
	"context"

	"social-media-service/internal/model"
	"social-media-service/pkg/db"
	"social-media-service/proto/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	user.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
	var u model.User
	if err := db.DB.First(&u, req.Id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &user.UserResponse{
		Id:       uint32(u.ID),
		Username: u.Username,
		Email:    u.Email,
	}, nil
}
