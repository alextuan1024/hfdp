package custom

type WeatherData struct {
	Observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.Observers = append(w.Observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {

	for i, v := range w.Observers {
		if v == o {
			if i < len(w.Observers)-1 {
				copy(w.Observers[i:], w.Observers[i+1:])
			}
			w.Observers[len(w.Observers)-1] = nil
			w.Observers = w.Observers[:len(w.Observers)-1]
		}
	}
}

func (w *WeatherData) notifyObservers() {
	for _, ob := range w.Observers {
		ob.update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementChanged() {
	w.notifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	w.temperature, w.humidity, w.pressure = temperature, humidity, pressure
	w.MeasurementChanged()
}
