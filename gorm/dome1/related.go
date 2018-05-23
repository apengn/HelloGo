package main

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Name string
}
type Person struct {
	gorm.Model
	Profile      Profile `gorm:"ForeignKey:ProfileRefer"` //指定ProfileRefer为外键
	ProfileRefer int
}
