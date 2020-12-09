package main

import (
	"fyne.io/fyne/app"
	"github.com/onix3/train-timetable/resource"
	"github.com/onix3/train-timetable/src"
)

func main() {
	src.A = app.NewWithID("train-timetable")
	src.A.SetIcon(resource.TrainPng)

	// запросы кэшируются: при повторном запросе тех же рейсов запрос к API не осуществляется
	src.LoadCache()
	src.Last1,src.Last2 = src.LoadLastCities()

	src.W = src.A.NewWindow("Расписаньице")
	src.ScreenHeight = src.GetScreenHeight()
	src.W.SetContent(src.Content())
	src.W.CenterOnScreen()
	src.W.ShowAndRun()
}