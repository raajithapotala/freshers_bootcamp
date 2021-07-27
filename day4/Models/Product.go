package Models

import (
	"fmt"
	"freshers_bootcamp/day4/Config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllProducts.... Fetch all products data
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateProduct ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetProductByID ... Fetch only one product by Id
func GetProductByID(product *Product, Id string) (err error) {
	if err = Config.DB.Where("id = ?", Id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProduct ... Update product
func UpdateProduct(product *Product, Id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

//DeleteProduct ... Delete product
func DeleteProduct(product *Product, Id string) (err error) {
	Config.DB.Where("id = ?", Id).Delete(product)
	return nil
}
