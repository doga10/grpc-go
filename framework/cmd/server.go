package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	oauth2 "grpc-go/application/oauth"
	"grpc-go/application/users"
	"grpc-go/framework/database"
	"grpc-go/framework/pb"
	"grpc-go/framework/servers"
	"log"
	"net"
)

var db *mongo.Database

func main() {
	db = database.ConnectMongoDB()

	userServer := setUpUserServer()
	oauthServer := setUpOAuthServer()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	pb.RegisterOAuthServiceServer(grpcServer, oauthServer)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", 8080)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	fmt.Println("Server running on tpc://0.0.0.0:8080")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

func setUpUserServer() *servers.UserServer {
	userServer := servers.NewUserServer()
	userRepository := users.UserRepository{Db: db}
	userServer.UserService = users.UserService{UserRepository: userRepository}
	return userServer
}

func setUpOAuthServer() *servers.OAuthServer {
	oauthServer := servers.NewOAuthServer()
	userRepository := users.UserRepository{Db: db}
	oauthServer.OAuthService = oauth2.OAuthService{OAuthRepository: userRepository}
	return oauthServer
}
