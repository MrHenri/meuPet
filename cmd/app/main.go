package main

import (
	"fmt"
	"net"

	"github.com/MrHenri/meuPet/configs"
	"github.com/MrHenri/meuPet/internal/infra/database"
	"github.com/MrHenri/meuPet/internal/infra/grpc/pb"
	"github.com/MrHenri/meuPet/internal/infra/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf, err := configs.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.Init(conf)
	if err != nil {
		panic(err)
	}

	userUseCase := NewUserUseCase(db)
	userService := service.NewUserService(*userUseCase)

	grpcServer := grpc.NewServer()
	pb.RegisterManageUserServer(grpcServer, userService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", conf.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCServerPort))
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(lis)
}
