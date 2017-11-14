package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	//"fmt"
)

//type Customer struct {
//	id       int
//	name     string
//	email    string
//	phone_no string
//}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:root@/invoice_app")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//// Migrate the schema
	//db.AutoMigrate(&Customer{})
	//
	//// Create
	//db.Create(&Customer{name: "Ankita palekar", email: "ankita@cdmx.in", phone_no: "9158381548"})
	//
	//// Read
	//var cus Customer
	//firstCust := db.First(&cus, 1) // find product with id 1
	//fmt.Println(firstCust)
	////db.First(&cus, "email = ?", "ankita@cdmx.in") // find product with code l1212
	//
	//
	//
	//// Update - update product's price to 2000
	////db.Model(&cus).Update("Price", 2000)
	//
	//// Delete - delete product
	////db.Delete(&cus)

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
}
