package servers

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	application "grpc-go/application/oauth"
	"grpc-go/framework/pb"
)

type OAuthServer struct {
	OAuthService application.OAuthService
}

func NewOAuthServer() *OAuthServer {
	return &OAuthServer{}
}

func (server *OAuthServer) Login(ctx context.Context, req *pb.OAuthRequest) (*pb.OAuthResponse, error) {
	token, err := server.OAuthService.DoLogin(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error persisting information: %v", err)
	}

	return &pb.OAuthResponse{
		Token: token.Token,
	}, nil
}
