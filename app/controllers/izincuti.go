package controllers

import (
	"github.com/revel/revel"
	m "finger/app/models"
	"time"
	"finger/app/routes"
	"finger/app"
	"log"
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

func(c Izincuti) Save(mulai string) revel.Result{
	a, _ := time.Parse(FORM_DATE_CON, mulai)
	zz := a.Format(SQL_DATE_CON)
	var izincuti m.Izincuti
	izincuti.PegawaiId = 20
	izincuti.JenisIzin = "ayam bakar"
	izincuti.Mulai = zz
	izincuti.Berakhir = zz
	log.Println("_________________________")
	log.Println(zz)
	log.Println(a)
	log.Println("__________________________")
	err := app.DB.Save(&izincuti)
	if err.Error !=nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Izincuti.Index())
}
