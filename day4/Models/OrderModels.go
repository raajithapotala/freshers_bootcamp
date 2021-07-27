//Models/UserModel.go
package Models

type Order struct {
	OrderId     uint     `json:"orderId"`
	Customer    Customer `gorm:"foreignKey:CustomerId:"`
	CustomerID  uint     `gorm:"default:null"`
	Product     Product  `gorm:"foreignKey:ProductId:"`
	ProductID   uint     `gorm:"default:null"`
	OrderStatus string
}

func (b *Order) TableName() string {
	return "orders"
}
