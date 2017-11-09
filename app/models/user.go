package models

import "time"

type User struct {
	Id int `gorm:"auto_increment, primary_key"`
	Login string `gorm:"varchar(50)"`
	Password string `gorm:"varchar(50)"`
	Nama string `gorm:"varchar(50)"`
	LastLogin time.Time `gorm:"datetime"`
}