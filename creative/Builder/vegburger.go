package Builder

type VegBurger struct {
	BurgerBase
}

func (b *VegBurger) Name() string {
	return "Veg Burger"
}

func (b *VegBurger) Price() float32 {
	return 50.5
}

