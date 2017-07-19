package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/prepare?charset=utf8")
	// 这个地方处理的有点生猛
	if err != nil {
		panic(err)
	}
	return db
}
