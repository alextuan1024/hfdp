package pizza

import "fmt"

type Store interface {
	OrderPizza(string) Pizza
	CreatePizza(string) Pizza
}

type Preparer interface {
	Prepare()
}

type Baker interface {
	Bake()
}

type Cutter interface {
	Cut()
}

type Boxer interface {
	Box()
}

type Pizza interface {
	Preparer
	Baker
	Cutter
	Boxer
	GetName() string
}

func defaultPrepare(name string) {
	fmt.Printf("Preparing %s\n", name)
}

func defaultBake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func defaultCut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func defaultBox() {
	fmt.Println("Place pizza in official PizzaStore box")
}
