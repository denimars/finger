package models

import "github.com/revel/revel"

type Mesin struct{
	Id int `gorm:"int(10); auto_increment; primary_key"`
	Kode string `gorm:"varchar(10)"`
	Ip string `gorm:"varchar(15)"`
}

func(mesin *Mesin) Validate(v *revel.Validation){
	v.Required(mesin.Kode).Message("*harus diisi")
	v.MaxSize(mesin.Kode, 10).Message("*max 10")

	v.Required(mesin.Ip).Message("*harus diisi")
	v.MinSize(mesin.Ip, 8).Message("*min 7")
	v.MaxSize(mesin.Ip, 15).Message("*max 15")
}