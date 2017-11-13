package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

func create(token string, alg string, time_expiry string) {
	db, err := sql.Open("mysql", "root:root@/invoice_app")
	checkErr(err)
	stmt, err := db.Prepare("INSERT into jwt(id, user_id, token, alg, time_expiry) VALUES (?, ?, ?, ?, ?) ")
	checkErr(err)

	res, err := stmt.Exec(nil, "1", token, alg, time_expiry)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}
