package shape

import "fmt"

type Rectangle struct {
}

func (c *Rectangle) Draw() {
	fmt.Println("This is Retangle.")
}
