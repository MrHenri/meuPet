package usecase

import (
	"github.com/MrHenri/meuPet/internal/entities"
	"github.com/MrHenri/meuPet/pkg/events"
)

type UserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserUseCase struct {
	UserRepository    entities.UserRepositoryInterface
	UserEvent         events.EventInterface
	UserEventDispatch events.EventDispatcherInterface
}

//TODO: Event
// func NewUserUseCase(userRepository entities.UserRepositoryInterface,
// 	userEvent events.EventInterface,
// 	userEventDispatch events.EventDispatcherInterface) *UserUseCase {
// 	return &UserUseCase{UserRepository: userRepository,
// 		UserEvent:         userEvent,
// 		UserEventDispatch: userEventDispatch}
// }

func NewUserUseCase(userRepository entities.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{UserRepository: userRepository}
}

func (u *UserUseCase) Register(userInputDTO UserInputDTO) error {
	user, err := entities.NewUser(userInputDTO.Name, userInputDTO.Email, userInputDTO.Password, userInputDTO.Phone)
	if err != nil {
		return err
	}

	if err = u.UserRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
	//TODO: Event
	// return u.UserEventDispatch.Dispatch(u.UserEvent)
}
