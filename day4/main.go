//main.go
package main

import (
	"fmt"

	"freshers_bootcamp/day4/Config"
	"freshers_bootcamp/day4/Models"
	"freshers_bootcamp/day4/Routes"

	"github.com/jinzhu/gorm"
)

var (
	err error
)

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.Order{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
