package models


type Gurukaryawan struct{
	Id int `gorm:"int(10); not_null; auto_increment; primary_key"`
	Nama string `gorm:"varchar(50)"`
	JurusanId int `gorm:"int(11)"`
	Uid int `gorm:"int(11)"`
	Jabatan string `gorm:"varchar(30)"`
}