package  models

import "time"

type Attlog struct{
	Id int `gorm:"auto_increment;not_null;primary_key"`
	Uid int
	ScanTime time.Time `gorm:"datetime"`
	Mesin_id int
}