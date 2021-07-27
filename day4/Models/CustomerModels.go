//Models/UserModel.go
package Models

type Customer struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

func (b *Customer) TableName() string {
	return "Customer"
}
