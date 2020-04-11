package custom

import "fmt"

type CurrentConditionsDisplay struct {
	Subject                         Subject
	temperature, humidity, pressure float64
}

func (c *CurrentConditionsDisplay) display() {
	fmt.Printf("Current conditions: %.2f F degrees and humidity: %.2f%%\n", c.temperature, c.humidity)
}

func (c *CurrentConditionsDisplay) update(temperature, humidity, pressure float64) {
	c.temperature, c.humidity, c.pressure = temperature, humidity, pressure
	c.display()
}
