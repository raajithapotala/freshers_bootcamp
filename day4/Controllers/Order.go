package Controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"freshers_bootcamp/day4/Models"
	"github.com/gin-gonic/gin"
)

//GetOrders ... Get all orders
func GetOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//Create Order

func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.GetProductByID(& order.Product, strconv.Itoa(int(order.ProductID)))
	err= Models.GetCustomerByID(& order.Customer,strconv.Itoa(int(order.CustomerID)))
	err = Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//Get the Order by id

func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderByID(&order , id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
