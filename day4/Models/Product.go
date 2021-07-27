package Models

import (
	"freshers_bootcamp/day4/Config"
	global "freshers_bootcamp/day4/Global"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllProducts.... Fetch all products data
func GetAllProducts(product *[]Product) (err error) {
	global.Mutex.Lock()
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	defer global.Mutex.Unlock()
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
	global.Mutex.Lock()
	Config.DB.Save(product)
	defer global.Mutex.Unlock()
	return nil
}

//DeleteProduct ... Delete product
func DeleteProduct(product *Product, Id string) (err error) {
	global.Mutex.Lock()
	Config.DB.Where("id = ?", Id).Delete(product)
	defer global.Mutex.Unlock()
	return nil
}
