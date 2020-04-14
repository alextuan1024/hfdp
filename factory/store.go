package main

import (
	"fmt"
	"github.com/alextuan1024/hfdp/factory/pizza"
)

func main() {
	var nyStore, chicagoStore pizza.Store
	nyStore = new(pizza.NYPizzaStore)
	chicagoStore = new(pizza.ChicagoPizzaStore)

	var nyCheese, chicagoCheese pizza.Pizza
	nyCheese = nyStore.OrderPizza("cheese")
	fmt.Printf("Ethan ordered a %s\n", nyCheese.GetName())

	chicagoCheese = chicagoStore.OrderPizza("cheese")
	fmt.Printf("Joel orderd a %s\n", chicagoCheese.GetName())
}
