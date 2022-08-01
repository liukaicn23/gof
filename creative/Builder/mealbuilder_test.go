package Builder

import (
	"testing"
)

func TestMealBuilder_Main(t *testing.T) {
	m := MealBuilder{}
	if got := m.PrepareNoVegMeal(); got == nil {
		t.Errorf("PrepareNoVegMeal() error: %v\n", got)
	} else {
		t.Logf("PrepareNoVegMeal() get Cost: %v\n", got.getCost())
	}

	if got := m.PrepareVegMeal(); got == nil {
		t.Errorf("PrepareVegMeal() error: %v\n", got)
	} else {
		t.Logf("PrepareVegMeal() get Cost: %v\n", got.getCost())
	}
}
