package models


type Jurusan struct{
	Id int `gorm:"int(10); auto_increment; primary_key"`
	Kode string `gorm:"varchar(10)"`
	Nama string `gorm:"varchar(50)"`
}