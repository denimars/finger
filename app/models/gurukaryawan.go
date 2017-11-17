package models

import "github.com/revel/revel"


type Gurukaryawan struct{
	Id int `gorm:"int(10); not_null; auto_increment; primary_key"`
	Nama string `gorm:"varchar(50)"`
	JurusanId int `gorm:"int(11)"`
	Uid int `gorm:"int(11)"`
	Jabatan string `gorm:"varchar(30)"`
}

func(gurukaryawan Gurukaryawan) Validate(v *revel.Validation){
	v.Required(gurukaryawan.Nama).Message("*harus diisi")
	v.MinSize(gurukaryawan.Nama, 3).Message("*min 3")
	v.MaxSize(gurukaryawan.Nama, 50).Message("*min 50")
	
	v.Required(gurukaryawan.Uid).Message("*harus diisi")
	v.MaxSize(gurukaryawan.Uid, 11).Message("*max 11")
}


