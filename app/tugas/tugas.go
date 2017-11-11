package tugas

import "time"

const(
	SQL_DATE_TIME = "2006-01-02 15:04:05"
	VIEW_DATE_TIME = "02/01/2006 3:04 PM"
	SQL_DATE = "2006-01-02"
	VIEW_DATE = "02/01/2006"
)

func ToSQLDT(t string) string{
	tt, _ := time.Parse(VIEW_DATE_TIME, t)
	return tt.Format(SQL_DATE_TIME)
}

func ToSQLD(t string) string{
	tt, _ := time.Parse(VIEW_DATE, t)
	return tt.Format(SQL_DATE)
}