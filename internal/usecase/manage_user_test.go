package usecase

import (
	"errors"
	"testing"

	"github.com/MrHenri/meuPet/internal/entities"
	"github.com/MrHenri/meuPet/internal/events"
	pkgEvent "github.com/MrHenri/meuPet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

type MockEventDispatcher struct {
	mock.Mock
}

func (m *MockEventDispatcher) Register(eventName pkgEvent.EventName, handler pkgEvent.EventHandlerInterface) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

func (m *MockEventDispatcher) Dispatch(event pkgEvent.EventInterface) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventDispatcher) Has(eventName pkgEvent.EventName, handler pkgEvent.EventHandlerInterface) bool {
	args := m.Called(eventName)
	return args.Bool(0)
}

func TestRegister(t *testing.T) {
	var mockRepo MockUserRepository

	userCreated := events.NewUserEvent(events.UserCreated)
	var mockEventDispatcher MockEventDispatcher
	userUseCase := NewUserUseCase(&mockRepo, userCreated, &mockEventDispatcher)

	mockRepoCall := mockRepo.On("CreateUser", mock.Anything).Return(nil)
	mockEventDispatcher.On("Dispatch", mock.Anything).Return(nil)

	t.Run("Success", func(t *testing.T) {
		userInputDTO := UserInputDTO{
			Name:     "test",
			Email:    "test@example.com",
			Password: "password",
			Phone:    "123456789",
		}

		err := userUseCase.Register(userInputDTO)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockEventDispatcher.AssertExpectations(t)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		userInputDTO := UserInputDTO{
			Name:     "",
			Email:    "",
			Password: "",
			Phone:    "",
		}

		err := userUseCase.Register(userInputDTO)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockEventDispatcher.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userInputDTO := UserInputDTO{
			Name:     "test",
			Email:    "test@test.com",
			Password: "password",
			Phone:    "12345678",
		}

		mockRepoCall.Unset()

		mockRepo.On("CreateUser", mock.Anything).Return(errors.New("repository Error"))

		err := userUseCase.Register(userInputDTO)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockEventDispatcher.AssertExpectations(t)
	})
}
