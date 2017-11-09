package models

import "github.com/revel/revel"

type Jurusan struct{
	Id int `gorm:"int(10); auto_increment; primary_key"`
	Kode string `gorm:"varchar(10)"`
	Nama string `gorm:"varchar(50)"`
}

func (jurusan *Jurusan) Validate(v *revel.Validation){
	v.Check(
		jurusan.Kode,
		revel.MinSize{3},
		revel.MaxSize{50},
	)

	v.Check(
		jurusan.Nama,
		revel.MinSize{5},
		revel.MaxSize{50},
	)
}

