package Models

import (
	"freshers_bootcamp/day4/Config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllOrders.... Fetch all order data
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

//CreateOrder ... Insert New data
func CreateOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

//GetOrderByID ... Fetch only one order by Id
func GetOrderByID(order *Order, Id string) (err error) {
	if err = Config.DB.Where("id = ?", Id).First(order).Error; err != nil {
		return err
	}
	return nil
}
