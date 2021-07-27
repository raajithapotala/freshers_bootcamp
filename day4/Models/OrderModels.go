//Models/UserModel.go
package Models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model

	OrderId     uint     `json:"orderId"`
	Customer    Customer `gorm:"foreignKey:CustomerId:"`
	CustomerID  uint     `gorm:"default:null"`
	Product     Product  `gorm:"foreignKey:ProductId:"`
	ProductID   uint     `gorm:"default:null"`
	Quantity    uint
	OrderStatus string
}

func (b *Order) TableName() string {
	return "orders"
}
