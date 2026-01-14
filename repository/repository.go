package repository

import (
	"factory/configuration"
	"factory/repository/mysql"
	"factory/repository/sqlserver"
	"fmt"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
}

func New(configuration configuration.Configuration) (Repository, error) {
	var repository Repository

	switch configuration.Engine {
	case "MySql":
		repository = mysql.NewMysql()
	case "SqlServer":
		repository = sqlserver.NewSqlServer()
	default:
		return nil, fmt.Errorf("unsupported engine %s", configuration.Engine)
	}

	return repository, nil
}
