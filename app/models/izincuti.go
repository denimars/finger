package models

import "time"

type Izincuti struct{
	Id int `gorm:"int(11); primary_key; not_null; auto_increment"`
	PegawaiId int `gorm:"int(11)"`
	JenisIzin string `gorm:"varchar(50)"`
	Mulai time.Time `gorm:"datetime"`
	Berakhir time.Time `gorm:"datetime"`
}