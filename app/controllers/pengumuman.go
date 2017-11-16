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

func(c Pengumuman) Push(id int, judul string, pesan string, mulai string, berakhir string) revel.Result{
	pengumuman := c.GetPengumuman(id)
	pengumuman.Judul = strings.ToUpper(strings.TrimSpace(judul))
	pengumuman.Pesan = strings.ToUpper(strings.TrimSpace(pesan))
	pengumuman.StartTime = task.ToSQLDT(strings.TrimSpace(mulai))
	pengumuman.EndTime	= task.ToSQLDT(strings.TrimSpace(berakhir))
	err := app.DB.Save(&pengumuman)
	if err.Error != nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Pengumuman.Index())
}

func(c Pengumuman) Edit(id int) revel.Result{
	pengumuman := c.GetPengumuman(id)
	return c.Render(pengumuman)
}

func(c Pengumuman) Delete(id int) revel.Result{
	var pengumuman m.Pengumuman
	err := app.DB.Where("Id=?", id).Delete(&pengumuman)
	if err.Error != nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Pengumuman.Index())
}

func(c Pengumuman) GetPengumuman(id int) *m.Pengumuman{
	var pengumuman m.Pengumuman
	app.DB.Find(&pengumuman, id)
	return &pengumuman
}