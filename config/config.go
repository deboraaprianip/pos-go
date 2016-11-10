package config

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	MySqlDB   string
	MysqlUser string
	MysqlPass string
	MysqlHost string
	MysqlPort string
}

func LoadMysql() (*sql.DB, error) {
	DB, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/pos")
	if err != nil {
		return nil, errors.New("Not Connect")
	}
	return DB, nil
}
