package pizza

import "fmt"

type ChicagoStyleCheesePizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (cc *ChicagoStyleCheesePizza) Prepare() {
	defaultPrepare(cc.Name)
}

func (cc *ChicagoStyleCheesePizza) Bake() {
	defaultBake()
}

func (cc *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}

func (cc *ChicagoStyleCheesePizza) Box() {
	defaultBox()
}
func (cc *ChicagoStyleCheesePizza) GetName() string {
	return cc.Name
}

type ChicagoPizzaStore struct {
}

func (cs *ChicagoPizzaStore) OrderPizza(s string) Pizza {
	pizza := cs.CreatePizza(s)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (cs *ChicagoPizzaStore) CreatePizza(s string) Pizza {
	if s == "cheese" {
		return &ChicagoStyleCheesePizza{
			Name:     "Chicago Style Deep Dish Cheese Pizza",
			Dough:    "Extra Thick Crust Dough",
			Sauce:    "Plum Tomato Sauce",
			Toppings: []string{"Shredded Mozzarella Cheese"},
		}
	}
	return nil
}
