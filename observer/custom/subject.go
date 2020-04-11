package custom

type Observer interface {
	update(temperature, humidity, pressure float64)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	notifyObservers()
}

type Displayer interface {
	display()
}
