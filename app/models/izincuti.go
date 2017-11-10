package models

import (
	"github.com/revel/revel"
)

type Izincuti struct{
	Id int `gorm:"int(11); primary_key; not_null; auto_increment"`
	PegawaiId int `gorm:"int(11)"`
	JenisIzin string `gorm:"varchar(50)"`
	Mulai string `gorm:"datetime"`
	Berakhir string `gorm:"datetime"`
}

func (izincuti Izincuti) Validate(v *revel.Validation){
	v.Required(izincuti.Mulai).Message("*harus diisi")

	v.Required(izincuti.Berakhir).Message("*harus diisi")
}