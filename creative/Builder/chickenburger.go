package Builder

type ChickenBurger struct {
	BurgerBase
}

func (b *ChickenBurger) Name() string {
	return "Chicken Burger"
}

func (b *ChickenBurger) Price() float32 {
	return 5.5
}
