package usecase

import (
	"github.com/MrHenri/meuPet/internal/entities"
	"github.com/MrHenri/meuPet/pkg/events"
)

type UserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserOutputDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserUseCase struct {
	UserRepository    entities.UserRepositoryInterface
	UserEvent         events.EventInterface
	UserEventDispatch events.EventDispatcherInterface
}

func NewUserUseCase(userRepository entities.UserRepositoryInterface,
	userEvent events.EventInterface,
	userEventDispatch events.EventDispatcherInterface) *UserUseCase {
	return &UserUseCase{UserRepository: userRepository,
		UserEvent:         userEvent,
		UserEventDispatch: userEventDispatch}
}

func (u *UserUseCase) Register(userInputDTO UserInputDTO) error {
	user, err := entities.NewUser(userInputDTO.Email, userInputDTO.Password, userInputDTO.Phone)
	if err != nil {
		return err
	}

	if err = u.UserRepository.CreateUser(user); err != nil {
		return err
	}

	return u.UserEventDispatch.Dispatch(u.UserEvent)
}
