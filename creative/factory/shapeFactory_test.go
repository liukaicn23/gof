package factory

import (
	"gof/creative/factory/shape"
	"reflect"
	"testing"
)

func TestShapeFactory_getShape(t *testing.T) {
	type args struct {
		shapeName string
	}
	var circle shape.Shape = &shape.Cricle{}
	var tests = []struct {
		name string
		args args
		want shape.Shape
	}{
		// TODO: Add test cases.
		{
			"Create circle is ok.",
			args{SHAPE_CIRCLE},
			circle,
		},
		{
			"Create circle is error.",
			args{SHAPE_RECTANGLE},
			circle,
		},
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
