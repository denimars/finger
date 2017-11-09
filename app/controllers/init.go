package controllers

import(
	"github.com/revel/revel"
	"time"
)

const(
	SQL_DATE = "2006-01-02T15:04:05Z07:00"
	SHOW_DATE = "02-01-2006"
	SQL_DATETIME = "2006-01-02 15:04:05"
	SHOW_DATETIME = "15:04:05"
)

func init(){
	revel.TemplateFuncs["indonesiadate"] = func(value string) string{
		a, _ := time.Parse(SQL_DATE, value)
		return a.Format(SHOW_DATE)
	}

	revel.TemplateFuncs["indonesiadatetime"] = func(value time.Time) string{
		a,_:= time.Parse(SQL_DATETIME, value.String())
		return a.Format(SHOW_DATETIME)
	}
	
}