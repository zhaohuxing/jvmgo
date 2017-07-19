package models

import "prepare/databases"
import "fmt"

type User struct {
	Id          int64  `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func InsertUser(phoneNumber, password string) bool {
	// db := db.Open() 糟糕命名方式
	db := databases.Open()
	stmt, err := db.Prepare("insert into user (phone_number, password)values(?,?)")
	if err != nil {
		panic("insert prepare failed.")
	}

	_, err = stmt.Exec(phoneNumber, password)
	if err != nil {
		panic("insert exec failed.")
	}

	return true
}

func GetUser(phoneNumber, password string) bool {
	db := databases.Open()
	defer db.Close()
	sql := "select * from user where phone_number = ? and password = ?"

	stmt, err := db.Prepare(sql)
	defer stmt.Close()

	if err != nil {
		panic(err)
	}
	var id int
	var pn string
	var pwd string
	row := stmt.QueryRow(phoneNumber, password)
	err = row.Scan(&id, &pn, &pwd)
	if err != nil {
		return false
	}
	fmt.Println("id:", id, "pn:", pn, "pwd:", pwd)
	return true
}

/*
func GetUser(number, password string) bool {
	// 1. 获取所有的记录
	// 2. 使用预编译获取单行数据、
	// 3. 直接获取单行数据
	db := databases.Open()
	defer db.Close()
	// 获取所有记录
	sql := "select * from user"
	rows, err := db.Query(sql)
	for rows.Next() {
		var id int
		var pn string
		var pwd string
		err = rows.Scan(&id, &pn, &pwd)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "pn:", pn, "pwd:", pwd)
	}
	// 直接获取单行数据
	sql = "select * from user where phone_number = ? and password = ?"
	row := db.QueryRow(sql, number, password)
	var id int
	var pn string
	var pwd string
	err = row.Scan(&id, &pn, &pwd)
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", id, "pn:", pn, "pwd:", pwd)

	// 预编译获取
	stmt, err := db.Prepare(sql)
	defer stmt.Close()

	if err != nil {
		panic(err)
	}

	row = stmt.QueryRow(number, password)
	err = row.Scan(&id, &pn, &pwd)
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", id, "pn:", pn, "pwd:", pwd)
	return true
}*/
