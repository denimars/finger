package controllers

import(
	"github.com/revel/revel"
	m "finger/app/models"
	"finger/app/routes"
	"finger/app"
	"strings"
	//"log"
)

type Jurusan struct{
	App
}

func(c Jurusan) Index() revel.Result{
	jurusan := GetAllJurusan()
	return c.Render(jurusan)
}

func(c Jurusan) Save(jurusan m.Jurusan) revel.Result{
	jurusan.Validate(c.Validation)
	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Jurusan.Index())
	}
	jurusan.Kode = strings.ToUpper(strings.TrimSpace(jurusan.Kode))
	jurusan.Nama = strings.ToUpper(strings.TrimSpace(jurusan.Nama))
	err := app.DB.Save(&jurusan)
	if err.Error !=nil{
		panic(err.Error)
	}
	//c.Flash.Success("Berhasil Disimpan")
	//c.Flash.Out["pesan"]="Berhasil Disimpan"
	return c.Redirect(routes.Jurusan.Index())
}

func(c Jurusan) Edit(id int) revel.Result{
	jurusan := c.GetJurusan(id)
	return c.Render(jurusan)
}

func(c Jurusan) Push(id int, kode string, nama string) revel.Result{
	jurusan := c.GetJurusan(id)
	jurusan.Kode = strings.ToUpper(strings.TrimSpace(kode))
	jurusan.Nama = strings.ToUpper(strings.TrimSpace(nama))
	err := app.DB.Save(&jurusan)
	if err.Error!=nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Jurusan.Index())
}

func(c Jurusan) Delete(id int) revel.Result{
	var jurusan m.Jurusan
	err := app.DB.Where("Id=?", id).Delete(&jurusan)
	if err.Error != nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Jurusan.Index())
}

func(c Jurusan) GetJurusan(id int) *m.Jurusan{
	var jurusan m.Jurusan
	app.DB.Find(&jurusan, id)
	return &jurusan
}

func GetAllJurusan() []m.Jurusan{
	var jurusan []m.Jurusan
	app.DB.Find(&jurusan)
	return jurusan
}