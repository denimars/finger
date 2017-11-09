package models


type Mesin struct{
	Id int `gorm:"int(10); auto_increment; primary_key"`
	Kode string `gorm:"varchar(10)"`
	Ip string `gorm:"varchar(15)"`
}