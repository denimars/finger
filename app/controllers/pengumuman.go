package controllers


import(
	"github.com/revel/revel"
	"finger/app/routes"
	"finger/app"
	m "finger/app/models"
	"finger/app/tugas"
	"strings"

)

type Pengumuman struct{
	App
}

func(c Pengumuman) Index() revel.Result{
	return c.Render()
}

func(c Pengumuman) Save(pengumuman m.Pengumuman) revel.Result{
	pengumuman.Judul = strings.ToUpper(strings.TrimSpace(pengumuman.Judul))
	pengumuman.Pesan = strings.ToUpper(strings.TrimSpace(pengumuman.Pesan))
	pengumuman.StartTime = tugas.ToSQLDT(strings.TrimSpace(pengumuman.StartTime))
	pengumuman.EndTime	= tugas.ToSQLDT(strings.TrimSpace(pengumuman.EndTime))
	err := app.DB.Save(&pengumuman)
	if err.Error != nil{
		panic(err.Error)
	}
	return c.Redirect(routes.Pengumuman.Index())
}