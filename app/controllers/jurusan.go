package controllers

import(
	"github.com/revel/revel"
	m "finger/app/models"
	"finger/app/routes"
	"finger/app"
	"strings"
)

type Jurusan struct{
	App
}

func(c Jurusan) Index() revel.Result{
	return c.Render()
}

func(c Jurusan) Save(jurusan m.Jurusan) revel.Result{
	jurusan.Validate(c.Validation)
	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/jurusan?error=notnull")
	}
	jurusan.Kode = strings.ToUpper(jurusan.Kode)
	jurusan.Nama = strings.ToUpper(jurusan.Nama)
	err := app.DB.Save(&jurusan)
	if err.Error !=nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Jurusan.Index())
}