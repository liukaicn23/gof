package Builder

type Coke struct {
	ColdDrinkBase
}

func (c Coke) Name() string {
	return "coke"
}

func (c Coke) Price() float32 {
	return 50.0
}
