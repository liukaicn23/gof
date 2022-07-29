package shape

import "fmt"

type Rectangle struct {
}

func (c *Rectangle) draw() {
	fmt.Println("This is Retangle.")
}
