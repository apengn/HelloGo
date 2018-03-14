package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"os"
)

//import (
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jinzhu/gorm"
//	"os"
//	"fmt"
//	"net/http"
//)
//
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
}

func main1() {
	db, err := gorm.Open("mysql", "root:root@/gorm?charset=utf8&parseTime=True&loc=Local")
	if os.IsExist(err) {
		panic(err)
	}

	//defer db.Close()
	//建表
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})

	//创建
	db.Create(&Product{Code: "L1212", Price: 10000})

	var product Product
	db.First(&product, 1) //查询Id =1的Product

	//db.First(&product, "code=?", "L1212") //查询code为L1212的Product
	//
	////更新
	//db.Model(&product).Update("Price", 2000)
	//
	////删除
	//db.Delete(&product)

	fmt.Println(product)


}

func main() {
	http.HandleFunc("/hello", msgHand)
	http.ListenAndServe(":8089", nil)
}

func msgHand(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello HandlerFunc")
}
