package color

import "fmt"

type Yellow struct {
}

func (y *Yellow) Fill() {
	fmt.Println("Fill yellow color.")
}
