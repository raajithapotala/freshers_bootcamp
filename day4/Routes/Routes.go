package Routes

import (
	"freshers_bootcamp/day4/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//For accessing the Product Database
	prod := r.Group("/product-api")
	{
		prod.GET("product", Controllers.GetProducts)
		prod.POST("product", Controllers.CreateProduct)
		prod.GET("product/:id", Controllers.GetProductByID)
		prod.PATCH("product/:id", Controllers.UpdateProduct)
		prod.DELETE("product/:id", Controllers.DeleteProduct)
	}

    // For accessing the Customer Database
	cust := r.Group("/customer-api")
	{
		cust.GET("customer", Controllers.GetCustomers)
		cust.POST("customer", Controllers.CreateCustomer)
		cust.GET("customer/:id", Controllers.GetCustomerByID)
	}

	//For accessing the Orders Database
	order := r.Group("/order-api")
	{
		order.GET("order", Controllers.GetOrders)
		order.POST("order", Controllers.CreateOrder)
		order.GET("order/:id", Controllers.GetOrderByID)
	}

	return r
}
