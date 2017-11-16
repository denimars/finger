package controllers


import(
	"github.com/revel/revel"
	"finger/app/routes"
	"finger/app"
	m "finger/app/models"
	"strings"
	"finger/app/task"

)

type Pengumuman struct{
	App
}

func(c Pengumuman) Index() revel.Result{
	var pengumuman []m.Pengumuman
	err := app.DB.Find(&pengumuman)
	if err.Error !=nil{
		panic(err.Error)
	}
	return c.Render(pengumuman)
}

func(c Pengumuman) Save(pengumuman m.Pengumuman) revel.Result{
	pengumuman.Validate(c.Validation)
	if c.Validation.HasErrors(){
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(routes.Pengumuman.Index())
	}
	
	pengumuman.Judul = strings.ToUpper(strings.TrimSpace(pengumuman.Judul))
	pengumuman.Pesan = strings.ToUpper(strings.TrimSpace(pengumuman.Pesan))
	pengumuman.StartTime = task.ToSQLDT(strings.TrimSpace(pengumuman.StartTime))
	pengumuman.EndTime	= task.ToSQLDT(strings.TrimSpace(pengumuman.EndTime))
	pengumuman.Uid = 0

	err := app.DB.Save(&pengumuman)
	if err.Error != nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Pengumuman.Index())
}