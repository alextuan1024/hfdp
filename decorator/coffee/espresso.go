package coffee

type Espresso struct {
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) Cost() float64 {
	return 1.99
}
