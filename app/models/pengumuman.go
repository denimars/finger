package models

import "github.com/revel/revel"


type Pengumuman struct{
	Id int `gorm:"int(11); primary_key; not_null; auto_increment"`
	Judul string `gorm:"varchar(50)"`
	Pesan string `gorm:"text"`
	StartTime string `gorm:"datetime"`
	EndTime string `gorm:"datetime"`
	Uid	int 
}

func(pengumuman *Pengumuman) Validate(v *revel.Validation){
	v.Required(pengumuman.Judul).Message("*harus diisi")
	v.MinSize(pengumuman.Judul, 3).Message("*min 3")
	v.MaxSize(pengumuman.Judul, 50).Message("*max 50")

	v.Required(pengumuman.Pesan).Message("*harus diisi")
	v.MinSize(pengumuman.Pesan, 3).Message("*min 3")
	v.MaxSize(pengumuman.Pesan, 255).Message("*max 255")

	v.Required(pengumuman.StartTime).Message("*harus diisi")
	v.Required(pengumuman.EndTime).Message("*harus diisi")

}