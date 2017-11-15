package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	//"fmt"
	//"time"
	"fmt"
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

type User struct {
	gorm.Model
	Name      string
	Profile   Profile
	ProfileId int
}

type Profile struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/invoice_app")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	//db.AutoMigrate(&User{}, &Profile{})
	//createProfile(db, err, "Ankita Profile")

	//createProfile(db, err, "anki")
	//createUser(db, err, "ankita user")

	//var user User
	//
	//var profile Profile
	//db.Model(&user).Related(&profile)
	//fmt.Println(user)
	//fmt.Println(profile)
	user := Profile{}


	db.First(&user)
	fmt.Println(user)
}

func createProfile(db *gorm.DB, err error, name string) {
	db.Create(&Profile{Name: name})
}

func createUser(db *gorm.DB, err error, name string) {
	db.Create(&User{Name: name, ProfileId: 1})
}
