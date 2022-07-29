package Builder

type ColdDrink interface {
	Item
}

type ColdDrinkBase struct {
}

func (c *ColdDrinkBase) Packing() Packing {
	return &Bottle{}
}
