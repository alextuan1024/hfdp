package condiment

import (
	"github.com/alextuan1024/hfdp/decorator/interfaces"
)

type Soy struct {
	interfaces.Beverage
}

func (s *Soy) GetDescription() string {
	return s.Beverage.GetDescription() + ", Soy"
}

func (s *Soy) Cost() float64 {
	return .15 + s.Beverage.Cost()
}
