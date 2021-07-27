package Routes

import (
	"freshers_bootcamp/day4/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/product-api")
	{
		grp1.GET("product", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PATCH("product/:id", Controllers.UpdateProduct)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)
	}

	grp2 := r.Group("/customer-api")
	{
		grp2.GET("customer", Controllers.GetCustomers)
		grp2.POST("customer", Controllers.CreateCustomer)
		grp2.GET("customer/:id", Controllers.GetCustomerByID)
	}

	grp3 := r.Group("/order-api")
	{
		grp3.GET("order", Controllers.GetOrders)
		grp3.POST("order", Controllers.CreateOrder)
		grp3.GET("order/:id", Controllers.GetOrderByID)
	}

	return r
}