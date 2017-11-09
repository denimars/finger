package models


import "time"

type Attdatangpulangbydate struct{
	Uid int `gorm:"int(11)"`
	Tanggal string `gorm:"date"`
	Datang time.Time `gorm:"datetime"`
	Pulang time.Time `gorm:"datetime"`
}