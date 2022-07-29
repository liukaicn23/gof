package singleton

import (
	"fmt"
	"reflect"
	"testing"
)

type TestSingleton struct {
}

func (singleton *TestSingleton) Eating() {
	fmt.Println("Eating")
}

func add[T any](a, b T) (c T) {
	fmt.Println(a)
	fmt.Println(b)
	return
}

func TestGetInstance(t *testing.T) {
	add[int](1, 2)
	t.Run("Eating", func(t *testing.T) {
		var got *Singleton
		if got = GetInstance(); !reflect.DeepEqual(got, &Singleton{}) {
			t.Errorf("GetInstance() = %v, want %v", got, &Singleton{})
		}
	})
}
