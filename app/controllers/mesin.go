package controllers

import(
	"github.com/revel/revel"
	m "finger/app/models"
	"strings"
	"finger/app"
	"finger/app/routes"
)

type Mesin struct{
	App
}

func(c Mesin) Index() revel.Result{
	var mesin []m.Mesin
	app.DB.Find(&mesin)
	return c.Render(mesin)
}

func(c Mesin) Save(mesin m.Mesin) revel.Result{
	mesin.Validate(c.Validation)
	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Mesin.Index())
	}
	mesin.Kode = strings.ToUpper(strings.TrimSpace(mesin.Kode))
	mesin.Ip = strings.ToUpper(strings.TrimSpace(mesin.Ip))
	err := app.DB.Save(&mesin)
	if err.Error != nil{
		panic(err.Error)
	} 
	return c.Redirect(routes.Mesin.Index())
}