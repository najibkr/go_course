package store

type Product struct {
	ID       string
	Name     string
	Price    float64
	Discount float64
}

func (p Product) GetDiscountedPrice() float64 {
	if p.Discount == 0 {
		return p.Price
	}
	return p.Price - (p.Price * p.Discount / 100)
}
