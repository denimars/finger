package controllers

import(
	"github.com/revel/revel"
	"finger/app"
	"finger/app/models"
	"log"
)


type Allaktivitastoday struct{
	App
}

func(c Allaktivitastoday) Index() revel.Result{
	var aktivitas []models.Allaktivitastoday
	err := app.DB.Find(&aktivitas)
	if err != nil{
		log.Println(err)
	}
	return c.Render(aktivitas)
}