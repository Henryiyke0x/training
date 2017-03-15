package lexus

type Lexus struct {
	Name     string
	Price    float64
	Discount float64
	Color    string
}

func (l Lexus) ShowName() string {
	return l.Name
}

func (l Lexus) ShowPrice() float64 {
	return l.Price * l.Discount
}

func (l Lexus) ShowDiscount() float64 {
	return l.Discount
}

// func (l Lexus) ShowColor() string {
// 	return l.Color
// }