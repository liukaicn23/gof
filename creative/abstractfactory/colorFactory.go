package abstractfactory

import (
	"gof/creative/abstractfactory/color"
	"gof/creative/abstractfactory/shape"
)

type ColorFactory struct {
}

const COLOR_RED = "RED"
const COLOR_GREEN = "GREEN"
const COLOR_YELLOW = "YELLOW"

func (s *ColorFactory) getShape(shapeName string) shape.Shape {
	return nil
}

func (s *ColorFactory) getColor(colorName string) color.Color {
	var col color.Color

	switch colorName {
	case COLOR_RED:
		col = new(color.Red)
	case COLOR_GREEN:
		col = new(color.Green)
	case COLOR_YELLOW:
		col = new(color.Yellow)
	}

	return col
}
