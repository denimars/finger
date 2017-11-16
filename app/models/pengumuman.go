package models


type Pengumuman struct{
	Id int `gorm:"int(11); primary_key; not_null; auto_increment"`
	Judul string `gorm:"varchar(50)"`
	Pesan string `gorm:"text"`
	StartTime string `gorm:"datetime"`
	EndTime string `gorm:"datetime"`
}