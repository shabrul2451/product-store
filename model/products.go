package model

type Product struct {
	Id             string  `bson:"id" json:"id"`
	Name           string  `bson:"name" json:"name"`
	Description    string  `bson:"description" json:"description"`
	Specifications string  `bson:"specifications" json:"specifications"`
	BrandId        string  `bson:"brand_id" json:"brand_id"`
	CategoryId     string  `bson:"category_id" json:"category_id"`
	SupplierId     string  `bson:"supplier_id" json:"supplier_id"`
	UnitPrice      float64 `bson:"unit_price" json:"unit_price"`
	DiscountPrice  float64 `bson:"discount_price" json:"discount_price"`
	Tags           string  `bson:"tags" json:"tags"`
	StatusId       int     `bson:"status_id" json:"status_id"`
}

type ProductDto struct {
	Id             string  `bson:"id" json:"id"`
	Name           string  `bson:"name" json:"name"`
	Description    string  `bson:"description" json:"description"`
	Specifications string  `bson:"specifications" json:"specifications"`
	BrandId        string  `bson:"brand_id" json:"brand_id"`
	CategoryId     string  `bson:"category_id" json:"category_id"`
	SupplierId     string  `bson:"supplier_id" json:"supplier_id"`
	UnitPrice      float64 `bson:"unit_price" json:"unit_price"`
	DiscountPrice  float64 `bson:"discount_price" json:"discount_price"`
	Tags           string  `bson:"tags" json:"tags"`
	StatusId       int     `bson:"status_id" json:"status_id"`
	StockQuantity  int64   `json:"stock_quantity"`
}

func (p ProductDto) ConvertToProduct() Product {
	product := Product{
		Id:             p.Id,
		Name:           p.Name,
		Description:    p.Description,
		Specifications: p.Specifications,
		BrandId:        p.BrandId,
		CategoryId:     p.CategoryId,
		SupplierId:     p.SupplierId,
		UnitPrice:      p.UnitPrice,
		DiscountPrice:  p.DiscountPrice,
		Tags:           p.Tags,
		StatusId:       p.StatusId,
	}
	return product
}
