package Builder

type Item interface {
	Name() string
	Packing() Packing
	Price() float32
}
