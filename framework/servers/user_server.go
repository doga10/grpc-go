package servers

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	application "grpc-go/application/users"
	domain "grpc-go/domain/users"
	"grpc-go/framework/pb"
)

type UserServer struct {
	UserService application.UserService
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (server *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := domain.NewUser(req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error during the validation: %v", err)
	}

	cur, err := server.UserService.Save(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error persisting information: %v", err)
	}

	return &pb.UserResponse{
		Id: cur.ID.String(),
		Name: cur.Name,
		Email: cur.Email,
		Password: cur.Password,
		CreatedAt: cur.CreatedAt.String(),
		UpdatedAt: cur.UpdatedAt.String(),
	}, nil
}
