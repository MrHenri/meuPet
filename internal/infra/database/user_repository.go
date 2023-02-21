package database

import (
	"database/sql"

	"github.com/MrHenri/meuPet/internal/entities"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u *UserRepository) CreateUser(user *entities.User) error {
	_, err := u.Db.Exec("INSERT INTO users (id, name, email, password, phone) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Name, user.Email, user.Password, user.Phone)
	return err
}
