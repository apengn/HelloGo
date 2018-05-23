package main

import (
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Dog struct {
	AnimalID int `gorm:"primary_key"` //设置主键
}

type Shop struct {
	gorm.Model
}
type Product struct {
	gorm.Model
	Code  string
	Price uint `gorm:"column:price"` //重重新设置列名
}

type User1 struct {
	//gorm.Model
	ID    int
	Code  string
	Price uint `gorm:"column:price"` //重重新设置列名
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local")

	if os.IsExist(err) {
		panic(err)
	}
	//defer db.Close()
	//自动迁移模式   注意点：自动迁移模式将保持更新到最新
	//自动迁移模式仅仅会创建表，缺少列和索引，并不会 改变现有列的类型或删除未使用的列以保护数据
	//db.AutoMigrate(&Product{}, &User{})
	//db.CreateTable(&User{})

	//db.HasTable("users") //db.HasTable(&User{}) //判断是否含有user模型表

	//db.CreateTable(&User1{})

	//db.Create(&User1{5,"3232",10020})
	//fmt.Println()

	err = db.Where("ID=?",2).Delete(&User1{}).Error
	fmt.Println(err)
	////创建表  并把    ENGINE=INNODB附加到sql语句 中
	//db.Set("gorm:table_options", "ENGINE=INNODB").CreateTable(&User{})
	//
	//db.DropTable("users") //db.DropTable(&User{})//删除表
	//
	//db.DropTableIfExists("users", &Product{}) //如果表存在的就   删除多张表
	//
	//db.Model(&Product{}).ModifyColumn("code", "text") //修改列的数据类型
	//
	//db.Model(&Product{}).DropColumn("code") //删除列
	//
	////添加外键       product_id 外键的id       products(id)  所对应表的主键id   添加表的约束   ，添加表的约束
	//db.Model(&Shop{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	//
	////删除外键
	//db.Model(&Product{}).RemoveForeignKey("product_id", "products(id)")
	//
	////添加索引
	//db.Model(&Product{}).AddIndex("idx_product_code", "code")
	////为多列添加索引
	//db.Model(&Product{}).AddIndex("idx_product_code_price", "code", "price")
	////为多列添加唯一索引
	//db.Model(&Product{}).AddUniqueIndex("idx_product_code_price", "code", "price")
	////添加唯一索引
	//db.Model(&Product{}).AddUniqueIndex("idx_product_code", "code")
	////删除索引
	//db.Model(&Product{}).RemoveIndex("idx_product_code")
	//
	//db.SingularTable(true) //全局禁用表名复数
	//
	////更改默认的表名
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "prefix_" + defaultTableName
	//}
}
