package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
	"fmt"
)

var engine *xorm.Engine

type U struct {
	User
	Count int
}
type User struct {
	Id   int64
	Name string
}

type Home struct {
	Id       int64
	HomeName string
	UserId   int64
}

var err error

func main() {
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	engine.ShowSQL(true)
	engine.Sync2(new(User), new(Home))
	//user := new(User)
	//user.Name = "myname"

	has, err := engine.Exist(&User{Name:"myname",})
	fmt.Println(has, err)
	//_, err = engine.Insert(user)
	//
	//home := new(Home)
	//home.UserId=user.Id
	//home.HomeName="xxxxxxxxxxxx"
	//_, err = engine.Insert(home)
	//
	//home2 := new(Home)
	//home2.UserId=user.Id
	//home2.HomeName="xxxxxxxxxxxx"
	//_, err = engine.Insert(home2)
	//
	//home3 := new(Home)
	//home3.UserId=user.Id
	//home3.HomeName="xxxxxxxxxxxx"
	//_, err = engine.Insert(home3)
	//
	//fmt.Println(err)
	//u := []*U{}
	//engine.Table("user").Alias("u").Join("INNER", []string{"home", "h"}, "h.user_id=u.id").Where("u.name='myname'").Select("u.id")
	//fmt.Println(u)

}
