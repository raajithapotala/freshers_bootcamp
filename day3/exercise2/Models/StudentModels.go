//Models/UserModel.go
package Models
type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	LastName string `json: "lastName"`
	DOB      string ` json: "DOB"`
	Address string `json:"address"`
	Subject   string `json:"subject"`
	Marks    int `json:"marks"`
}
func (b *User) TableName() string {
	return "user"
}