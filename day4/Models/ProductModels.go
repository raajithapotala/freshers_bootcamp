//Models/UserModel.go
package Models

type Product struct {
	Id            uint   `json:"prodId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	QuantityAvail uint   `json:"quantityAvail"`
	Price         uint   `json:"price"`
}

func (b *Product) TableName() string {
	return "product"
}
