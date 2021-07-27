package Controllers

import (
	"fmt"
	"net/http"

	"freshers_bootcamp/day4/Models"
	"github.com/gin-gonic/gin"
)

//GetCustomers ... Get all customers
func GetCustomers(c *gin.Context) {
	var customer []Models.Customer
	err := Models.GetAllCustomers(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,customer)
	}
}

//Create Customer

func CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Models.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,customer)
	}
}

//Get the Order by id

func GetCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer Models.Customer
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
