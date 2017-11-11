package controllers

import (
	"github.com/revel/revel"
	m "finger/app/models"
	"finger/app/routes"
	"finger/app"
	"finger/app/tugas"
	"strings"
)

const(
	SQL_DATE_CON = "2006-01-02 15:04:05"
	FORM_DATE_CON = "02/01/2006 3:04 PM"
)



type Izincuti struct{
	App
}

func(c Izincuti) Index() revel.Result{
	return c.Render()
}

func(c Izincuti) Save(mulai string, akhir string) revel.Result{
	var izincuti m.Izincuti
	izincuti.PegawaiId = 20
	izincuti.JenisIzin = "ayam bakar"
	izincuti.Mulai = tugas.ToSQLDT(strings.TrimSpace(mulai))
	izincuti.Berakhir = tugas.ToSQLDT(strings.TrimSpace(akhir))
	err := app.DB.Save(&izincuti)
	if err.Error !=nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Izincuti.Index())
}