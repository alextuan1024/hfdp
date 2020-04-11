package coffee

type HouseBlend struct {
}

func (h *HouseBlend) GetDescription() string {
	return "House Blend"
}

func (h *HouseBlend) Cost() float64 {
	return .89
}
