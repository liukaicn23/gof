package color

import "fmt"

type Green struct {
}

func (g *Green) Fill() {
	fmt.Println("Fill green color.")
}
