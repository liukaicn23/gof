package abstractfactory

type FactoryProducer struct {
}

const SHAPE = "SHAPE"
const COLOR = "COLOR"

func (fp *FactoryProducer) getFactory(name string) AbstractFactory {
	var af AbstractFactory

	switch name {
	case "SHAPE":
		af = new(ShapeFactory)
	case "COLOR":
		af = new(ColorFactory)
	}

	return af
}
