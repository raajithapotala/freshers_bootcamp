package Controllers

import (
	"fmt"
	"time"

	"freshers_bootcamp/day4/Config"
	"freshers_bootcamp/day4/Models"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	// add getbyuserid
}

//Create Order

func CreateOrder(c *gin.Context) {
	var order Models.Order
	err := c.BindJSON(&order)
	if err != nil {
		return
	}

	coolDown := CheckCoolDown(int(order.CustomerID))
	if coolDown == false{
		c.JSON(http.StatusOK,gin.H{
			"message":"Please wait till Cooldown time of 1 minute",
		})
		return
	}
	possible := checkfeasibilty(order,c)
	if possible == false{
		c.JSON(http.StatusOK,gin.H{
			"status":"order failed",
		})
		return
	}
	order.OrderStatus = "Order placed"
	err = Models.GetProductByID(&order.Product, strconv.Itoa(int(order.ProductID)))
	err = Models.GetCustomerByID(&order.Customer, strconv.Itoa(int(order.CustomerID)))
	err = Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":order.OrderId,
			"product_id":order.ProductID,
			"quantity":order.Quantity,
			"status":order.OrderStatus,
		})
		//t = c
		//channel <- order.Id
	}
}

//Get the Order by id

func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order

	err := Models.GetProductByID(&order.Product, strconv.Itoa(int(order.ProductID)))
	err = Models.GetCustomerByID(&order.Customer, strconv.Itoa(int(order.CustomerID)))
	err = Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


//Checking cool down
func CheckCoolDown(customerId int) bool{
	var order Models.Order
	Config.DB.Model(&order).Where("customer_id = ?",customerId).Last(&order)

	if order.OrderId == 0{
		return true
	}

	currentTime := time.Now()
	diffTime := currentTime.Sub(order.CreatedAt).Seconds()

	if diffTime <= 60{
		return false
	}
	return true
}

//Checking if it is feasible to place order
func checkfeasibilty(order Models.Order, c *gin.Context) bool{
	id := order.ProductID
	var prod Models.Product
	err := Models.GetProductByID(&prod,strconv.Itoa(int(id)))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}

	if prod.QuantityAvail < order.Quantity{
		order.OrderStatus = "Failed"
		return false
	}

	prod.QuantityAvail -= order.Quantity
	order.OrderStatus = "Processed"
	err = Models.UpdateProduct(&prod,strconv.Itoa(int(id)))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}
	return true
}