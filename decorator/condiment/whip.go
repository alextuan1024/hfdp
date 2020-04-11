package condiment

import (
	"github.com/alextuan1024/hfdp/decorator/interfaces"
)

type Whip struct {
	interfaces.Beverage
}

func (w *Whip) GetDescription() string {
	return w.Beverage.GetDescription() + ", Whip"
}

func (w *Whip) Cost() float64 {
	return .10 + w.Beverage.Cost()
}
