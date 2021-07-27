package Models

import (
	"freshers_bootcamp/day4/Global"
	_ "github.com/go-sql-driver/mysql"

	"freshers_bootcamp/day4/Config"
)

//GetAllOrders.... Fetch all order data
func GetAllOrders(order *[]Order) (err error) {
	global.Mutex.Lock()
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	defer global.Mutex.Unlock()
	return nil
}

//CreateOrder ... Insert New data
func CreateOrder(order *Order) (err error) {
	global.Mutex.Lock()
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	defer global.Mutex.Unlock()
	return nil
}

//GetOrderByID ... Fetch only one order by Id
func GetOrderByID(order *Order, Id string) (err error) {
	if err = Config.DB.Where("id = ?", Id).First(order).Error; err != nil {
		return err
	}
	return nil
}
