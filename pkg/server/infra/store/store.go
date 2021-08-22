package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/toaru/clean-arch-api/config"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(appEnv string) *SqlHandler {
	source := getSource(appEnv)
	con := createConnection(source)
	sqlHandler := &SqlHandler{
		Conn: con,
	}
	return sqlHandler
}

func getSource(appEnv string) string {
	name := config.Configs.DB_NAME
	user := config.Configs.DB_USER
	pass := config.Configs.DB_PASS
	host := config.Configs.DB_HOST
	port := config.Configs.DB_PORT
	instance := config.Configs.DB_INSTANCE_NAME
	socket := "/cloudsql"

	var source string
	switch appEnv {
	case "production":
		source = fmt.Sprintf("%s:%s@unix(%s/%s)/%s?parseTime=true&collation=utf8mb4_bin",
			user, pass, socket, instance, name)
	default:
		// set only collation and not charset, See: https://github.com/go-sql-driver/mysql/issues/745
		source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&collation=utf8mb4_bin",
			user, pass, host, port, name)
	}

	return source
}

func createConnection(Source string) *sql.DB {
	con, err := sql.Open("mysql", Source)
	// https://blog.nownabe.com/2017/01/16/570.html#accessing-the-database
	// defer con.Close()
	if err != nil {
		log.Fatal("DB Open Error: ", err)
	}
	return con
}
