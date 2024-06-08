package models

type Product struct {
	Id               int32              `json:"id" gorm:"primaryKey"`
	Name             string             `json:"name"`
	Sku              string             `json:"sku"`
	Category         string             `json:"category"`
	Description      string             `json:"description"`
	Stock            int32              `json:"stock"`
	Price            int32              `json:"price"`
	StockDecreaseLog []StockDecreaseLog `gorm:"foreignKey:Product"`
}
