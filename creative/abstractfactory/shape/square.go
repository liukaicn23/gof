package shape

import "fmt"

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("This is Square.")
}
