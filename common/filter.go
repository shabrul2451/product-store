package common

type ProductFilter struct {
	Name       string
	MaxPrice   float64
	MinPrice   float64
	Brands     []string
	Category   string
	Supplier   string
	IsVerified bool
	Page       int64
	Limit      int64
}
