package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql"
)

type User struct {
	gorm.Model
	Birthday   time.Timer
	Age        int
	Name       string `gorm:"size:255"`
	Num        int    `gorm:"AUTO_INCREMENT"`
	CreditCard CreditCard //one to one
	Email      []Email

	BillingAddress Address

	BillingAddressID sql.NullInt64

	ShippingAddress   Address
	shippingAddressID int
	IgnoreMe          int        `gorm:"-"`
	Languages         []Language `gorm:"many2many:user_languages;"` //Many-To-Many , 'user_languages'是连接表
}

type Language struct {
	ID   int                                     //字段ID 默认为主键
	Name string `gorm:"index:idx_language_name"` //创建索引并命名，如果找到其他相同名称的索引则会创建组合索引
	Code string `gorm:"index:idx_language_code"` //unique_index also works
}
type CreditCard struct {
	gorm.Model
	UserID int
	Number string
}
type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          //外键（属于） ，tag `index` 是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` //type设置为sql类型  unique_index 为该列的唯一索引
	Subscribed bool
}
type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` //设置字段不能为空，并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}
