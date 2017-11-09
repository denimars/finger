package models

import "github.com/revel/revel"

type Jurusan struct{
	Id int `gorm:"int(10); auto_increment; primary_key"`
	Kode string `gorm:"varchar(10)"`
	Nama string `gorm:"varchar(50)"`
}

func (jurusan *Jurusan) Validate(v *revel.Validation){
	v.Required(jurusan.Kode).Message("*Harus diisi")
	v.MinSize(jurusan.Kode, 2).Message("*Min 2")
	v.MaxSize(jurusan.Kode, 10).Message("*Max 10")

	v.Required(jurusan.Nama).Message("*Harus diisi")
	v.MinSize(jurusan.Nama, 5).Message("*Min 5")
	v.MaxSize(jurusan.Nama, 50).Message("*Max 10")
}

