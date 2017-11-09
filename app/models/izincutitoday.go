package models

import "time"

type Izincutitoday struct{
	Id int `gorm:"int(10); not_null; auto_increment"`
	Tanggal time.Time `gorm:"datetime"`
	JenisIzin string `gorm:"varchar(50)"`
}