package Builder

type Meal struct {
	items []Item
}

func (m *Meal) addItem(item Item) {
	m.items = append(m.items, item)
}

func (m *Meal) getCost() float32 {
	var cost float32 = 0.0
	for _, item := range m.items {
		if item != nil {
			cost = cost + item.Price()
		}
	}
	return cost
}

func InitMeal() *Meal {
	meal := Meal{
		items: make([]Item, 0, 100),
	}
	return &meal
}
