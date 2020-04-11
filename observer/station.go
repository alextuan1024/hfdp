package main

import "github.com/alextuan1024/hfdp/observer/custom"

func main() {
	wd := new(custom.WeatherData)
	wd.Observers = make([]custom.Observer, 0) // empty slice

	dis := new(custom.CurrentConditionsDisplay)
	dis.Subject = wd
	wd.RegisterObserver(dis)

	wd.SetMeasurements(80, 65, 30.4)
	wd.SetMeasurements(82, 70, 29.2)
	wd.SetMeasurements(78, 90, 29.2)
}
