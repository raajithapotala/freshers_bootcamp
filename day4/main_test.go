package main

import (
	"bytes"
	"encoding/json"
	"freshers_bootcamp/day4/Models"
	"net/http"
	"net/http/httptest"
	"testing"

	"freshers_bootcamp/day4/Config"
	"freshers_bootcamp/day4/Controllers"
	"freshers_bootcamp/day4/Routes"

	"github.com/jinzhu/gorm"
)

//Testing the getProduct api
func TestGetProduct(t *testing.T) {
	//SQL Connection using GORM
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//Setting the router
	router := Routes.SetupRouter()
	router.GET("/product-api/product/", Controllers.GetProducts)

	//Get request
	request, _ := http.NewRequest("GET", "/product-api/product/", nil)

	//Recording the response
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `[{"prodId":1,"name":"bottle","description":"color-grey","quantityAvail":3,"price":250},{"prodId":2,"name":"","description":"","quantityAvail":0,"price":0},{"prodId":56,"name":"charger","description":"color-white","quantityAvail":5,"price":1250},{"prodId":76,"name":"charger","description":"color-white","quantityAvail":5,"price":1250},{"prodId":78,"name":"xyz","description":"","quantityAvail":67,"price":0},{"prodId":89,"name":"charger","description":"color-white","quantityAvail":7,"price":1250},{"prodId":90,"name":"charger","description":"color-white","quantityAvail":7,"price":250},{"prodId":159,"name":"bottle","description":"color-grey","quantityAvail":3,"price":250},{"prodId":190,"name":"charger","description":"color-white","quantityAvail":71,"price":250},{"prodId":191,"name":"","description":"","quantityAvail":0,"price":0}]`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

//Testing the creating product data
func TestCreateProd(t *testing.T) {
	//SQL Connection using GORM
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.Product{})

	//Setting the router
	router := Routes.SetupRouter()
	router.POST("/product-api/product/", Controllers.CreateProduct)

	//send request
	newprod := Models.Product{
		Id: 130,
		Name: "Sample",
		Description: "test",
		QuantityAvail: 5,
		Price: 128,
	}

	responseBody, _ := json.Marshal(newprod)
	req, _ := http.NewRequest("POST", "/product-api/product/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `{"description":"test","message":"product added successfully","name":"Sample","price":128,"prodId":130,"quantityAvail":5}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

// Testing the updation api
func TestUpdateProduct(t *testing.T) {
	//SQL database using GORM
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.Product{})

	//setting up router
	router := Routes.SetupRouter()
	router.PATCH("product-api/product/130/", Controllers.UpdateProduct)

	//send request
	updated := Models.Product{
		Price : 147,
	}

	responseBody, _ := json.Marshal(updated)
	req, _ := http.NewRequest("PATCH", "/product-api/product/130/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	expectedOutput := `{"prodId":130,"name":"Sample","description":"test","quantityAvail":5,"price":149}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}
}

//Testing the delete API
func TestDeleteProduct(t *testing.T) {
	//SQL Connection using GORM
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.DELETE("/product-api/product/", Controllers.DeleteProduct)

	//Get request
	req, _ := http.NewRequest("DELETE", "/product-api/product/130/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	//checking test case
	if response.Code != 307 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Code, 200)
	}
}

//Testing the get customer api
func TestGetCustomer(t *testing.T) {
	//SQL Connection using GORM
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//Setting the router
	router := Routes.SetupRouter()
	router.GET("/customer-api/customer/", Controllers.GetCustomers)

	//Get request
	request, _ := http.NewRequest("GET", "/customer-api/customer/", nil)

	//Recording the response
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `[{"id":123,"name":"Raaj","email":"xyz","phone":"7014401231","location":"India"},{"id":1234,"name":"Raaj","email":"xyz","phone":"7014401231","location":"India"}]`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

//Testing the create api for customer database
func TestCreateCust(t *testing.T) {
	//SQL Connection using GORM
	Config.DB, _ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.Customer{})

	//Setting the router
	router := Routes.SetupRouter()
	router.POST("/customer-api/customer/", Controllers.CreateCustomer)

	//send request
	newcust := Models.Customer{
		Id: 131,
		Name: "Sample",
		Email : "hello",
		Phone : "898764567",
		Location : "US",
	}

	responseBody, _ := json.Marshal(newcust)
	req, _ := http.NewRequest("POST", "/customer-api/customer/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `{"id":131,"name":"Sample","email":"hello","phone":"898764567","location":"US"}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}