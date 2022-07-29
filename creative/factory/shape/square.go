package shape

import "fmt"

type Square struct {
}

func (s *Square) draw() {
	fmt.Println("This is Square.")
}
