package controllers

import(
	"github.com/revel/revel"
)

type Jurusan struct{
	App
}

func(c Jurusan) Index() revel.Result{
	return c.Render()
}