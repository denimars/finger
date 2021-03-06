package models

import "github.com/revel/revel"

type Jurusan struct{
	Id int `gorm:"int(10); auto_increment; primary_key"`
	Kode string `gorm:"varchar(10)"`
	Nama string `gorm:"varchar(50)"`
}

func (jurusan *Jurusan) Validate(v *revel.Validation){
	v.Required(jurusan.Kode).Message("*harus diisi")
	v.MaxSize(jurusan.Kode, 10).Message("*max 10")

	v.Required(jurusan.Nama).Message("*harus diisi")
	v.MinSize(jurusan.Nama, 2).Message("*min 2")
	v.MaxSize(jurusan.Nama, 50).Message("*max 10")
}

