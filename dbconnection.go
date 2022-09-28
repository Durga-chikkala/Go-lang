package Driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	user   string
	pass   string
	host   string
	port   string
	dbName string
}

func Connstr(config MysqlConfig) string {
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.user, config.pass, config.host, config.port, config.dbName)
	return connStr
}
func Connection() (*sql.DB, error) {
	a := MysqlConfig{"Durga", "abc", "localhost", "3306", "SIMPLE"}
	db, err := sql.Open("mysql", Connstr(a))

	return db, err
}
