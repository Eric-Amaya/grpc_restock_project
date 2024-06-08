package models

type StockDecreaseLog struct {
	Id      int32 `json:"id" gorm:"primaryKey"`
	OrderId int32 `json:"orderId"`
	Product int32 `json:"product_id"`
}
