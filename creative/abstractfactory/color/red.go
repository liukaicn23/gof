package color

import "fmt"

type Red struct {
}

func (r *Red) Fill() {
	fmt.Println("Fill red color.")
}
