package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConnection() (*sql.DB, error) {
	mysqlUrl := os.Getenv("MYSQL_URL")
	mysqlDB := os.Getenv("MYSQL_DB")
	dataSourceName := fmt.Sprintf("root:root@tcp(%s)/%s", mysqlUrl, mysqlDB)
	db, err := sql.Open("mysql", dataSourceName)
	return db, err
}
