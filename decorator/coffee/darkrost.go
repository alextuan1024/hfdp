package coffee

type DarkRoast struct {
}

func (d *DarkRoast) GetDescription() string {
	return "Dark Roast"
}

func (d *DarkRoast) Cost() float64 {
	return .99
}
