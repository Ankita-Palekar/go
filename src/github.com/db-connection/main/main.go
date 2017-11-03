package main

import (
	_ "github.com/go-sql-driver/mysql"
	//"database/sql"
	"fmt"
	"encoding/json"
	"database/sql"
)

type customer struct {
	id       int
	name     string
	email    string
	phone_no string
}

type Server struct {
	ServerName string  `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}




func main() {
	db, err := sql.Open("mysql", "root:root@/invoice_app")
	checkErr(err)
	stmt, err := db.Prepare("INSERT into customers  (id, name, email, phone_no) VALUES (?, ?, ?, ?) ")
	checkErr(err)

	res, err := stmt.Exec(nil, "Ankia", "ankita@gmail.com", "9158381548")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	rows, err := db.Query("SELECT * FROM customers")

	checkErr(err)

	fmt.Println("-----------------")
	fmt.Println(rows)
	fmt.Println("-----------------")

	for rows.Next() {
		var uid int
		var name string
		var email string
		var phone_no string
		err = rows.Scan(&uid, &name, &email, &phone_no)
		var user customer
		user.id = uid
		user.name = name
		user.email = email
		user.phone_no = phone_no
		//checkErr(err)
		//fmt.Println(uid)
		//fmt.Println(username)
		//fmt.Println(department)
		//fmt.Println(created)


		fmt.Println("user =>", user)
		fmt.Println(json.Marshal(user))
	}

	db.Close()

	var s Serverslice
	str := `{"servers":[{"serverName":"INDIA","serverIP":"127.0.0.1"},{"serverName":"INDIA VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)


}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
