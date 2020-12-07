package src

import (
	"fyne.io/fyne"
	"time"
)

func thisDay0000(t time.Time) time.Time {
	y,m,d := t.Year(), t.Month(), t.Day()
	return time.Date(y,m,d,0,0,0,0,time.Local)
}

var today0000 = thisDay0000(time.Now())

var (
	A fyne.App      // приложение
	W fyne.Window   // главное окно
)