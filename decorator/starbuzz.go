package main

import (
	"fmt"
	"github.com/alextuan1024/hfdp/decorator/coffee"
	"github.com/alextuan1024/hfdp/decorator/condiment"
	"github.com/alextuan1024/hfdp/decorator/interfaces"
)

func main() {

	var esp interfaces.Beverage
	esp = new(coffee.Espresso)
	esp = &condiment.Soy{
		Beverage: esp,
	}
	fmt.Printf("%s $%f\n", esp.GetDescription(), esp.Cost())

	var bev interfaces.Beverage
	bev = &coffee.DarkRoast{}
	bev = &condiment.Mocha{
		Beverage: bev,
	}
	bev = &condiment.Whip{
		Beverage: bev,
	}
	bev = &condiment.Mocha{
		Beverage: bev,
	}

	fmt.Printf("%s $%f\n", bev.GetDescription(), bev.Cost())
}
