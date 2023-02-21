package database

import (
	"database/sql"

	"github.com/MrHenri/meuPet/configs"
	_ "github.com/go-sql-driver/mysql"
)

func Init(conf *configs.Conf) (*sql.DB, error) {
	dns := conf.DBUser + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ")/" + conf.DBName
	db, err := sql.Open(conf.DBDriver, dns)
	return db, err
}
