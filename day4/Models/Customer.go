package Models

import (
	"freshers_bootcamp/day4/Config"
	"freshers_bootcamp/day4/Global"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllCustomers.... Fetch all customers data
func GetAllCustomers(customer *[]Customer) (err error) {
	global.Mutex.Lock()
	if err = Config.DB.Find(customer).Error; err != nil {
		return err
	}
	defer global.Mutex.Unlock()
	return nil
}

//CreateCustpmer...Insert New data
func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

//GetCustomerByID ... Fetch only one product by Id
func GetCustomerByID(customer *Customer, Id string) (err error) {
	if err = Config.DB.Where("id = ?", Id).First(customer).Error; err != nil {
		return err
	}
	return nil
}
