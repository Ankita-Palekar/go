package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

var db, err = sql.Open("mysql", "root:root@/invoice_app")

func connectDB() {
	checkErr(err)
}

func insertJWTToDB() {
	stmt, err := db.Prepare("INSERT into customers  (id, name, email, phone_no) VALUES (?, ?, ?, ?) ")
	checkErr(err)

	res, err := stmt.Exec(nil, "Ankia", "ankita@gmail.com", "9158381548")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

