package abstractfactory

import (
	"gof/creative/abstractfactory/color"
	"gof/creative/abstractfactory/shape"
	"reflect"
	"testing"
)

func TestColorFactory_getColor(t *testing.T) {
	type args struct {
		colorName string
	}
	tests := []struct {
		name string
		args args
		want color.Color
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ColorFactory{}
			if got := s.getColor(tt.args.colorName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorFactory_getShape(t *testing.T) {
	type args struct {
		shapeName string
	}
	tests := []struct {
		name string
		args args
		want shape.Shape
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ColorFactory{}
			if got := s.getShape(tt.args.shapeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getShape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactoryProducer_factory(t *testing.T) {
	type args struct {
		name string
	}
	t.Run("", func(t *testing.T) {
		fp := &FactoryProducer{}
		var got1 AbstractFactory
		if got1 = fp.getFactory("SHAPE"); !reflect.DeepEqual(got1, &ShapeFactory{}) {
			t.Errorf("getFactory() = %v, want %v", got1, &ShapeFactory{})
		}
		var shape1 shape.Shape
		if shape1 = got1.getShape("CIRCLE"); !reflect.DeepEqual(shape1, &shape.Cricle{}) {
			t.Errorf("getShape() = %v, want %v", shape1, &shape.Cricle{})
		}

		shape1.Draw()

		var got2 AbstractFactory
		if got2 = fp.getFactory("COLOR"); !reflect.DeepEqual(got2, &ColorFactory{}) {
			t.Errorf("getFactory() = %v, want %v", got2, &ColorFactory{})
		}
		var color1 color.Color
		if color1 = got2.getColor("RED"); !reflect.DeepEqual(color1, &color.Red{}) {
			t.Errorf("getColor() = %v, want %v", color1, &color.Red{})
		}

		color1.Fill()

	})

}

func TestFactoryProducer_getFactory(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want AbstractFactory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := &FactoryProducer{}
			if got := fp.getFactory(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShapeFactory_getColor(t *testing.T) {
	type args struct {
		colorName string
	}
	tests := []struct {
		name string
		args args
		want color.Color
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShapeFactory{}
			if got := s.getColor(tt.args.colorName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShapeFactory_getShape(t *testing.T) {
	type args struct {
		shapeName string
	}
	tests := []struct {
		name string
		args args
		want shape.Shape
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShapeFactory{}
			if got := s.getShape(tt.args.shapeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getShape() = %v, want %v", got, tt.want)
			}
		})
	}
}
