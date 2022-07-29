package abstractfactory

import (
	"gof/creative/abstractfactory/color"
	"gof/creative/abstractfactory/shape"
)

type AbstractFactory interface {
	getShape(name string) shape.Shape
	getColor(name string) color.Color
}