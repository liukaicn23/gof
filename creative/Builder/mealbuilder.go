package Builder

type MealBuilder struct {
}

func (m MealBuilder) PrepareVegMeal() *Meal {
	meal := InitMeal()
	meal.addItem(&VegBurger{})
	meal.addItem(&Coke{})
	return meal
}

func (m MealBuilder) PrepareNoVegMeal() *Meal {
	meal := InitMeal()
	meal.addItem(&ChickenBurger{})
	meal.addItem(&Coke{})
	return meal
}
