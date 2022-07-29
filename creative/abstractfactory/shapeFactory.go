package abstractfactory

import (
	"gof/creative/abstractfactory/color"
	"gof/creative/abstractfactory/shape"
)

type ShapeFactory struct {
}

const SHAPE_CIRCLE = "CIRCLE"
const SHAPE_RECTANGLE = "RECTANGLE"
const SHAPE_SQUARE = "SQUARE"

func (s *ShapeFactory) getColor(colorName string) color.Color {
	return nil
}

func (s *ShapeFactory) getShape(shapeName string) shape.Shape {
	var sha shape.Shape

	switch shapeName {
	case SHAPE_CIRCLE:
		sha = new(shape.Cricle)
	case SHAPE_RECTANGLE:
		sha = new(shape.Rectangle)
	case SHAPE_SQUARE:
		sha = new(shape.Square)
	}

	return sha
}
