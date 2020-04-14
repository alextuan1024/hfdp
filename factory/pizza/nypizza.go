package pizza

type NYPizzaStore struct {
}

func (ns *NYPizzaStore) OrderPizza(s string) Pizza {
	pizza := ns.CreatePizza(s)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (ns *NYPizzaStore) CreatePizza(s string) Pizza {
	if s == "cheese" {
		return &NYStyleCheesePizza{
			Name:     "NY Style Sauce and Cheese Pizza",
			Dough:    "Thin Crust Dough",
			Sauce:    "Marinara Sauce",
			Toppings: []string{"Grated Reggiano Cheese"},
		}
	}
	return nil
}

type NYStyleCheesePizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (nc *NYStyleCheesePizza) Prepare() {
	defaultPrepare(nc.Name)
}

func (nc *NYStyleCheesePizza) Bake() {
	defaultBake()
}

func (nc *NYStyleCheesePizza) Cut() {
	defaultCut()
}

func (nc *NYStyleCheesePizza) Box() {
	defaultBox()
}

func (nc *NYStyleCheesePizza) GetName() string {
	return nc.Name
}
