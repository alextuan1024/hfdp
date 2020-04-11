package condiment

import (
	"github.com/alextuan1024/hfdp/decorator/interfaces"
)

type Mocha struct {
	interfaces.Beverage
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float64 {
	return .20 + m.Beverage.Cost()
}
