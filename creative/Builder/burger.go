package Builder

type Burger interface {
	Item
}
type BurgerBase struct {
}

func (b BurgerBase) Packing() Packing {
	return Wrapper{}
}
