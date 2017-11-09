package models

import "time"

type Pengumuman struct{
	Id int `gorm:"int(11); primary_key; not_null; auto_increment"`
	Judul string `gorm:"varchar(50)"`
	Pesan string `gorm:"text"`
	StartTime time.Time `gorm:"datetime"`
	EndTime time.Time `gorm:"datetime"`
}