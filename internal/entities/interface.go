package entities

type UserRepositoryInterface interface {
	CreateUser(user *User) error
}
