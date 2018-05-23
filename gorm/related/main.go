package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//
//type User struct {
//	gorm.Model
//	Profile   Profile `gorm:"ForeignKey:ProfileID"` //指定ProfileID作为外键
//	ProfileID int
//}
//type Profile struct {
//	gorm.Model
//	Name string
//}

//type Profile struct {
//	gorm.Model
//	Name   string
//	UserID uint
//}
//
//type User struct {
//	gorm.Model
//	Refer   string
//	Profile Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
//}

type User struct {
	ID      int
	Name    string
	Email   Email `gorm:"ForeignKey:EmailID"`
	EmailID int
}

type Email struct {
	ID     int
	EmailS string
}

type Language struct {
	ID int
	//Language []Language `gorm:"many2many:email_language"` //多对多
	Name string
}

type Person struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{}, &Email{}, &Person{})
	//db.Model(&User{}).Related(&Language{}, "Languages")
	//u := User{Name: "NAME", Email: Email{EmailS: "2131@QQ.COM"}}
	//
	//if e := db.Save(&u).Error; e != nil {
	//	fmt.Println(e.Error())
	//}

	// Query
	//var email Email
	//db.Model(&u).Related(&email)

	//var email1 Email
	//db.Model(&u).Association("Email").Find(&email1)
	//fmt.Println(email1)
	//u:=User{}
	//fmt.Println(db.Model(&u).Association("Email").Error)
	//fmt.Println(db.Model(&u).Association("Email").Find(&email1).Error)
	//fmt.Println(email1)

	//db.Find(&user)
	//db.Last(&user)

	//db.Where("id = ?","8").Find(&user)

	//不等于
	//db.Where("name <> ?","NAME1").Find(&user)

	//db.Where("id in (?)",[]string{"7","9"}).Find(&user)

	//db.Where("name like ?","%N%").Find(&user)

	//db.Where("name like ? And id>=?","N%",8).Find(&user)
	//db.Where("id BETWEEN ? AND ?",6,9).Find(&user)

	//db.Where(&User{ID:9}).Find(&user)

	//db.Where(map[string]interface{}{"id": 9, "name": "NAME"}).Find(&user)

	//db.Where([]int64{7,9}).Find(&user).Offset(99).Limit(4).Order()
	//var p = []Person{}
	////db.Select("name,age").Find(&p)
	//var count int
	//db.Table("people").Select("COALESCE(age,?)", 32).Count(&count)
	//fmt.Println(count)

	//var name []string
	//db.Model(&Person{}).Pluck("name", &name)
	//
	//fmt.Println(name)

	//r := Result{}
	//db.Model(&Person{}).Select("name,age").Where("id = ?", 2).Scan(&r)
	//fmt.Println(r)

	r, err := db.Table("users").Select("users.name,emails.email_s").Joins("left join emails on emails.id=users.email_id").Rows()

	if err != nil {
		fmt.Println(err.Error())
	}
	for r.Next() {

		r1 := Result1{}
		r.Scan(&r1.Name, &r1.EmailS)

		fmt.Println(r1)

	}

	db.Scopes()

	db.Table("result_table").CreateTable(&Result{})

	db.Preload()
}

type Result struct {
	Name string
	Age  int
}
type Result1 struct {
	Name   string
	EmailS string
}
