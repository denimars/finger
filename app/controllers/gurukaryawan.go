package controllers

import(
	"github.com/revel/revel"
	"finger/app"
	"finger/app/models"
	"finger/app/routes"
)

type Gurukaryawan struct{
	App
}

func(c Gurukaryawan) Index() revel.Result{
	return c.Render()
}

func(c Gurukaryawan) CreateGurukaryawan(nama string, jurusanid int, jabatan string) revel.Result{
	var gurukaryawan models.Gurukaryawan
	gurukaryawan.Nama = nama
	gurukaryawan.JurusanId=jurusanid
	gurukaryawan.Jabatan = jabatan
	app.DB.Create(&gurukaryawan)
	return c.Redirect(routes.Gurukaryawan.Index())
	
}