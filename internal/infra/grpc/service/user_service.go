package service

import (
	"context"

	"github.com/MrHenri/meuPet/internal/infra/grpc/pb"
	"github.com/MrHenri/meuPet/internal/usecase"
)

type UserService struct {
	pb.UnimplementedManageUserServer
	UserUseCase usecase.UserUseCase
}

func NewUserService(userUseCase usecase.UserUseCase) *UserService {
	return &UserService{UserUseCase: userUseCase}
}

func (u *UserService) RegisterUser(ctx context.Context, in *pb.UserCreationInput) (*pb.ResponseMessage, error) {
	dto := usecase.UserInputDTO{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
		Phone:    in.Phone,
	}

	err := u.UserUseCase.Register(dto)
	return &pb.ResponseMessage{Success: err != nil}, err
}
