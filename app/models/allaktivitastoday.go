package models


import "time"

type Allaktivitastoday struct{
	Id int `gorm:"int(10)"`
	Nama string `gorm:"varchar(50)"`
	Jabatan string `gorm:"varchar(50)"`
	Tanggal string `gorm:"date"`
	Datang time.Time `gorm:"datetime"`
	Pulang time.Time `gorm:"datetime"`
	Keterangan string `gorm:"varchar(50)"`
}